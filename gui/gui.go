//go:generate genqrc qml

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/andres-erbsen/chatterbox/client/persistence"
	"github.com/andres-erbsen/chatterbox/proto"
	"gopkg.in/fsnotify.v1"
	"gopkg.in/qml.v1"
)

var root = flag.String("root", "", "chatterbox root directory")

type gui struct {
	persistence.Paths
	engine *qml.Engine

	conversations        []*proto.ConversationMetadata
	conversationsIndex   map[string]int
	conversationsDisplay qml.Object

	watcher *fsnotify.Watcher

	openConversations map[string]*qml.Window

	stop chan struct{}
}

func main() {
	flag.Parse()
	if *root == "" {
		fmt.Fprintf(os.Stderr, "USAGE: %s -root=ROOTDIR", os.Args[0])
		os.Exit(1)
	}

	g := &gui{stop: make(chan struct{}), conversationsIndex: make(map[string]int), openConversations: make(map[string]*qml.Window)}
	g.Paths = persistence.Paths{
		RootDir:     *root,
		Application: "chat-create",
	}
	go g.watch()
	if err := qml.Run(g.run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
}

type Conversation struct {
	Subject     string
	Users       []string
	LastMessage string
}

func toJson(v interface{}) string {
	rawJson, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(rawJson)
}

func newConversation(engine *qml.Engine) error {
	controls, err := engine.LoadFile("qml/new-conversation.qml")
	if err != nil {
		return err
	}
	window := controls.CreateWindow(nil)

	window.On("sendMessage", func(to, subject, message string) {
		println("To: " + to)
		println("Subject: " + subject)
		println("Message: " + message)
	})

	return nil
}

func addMessage (window *qml.Window, msg *persistence.Message) {
	messageModel := window.ObjectByName("messageModel")
	msg.Content = strings.TrimSpace(msg.Content)
	messageModel.Call("addItem", toJson(msg))
	window.ObjectByName("messageView").Call("positionViewAtEnd")
}

func (g *gui) openConversation(idx int) error {
	controls, err := g.engine.LoadFile("qml/old-conversation.qml")
	if err != nil {
		return err
	}
	window := controls.CreateWindow(nil)

	conv := g.conversations[idx]

	//TODO: if an open conversation is selected again, focus that window

	qml.Lock()
	g.openConversations[persistence.ConversationName(conv)] = window;
	qml.Unlock()

	msgs, err := g.LoadMessages(conv)
	if err != nil {
		panic(err)
	}
	for _, msg := range msgs {
		addMessage(window, msg)
	}

	ctx := g.engine.Context()
	ctx.SetVar("textAreaCleared", false)

	messageArea := window.ObjectByName("messageArea")

	window.ObjectByName("messageArea").On("focusChanged", func() {
		if !(ctx.Var("textAreaCleared").(bool)) {
			messageArea.Call("selectAll")
			ctx.SetVar("textAreaCleared", true)
		}
	})

	window.On("sendMessage", func(message string) {
		println("Send: " + message)
		err := g.MessageToOutbox(persistence.ConversationName(conv), message)
		if err != nil {
			panic(err)
		}
	})

	window.On("closing", func() {
		qml.Lock()
		delete(g.openConversations, persistence.ConversationName(conv))
		qml.Unlock()
	})

	return nil
}

func (g *gui) run() error {
	defer close(g.stop)
	g.engine = qml.NewEngine()
	controls, err := g.engine.LoadFile("qml/history.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)
	g.conversationsDisplay = window.ObjectByName("listModel")
	convs, err := g.ListConversations()
	if err != nil {
		return err
	}
	for _, con := range convs {
		g.handleConversation(con)
	}

	table := window.ObjectByName("table")
	table.On("activated", g.openConversation)
	table.Set("focus", "true")

	window.Show()
	window.Wait()
	return nil
}

func (g *gui) handleConversation(con *proto.ConversationMetadata) {
	if _, already := g.conversationsIndex[persistence.ConversationName(con)]; already {
		return
	}
	qml.Lock()
	defer qml.Unlock()
	g.conversationsIndex[persistence.ConversationName(con)] = len(g.conversations)
	g.conversations = append(g.conversations, con)
	c := Conversation{Subject: con.Subject, Users: con.Participants}
	g.conversationsDisplay.Call("addItem", toJson(c))
}

func (g *gui) watch() {
	var err error
	g.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer g.watcher.Close()
	err = g.watcher.Add(g.ConversationDir())
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case <-g.stop:
			return
		case err := <-g.watcher.Errors:
			fmt.Println("error:", err)
		case e := <-g.watcher.Events:
			rpath, err := filepath.Rel(g.ConversationDir(), e.Name)
			if err != nil {
				panic(err)
			}
			if !(e.Op == fsnotify.Create || e.Op == fsnotify.Rename) {
				// TODO: handle move, delete
				continue
			}
			if match, _ := filepath.Match("*", rpath); match {
				// when a conversation is created it MUST have a metadata file when
				// it is moved to the conversations directory
				err = g.watcher.Add(g.ConversationDir())
				if err != nil {
					log.Printf("error watching conversation %s: %s\n", rpath, err)
					// continue after error
				}
				c, err := persistence.ReadConversationMetadata(e.Name)
				if err != nil {
					log.Printf("error reading metadata of %s: %s\n", rpath, err)
					continue
				}
				g.handleConversation(c)
			} else if match, _ := filepath.Match("*/*", rpath); match {
				// TODO: handle incoming message
				win := g.openConversations[filepath.Base(rpath)]
				addMessage(win, persistence.ReadMessageFromFile(rpath))

			} else {
				log.Printf("event at unknown path: %s", rpath)
			}
		}
	}
}
