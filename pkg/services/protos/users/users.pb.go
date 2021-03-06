// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/services/protos/users/users.proto

/*
	Package users is a generated protocol buffer package.

	It is generated from these files:
		pkg/services/protos/users/users.proto

	It has these top-level messages:
		User
*/
package users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import certificates "github.com/unerror/waffy/pkg/services/protos/certificates"

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

// Role represents the access level of the user
type Role int32

const (
	Role_USER  Role = 0
	Role_ADMIN Role = 1
)

var Role_name = map[int32]string{
	0: "USER",
	1: "ADMIN",
}
var Role_value = map[string]int32{
	"USER":  0,
	"ADMIN": 1,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}
func (Role) EnumDescriptor() ([]byte, []int) { return fileDescriptorUsers, []int{0} }

// User is a user who can access the system
type User struct {
	Email       string                    `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name        string                    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Role        Role                      `protobuf:"varint,3,opt,name=role,proto3,enum=users.Role" json:"role,omitempty"`
	Certificate *certificates.Certificate `protobuf:"bytes,4,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorUsers, []int{0} }

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_USER
}

func (m *User) GetCertificate() *certificates.Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "users.User")
	proto.RegisterEnum("users.Role", Role_name, Role_value)
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Email) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUsers(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUsers(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Role != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUsers(dAtA, i, uint64(m.Role))
	}
	if m.Certificate != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintUsers(dAtA, i, uint64(m.Certificate.Size()))
		n1, err := m.Certificate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeFixed64Users(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Users(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintUsers(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUsers(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUsers(uint64(l))
	}
	if m.Role != 0 {
		n += 1 + sovUsers(uint64(m.Role))
	}
	if m.Certificate != nil {
		l = m.Certificate.Size()
		n += 1 + l + sovUsers(uint64(l))
	}
	return n
}

func sovUsers(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUsers(x uint64) (n int) {
	return sovUsers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsers
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
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
				return ErrInvalidLengthUsers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
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
				return ErrInvalidLengthUsers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= (Role(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
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
				return ErrInvalidLengthUsers
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Certificate == nil {
				m.Certificate = &certificates.Certificate{}
			}
			if err := m.Certificate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUsers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUsers
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
func skipUsers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUsers
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
					return 0, ErrIntOverflowUsers
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
					return 0, ErrIntOverflowUsers
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
				return 0, ErrInvalidLengthUsers
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUsers
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
				next, err := skipUsers(dAtA[start:])
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
	ErrInvalidLengthUsers = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUsers   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pkg/services/protos/users/users.proto", fileDescriptorUsers) }

var fileDescriptorUsers = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0xc8, 0x4e, 0xd7,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6,
	0x2f, 0x2d, 0x4e, 0x2d, 0x82, 0x92, 0x7a, 0x60, 0x21, 0x21, 0x56, 0x30, 0x47, 0xca, 0x27, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x34, 0x2f, 0xb5, 0xa8, 0x28, 0xbf,
	0x48, 0xbf, 0x3c, 0x31, 0x2d, 0xad, 0x52, 0x1f, 0x9b, 0x31, 0xc9, 0xa9, 0x45, 0x25, 0x99, 0x69,
	0x99, 0xc9, 0x89, 0x25, 0xa9, 0xa8, 0x1c, 0x88, 0xa1, 0x4a, 0x7d, 0x8c, 0x5c, 0x2c, 0xa1, 0xc5,
	0xa9, 0x45, 0x42, 0x22, 0x5c, 0xac, 0xa9, 0xb9, 0x89, 0x99, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x10, 0x8e, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x10,
	0xcc, 0x16, 0x92, 0xe7, 0x62, 0x29, 0xca, 0xcf, 0x49, 0x95, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x33,
	0xe2, 0xd6, 0x83, 0xb8, 0x31, 0x28, 0x3f, 0x27, 0x35, 0x08, 0x2c, 0x21, 0x64, 0xcd, 0xc5, 0x8d,
	0x64, 0x93, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xa4, 0x1e, 0x8a, 0xed, 0xce, 0x08, 0x4e,
	0x10, 0xb2, 0x6a, 0x2d, 0x69, 0x2e, 0x16, 0x90, 0x51, 0x42, 0x1c, 0x5c, 0x2c, 0xa1, 0xc1, 0xae,
	0x41, 0x02, 0x0c, 0x42, 0x9c, 0x5c, 0xac, 0x8e, 0x2e, 0xbe, 0x9e, 0x7e, 0x02, 0x8c, 0x4e, 0x02,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72,
	0x0c, 0x49, 0x6c, 0x60, 0x6f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x54, 0xad, 0x63, 0x91,
	0x44, 0x01, 0x00, 0x00,
}
