// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/iscn/store.proto

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

type StoreRecord struct {
	IscnId   IscnId    `protobuf:"bytes,1,opt,name=iscn_id,json=iscnId,proto3" json:"iscn_id"`
	CidBytes []byte    `protobuf:"bytes,2,opt,name=cid_bytes,json=cidBytes,proto3" json:"cid_bytes,omitempty"`
	Data     IscnInput `protobuf:"bytes,3,opt,name=data,proto3,customtype=IscnInput" json:"data"`
}

func (m *StoreRecord) Reset()         { *m = StoreRecord{} }
func (m *StoreRecord) String() string { return proto.CompactTextString(m) }
func (*StoreRecord) ProtoMessage()    {}
func (*StoreRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9d9faad3d449872, []int{0}
}
func (m *StoreRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRecord.Merge(m, src)
}
func (m *StoreRecord) XXX_Size() int {
	return m.Size()
}
func (m *StoreRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRecord.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRecord proto.InternalMessageInfo

func (m *StoreRecord) GetIscnId() IscnId {
	if m != nil {
		return m.IscnId
	}
	return IscnId{}
}

func (m *StoreRecord) GetCidBytes() []byte {
	if m != nil {
		return m.CidBytes
	}
	return nil
}

type ContentIdRecord struct {
	OwnerAddressBytes []byte `protobuf:"bytes,1,opt,name=owner_address_bytes,json=ownerAddressBytes,proto3" json:"owner_address_bytes,omitempty"`
	LatestVersion     uint64 `protobuf:"varint,2,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
}

func (m *ContentIdRecord) Reset()         { *m = ContentIdRecord{} }
func (m *ContentIdRecord) String() string { return proto.CompactTextString(m) }
func (*ContentIdRecord) ProtoMessage()    {}
func (*ContentIdRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9d9faad3d449872, []int{1}
}
func (m *ContentIdRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentIdRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentIdRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentIdRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentIdRecord.Merge(m, src)
}
func (m *ContentIdRecord) XXX_Size() int {
	return m.Size()
}
func (m *ContentIdRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentIdRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ContentIdRecord proto.InternalMessageInfo

func (m *ContentIdRecord) GetOwnerAddressBytes() []byte {
	if m != nil {
		return m.OwnerAddressBytes
	}
	return nil
}

func (m *ContentIdRecord) GetLatestVersion() uint64 {
	if m != nil {
		return m.LatestVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*StoreRecord)(nil), "likechain.iscn.StoreRecord")
	proto.RegisterType((*ContentIdRecord)(nil), "likechain.iscn.ContentIdRecord")
}

func init() { proto.RegisterFile("likechain/iscn/store.proto", fileDescriptor_c9d9faad3d449872) }

var fileDescriptor_c9d9faad3d449872 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x27, 0x5a, 0xaa, 0x4d, 0xb5, 0xd2, 0x51, 0xa4, 0xb4, 0x90, 0x96, 0x82, 0xd0, 0x8d,
	0x19, 0x68, 0xf1, 0x00, 0x8e, 0xab, 0x82, 0xab, 0x11, 0x5c, 0xb8, 0x19, 0xa6, 0x49, 0x68, 0x83,
	0x35, 0x19, 0x92, 0xb4, 0xda, 0x1b, 0xb8, 0xf4, 0x58, 0x5d, 0x76, 0x29, 0x2e, 0x8a, 0xcc, 0x5c,
	0x44, 0x92, 0xd4, 0x82, 0x6e, 0xc2, 0x7b, 0xef, 0x4b, 0xfe, 0xff, 0xe5, 0x87, 0xed, 0x39, 0x7f,
	0x66, 0x64, 0x96, 0x71, 0x11, 0x71, 0x4d, 0x44, 0xa4, 0x8d, 0x54, 0x0c, 0xe7, 0x4a, 0x1a, 0x19,
	0x36, 0xf6, 0x0c, 0x5b, 0xd6, 0xbe, 0x98, 0xca, 0xa9, 0x74, 0x28, 0xb2, 0x95, 0xbf, 0xd5, 0xee,
	0xfc, 0x53, 0xb0, 0x07, 0xa7, 0x1e, 0xf6, 0xdf, 0x01, 0xac, 0x3f, 0x58, 0xc9, 0x84, 0x11, 0xa9,
	0x68, 0x78, 0x03, 0x8f, 0x2c, 0x4f, 0x39, 0x6d, 0x81, 0x1e, 0x18, 0xd4, 0x87, 0x97, 0xf8, 0xaf,
	0x09, 0x1e, 0x6b, 0x22, 0xc6, 0x34, 0xae, 0xac, 0xb7, 0xdd, 0x20, 0xa9, 0x72, 0xd7, 0x85, 0x1d,
	0x58, 0x23, 0x9c, 0xa6, 0x93, 0x95, 0x61, 0xba, 0x75, 0xd0, 0x03, 0x83, 0x93, 0xe4, 0x98, 0x70,
	0x1a, 0xdb, 0x3e, 0xbc, 0x82, 0x15, 0x9a, 0x99, 0xac, 0x75, 0x68, 0xe7, 0x71, 0xd3, 0x3e, 0xfc,
	0xda, 0x76, 0x6b, 0x4e, 0x48, 0xe4, 0x0b, 0x93, 0x38, 0xdc, 0x9f, 0xc1, 0xb3, 0x3b, 0x29, 0x0c,
	0x13, 0x66, 0x4c, 0x77, 0xdb, 0x60, 0x78, 0x2e, 0x5f, 0x05, 0x53, 0x69, 0x46, 0xa9, 0x62, 0x5a,
	0xef, 0x0c, 0x80, 0x33, 0x68, 0x3a, 0x74, 0xeb, 0xc9, 0xaf, 0x53, 0x63, 0x9e, 0x19, 0xa6, 0x4d,
	0xba, 0x64, 0x4a, 0x73, 0x29, 0xdc, 0x2e, 0x95, 0xe4, 0xd4, 0x4f, 0x1f, 0xfd, 0x30, 0xbe, 0x5f,
	0x17, 0x08, 0x6c, 0x0a, 0x04, 0xbe, 0x0b, 0x04, 0x3e, 0x4a, 0x14, 0x6c, 0x4a, 0x14, 0x7c, 0x96,
	0x28, 0x78, 0x1a, 0x4e, 0xb9, 0x99, 0x2d, 0x26, 0x98, 0xc8, 0x97, 0xc8, 0xfd, 0x5b, 0x72, 0xb1,
	0x2f, 0xae, 0x7d, 0x88, 0xcb, 0x51, 0xf4, 0xe6, 0x93, 0x34, 0xab, 0x9c, 0xe9, 0x49, 0xd5, 0x25,
	0x39, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x67, 0x29, 0x0a, 0xaa, 0x01, 0x00, 0x00,
}

func (m *StoreRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Data.Size()
		i -= size
		if _, err := m.Data.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStore(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.CidBytes) > 0 {
		i -= len(m.CidBytes)
		copy(dAtA[i:], m.CidBytes)
		i = encodeVarintStore(dAtA, i, uint64(len(m.CidBytes)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.IscnId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStore(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ContentIdRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentIdRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentIdRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestVersion != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.LatestVersion))
		i--
		dAtA[i] = 0x10
	}
	if len(m.OwnerAddressBytes) > 0 {
		i -= len(m.OwnerAddressBytes)
		copy(dAtA[i:], m.OwnerAddressBytes)
		i = encodeVarintStore(dAtA, i, uint64(len(m.OwnerAddressBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStore(dAtA []byte, offset int, v uint64) int {
	offset -= sovStore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoreRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.IscnId.Size()
	n += 1 + l + sovStore(uint64(l))
	l = len(m.CidBytes)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = m.Data.Size()
	n += 1 + l + sovStore(uint64(l))
	return n
}

func (m *ContentIdRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddressBytes)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if m.LatestVersion != 0 {
		n += 1 + sovStore(uint64(m.LatestVersion))
	}
	return n
}

func sovStore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStore(x uint64) (n int) {
	return sovStore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: StoreRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IscnId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IscnId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CidBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CidBytes = append(m.CidBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.CidBytes == nil {
				m.CidBytes = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *ContentIdRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: ContentIdRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentIdRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddressBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddressBytes = append(m.OwnerAddressBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.OwnerAddressBytes == nil {
				m.OwnerAddressBytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestVersion", wireType)
			}
			m.LatestVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
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
func skipStore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
				return 0, ErrInvalidLengthStore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStore = fmt.Errorf("proto: unexpected end of group")
)
