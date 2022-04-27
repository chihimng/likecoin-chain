// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likenft/claimable_nft.proto

package types

import (
	fmt "fmt"
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

type MintableNFT struct {
	ClassId string   `protobuf:"bytes,1,opt,name=classId,proto3" json:"classId,omitempty"`
	Id      string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Input   NFTInput `protobuf:"bytes,3,opt,name=input,proto3" json:"input"`
}

func (m *MintableNFT) Reset()         { *m = MintableNFT{} }
func (m *MintableNFT) String() string { return proto.CompactTextString(m) }
func (*MintableNFT) ProtoMessage()    {}
func (*MintableNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_e273e0179f129bfd, []int{0}
}
func (m *MintableNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintableNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintableNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintableNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintableNFT.Merge(m, src)
}
func (m *MintableNFT) XXX_Size() int {
	return m.Size()
}
func (m *MintableNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_MintableNFT.DiscardUnknown(m)
}

var xxx_messageInfo_MintableNFT proto.InternalMessageInfo

func (m *MintableNFT) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *MintableNFT) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MintableNFT) GetInput() NFTInput {
	if m != nil {
		return m.Input
	}
	return NFTInput{}
}

func init() {
	proto.RegisterType((*MintableNFT)(nil), "likecoin.likechain.likenft.MintableNFT")
}

func init() { proto.RegisterFile("likenft/claimable_nft.proto", fileDescriptor_e273e0179f129bfd) }

var fileDescriptor_e273e0179f129bfd = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xc9, 0xcc, 0x4e,
	0xcd, 0x4b, 0x2b, 0xd1, 0x4f, 0xce, 0x49, 0xcc, 0xcc, 0x4d, 0x4c, 0xca, 0x49, 0x8d, 0xcf, 0x4b,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x02, 0x49, 0x26, 0xe7, 0x67, 0xe6, 0xe9,
	0x81, 0x19, 0x19, 0x89, 0x50, 0x56, 0x5e, 0x5a, 0x89, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58,
	0x99, 0x3e, 0x88, 0x05, 0xd1, 0x21, 0x25, 0x0e, 0x33, 0x2e, 0x2f, 0xad, 0x24, 0x3e, 0x33, 0xaf,
	0xa0, 0x14, 0x6a, 0x94, 0x52, 0x25, 0x17, 0xb7, 0x6f, 0x66, 0x5e, 0x09, 0xc8, 0x02, 0x3f, 0xb7,
	0x10, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0x9c, 0xc4, 0xe2, 0x62, 0xcf, 0x14, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x18, 0x57, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x09, 0x2c, 0xc8, 0x94,
	0x99, 0x22, 0xe4, 0xc0, 0xc5, 0x0a, 0x36, 0x47, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x45,
	0x0f, 0xb7, 0x9b, 0xf4, 0xfc, 0xdc, 0x42, 0x3c, 0x41, 0x6a, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67,
	0x08, 0x82, 0x68, 0x74, 0x72, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0xdd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0xb1, 0xfa, 0x70,
	0x63, 0xf5, 0x2b, 0xf4, 0x61, 0xbe, 0x29, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x7b, 0xc5,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x66, 0xa1, 0xa3, 0x34, 0x01, 0x00, 0x00,
}

func (m *MintableNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintableNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintableNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Input.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClaimableNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
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
func (m *MintableNFT) Size() (n int) {
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
	l = m.Input.Size()
	n += 1 + l + sovClaimableNft(uint64(l))
	return n
}

func sovClaimableNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimableNft(x uint64) (n int) {
	return sovClaimableNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintableNFT) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MintableNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintableNFT: illegal tag %d (wire type %d)", fieldNum, wire)
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
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimableNft
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
				return ErrInvalidLengthClaimableNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimableNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Input.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
