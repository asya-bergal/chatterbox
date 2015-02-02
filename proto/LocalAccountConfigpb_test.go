// Code generated by protoc-gen-gogo.
// source: LocalAccountConfig.proto
// DO NOT EDIT!

package proto

import testing12 "testing"
import math_rand12 "math/rand"
import time12 "time"
import code_google_com_p_gogoprotobuf_proto9 "code.google.com/p/gogoprotobuf/proto"
import testing13 "testing"
import math_rand13 "math/rand"
import time13 "time"
import encoding_json3 "encoding/json"
import testing14 "testing"
import math_rand14 "math/rand"
import time14 "time"
import code_google_com_p_gogoprotobuf_proto10 "code.google.com/p/gogoprotobuf/proto"
import math_rand15 "math/rand"
import time15 "time"
import testing15 "testing"
import code_google_com_p_gogoprotobuf_proto11 "code.google.com/p/gogoprotobuf/proto"

func TestLocalAccountConfigProto(t *testing12.T) {
	popr := math_rand12.New(math_rand12.NewSource(time12.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto9.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &LocalAccountConfig{}
	if err := code_google_com_p_gogoprotobuf_proto9.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestLocalAccountConfigMarshalTo(t *testing12.T) {
	popr := math_rand12.New(math_rand12.NewSource(time12.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &LocalAccountConfig{}
	if err := code_google_com_p_gogoprotobuf_proto9.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkLocalAccountConfigProtoMarshal(b *testing12.B) {
	popr := math_rand12.New(math_rand12.NewSource(616))
	total := 0
	pops := make([]*LocalAccountConfig, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedLocalAccountConfig(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto9.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkLocalAccountConfigProtoUnmarshal(b *testing12.B) {
	popr := math_rand12.New(math_rand12.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto9.Marshal(NewPopulatedLocalAccountConfig(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &LocalAccountConfig{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto9.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestLocalAccountConfigJSON(t *testing13.T) {
	popr := math_rand13.New(math_rand13.NewSource(time13.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, true)
	jsondata, err := encoding_json3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &LocalAccountConfig{}
	err = encoding_json3.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestLocalAccountConfigProtoText(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, true)
	data := code_google_com_p_gogoprotobuf_proto10.MarshalTextString(p)
	msg := &LocalAccountConfig{}
	if err := code_google_com_p_gogoprotobuf_proto10.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestLocalAccountConfigProtoCompactText(t *testing14.T) {
	popr := math_rand14.New(math_rand14.NewSource(time14.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, true)
	data := code_google_com_p_gogoprotobuf_proto10.CompactTextString(p)
	msg := &LocalAccountConfig{}
	if err := code_google_com_p_gogoprotobuf_proto10.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestLocalAccountConfigSize(t *testing15.T) {
	popr := math_rand15.New(math_rand15.NewSource(time15.Now().UnixNano()))
	p := NewPopulatedLocalAccountConfig(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto11.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto11.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto11.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkLocalAccountConfigSize(b *testing15.B) {
	popr := math_rand15.New(math_rand15.NewSource(616))
	total := 0
	pops := make([]*LocalAccountConfig, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedLocalAccountConfig(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
