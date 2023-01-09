// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mage/profiles/v3/models_packets.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// LinkChainAccountPacketData defines the object that should be sent inside a
// MsgSendPacket when wanting to link an external chain to a Mage profile
// using IBC
type LinkChainAccountPacketData struct {
	// SourceAddress contains the details of the external chain address
	SourceAddress *types.Any `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty" yaml:"source_address"`
	// SourceProof represents the proof of ownership of the source address
	SourceProof Proof `protobuf:"bytes,2,opt,name=source_proof,json=sourceProof,proto3" json:"source_proof" yaml:"source_proof"`
	// SourceChainConfig contains the details of the source chain
	SourceChainConfig ChainConfig `protobuf:"bytes,3,opt,name=source_chain_config,json=sourceChainConfig,proto3" json:"source_chain_config" yaml:"source_chain_config"`
	// DestinationAddress represents the Mage address of the profile that should
	// be linked with the external account
	DestinationAddress string `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty" yaml:"destination_address"`
	// DestinationProof contains the proof of ownership of the DestinationAddress
	DestinationProof Proof `protobuf:"bytes,5,opt,name=destination_proof,json=destinationProof,proto3" json:"destination_proof" yaml:"destination_proof"`
}

func (m *LinkChainAccountPacketData) Reset()         { *m = LinkChainAccountPacketData{} }
func (m *LinkChainAccountPacketData) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountPacketData) ProtoMessage()    {}
func (*LinkChainAccountPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecd3778626d03f03, []int{0}
}
func (m *LinkChainAccountPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountPacketData.Merge(m, src)
}
func (m *LinkChainAccountPacketData) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountPacketData proto.InternalMessageInfo

// LinkChainAccountPacketAck defines a struct for the packet acknowledgment
type LinkChainAccountPacketAck struct {
	// SourceAddress contains the external address that has been linked properly
	// with the profile
	SourceAddress string `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
}

func (m *LinkChainAccountPacketAck) Reset()         { *m = LinkChainAccountPacketAck{} }
func (m *LinkChainAccountPacketAck) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountPacketAck) ProtoMessage()    {}
func (*LinkChainAccountPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecd3778626d03f03, []int{1}
}
func (m *LinkChainAccountPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountPacketAck.Merge(m, src)
}
func (m *LinkChainAccountPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountPacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LinkChainAccountPacketData)(nil), "mage.profiles.v3.LinkChainAccountPacketData")
	proto.RegisterType((*LinkChainAccountPacketAck)(nil), "mage.profiles.v3.LinkChainAccountPacketAck")
}

func init() {
	proto.RegisterFile("mage/profiles/v3/models_packets.proto", fileDescriptor_ecd3778626d03f03)
}

var fileDescriptor_ecd3778626d03f03 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x4d, 0xa0, 0x20, 0x4d, 0x0a, 0x88, 0x49, 0x07, 0xa9, 0x2d, 0x52, 0x52, 0x45, 0x42, 0x54,
	0x42, 0x63, 0x0b, 0xca, 0x6a, 0x76, 0xcd, 0xb0, 0x40, 0x02, 0x89, 0x51, 0x77, 0xb0, 0xa9, 0x5c,
	0xc7, 0x4d, 0xad, 0xa6, 0xfe, 0x51, 0xec, 0x56, 0xf4, 0x06, 0x2c, 0x59, 0x70, 0x00, 0x0e, 0xc1,
	0x21, 0x46, 0xac, 0x66, 0xc9, 0xaa, 0x42, 0xed, 0x0d, 0xe6, 0x04, 0x28, 0xb6, 0xcb, 0xa4, 0xb4,
	0x68, 0x76, 0xfe, 0xff, 0xbd, 0xf7, 0xdf, 0xcf, 0xcb, 0xf7, 0x9e, 0x27, 0x4c, 0xce, 0x40, 0xe2,
	0xbc, 0x80, 0x31, 0xcf, 0x98, 0xc4, 0x8b, 0x1e, 0x9e, 0x41, 0xc2, 0x32, 0x39, 0xcc, 0x09, 0x9d,
	0x32, 0x25, 0x51, 0x5e, 0x80, 0x02, 0xdf, 0x37, 0x44, 0xb4, 0x25, 0xa2, 0x45, 0xaf, 0x7d, 0x92,
	0x42, 0x0a, 0x1a, 0xc6, 0xe5, 0xcb, 0x30, 0xdb, 0xad, 0x14, 0x20, 0xcd, 0x18, 0xd6, 0xd5, 0x68,
	0x3e, 0xc6, 0x44, 0x2c, 0xb7, 0x10, 0x85, 0x72, 0xc8, 0xd0, 0x68, 0x4c, 0x61, 0xa1, 0x17, 0xff,
	0x5f, 0x84, 0x4e, 0x08, 0x17, 0xc3, 0x8c, 0x8b, 0xa9, 0x25, 0x47, 0xdf, 0x6a, 0x5e, 0xfb, 0x3d,
	0x17, 0xd3, 0xf3, 0x12, 0xe9, 0x53, 0x0a, 0x73, 0xa1, 0x2e, 0xf4, 0xba, 0x6f, 0x88, 0x22, 0x3e,
	0xf3, 0x1e, 0x49, 0x98, 0x17, 0x94, 0x0d, 0x49, 0x92, 0x14, 0x4c, 0xca, 0xa6, 0xdb, 0x71, 0xbb,
	0xf5, 0x57, 0x27, 0xc8, 0xac, 0x86, 0xb6, 0xab, 0xa1, 0xbe, 0x58, 0xc6, 0xdd, 0xeb, 0x55, 0xf8,
	0x64, 0x49, 0x66, 0xd9, 0x59, 0xb4, 0xab, 0x8a, 0x7e, 0xfe, 0x38, 0xad, 0xf7, 0xcd, 0xbb, 0x9c,
	0x3b, 0x78, 0x68, 0x70, 0xdb, 0xf2, 0x3f, 0x7a, 0x0f, 0xac, 0x20, 0x2f, 0x00, 0xc6, 0xcd, 0x3b,
	0xda, 0xa4, 0x85, 0xf6, 0x93, 0x42, 0x17, 0x25, 0x21, 0x7e, 0x7a, 0xb9, 0x0a, 0x9d, 0xeb, 0x55,
	0xd8, 0xd8, 0x71, 0xd3, 0xe2, 0x68, 0x50, 0x37, 0xa5, 0x66, 0xfa, 0xd2, 0x6b, 0x58, 0xd4, 0x7c,
	0x3c, 0x05, 0x31, 0xe6, 0x69, 0xf3, 0xae, 0x76, 0x08, 0x0f, 0x39, 0xe8, 0x28, 0xce, 0x35, 0x2d,
	0x8e, 0xac, 0x4f, 0x7b, 0xc7, 0xa7, 0x3a, 0x29, 0x1a, 0x1c, 0x9b, 0x6e, 0x45, 0xe6, 0x7f, 0xf0,
	0x1a, 0x09, 0x93, 0x8a, 0x0b, 0xa2, 0x38, 0x88, 0xbf, 0xd9, 0xd5, 0x3a, 0x6e, 0xf7, 0x28, 0x0e,
	0x6e, 0xe6, 0x1d, 0x20, 0x45, 0x03, 0xbf, 0xd2, 0xdd, 0x06, 0x34, 0xf1, 0x8e, 0xab, 0x5c, 0x93,
	0xd2, 0xbd, 0xdb, 0x52, 0xea, 0xd8, 0xed, 0x9b, 0xfb, 0x6e, 0x36, 0xaa, 0xc7, 0x95, 0x9e, 0xd6,
	0x9c, 0xd5, 0xbe, 0x7c, 0x0f, 0x9d, 0xe8, 0xad, 0xd7, 0x3a, 0x7c, 0x15, 0x7d, 0x3a, 0xf5, 0x9f,
	0x1d, 0x3c, 0x8a, 0xa3, 0x7f, 0x7e, 0xaa, 0x99, 0x14, 0xbf, 0xbb, 0x5c, 0x07, 0xee, 0xd5, 0x3a,
	0x70, 0x7f, 0xaf, 0x03, 0xf7, 0xeb, 0x26, 0x70, 0xae, 0x36, 0x81, 0xf3, 0x6b, 0x13, 0x38, 0x9f,
	0x5e, 0xa6, 0x5c, 0x4d, 0xe6, 0x23, 0x44, 0x61, 0x86, 0xcd, 0x27, 0x9c, 0x66, 0x64, 0x24, 0xed,
	0x1b, 0x2f, 0x5e, 0xe3, 0xcf, 0x37, 0x37, 0xac, 0x96, 0x39, 0x93, 0xa3, 0xfb, 0xfa, 0xdc, 0x7a,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x12, 0xb2, 0x30, 0x23, 0x6c, 0x03, 0x00, 0x00,
}

func (m *LinkChainAccountPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DestinationProof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintModelsPackets(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.SourceChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.SourceProof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SourceAddress != nil {
		{
			size, err := m.SourceAddress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsPackets(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LinkChainAccountPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintModelsPackets(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsPackets(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsPackets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LinkChainAccountPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceAddress != nil {
		l = m.SourceAddress.Size()
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	l = m.SourceProof.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	l = m.SourceChainConfig.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	l = m.DestinationProof.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	return n
}

func (m *LinkChainAccountPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	return n
}

func sovModelsPackets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsPackets(x uint64) (n int) {
	return sovModelsPackets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LinkChainAccountPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsPackets
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
			return fmt.Errorf("proto: LinkChainAccountPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SourceAddress == nil {
				m.SourceAddress = &types.Any{}
			}
			if err := m.SourceAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SourceProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SourceChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DestinationProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsPackets
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
func (m *LinkChainAccountPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsPackets
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
			return fmt.Errorf("proto: LinkChainAccountPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsPackets
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
func skipModelsPackets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsPackets
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
					return 0, ErrIntOverflowModelsPackets
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
					return 0, ErrIntOverflowModelsPackets
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
				return 0, ErrInvalidLengthModelsPackets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsPackets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsPackets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsPackets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsPackets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsPackets = fmt.Errorf("proto: unexpected end of group")
)
