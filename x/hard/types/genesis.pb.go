// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo/hard/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the hard module's genesis state.
type GenesisState struct {
	Params                    Params                                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PreviousAccumulationTimes GenesisAccumulationTimes                 `protobuf:"bytes,2,rep,name=previous_accumulation_times,json=previousAccumulationTimes,proto3,castrepeated=GenesisAccumulationTimes" json:"previous_accumulation_times"`
	Deposits                  Deposits                                 `protobuf:"bytes,3,rep,name=deposits,proto3,castrepeated=Deposits" json:"deposits"`
	Borrows                   Borrows                                  `protobuf:"bytes,4,rep,name=borrows,proto3,castrepeated=Borrows" json:"borrows"`
	TotalSupplied             github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=total_supplied,json=totalSupplied,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_supplied"`
	TotalBorrowed             github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=total_borrowed,json=totalBorrowed,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_borrowed"`
	TotalReserves             github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=total_reserves,json=totalReserves,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_reserves"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c49fc6faba1ed7d5, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPreviousAccumulationTimes() GenesisAccumulationTimes {
	if m != nil {
		return m.PreviousAccumulationTimes
	}
	return nil
}

func (m *GenesisState) GetDeposits() Deposits {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetBorrows() Borrows {
	if m != nil {
		return m.Borrows
	}
	return nil
}

func (m *GenesisState) GetTotalSupplied() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalSupplied
	}
	return nil
}

func (m *GenesisState) GetTotalBorrowed() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalBorrowed
	}
	return nil
}

func (m *GenesisState) GetTotalReserves() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalReserves
	}
	return nil
}

// GenesisAccumulationTime stores the previous distribution time and its corresponding denom.
type GenesisAccumulationTime struct {
	CollateralType           string                                 `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty"`
	PreviousAccumulationTime time.Time                              `protobuf:"bytes,2,opt,name=previous_accumulation_time,json=previousAccumulationTime,proto3,stdtime" json:"previous_accumulation_time"`
	SupplyInterestFactor     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=supply_interest_factor,json=supplyInterestFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"supply_interest_factor"`
	BorrowInterestFactor     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=borrow_interest_factor,json=borrowInterestFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"borrow_interest_factor"`
}

func (m *GenesisAccumulationTime) Reset()         { *m = GenesisAccumulationTime{} }
func (m *GenesisAccumulationTime) String() string { return proto.CompactTextString(m) }
func (*GenesisAccumulationTime) ProtoMessage()    {}
func (*GenesisAccumulationTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_c49fc6faba1ed7d5, []int{1}
}
func (m *GenesisAccumulationTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccumulationTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccumulationTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccumulationTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccumulationTime.Merge(m, src)
}
func (m *GenesisAccumulationTime) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccumulationTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccumulationTime.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccumulationTime proto.InternalMessageInfo

func (m *GenesisAccumulationTime) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func (m *GenesisAccumulationTime) GetPreviousAccumulationTime() time.Time {
	if m != nil {
		return m.PreviousAccumulationTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nemo.hard.v1beta1.GenesisState")
	proto.RegisterType((*GenesisAccumulationTime)(nil), "nemo.hard.v1beta1.GenesisAccumulationTime")
}

func init() { proto.RegisterFile("nemo/hard/v1beta1/genesis.proto", fileDescriptor_c49fc6faba1ed7d5) }

var fileDescriptor_c49fc6faba1ed7d5 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0xa6, 0xff, 0xd8, 0x42, 0x0b, 0x56, 0x05, 0x6e, 0x40, 0x76, 0xd4, 0x03, 0x44,
	0x88, 0xd8, 0xb4, 0x1c, 0xb8, 0x70, 0xc1, 0x44, 0x05, 0x6e, 0xc8, 0xcd, 0x89, 0x8b, 0xb5, 0x76,
	0xa6, 0xee, 0xaa, 0xb6, 0xd7, 0xda, 0x5d, 0xa7, 0xe4, 0x1d, 0x10, 0xea, 0x73, 0x70, 0xe6, 0x09,
	0x38, 0xf5, 0x58, 0x71, 0x42, 0x1c, 0x5a, 0x94, 0xbc, 0x08, 0xf2, 0xee, 0x26, 0x29, 0x4a, 0x22,
	0x71, 0xa0, 0xa7, 0x64, 0x77, 0xbf, 0xf9, 0x7e, 0xe3, 0xd9, 0x99, 0x45, 0x4e, 0x0e, 0x19, 0xf5,
	0x8e, 0x31, 0xeb, 0x79, 0xfd, 0xbd, 0x08, 0x04, 0xde, 0xf3, 0x12, 0xc8, 0x81, 0x13, 0xee, 0x16,
	0x8c, 0x0a, 0x6a, 0xde, 0xab, 0x04, 0x6e, 0x25, 0x70, 0xb5, 0xa0, 0x61, 0xc7, 0x94, 0x67, 0x94,
	0x7b, 0x11, 0xe6, 0x30, 0x89, 0x8a, 0x29, 0xc9, 0x55, 0x48, 0x63, 0x47, 0x9d, 0x87, 0x72, 0xe5,
	0xa9, 0x85, 0x3e, 0xda, 0x4e, 0x68, 0x42, 0xd5, 0x7e, 0xf5, 0x4f, 0xef, 0x3a, 0x09, 0xa5, 0x49,
	0x0a, 0x9e, 0x5c, 0x45, 0xe5, 0x91, 0x27, 0x48, 0x06, 0x5c, 0xe0, 0xac, 0xd0, 0x82, 0x47, 0xb3,
	0x59, 0xca, 0x8c, 0xe4, 0xe9, 0xee, 0xf7, 0x15, 0x74, 0xfb, 0xad, 0x4a, 0xfa, 0x50, 0x60, 0x01,
	0xe6, 0x4b, 0xb4, 0x5a, 0x60, 0x86, 0x33, 0x6e, 0x19, 0x4d, 0xa3, 0xb5, 0xb1, 0xbf, 0xe3, 0xce,
	0x7c, 0x84, 0xfb, 0x41, 0x0a, 0xfc, 0xe5, 0xf3, 0x4b, 0xa7, 0x16, 0x68, 0xb9, 0xf9, 0xd9, 0x40,
	0x0f, 0x0b, 0x06, 0x7d, 0x42, 0x4b, 0x1e, 0xe2, 0x38, 0x2e, 0xb3, 0x32, 0xc5, 0x82, 0xd0, 0x3c,
	0x94, 0x19, 0x59, 0x4b, 0xcd, 0x7a, 0x6b, 0x63, 0xff, 0xe9, 0x1c, 0x3b, 0xcd, 0x7f, 0x7d, 0x2d,
	0xa6, 0x4b, 0x32, 0xf0, 0x9b, 0x95, 0xff, 0xd7, 0x2b, 0xc7, 0x5a, 0x20, 0xe0, 0xc1, 0xce, 0x18,
	0x38, 0x73, 0x64, 0xbe, 0x43, 0xeb, 0x3d, 0x28, 0x28, 0x27, 0x82, 0x5b, 0x75, 0x89, 0x6e, 0xcc,
	0x41, 0x77, 0x94, 0xc4, 0xbf, 0xab, 0x51, 0xeb, 0x7a, 0x83, 0x07, 0x93, 0x68, 0xb3, 0x83, 0xd6,
	0x22, 0xca, 0x18, 0x3d, 0xe5, 0xd6, 0xb2, 0x34, 0x9a, 0x57, 0x12, 0x5f, 0x2a, 0xfc, 0x2d, 0xed,
	0xb3, 0xa6, 0xd6, 0x3c, 0x18, 0x87, 0x9a, 0x0c, 0x6d, 0x0a, 0x2a, 0x70, 0x1a, 0xf2, 0xb2, 0x28,
	0x52, 0x02, 0x3d, 0x6b, 0x45, 0x9b, 0xe9, 0x4b, 0xae, 0x3a, 0x62, 0x62, 0xf7, 0x86, 0x92, 0xdc,
	0x7f, 0xae, 0xcd, 0x5a, 0x09, 0x11, 0xc7, 0x65, 0xe4, 0xc6, 0x34, 0xd3, 0x1d, 0xa1, 0x7f, 0xda,
	0xbc, 0x77, 0xe2, 0x89, 0x41, 0x01, 0x5c, 0x06, 0xf0, 0xe0, 0x8e, 0x44, 0x1c, 0x6a, 0xc2, 0x94,
	0xa9, 0x92, 0x80, 0x9e, 0xb5, 0x7a, 0x53, 0x4c, 0x5f, 0x13, 0xa6, 0x4c, 0x06, 0x1c, 0x58, 0x1f,
	0xb8, 0xb5, 0x76, 0x53, 0xcc, 0x40, 0x13, 0x76, 0xbf, 0xd4, 0xd1, 0x83, 0x05, 0x3d, 0x62, 0x3e,
	0x41, 0x5b, 0x31, 0x4d, 0x53, 0x2c, 0x80, 0xe1, 0x34, 0xac, 0x4c, 0x64, 0x63, 0xdf, 0x0a, 0x36,
	0xa7, 0xdb, 0xdd, 0x41, 0x01, 0x66, 0x84, 0x1a, 0x8b, 0xdb, 0xd7, 0x5a, 0x92, 0xc3, 0xd0, 0x70,
	0xd5, 0xb4, 0xb9, 0xe3, 0x69, 0x73, 0xbb, 0xe3, 0x69, 0xf3, 0xd7, 0xab, 0xaf, 0x38, 0xbb, 0x72,
	0x8c, 0xc0, 0x5a, 0xd4, 0x95, 0x26, 0x43, 0xf7, 0xe5, 0xf5, 0x0f, 0x42, 0x92, 0x0b, 0x60, 0xc0,
	0x45, 0x78, 0x84, 0x63, 0x41, 0x99, 0x55, 0xaf, 0x72, 0xf2, 0x5f, 0x55, 0x1e, 0xbf, 0x2e, 0x9d,
	0xc7, 0xff, 0x50, 0x89, 0x0e, 0xc4, 0x3f, 0xbe, 0xb5, 0x91, 0xae, 0x6a, 0x07, 0xe2, 0x60, 0x5b,
	0x79, 0xbf, 0xd7, 0xd6, 0x07, 0xd2, 0xb9, 0x62, 0xaa, 0xeb, 0x9f, 0x61, 0x2e, 0xff, 0x0f, 0xa6,
	0xf2, 0xfe, 0x9b, 0xe9, 0x1f, 0x9c, 0x0f, 0x6d, 0xe3, 0x62, 0x68, 0x1b, 0xbf, 0x87, 0xb6, 0x71,
	0x36, 0xb2, 0x6b, 0x17, 0x23, 0xbb, 0xf6, 0x73, 0x64, 0xd7, 0x3e, 0x3e, 0xbb, 0x46, 0x21, 0x79,
	0x5c, 0x46, 0x25, 0x6f, 0xe7, 0x20, 0x4e, 0x29, 0x3b, 0xf1, 0xe4, 0x43, 0xf5, 0x49, 0x3d, 0x55,
	0x92, 0x17, 0xad, 0xca, 0x3a, 0xbf, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x35, 0xcf, 0xd9, 0xdb,
	0x6a, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalReserves) > 0 {
		for iNdEx := len(m.TotalReserves) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalReserves[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.TotalBorrowed) > 0 {
		for iNdEx := len(m.TotalBorrowed) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalBorrowed[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.TotalSupplied) > 0 {
		for iNdEx := len(m.TotalSupplied) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalSupplied[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Borrows) > 0 {
		for iNdEx := len(m.Borrows) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Borrows[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PreviousAccumulationTimes) > 0 {
		for iNdEx := len(m.PreviousAccumulationTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PreviousAccumulationTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisAccumulationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccumulationTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccumulationTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BorrowInterestFactor.Size()
		i -= size
		if _, err := m.BorrowInterestFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SupplyInterestFactor.Size()
		i -= size
		if _, err := m.SupplyInterestFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PreviousAccumulationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PreviousAccumulationTimes) > 0 {
		for _, e := range m.PreviousAccumulationTimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Borrows) > 0 {
		for _, e := range m.Borrows {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TotalSupplied) > 0 {
		for _, e := range m.TotalSupplied {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TotalBorrowed) > 0 {
		for _, e := range m.TotalBorrowed {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TotalReserves) > 0 {
		for _, e := range m.TotalReserves {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisAccumulationTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SupplyInterestFactor.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.BorrowInterestFactor.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousAccumulationTimes = append(m.PreviousAccumulationTimes, GenesisAccumulationTime{})
			if err := m.PreviousAccumulationTimes[len(m.PreviousAccumulationTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Borrows = append(m.Borrows, Borrow{})
			if err := m.Borrows[len(m.Borrows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupplied", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalSupplied = append(m.TotalSupplied, types.Coin{})
			if err := m.TotalSupplied[len(m.TotalSupplied)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBorrowed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalBorrowed = append(m.TotalBorrowed, types.Coin{})
			if err := m.TotalBorrowed[len(m.TotalBorrowed)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalReserves", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalReserves = append(m.TotalReserves, types.Coin{})
			if err := m.TotalReserves[len(m.TotalReserves)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisAccumulationTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisAccumulationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccumulationTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PreviousAccumulationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupplyInterestFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SupplyInterestFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowInterestFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BorrowInterestFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
