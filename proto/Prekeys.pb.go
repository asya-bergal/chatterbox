// Code generated by protoc-gen-gogo.
// source: Prekeys.proto
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io5 "io"
import fmt5 "fmt"
import code_google_com_p_gogoprotobuf_proto5 "code.google.com/p/gogoprotobuf/proto"

import bytes5 "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type Prekeys struct {
	PrekeySecrets    []Byte32 `protobuf:"bytes,1,rep,customtype=Byte32" json:"PrekeySecrets"`
	PrekeyPublics    []Byte32 `protobuf:"bytes,2,rep,customtype=Byte32" json:"PrekeyPublics"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Prekeys) Reset()         { *m = Prekeys{} }
func (m *Prekeys) String() string { return proto1.CompactTextString(m) }
func (*Prekeys) ProtoMessage()    {}

func init() {
}
func (m *Prekeys) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io5.ErrUnexpectedEOF
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
			if wireType != 2 {
				return fmt5.Errorf("proto: wrong wireType = %d for field PrekeySecrets", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
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
				return io5.ErrUnexpectedEOF
			}
			m.PrekeySecrets = append(m.PrekeySecrets, Byte32{})
			m.PrekeySecrets[len(m.PrekeySecrets)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt5.Errorf("proto: wrong wireType = %d for field PrekeyPublics", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io5.ErrUnexpectedEOF
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
				return io5.ErrUnexpectedEOF
			}
			m.PrekeyPublics = append(m.PrekeyPublics, Byte32{})
			m.PrekeyPublics[len(m.PrekeyPublics)-1].Unmarshal(data[index:postIndex])
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
			skippy, err := code_google_com_p_gogoprotobuf_proto5.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io5.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Prekeys) Size() (n int) {
	var l int
	_ = l
	if len(m.PrekeySecrets) > 0 {
		for _, e := range m.PrekeySecrets {
			l = e.Size()
			n += 1 + l + sovPrekeys(uint64(l))
		}
	}
	if len(m.PrekeyPublics) > 0 {
		for _, e := range m.PrekeyPublics {
			l = e.Size()
			n += 1 + l + sovPrekeys(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPrekeys(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPrekeys(x uint64) (n int) {
	return sovPrekeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedPrekeys(r randyPrekeys, easy bool) *Prekeys {
	this := &Prekeys{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(10)
		this.PrekeySecrets = make([]Byte32, v1)
		for i := 0; i < v1; i++ {
			v2 := NewPopulatedByte32(r)
			this.PrekeySecrets[i] = *v2
		}
	}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.PrekeyPublics = make([]Byte32, v3)
		for i := 0; i < v3; i++ {
			v4 := NewPopulatedByte32(r)
			this.PrekeyPublics[i] = *v4
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedPrekeys(r, 3)
	}
	return this
}

type randyPrekeys interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePrekeys(r randyPrekeys) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringPrekeys(r randyPrekeys) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RunePrekeys(r)
	}
	return string(tmps)
}
func randUnrecognizedPrekeys(r randyPrekeys, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldPrekeys(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldPrekeys(data []byte, r randyPrekeys, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulatePrekeys(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulatePrekeys(data, uint64(v6))
	case 1:
		data = encodeVarintPopulatePrekeys(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulatePrekeys(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulatePrekeys(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulatePrekeys(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulatePrekeys(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *Prekeys) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Prekeys) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PrekeySecrets) > 0 {
		for _, msg := range m.PrekeySecrets {
			data[i] = 0xa
			i++
			i = encodeVarintPrekeys(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.PrekeyPublics) > 0 {
		for _, msg := range m.PrekeyPublics {
			data[i] = 0x12
			i++
			i = encodeVarintPrekeys(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64Prekeys(data []byte, offset int, v uint64) int {
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
func encodeFixed32Prekeys(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPrekeys(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *Prekeys) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Prekeys)
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
	if len(this.PrekeySecrets) != len(that1.PrekeySecrets) {
		return false
	}
	for i := range this.PrekeySecrets {
		if !this.PrekeySecrets[i].Equal(that1.PrekeySecrets[i]) {
			return false
		}
	}
	if len(this.PrekeyPublics) != len(that1.PrekeyPublics) {
		return false
	}
	for i := range this.PrekeyPublics {
		if !this.PrekeyPublics[i].Equal(that1.PrekeyPublics[i]) {
			return false
		}
	}
	if !bytes5.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
