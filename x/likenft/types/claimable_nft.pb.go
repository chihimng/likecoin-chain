// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likenft/claimable_nft.proto

package types

import (
	fmt "fmt"
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

type ClaimableNFT struct {
	ClassId string `protobuf:"bytes,1,opt,name=classId,proto3" json:"classId,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Input   string `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
}

func (m *ClaimableNFT) Reset()         { *m = ClaimableNFT{} }
func (m *ClaimableNFT) String() string { return proto.CompactTextString(m) }
func (*ClaimableNFT) ProtoMessage()    {}
func (*ClaimableNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_e273e0179f129bfd, []int{0}
}
func (m *ClaimableNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimableNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimableNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimableNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimableNFT.Merge(m, src)
}
func (m *ClaimableNFT) XXX_Size() int {
	return m.Size()
}
func (m *ClaimableNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimableNFT.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimableNFT proto.InternalMessageInfo

func (m *ClaimableNFT) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *ClaimableNFT) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ClaimableNFT) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func init() {
	proto.RegisterType((*ClaimableNFT)(nil), "likecoin.likechain.likenft.ClaimableNFT")
}

func init() { proto.RegisterFile("likenft/claimable_nft.proto", fileDescriptor_e273e0179f129bfd) }

var fileDescriptor_e273e0179f129bfd = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xc9, 0xcc, 0x4e,
	0xcd, 0x4b, 0x2b, 0xd1, 0x4f, 0xce, 0x49, 0xcc, 0xcc, 0x4d, 0x4c, 0xca, 0x49, 0x8d, 0xcf, 0x4b,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x02, 0x49, 0x26, 0xe7, 0x67, 0xe6, 0xe9,
	0x81, 0x19, 0x19, 0x89, 0x50, 0x56, 0x5e, 0x5a, 0x89, 0x92, 0x1f, 0x17, 0x8f, 0x33, 0x4c, 0x8b,
	0x9f, 0x5b, 0x88, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x4e, 0x62, 0x71, 0xb1, 0x67, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0x04, 0x16,
	0x64, 0xca, 0x4c, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x2b, 0x28, 0x2d, 0x91, 0x60, 0x06, 0x0b,
	0x41, 0x38, 0x4e, 0xee, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9b,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x73, 0x90, 0x3e, 0xdc, 0x41,
	0xfa, 0x15, 0xfa, 0x30, 0x2f, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xdd, 0x6e, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x6b, 0xb5, 0xa8, 0xda, 0x00, 0x00, 0x00,
}

func (m *ClaimableNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimableNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimableNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Input) > 0 {
		i -= len(m.Input)
		copy(dAtA[i:], m.Input)
		i = encodeVarintClaimableNft(dAtA, i, uint64(len(m.Input)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintClaimableNft(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintClaimableNft(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimableNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimableNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimableNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovClaimableNft(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovClaimableNft(uint64(l))
	}
	l = len(m.Input)
	if l > 0 {
		n += 1 + l + sovClaimableNft(uint64(l))
	}
	return n
}

func sovClaimableNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimableNft(x uint64) (n int) {
	return sovClaimableNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimableNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimableNft
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
			return fmt.Errorf("proto: ClaimableNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimableNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimableNft
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
				return ErrInvalidLengthClaimableNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimableNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimableNft
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
				return ErrInvalidLengthClaimableNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimableNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimableNft
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
				return ErrInvalidLengthClaimableNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimableNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Input = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaimableNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimableNft
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
func skipClaimableNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimableNft
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
					return 0, ErrIntOverflowClaimableNft
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
					return 0, ErrIntOverflowClaimableNft
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
				return 0, ErrInvalidLengthClaimableNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimableNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimableNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimableNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimableNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimableNft = fmt.Errorf("proto: unexpected end of group")
)
