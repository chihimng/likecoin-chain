// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/likenft/offer_expire_queue.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type OfferExpireQueueEntry struct {
	ExpireTime time.Time `protobuf:"bytes,1,opt,name=expire_time,json=expireTime,proto3,stdtime" json:"expire_time"`
	OfferKey   []byte    `protobuf:"bytes,2,opt,name=offer_key,json=offerKey,proto3" json:"offer_key,omitempty"`
}

func (m *OfferExpireQueueEntry) Reset()         { *m = OfferExpireQueueEntry{} }
func (m *OfferExpireQueueEntry) String() string { return proto.CompactTextString(m) }
func (*OfferExpireQueueEntry) ProtoMessage()    {}
func (*OfferExpireQueueEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_74ff4e28d7ee1f12, []int{0}
}
func (m *OfferExpireQueueEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OfferExpireQueueEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OfferExpireQueueEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OfferExpireQueueEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfferExpireQueueEntry.Merge(m, src)
}
func (m *OfferExpireQueueEntry) XXX_Size() int {
	return m.Size()
}
func (m *OfferExpireQueueEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_OfferExpireQueueEntry.DiscardUnknown(m)
}

var xxx_messageInfo_OfferExpireQueueEntry proto.InternalMessageInfo

func (m *OfferExpireQueueEntry) GetExpireTime() time.Time {
	if m != nil {
		return m.ExpireTime
	}
	return time.Time{}
}

func (m *OfferExpireQueueEntry) GetOfferKey() []byte {
	if m != nil {
		return m.OfferKey
	}
	return nil
}

func init() {
	proto.RegisterType((*OfferExpireQueueEntry)(nil), "likechain.likenft.OfferExpireQueueEntry")
}

func init() {
	proto.RegisterFile("likechain/likenft/offer_expire_queue.proto", fileDescriptor_74ff4e28d7ee1f12)
}

var fileDescriptor_74ff4e28d7ee1f12 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0xc9, 0xcc, 0x4e,
	0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb1, 0xf2, 0xd2, 0x4a, 0xf4, 0xf3, 0xd3, 0xd2, 0x52,
	0x8b, 0xe2, 0x53, 0x2b, 0x0a, 0x32, 0x8b, 0x52, 0xe3, 0x0b, 0x4b, 0x53, 0x4b, 0x53, 0xf5, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xe1, 0x6a, 0xf5, 0xa0, 0x6a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44, 0xa1, 0x94, 0x7c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa,
	0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x5f, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b, 0x00,
	0x51, 0xa0, 0x54, 0xcd, 0x25, 0xea, 0x0f, 0xb2, 0xc5, 0x15, 0x6c, 0x49, 0x20, 0xc8, 0x0e, 0xd7,
	0xbc, 0x92, 0xa2, 0x4a, 0x21, 0x57, 0x2e, 0x6e, 0xa8, 0xc5, 0x20, 0x2d, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x10, 0xf3, 0xf4, 0x60, 0xe6, 0xe9, 0x85, 0xc0, 0xcc, 0x73, 0xe2,
	0x38, 0x71, 0x4f, 0x9e, 0x61, 0xc2, 0x7d, 0x79, 0xc6, 0x20, 0x2e, 0x88, 0x46, 0x90, 0x94, 0x90,
	0x34, 0x17, 0x27, 0xc4, 0x17, 0xd9, 0xa9, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x1c,
	0x60, 0x01, 0xef, 0xd4, 0x4a, 0x27, 0xff, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0x32, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x05, 0x87, 0x46, 0x72,
	0x3e, 0x34, 0x58, 0x40, 0x0c, 0x5d, 0x48, 0x28, 0x95, 0x19, 0xeb, 0x57, 0xc0, 0x83, 0xaa, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x2e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3,
	0x4f, 0xce, 0xf3, 0x4c, 0x01, 0x00, 0x00,
}

func (m *OfferExpireQueueEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OfferExpireQueueEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OfferExpireQueueEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OfferKey) > 0 {
		i -= len(m.OfferKey)
		copy(dAtA[i:], m.OfferKey)
		i = encodeVarintOfferExpireQueue(dAtA, i, uint64(len(m.OfferKey)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpireTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintOfferExpireQueue(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintOfferExpireQueue(dAtA []byte, offset int, v uint64) int {
	offset -= sovOfferExpireQueue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OfferExpireQueueEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireTime)
	n += 1 + l + sovOfferExpireQueue(uint64(l))
	l = len(m.OfferKey)
	if l > 0 {
		n += 1 + l + sovOfferExpireQueue(uint64(l))
	}
	return n
}

func sovOfferExpireQueue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOfferExpireQueue(x uint64) (n int) {
	return sovOfferExpireQueue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OfferExpireQueueEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOfferExpireQueue
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
			return fmt.Errorf("proto: OfferExpireQueueEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OfferExpireQueueEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOfferExpireQueue
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
				return ErrInvalidLengthOfferExpireQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOfferExpireQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpireTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOfferExpireQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOfferExpireQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOfferExpireQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OfferKey = append(m.OfferKey[:0], dAtA[iNdEx:postIndex]...)
			if m.OfferKey == nil {
				m.OfferKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOfferExpireQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOfferExpireQueue
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
func skipOfferExpireQueue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOfferExpireQueue
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
					return 0, ErrIntOverflowOfferExpireQueue
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
					return 0, ErrIntOverflowOfferExpireQueue
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
				return 0, ErrInvalidLengthOfferExpireQueue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOfferExpireQueue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOfferExpireQueue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOfferExpireQueue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOfferExpireQueue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOfferExpireQueue = fmt.Errorf("proto: unexpected end of group")
)
