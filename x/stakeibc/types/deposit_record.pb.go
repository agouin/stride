// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/deposit_record.proto

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

type DepositRecord_Status int32

const (
	// pending transfer to delegate account
	DepositRecord_TRANSFER DepositRecord_Status = 0
	// pending staking on delegate account
	DepositRecord_STAKE DepositRecord_Status = 1
)

var DepositRecord_Status_name = map[int32]string{
	0: "TRANSFER",
	1: "STAKE",
}

var DepositRecord_Status_value = map[string]int32{
	"TRANSFER": 0,
	"STAKE":    1,
}

func (x DepositRecord_Status) String() string {
	return proto.EnumName(DepositRecord_Status_name, int32(x))
}

func (DepositRecord_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cc20567faf76ca44, []int{0, 0}
}

type DepositRecord struct {
	Id         uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount     int64                `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom      string               `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	HostZoneId string               `protobuf:"bytes,4,opt,name=hostZoneId,proto3" json:"hostZoneId,omitempty"`
	Status     DepositRecord_Status `protobuf:"varint,6,opt,name=status,proto3,enum=Stridelabs.stride.stakeibc.DepositRecord_Status" json:"status,omitempty"`
}

func (m *DepositRecord) Reset()         { *m = DepositRecord{} }
func (m *DepositRecord) String() string { return proto.CompactTextString(m) }
func (*DepositRecord) ProtoMessage()    {}
func (*DepositRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc20567faf76ca44, []int{0}
}
func (m *DepositRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositRecord.Merge(m, src)
}
func (m *DepositRecord) XXX_Size() int {
	return m.Size()
}
func (m *DepositRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DepositRecord proto.InternalMessageInfo

func (m *DepositRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DepositRecord) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DepositRecord) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DepositRecord) GetHostZoneId() string {
	if m != nil {
		return m.HostZoneId
	}
	return ""
}

func (m *DepositRecord) GetStatus() DepositRecord_Status {
	if m != nil {
		return m.Status
	}
	return DepositRecord_TRANSFER
}

func init() {
	proto.RegisterEnum("Stridelabs.stride.stakeibc.DepositRecord_Status", DepositRecord_Status_name, DepositRecord_Status_value)
	proto.RegisterType((*DepositRecord)(nil), "Stridelabs.stride.stakeibc.DepositRecord")
}

func init() { proto.RegisterFile("stakeibc/deposit_record.proto", fileDescriptor_cc20567faf76ca44) }

var fileDescriptor_cc20567faf76ca44 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x33, 0xfd, 0x19, 0xda, 0x8b, 0x96, 0x30, 0x88, 0x04, 0xc1, 0x21, 0x76, 0x95, 0x8d,
	0x13, 0xd1, 0x27, 0xa8, 0x58, 0xa9, 0x3f, 0xb8, 0x98, 0x74, 0xd5, 0x8d, 0x24, 0x99, 0xc1, 0x0e,
	0x9a, 0x4c, 0xc8, 0x4c, 0x40, 0xdf, 0xc2, 0xc7, 0x72, 0xd9, 0xa5, 0xcb, 0x92, 0xbc, 0x88, 0x90,
	0x54, 0xd1, 0x85, 0xbb, 0x7b, 0xce, 0xe5, 0x83, 0x8f, 0x03, 0xc7, 0xc6, 0xc6, 0xcf, 0x52, 0x25,
	0x69, 0x28, 0x64, 0xa1, 0x8d, 0xb2, 0x8f, 0xa5, 0x4c, 0x75, 0x29, 0x58, 0x51, 0x6a, 0xab, 0xc9,
	0x51, 0x64, 0x4b, 0x25, 0xe4, 0x4b, 0x9c, 0x18, 0x66, 0xda, 0x93, 0x7d, 0x03, 0xd3, 0x2d, 0x82,
	0xfd, 0xab, 0x0e, 0xe2, 0x2d, 0x43, 0x26, 0xd0, 0x53, 0xc2, 0x43, 0x3e, 0x0a, 0x06, 0xbc, 0xa7,
	0x04, 0x39, 0x04, 0x1c, 0x67, 0xba, 0xca, 0xad, 0xd7, 0xf3, 0x51, 0xd0, 0xe7, 0xbb, 0x44, 0x0e,
	0x60, 0x28, 0x64, 0xae, 0x33, 0xaf, 0xef, 0xa3, 0x60, 0xcc, 0xbb, 0x40, 0x28, 0xc0, 0x5a, 0x1b,
	0xbb, 0xd2, 0xb9, 0xbc, 0x11, 0xde, 0xa0, 0x7d, 0xfd, 0x6a, 0xc8, 0x02, 0xb0, 0xb1, 0xb1, 0xad,
	0x8c, 0x87, 0x7d, 0x14, 0x4c, 0xce, 0xcf, 0xd8, 0xff, 0x72, 0xec, 0x8f, 0x18, 0x8b, 0x5a, 0x8e,
	0xef, 0xf8, 0xe9, 0x09, 0xe0, 0xae, 0x21, 0x7b, 0x30, 0x5a, 0xf2, 0xd9, 0x43, 0x74, 0x3d, 0xe7,
	0xae, 0x43, 0xc6, 0x30, 0x8c, 0x96, 0xb3, 0xbb, 0xb9, 0x8b, 0x6e, 0x07, 0xa3, 0xa1, 0x8b, 0x2f,
	0x17, 0x1f, 0x35, 0x45, 0x9b, 0x9a, 0xa2, 0x6d, 0x4d, 0xd1, 0x7b, 0x43, 0x9d, 0x4d, 0x43, 0x9d,
	0xcf, 0x86, 0x3a, 0x2b, 0xf6, 0xa4, 0xec, 0xba, 0x4a, 0x58, 0xaa, 0xb3, 0xb0, 0xd3, 0x38, 0xbd,
	0x8f, 0x13, 0x13, 0x76, 0x1e, 0xe1, 0x6b, 0xf8, 0xb3, 0xab, 0x7d, 0x2b, 0xa4, 0x49, 0x70, 0xbb,
	0xe7, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xa0, 0x50, 0x5e, 0x70, 0x01, 0x00, 0x00,
}

func (m *DepositRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.HostZoneId) > 0 {
		i -= len(m.HostZoneId)
		copy(dAtA[i:], m.HostZoneId)
		i = encodeVarintDepositRecord(dAtA, i, uint64(len(m.HostZoneId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDepositRecord(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDepositRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovDepositRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovDepositRecord(uint64(m.Id))
	}
	if m.Amount != 0 {
		n += 1 + sovDepositRecord(uint64(m.Amount))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDepositRecord(uint64(l))
	}
	l = len(m.HostZoneId)
	if l > 0 {
		n += 1 + l + sovDepositRecord(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovDepositRecord(uint64(m.Status))
	}
	return n
}

func sovDepositRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDepositRecord(x uint64) (n int) {
	return sovDepositRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDepositRecord
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
			return fmt.Errorf("proto: DepositRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostZoneId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= DepositRecord_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDepositRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDepositRecord
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
func skipDepositRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDepositRecord
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
					return 0, ErrIntOverflowDepositRecord
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
					return 0, ErrIntOverflowDepositRecord
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
				return 0, ErrInvalidLengthDepositRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDepositRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDepositRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDepositRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDepositRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDepositRecord = fmt.Errorf("proto: unexpected end of group")
)
