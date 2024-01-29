// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cork/v2/cork.proto

package v2

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

type Cork struct {
	// call body containing the ABI encoded bytes to send to the contract
	EncodedContractCall []byte `protobuf:"bytes,1,opt,name=encoded_contract_call,json=encodedContractCall,proto3" json:"encoded_contract_call,omitempty"`
	// address of the contract to send the call
	TargetContractAddress string `protobuf:"bytes,2,opt,name=target_contract_address,json=targetContractAddress,proto3" json:"target_contract_address,omitempty"`
}

func (m *Cork) Reset()         { *m = Cork{} }
func (m *Cork) String() string { return proto.CompactTextString(m) }
func (*Cork) ProtoMessage()    {}
func (*Cork) Descriptor() ([]byte, []int) {
	return fileDescriptor_1219f78116b242b2, []int{0}
}
func (m *Cork) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cork.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cork.Merge(m, src)
}
func (m *Cork) XXX_Size() int {
	return m.Size()
}
func (m *Cork) XXX_DiscardUnknown() {
	xxx_messageInfo_Cork.DiscardUnknown(m)
}

var xxx_messageInfo_Cork proto.InternalMessageInfo

func (m *Cork) GetEncodedContractCall() []byte {
	if m != nil {
		return m.EncodedContractCall
	}
	return nil
}

func (m *Cork) GetTargetContractAddress() string {
	if m != nil {
		return m.TargetContractAddress
	}
	return ""
}

type ScheduledCork struct {
	Cork        *Cork  `protobuf:"bytes,1,opt,name=cork,proto3" json:"cork,omitempty"`
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Validator   string `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	Id          []byte `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ScheduledCork) Reset()         { *m = ScheduledCork{} }
func (m *ScheduledCork) String() string { return proto.CompactTextString(m) }
func (*ScheduledCork) ProtoMessage()    {}
func (*ScheduledCork) Descriptor() ([]byte, []int) {
	return fileDescriptor_1219f78116b242b2, []int{1}
}
func (m *ScheduledCork) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledCork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduledCork.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduledCork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledCork.Merge(m, src)
}
func (m *ScheduledCork) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledCork) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledCork.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledCork proto.InternalMessageInfo

func (m *ScheduledCork) GetCork() *Cork {
	if m != nil {
		return m.Cork
	}
	return nil
}

func (m *ScheduledCork) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ScheduledCork) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *ScheduledCork) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type CorkResult struct {
	Cork               *Cork  `protobuf:"bytes,1,opt,name=cork,proto3" json:"cork,omitempty"`
	BlockHeight        uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Approved           bool   `protobuf:"varint,3,opt,name=approved,proto3" json:"approved,omitempty"`
	ApprovalPercentage string `protobuf:"bytes,4,opt,name=approval_percentage,json=approvalPercentage,proto3" json:"approval_percentage,omitempty"`
}

func (m *CorkResult) Reset()         { *m = CorkResult{} }
func (m *CorkResult) String() string { return proto.CompactTextString(m) }
func (*CorkResult) ProtoMessage()    {}
func (*CorkResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1219f78116b242b2, []int{2}
}
func (m *CorkResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CorkResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CorkResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CorkResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CorkResult.Merge(m, src)
}
func (m *CorkResult) XXX_Size() int {
	return m.Size()
}
func (m *CorkResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CorkResult.DiscardUnknown(m)
}

var xxx_messageInfo_CorkResult proto.InternalMessageInfo

func (m *CorkResult) GetCork() *Cork {
	if m != nil {
		return m.Cork
	}
	return nil
}

func (m *CorkResult) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *CorkResult) GetApproved() bool {
	if m != nil {
		return m.Approved
	}
	return false
}

func (m *CorkResult) GetApprovalPercentage() string {
	if m != nil {
		return m.ApprovalPercentage
	}
	return ""
}

type CellarIDSet struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (m *CellarIDSet) Reset()         { *m = CellarIDSet{} }
func (m *CellarIDSet) String() string { return proto.CompactTextString(m) }
func (*CellarIDSet) ProtoMessage()    {}
func (*CellarIDSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_1219f78116b242b2, []int{3}
}
func (m *CellarIDSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CellarIDSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CellarIDSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CellarIDSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellarIDSet.Merge(m, src)
}
func (m *CellarIDSet) XXX_Size() int {
	return m.Size()
}
func (m *CellarIDSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CellarIDSet.DiscardUnknown(m)
}

var xxx_messageInfo_CellarIDSet proto.InternalMessageInfo

func (m *CellarIDSet) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*Cork)(nil), "cork.v2.Cork")
	proto.RegisterType((*ScheduledCork)(nil), "cork.v2.ScheduledCork")
	proto.RegisterType((*CorkResult)(nil), "cork.v2.CorkResult")
	proto.RegisterType((*CellarIDSet)(nil), "cork.v2.CellarIDSet")
}

func init() { proto.RegisterFile("cork/v2/cork.proto", fileDescriptor_1219f78116b242b2) }

var fileDescriptor_1219f78116b242b2 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x8a, 0xdb, 0x30,
	0x10, 0x87, 0xa3, 0xc4, 0xb4, 0xb1, 0x92, 0x94, 0xa2, 0x10, 0x6a, 0x4a, 0x71, 0x13, 0x9f, 0x72,
	0x28, 0x16, 0xb8, 0xd0, 0x9e, 0x5b, 0x97, 0xd2, 0xde, 0x8a, 0x73, 0xeb, 0xc5, 0x28, 0xd2, 0x60,
	0x7b, 0xad, 0x44, 0x46, 0x56, 0xcc, 0xe6, 0xbc, 0x2f, 0xb0, 0x6f, 0xb0, 0xaf, 0xb3, 0xc7, 0x1c,
	0xf7, 0xb8, 0x24, 0x2f, 0xb2, 0x58, 0xf9, 0xf7, 0x02, 0x7b, 0xf2, 0xf8, 0xf7, 0x31, 0x33, 0x9f,
	0x60, 0x30, 0xe1, 0x4a, 0x97, 0xb4, 0x89, 0x68, 0xfb, 0x0d, 0x2b, 0xad, 0x8c, 0x22, 0x6f, 0x6d,
	0xdd, 0x44, 0x81, 0xc6, 0x4e, 0xac, 0x74, 0x49, 0x22, 0x3c, 0x81, 0x35, 0x57, 0x02, 0x44, 0xca,
	0xd5, 0xda, 0x68, 0xc6, 0x4d, 0xca, 0x99, 0x94, 0x1e, 0x9a, 0xa2, 0xf9, 0x30, 0x19, 0x9f, 0x60,
	0x7c, 0x62, 0x31, 0x93, 0x92, 0x7c, 0xc3, 0x1f, 0x0c, 0xd3, 0x19, 0x98, 0x6b, 0x0b, 0x13, 0x42,
	0x43, 0x5d, 0x7b, 0xdd, 0x29, 0x9a, 0xbb, 0xc9, 0xe4, 0x88, 0xcf, 0x4d, 0x3f, 0x8e, 0x30, 0xb8,
	0x43, 0x78, 0xb4, 0xe0, 0x39, 0x88, 0x8d, 0x6c, 0x27, 0xea, 0x92, 0xcc, 0xb0, 0xd3, 0x0a, 0xd9,
	0x65, 0x83, 0x68, 0x14, 0x9e, 0xec, 0xc2, 0x16, 0x26, 0x16, 0x91, 0x19, 0x1e, 0x2e, 0xa5, 0xe2,
	0x65, 0x9a, 0x43, 0x91, 0xe5, 0xc6, 0x6e, 0x70, 0x92, 0x81, 0xcd, 0xfe, 0xd8, 0x88, 0x7c, 0xc2,
	0x6e, 0xc3, 0x64, 0x21, 0x98, 0x51, 0xda, 0xeb, 0x59, 0x83, 0x6b, 0x40, 0xde, 0xe1, 0x6e, 0x21,
	0x3c, 0xc7, 0x3e, 0xa7, 0x5b, 0x88, 0xe0, 0x01, 0x61, 0x6c, 0xe7, 0x43, 0xbd, 0x91, 0xe6, 0x95,
	0x14, 0x3e, 0xe2, 0x3e, 0xab, 0x2a, 0xad, 0x1a, 0x10, 0xd6, 0xa0, 0x9f, 0x5c, 0xfe, 0x09, 0xc5,
	0xe3, 0x63, 0xcd, 0x64, 0x5a, 0x81, 0xe6, 0xb0, 0x36, 0x2c, 0x03, 0x6b, 0xe4, 0x26, 0xe4, 0x8c,
	0xfe, 0x5d, 0x48, 0xf0, 0x19, 0x0f, 0x62, 0x90, 0x92, 0xe9, 0xbf, 0xbf, 0x16, 0x60, 0xc8, 0x7b,
	0xdc, 0x2b, 0x44, 0xed, 0xa1, 0x69, 0x6f, 0xee, 0x26, 0x6d, 0xf9, 0xf3, 0xf7, 0xe3, 0xde, 0x47,
	0xbb, 0xbd, 0x8f, 0x9e, 0xf7, 0x3e, 0xba, 0x3f, 0xf8, 0x9d, 0xdd, 0xc1, 0xef, 0x3c, 0x1d, 0xfc,
	0xce, 0xff, 0x2f, 0x59, 0x61, 0xf2, 0xcd, 0x32, 0xe4, 0x6a, 0x45, 0x2b, 0xc8, 0xb2, 0xed, 0x4d,
	0x43, 0x6b, 0xb5, 0x5a, 0x81, 0x2c, 0x40, 0xd3, 0xe6, 0x3b, 0xbd, 0xb5, 0xd7, 0x40, 0xcd, 0xb6,
	0x82, 0x7a, 0xf9, 0xc6, 0x1e, 0xc5, 0xd7, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0xc0, 0x22,
	0xce, 0x2a, 0x02, 0x00, 0x00,
}

func (m *Cork) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cork) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cork) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TargetContractAddress) > 0 {
		i -= len(m.TargetContractAddress)
		copy(dAtA[i:], m.TargetContractAddress)
		i = encodeVarintCork(dAtA, i, uint64(len(m.TargetContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EncodedContractCall) > 0 {
		i -= len(m.EncodedContractCall)
		copy(dAtA[i:], m.EncodedContractCall)
		i = encodeVarintCork(dAtA, i, uint64(len(m.EncodedContractCall)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ScheduledCork) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledCork) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledCork) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCork(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintCork(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintCork(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Cork != nil {
		{
			size, err := m.Cork.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCork(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CorkResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CorkResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CorkResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApprovalPercentage) > 0 {
		i -= len(m.ApprovalPercentage)
		copy(dAtA[i:], m.ApprovalPercentage)
		i = encodeVarintCork(dAtA, i, uint64(len(m.ApprovalPercentage)))
		i--
		dAtA[i] = 0x22
	}
	if m.Approved {
		i--
		if m.Approved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.BlockHeight != 0 {
		i = encodeVarintCork(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Cork != nil {
		{
			size, err := m.Cork.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCork(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CellarIDSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CellarIDSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CellarIDSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		for iNdEx := len(m.Ids) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Ids[iNdEx])
			copy(dAtA[i:], m.Ids[iNdEx])
			i = encodeVarintCork(dAtA, i, uint64(len(m.Ids[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCork(dAtA []byte, offset int, v uint64) int {
	offset -= sovCork(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cork) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EncodedContractCall)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	l = len(m.TargetContractAddress)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	return n
}

func (m *ScheduledCork) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cork != nil {
		l = m.Cork.Size()
		n += 1 + l + sovCork(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovCork(uint64(m.BlockHeight))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	return n
}

func (m *CorkResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cork != nil {
		l = m.Cork.Size()
		n += 1 + l + sovCork(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovCork(uint64(m.BlockHeight))
	}
	if m.Approved {
		n += 2
	}
	l = len(m.ApprovalPercentage)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	return n
}

func (m *CellarIDSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		for _, s := range m.Ids {
			l = len(s)
			n += 1 + l + sovCork(uint64(l))
		}
	}
	return n
}

func sovCork(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCork(x uint64) (n int) {
	return sovCork(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cork) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCork
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
			return fmt.Errorf("proto: Cork: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cork: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncodedContractCall", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncodedContractCall = append(m.EncodedContractCall[:0], dAtA[iNdEx:postIndex]...)
			if m.EncodedContractCall == nil {
				m.EncodedContractCall = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCork
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
func (m *ScheduledCork) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCork
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
			return fmt.Errorf("proto: ScheduledCork: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledCork: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cork", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cork == nil {
				m.Cork = &Cork{}
			}
			if err := m.Cork.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCork
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
func (m *CorkResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCork
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
			return fmt.Errorf("proto: CorkResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CorkResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cork", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cork == nil {
				m.Cork = &Cork{}
			}
			if err := m.Cork.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Approved = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovalPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApprovalPercentage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCork
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
func (m *CellarIDSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCork
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
			return fmt.Errorf("proto: CellarIDSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CellarIDSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ids = append(m.Ids, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCork
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
func skipCork(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCork
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
					return 0, ErrIntOverflowCork
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
					return 0, ErrIntOverflowCork
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
				return 0, ErrInvalidLengthCork
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCork
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCork
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCork        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCork          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCork = fmt.Errorf("proto: unexpected end of group")
)