// Code generated by protoc-gen-gogo.
// source: Messages.proto
// DO NOT EDIT!

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		Messages.proto
		Profile.proto

	It has these top-level messages:
		ServerToClient
		ClientToServer
*/
package proto

import proto1 "code.google.com/p/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type ServerToClient_StatusCode int32

const (
	ServerToClient_OK          ServerToClient_StatusCode = 0
	ServerToClient_PARSE_ERROR ServerToClient_StatusCode = 1
)

var ServerToClient_StatusCode_name = map[int32]string{
	0: "OK",
	1: "PARSE_ERROR",
}
var ServerToClient_StatusCode_value = map[string]int32{
	"OK":          0,
	"PARSE_ERROR": 1,
}

func (x ServerToClient_StatusCode) Enum() *ServerToClient_StatusCode {
	p := new(ServerToClient_StatusCode)
	*p = x
	return p
}
func (x ServerToClient_StatusCode) String() string {
	return proto1.EnumName(ServerToClient_StatusCode_name, int32(x))
}
func (x *ServerToClient_StatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(ServerToClient_StatusCode_value, data, "ServerToClient_StatusCode")
	if err != nil {
		return err
	}
	*x = ServerToClient_StatusCode(value)
	return nil
}

type ServerToClient struct {
	Status           *ServerToClient_StatusCode `protobuf:"varint,1,req,name=status,enum=proto.ServerToClient_StatusCode" json:"status,omitempty"`
	MessageList      [][]byte                   `protobuf:"bytes,3,rep,name=message_list" json:"message_list"`
	Envelope         []byte                     `protobuf:"bytes,4,opt,name=envelope" json:"envelope,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ServerToClient) Reset()         { *m = ServerToClient{} }
func (m *ServerToClient) String() string { return proto1.CompactTextString(m) }
func (*ServerToClient) ProtoMessage()    {}

type ClientToServer struct {
	CreateAccount    *bool                           `protobuf:"varint,1,opt,name=create_account" json:"create_account,omitempty"`
	DeliverEnvelope  *ClientToServer_DeliverEnvelope `protobuf:"bytes,2,opt,name=deliver_envelope" json:"deliver_envelope,omitempty"`
	DownloadEnvelope []byte                          `protobuf:"bytes,6,opt,name=download_envelope" json:"download_envelope,omitempty"`
	ListMessages     *bool                           `protobuf:"varint,5,opt,name=list_messages" json:"list_messages,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *ClientToServer) Reset()         { *m = ClientToServer{} }
func (m *ClientToServer) String() string { return proto1.CompactTextString(m) }
func (*ClientToServer) ProtoMessage()    {}

type ClientToServer_DeliverEnvelope struct {
	User             *Byte32 `protobuf:"bytes,3,req,customtype=Byte32" json:"User,omitempty"`
	Envelope         []byte  `protobuf:"bytes,4,req" json:"Envelope,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientToServer_DeliverEnvelope) Reset()         { *m = ClientToServer_DeliverEnvelope{} }
func (m *ClientToServer_DeliverEnvelope) String() string { return proto1.CompactTextString(m) }
func (*ClientToServer_DeliverEnvelope) ProtoMessage()    {}

func init() {
	proto1.RegisterEnum("proto.ServerToClient_StatusCode", ServerToClient_StatusCode_name, ServerToClient_StatusCode_value)
}
func (m *ServerToClient) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v ServerToClient_StatusCode
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (ServerToClient_StatusCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Status = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageList", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageList = append(m.MessageList, make([]byte, postIndex-index))
			copy(m.MessageList[len(m.MessageList)-1], data[index:postIndex])
			index = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envelope", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Envelope = append(m.Envelope, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *ClientToServer) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAccount", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.CreateAccount = &b
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverEnvelope", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeliverEnvelope == nil {
				m.DeliverEnvelope = &ClientToServer_DeliverEnvelope{}
			}
			if err := m.DeliverEnvelope.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownloadEnvelope", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DownloadEnvelope = append(m.DownloadEnvelope, data[index:postIndex]...)
			index = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMessages", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.ListMessages = &b
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *ClientToServer_DeliverEnvelope) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = &Byte32{}
			if err := m.User.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envelope", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Envelope = append(m.Envelope, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *ServerToClient) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		n += 1 + sovMessages(uint64(*m.Status))
	}
	if len(m.MessageList) > 0 {
		for _, b := range m.MessageList {
			l = len(b)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	if m.Envelope != nil {
		l = len(m.Envelope)
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *ClientToServer) Size() (n int) {
	var l int
	_ = l
	if m.CreateAccount != nil {
		n += 2
	}
	if m.DeliverEnvelope != nil {
		l = m.DeliverEnvelope.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.DownloadEnvelope != nil {
		l = len(m.DownloadEnvelope)
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.ListMessages != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *ClientToServer_DeliverEnvelope) Size() (n int) {
	var l int
	_ = l
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Envelope != nil {
		l = len(m.Envelope)
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessages(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedServerToClient(r randyMessages, easy bool) *ServerToClient {
	this := &ServerToClient{}
	v1 := ServerToClient_StatusCode([]int32{0, 1}[r.Intn(2)])
	this.Status = &v1
	if r.Intn(10) != 0 {
		v2 := r.Intn(100)
		this.MessageList = make([][]byte, v2)
		for i := 0; i < v2; i++ {
			v3 := r.Intn(100)
			this.MessageList[i] = make([]byte, v3)
			for j := 0; j < v3; j++ {
				this.MessageList[i][j] = byte(r.Intn(256))
			}
		}
	}
	if r.Intn(10) != 0 {
		v4 := r.Intn(100)
		this.Envelope = make([]byte, v4)
		for i := 0; i < v4; i++ {
			this.Envelope[i] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessages(r, 5)
	}
	return this
}

func NewPopulatedClientToServer(r randyMessages, easy bool) *ClientToServer {
	this := &ClientToServer{}
	if r.Intn(10) != 0 {
		v5 := bool(r.Intn(2) == 0)
		this.CreateAccount = &v5
	}
	if r.Intn(10) != 0 {
		this.DeliverEnvelope = NewPopulatedClientToServer_DeliverEnvelope(r, easy)
	}
	if r.Intn(10) != 0 {
		v6 := r.Intn(100)
		this.DownloadEnvelope = make([]byte, v6)
		for i := 0; i < v6; i++ {
			this.DownloadEnvelope[i] = byte(r.Intn(256))
		}
	}
	if r.Intn(10) != 0 {
		v7 := bool(r.Intn(2) == 0)
		this.ListMessages = &v7
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessages(r, 7)
	}
	return this
}

func NewPopulatedClientToServer_DeliverEnvelope(r randyMessages, easy bool) *ClientToServer_DeliverEnvelope {
	this := &ClientToServer_DeliverEnvelope{}
	this.User = NewPopulatedByte32(r)
	v8 := r.Intn(100)
	this.Envelope = make([]byte, v8)
	for i := 0; i < v8; i++ {
		this.Envelope[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessages(r, 5)
	}
	return this
}

type randyMessages interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMessages(r randyMessages) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringMessages(r randyMessages) string {
	v9 := r.Intn(100)
	tmps := make([]rune, v9)
	for i := 0; i < v9; i++ {
		tmps[i] = randUTF8RuneMessages(r)
	}
	return string(tmps)
}
func randUnrecognizedMessages(r randyMessages, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldMessages(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldMessages(data []byte, r randyMessages, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateMessages(data, uint64(key))
		v10 := r.Int63()
		if r.Intn(2) == 0 {
			v10 *= -1
		}
		data = encodeVarintPopulateMessages(data, uint64(v10))
	case 1:
		data = encodeVarintPopulateMessages(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateMessages(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateMessages(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateMessages(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateMessages(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *ServerToClient) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ServerToClient) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		data[i] = 0x8
		i++
		i = encodeVarintMessages(data, i, uint64(*m.Status))
	}
	if len(m.MessageList) > 0 {
		for _, b := range m.MessageList {
			data[i] = 0x1a
			i++
			i = encodeVarintMessages(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if m.Envelope != nil {
		data[i] = 0x22
		i++
		i = encodeVarintMessages(data, i, uint64(len(m.Envelope)))
		i += copy(data[i:], m.Envelope)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *ClientToServer) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClientToServer) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CreateAccount != nil {
		data[i] = 0x8
		i++
		if *m.CreateAccount {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.DeliverEnvelope != nil {
		data[i] = 0x12
		i++
		i = encodeVarintMessages(data, i, uint64(m.DeliverEnvelope.Size()))
		n1, err := m.DeliverEnvelope.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.DownloadEnvelope != nil {
		data[i] = 0x32
		i++
		i = encodeVarintMessages(data, i, uint64(len(m.DownloadEnvelope)))
		i += copy(data[i:], m.DownloadEnvelope)
	}
	if m.ListMessages != nil {
		data[i] = 0x28
		i++
		if *m.ListMessages {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *ClientToServer_DeliverEnvelope) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClientToServer_DeliverEnvelope) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.User != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintMessages(data, i, uint64(m.User.Size()))
		n2, err := m.User.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Envelope != nil {
		data[i] = 0x22
		i++
		i = encodeVarintMessages(data, i, uint64(len(m.Envelope)))
		i += copy(data[i:], m.Envelope)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64Messages(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Messages(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMessages(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *ServerToClient) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ServerToClient)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Status != nil && that1.Status != nil {
		if *this.Status != *that1.Status {
			return false
		}
	} else if this.Status != nil {
		return false
	} else if that1.Status != nil {
		return false
	}
	if len(this.MessageList) != len(that1.MessageList) {
		return false
	}
	for i := range this.MessageList {
		if !bytes.Equal(this.MessageList[i], that1.MessageList[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Envelope, that1.Envelope) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ClientToServer) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClientToServer)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.CreateAccount != nil && that1.CreateAccount != nil {
		if *this.CreateAccount != *that1.CreateAccount {
			return false
		}
	} else if this.CreateAccount != nil {
		return false
	} else if that1.CreateAccount != nil {
		return false
	}
	if !this.DeliverEnvelope.Equal(that1.DeliverEnvelope) {
		return false
	}
	if !bytes.Equal(this.DownloadEnvelope, that1.DownloadEnvelope) {
		return false
	}
	if this.ListMessages != nil && that1.ListMessages != nil {
		if *this.ListMessages != *that1.ListMessages {
			return false
		}
	} else if this.ListMessages != nil {
		return false
	} else if that1.ListMessages != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ClientToServer_DeliverEnvelope) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClientToServer_DeliverEnvelope)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if that1.User == nil {
		if this.User != nil {
			return false
		}
	} else if !this.User.Equal(*that1.User) {
		return false
	}
	if !bytes.Equal(this.Envelope, that1.Envelope) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
