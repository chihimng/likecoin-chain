// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/iscn/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	RegistryName string        `protobuf:"bytes,1,opt,name=registry_name,json=registryName,proto3" json:"registry_name,omitempty"`
	FeePerByte   types.DecCoin `protobuf:"bytes,2,opt,name=fee_per_byte,json=feePerByte,proto3" json:"fee_per_byte"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a4c68825ff5be8, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "likechain.iscn.Params")
}

func init() { proto.RegisterFile("likechain/iscn/params.proto", fileDescriptor_a4a4c68825ff5be8) }

var fileDescriptor_a4a4c68825ff5be8 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0xef, 0x44, 0x0a, 0xc6, 0x2a, 0x12, 0x1c, 0x4a, 0x95, 0x4b, 0xd0, 0xa5, 0x38, 0xdc,
	0x51, 0xdd, 0x1c, 0x63, 0x66, 0x29, 0x1d, 0x5d, 0xc2, 0xe5, 0x78, 0x4d, 0x0f, 0xcd, 0x5d, 0xb8,
	0x3b, 0xc5, 0x6c, 0x2e, 0x82, 0xa3, 0xa3, 0x63, 0x3e, 0x4e, 0xc7, 0x8e, 0x0e, 0x22, 0x92, 0x2c,
	0x7e, 0x0c, 0x49, 0xa2, 0x76, 0x7b, 0xbc, 0xdf, 0xe3, 0xfd, 0xf8, 0xff, 0xbd, 0xa3, 0x3b, 0x79,
	0x0b, 0x62, 0xc9, 0xa5, 0x62, 0xd2, 0x0a, 0xc5, 0x0a, 0x6e, 0x78, 0x6e, 0x69, 0x61, 0xb4, 0xd3,
	0xfe, 0xfe, 0x3f, 0xa4, 0x2d, 0x1c, 0x1f, 0x66, 0x3a, 0xd3, 0x1d, 0x62, 0xed, 0xd4, 0x5f, 0x8d,
	0x89, 0xd0, 0x36, 0xd7, 0x96, 0xa5, 0xdc, 0x02, 0x7b, 0x98, 0xa6, 0xe0, 0xf8, 0x94, 0x09, 0x2d,
	0x55, 0xcf, 0x4f, 0x9e, 0xb1, 0x37, 0x98, 0x75, 0x6f, 0xfd, 0x53, 0x6f, 0xcf, 0x40, 0x26, 0xad,
	0x33, 0x65, 0xa2, 0x78, 0x0e, 0x23, 0x1c, 0xe2, 0xc9, 0xce, 0x7c, 0xf8, 0xb7, 0xbc, 0xe6, 0x39,
	0xf8, 0xb1, 0x37, 0x5c, 0x00, 0x24, 0x05, 0x98, 0x24, 0x2d, 0x1d, 0x8c, 0xb6, 0x42, 0x3c, 0xd9,
	0x3d, 0x3f, 0xa6, 0xbd, 0x86, 0xb6, 0x1a, 0xfa, 0xab, 0xa1, 0x31, 0x88, 0x2b, 0x2d, 0x55, 0xb4,
	0xbd, 0xfa, 0x0c, 0xd0, 0xdc, 0x5b, 0x00, 0xcc, 0xc0, 0x44, 0xa5, 0x83, 0xcb, 0x83, 0x97, 0x2a,
	0x40, 0x6f, 0x55, 0x80, 0xbe, 0xab, 0x00, 0x3d, 0x7d, 0x84, 0x28, 0x8a, 0x57, 0x35, 0xc1, 0xeb,
	0x9a, 0xe0, 0xaf, 0x9a, 0xe0, 0xd7, 0x86, 0xa0, 0x75, 0x43, 0xd0, 0x7b, 0x43, 0xd0, 0xcd, 0x59,
	0x26, 0xdd, 0xf2, 0x3e, 0xa5, 0x42, 0xe7, 0xac, 0x8b, 0xac, 0xa5, 0x62, 0x9b, 0x62, 0x1e, 0xfb,
	0x6a, 0x5c, 0x59, 0x80, 0x4d, 0x07, 0x5d, 0xa8, 0x8b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c,
	0x5c, 0x9e, 0x55, 0x39, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FeePerByte.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RegistryName) > 0 {
		i -= len(m.RegistryName)
		copy(dAtA[i:], m.RegistryName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.RegistryName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RegistryName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.FeePerByte.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegistryName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePerByte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeePerByte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
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
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
