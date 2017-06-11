// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/services/protos/sites/sites.proto

/*
	Package sites is a generated protocol buffer package.

	It is generated from these files:
		pkg/services/protos/sites/sites.proto

	It has these top-level messages:
		Site
		Balancer
*/
package sites

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import nodes "github.com/unerror/waffy/pkg/services/protos/nodes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Site represents a Site that should be load balanced, and have Rules applied to it
type Site struct {
	Hostname    string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Alias       []string `protobuf:"bytes,2,rep,name=alias" json:"alias,omitempty"`
	Secure      bool     `protobuf:"varint,5,opt,name=secure,proto3" json:"secure,omitempty"`
	Autoencrypt bool     `protobuf:"varint,6,opt,name=autoencrypt,proto3" json:"autoencrypt,omitempty"`
}

func (m *Site) Reset()                    { *m = Site{} }
func (m *Site) String() string            { return proto.CompactTextString(m) }
func (*Site) ProtoMessage()               {}
func (*Site) Descriptor() ([]byte, []int) { return fileDescriptorSites, []int{0} }

func (m *Site) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Site) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Site) GetSecure() bool {
	if m != nil {
		return m.Secure
	}
	return false
}

func (m *Site) GetAutoencrypt() bool {
	if m != nil {
		return m.Autoencrypt
	}
	return false
}

// Balancer represents a Site load balancer
type Balancer struct {
	Proto string        `protobuf:"bytes,3,opt,name=proto,proto3" json:"proto,omitempty"`
	Port  string        `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Sites []*Site       `protobuf:"bytes,1,rep,name=sites" json:"sites,omitempty"`
	Notes []*nodes.Node `protobuf:"bytes,2,rep,name=notes" json:"notes,omitempty"`
}

func (m *Balancer) Reset()                    { *m = Balancer{} }
func (m *Balancer) String() string            { return proto.CompactTextString(m) }
func (*Balancer) ProtoMessage()               {}
func (*Balancer) Descriptor() ([]byte, []int) { return fileDescriptorSites, []int{1} }

func (m *Balancer) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *Balancer) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Balancer) GetSites() []*Site {
	if m != nil {
		return m.Sites
	}
	return nil
}

func (m *Balancer) GetNotes() []*nodes.Node {
	if m != nil {
		return m.Notes
	}
	return nil
}

func init() {
	proto.RegisterType((*Site)(nil), "sites.Site")
	proto.RegisterType((*Balancer)(nil), "sites.Balancer")
}
func (m *Site) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Site) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hostname) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSites(dAtA, i, uint64(len(m.Hostname)))
		i += copy(dAtA[i:], m.Hostname)
	}
	if len(m.Alias) > 0 {
		for _, s := range m.Alias {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Secure {
		dAtA[i] = 0x28
		i++
		if m.Secure {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Autoencrypt {
		dAtA[i] = 0x30
		i++
		if m.Autoencrypt {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Balancer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Balancer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sites) > 0 {
		for _, msg := range m.Sites {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSites(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Notes) > 0 {
		for _, msg := range m.Notes {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSites(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Proto) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSites(dAtA, i, uint64(len(m.Proto)))
		i += copy(dAtA[i:], m.Proto)
	}
	if len(m.Port) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSites(dAtA, i, uint64(len(m.Port)))
		i += copy(dAtA[i:], m.Port)
	}
	return i, nil
}

func encodeFixed64Sites(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Sites(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSites(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Site) Size() (n int) {
	var l int
	_ = l
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovSites(uint64(l))
	}
	if len(m.Alias) > 0 {
		for _, s := range m.Alias {
			l = len(s)
			n += 1 + l + sovSites(uint64(l))
		}
	}
	if m.Secure {
		n += 2
	}
	if m.Autoencrypt {
		n += 2
	}
	return n
}

func (m *Balancer) Size() (n int) {
	var l int
	_ = l
	if len(m.Sites) > 0 {
		for _, e := range m.Sites {
			l = e.Size()
			n += 1 + l + sovSites(uint64(l))
		}
	}
	if len(m.Notes) > 0 {
		for _, e := range m.Notes {
			l = e.Size()
			n += 1 + l + sovSites(uint64(l))
		}
	}
	l = len(m.Proto)
	if l > 0 {
		n += 1 + l + sovSites(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovSites(uint64(l))
	}
	return n
}

func sovSites(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSites(x uint64) (n int) {
	return sovSites(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Site) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSites
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Site: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Site: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = append(m.Alias, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secure", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Secure = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Autoencrypt", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Autoencrypt = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSites(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSites
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Balancer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSites
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Balancer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Balancer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sites", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sites = append(m.Sites, &Site{})
			if err := m.Sites[len(m.Sites)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notes = append(m.Notes, &nodes.Node{})
			if err := m.Notes[len(m.Notes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proto", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proto = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSites(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSites
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSites(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSites
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSites
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSites
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSites
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSites
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSites(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSites = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSites   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pkg/services/protos/sites/sites.proto", fileDescriptorSites) }

var fileDescriptorSites = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0x9d, 0x36, 0x09, 0xe9, 0x64, 0x23, 0x83, 0xc8, 0xd0, 0x45, 0x88, 0x05, 0x21, 0xab,
	0x0c, 0xe8, 0xde, 0x45, 0x0f, 0xe0, 0x22, 0x9e, 0x60, 0x9a, 0xfe, 0x6d, 0x83, 0xed, 0xfc, 0xe1,
	0x9f, 0x89, 0x52, 0x4f, 0xe2, 0x91, 0x5c, 0x7a, 0x04, 0x89, 0x17, 0x91, 0xcc, 0x04, 0x71, 0xe1,
	0x66, 0x98, 0xef, 0xbd, 0xc7, 0xcf, 0x7b, 0xfc, 0xb6, 0x7b, 0xde, 0x2b, 0x0b, 0xf4, 0xd2, 0x36,
	0x60, 0x55, 0x47, 0xe8, 0xd0, 0x2a, 0xdb, 0x3a, 0x98, 0xde, 0xca, 0x4b, 0x22, 0xf6, 0xb0, 0x7c,
	0xd8, 0xb7, 0xee, 0xd0, 0x6f, 0xaa, 0x06, 0x4f, 0xaa, 0x37, 0x40, 0x84, 0xa4, 0x5e, 0xf5, 0x6e,
	0x77, 0x56, 0xff, 0x9d, 0x31, 0xb8, 0x85, 0xe9, 0x0d, 0x67, 0x56, 0xc4, 0xa3, 0xa7, 0xd6, 0x81,
	0x58, 0xf2, 0xf4, 0x80, 0xd6, 0x19, 0x7d, 0x02, 0xc9, 0x0a, 0x56, 0x2e, 0xea, 0x5f, 0x16, 0x57,
	0x3c, 0xd6, 0xc7, 0x56, 0x5b, 0x39, 0x2b, 0xe6, 0xe5, 0xa2, 0x0e, 0x20, 0xae, 0x79, 0x62, 0xa1,
	0xe9, 0x09, 0x64, 0x5c, 0xb0, 0x32, 0xad, 0x27, 0x12, 0x05, 0xcf, 0x74, 0xef, 0x10, 0x4c, 0x43,
	0xe7, 0xce, 0xc9, 0xc4, 0x9b, 0x7f, 0xa5, 0xd5, 0x1b, 0x4f, 0xd7, 0xfa, 0xa8, 0x4d, 0x03, 0x24,
	0x6e, 0x78, 0x18, 0x22, 0x59, 0x31, 0x2f, 0xb3, 0xbb, 0xac, 0x0a, 0x1b, 0xc7, 0x4e, 0x75, 0x70,
	0xc6, 0x88, 0xc1, 0x31, 0x32, 0x9b, 0x22, 0xa1, 0xff, 0x23, 0x6e, 0xa1, 0x0e, 0xce, 0xd8, 0xd0,
	0xcf, 0x91, 0x73, 0x5f, 0x3d, 0x80, 0x10, 0x3c, 0xea, 0x90, 0x9c, 0x8c, 0xbc, 0xe8, 0xff, 0xeb,
	0xcb, 0x8f, 0x21, 0x67, 0x9f, 0x43, 0xce, 0xbe, 0x86, 0x9c, 0xbd, 0x7f, 0xe7, 0x17, 0x9b, 0xc4,
	0x87, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42, 0x16, 0xc9, 0x0a, 0x78, 0x01, 0x00, 0x00,
}
