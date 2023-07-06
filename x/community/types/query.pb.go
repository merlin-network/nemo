// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo/community/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryBalanceRequest defines the request type for querying x/community balance.
type QueryBalanceRequest struct {
}

func (m *QueryBalanceRequest) Reset()         { *m = QueryBalanceRequest{} }
func (m *QueryBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceRequest) ProtoMessage()    {}
func (*QueryBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae59ad859ec597e6, []int{0}
}
func (m *QueryBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceRequest.Merge(m, src)
}
func (m *QueryBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceRequest proto.InternalMessageInfo

// QueryBalanceResponse defines the response type for querying x/community balance.
type QueryBalanceResponse struct {
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *QueryBalanceResponse) Reset()         { *m = QueryBalanceResponse{} }
func (m *QueryBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceResponse) ProtoMessage()    {}
func (*QueryBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae59ad859ec597e6, []int{1}
}
func (m *QueryBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceResponse.Merge(m, src)
}
func (m *QueryBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceResponse proto.InternalMessageInfo

func (m *QueryBalanceResponse) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// QueryTotalBalanceRequest defines the request type for querying total community pool balance.
type QueryTotalBalanceRequest struct {
}

func (m *QueryTotalBalanceRequest) Reset()         { *m = QueryTotalBalanceRequest{} }
func (m *QueryTotalBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalBalanceRequest) ProtoMessage()    {}
func (*QueryTotalBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae59ad859ec597e6, []int{2}
}
func (m *QueryTotalBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalBalanceRequest.Merge(m, src)
}
func (m *QueryTotalBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalBalanceRequest proto.InternalMessageInfo

// QueryTotalBalanceResponse defines the response type for querying total
// community pool balance. This matches the x/distribution CommunityPool query response.
type QueryTotalBalanceResponse struct {
	// pool defines community pool's coins.
	Pool github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=pool,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"pool"`
}

func (m *QueryTotalBalanceResponse) Reset()         { *m = QueryTotalBalanceResponse{} }
func (m *QueryTotalBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTotalBalanceResponse) ProtoMessage()    {}
func (*QueryTotalBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae59ad859ec597e6, []int{3}
}
func (m *QueryTotalBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalBalanceResponse.Merge(m, src)
}
func (m *QueryTotalBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalBalanceResponse proto.InternalMessageInfo

func (m *QueryTotalBalanceResponse) GetPool() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Pool
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryBalanceRequest)(nil), "nemo.community.v1beta1.QueryBalanceRequest")
	proto.RegisterType((*QueryBalanceResponse)(nil), "nemo.community.v1beta1.QueryBalanceResponse")
	proto.RegisterType((*QueryTotalBalanceRequest)(nil), "nemo.community.v1beta1.QueryTotalBalanceRequest")
	proto.RegisterType((*QueryTotalBalanceResponse)(nil), "nemo.community.v1beta1.QueryTotalBalanceResponse")
}

func init() {
	proto.RegisterFile("nemo/community/v1beta1/query.proto", fileDescriptor_ae59ad859ec597e6)
}

var fileDescriptor_ae59ad859ec597e6 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x93, 0x42, 0x41, 0x32, 0x9c, 0x42, 0x41, 0x6d, 0x54, 0xa5, 0x10, 0x09, 0xb5, 0x52,
	0xa9, 0xdd, 0xb4, 0x6f, 0x50, 0xb8, 0x71, 0xa2, 0xe2, 0xc4, 0x05, 0x39, 0xc1, 0x0a, 0x51, 0x13,
	0x7f, 0x69, 0xed, 0x00, 0xb9, 0xf6, 0x8e, 0x84, 0xc4, 0x1b, 0x70, 0xe4, 0x19, 0xf6, 0x00, 0x3d,
	0x56, 0xda, 0x65, 0xa7, 0x6d, 0x6a, 0xf7, 0x20, 0x53, 0x1c, 0xaf, 0xea, 0xa6, 0x74, 0xea, 0x4e,
	0xb6, 0xec, 0xef, 0xff, 0xff, 0xff, 0xfc, 0x7d, 0x46, 0x2e, 0x67, 0x09, 0x90, 0x00, 0x92, 0x24,
	0xe3, 0x91, 0xcc, 0xc9, 0x0f, 0xcf, 0x67, 0x92, 0x7a, 0x64, 0x9e, 0xb1, 0x45, 0x8e, 0xd3, 0x05,
	0x48, 0xb0, 0x5e, 0x15, 0x35, 0x78, 0x57, 0x83, 0x75, 0x8d, 0xed, 0x04, 0x20, 0x12, 0x10, 0xc4,
	0xa7, 0x82, 0xed, 0x84, 0x01, 0x44, 0xbc, 0xd4, 0xd9, 0x8d, 0x10, 0x42, 0x50, 0x5b, 0x52, 0xec,
	0xf4, 0x69, 0x3b, 0x04, 0x08, 0x63, 0x46, 0x68, 0x1a, 0x11, 0xca, 0x39, 0x48, 0x2a, 0x23, 0xe0,
	0xa2, 0xbc, 0x75, 0x5f, 0xa2, 0x17, 0x9f, 0x8a, 0xe8, 0x09, 0x8d, 0x29, 0x0f, 0xd8, 0x94, 0xcd,
	0x33, 0x26, 0xa4, 0x9b, 0xa3, 0xc6, 0xed, 0x63, 0x91, 0x02, 0x17, 0xcc, 0xa2, 0xa8, 0x5e, 0x04,
	0x8a, 0xa6, 0xf9, 0xfa, 0x51, 0xef, 0xd9, 0xa8, 0x85, 0x4b, 0x24, 0x5c, 0x20, 0xdd, 0x70, 0xe2,
	0xf7, 0x10, 0xf1, 0xc9, 0x70, 0x75, 0xde, 0x31, 0xfe, 0x5f, 0x74, 0x7a, 0x61, 0x24, 0xbf, 0x67,
	0x7e, 0xf1, 0x1c, 0xa2, 0xf9, 0xcb, 0x65, 0x20, 0xbe, 0xcd, 0x88, 0xcc, 0x53, 0x26, 0x94, 0x40,
	0x4c, 0x4b, 0x67, 0xd7, 0x46, 0x4d, 0x15, 0xfd, 0x19, 0x24, 0x8d, 0xef, 0x60, 0x2d, 0x4d, 0xd4,
	0xaa, 0xb8, 0xd4, 0x70, 0x0c, 0x3d, 0x4e, 0x01, 0x62, 0xcd, 0xd6, 0xae, 0x64, 0xfb, 0xc0, 0x02,
	0x85, 0x37, 0xd6, 0x78, 0xfd, 0x23, 0xf0, 0xb4, 0x46, 0x4c, 0x95, 0xfd, 0xe8, 0xa4, 0x86, 0xea,
	0x0a, 0xc2, 0xfa, 0x6d, 0xa2, 0xa7, 0x1a, 0xc2, 0xea, 0xe3, 0xea, 0xa9, 0xe1, 0x8a, 0xf6, 0xda,
	0xef, 0x8e, 0x2b, 0x2e, 0xdf, 0xe5, 0x76, 0x97, 0xa7, 0x57, 0x7f, 0x6b, 0x6f, 0xac, 0x0e, 0x39,
	0xf0, 0x79, 0x7c, 0xcd, 0xf0, 0xcf, 0x44, 0xcf, 0xf7, 0x3b, 0x63, 0x0d, 0xef, 0xcd, 0xa9, 0xe8,
	0xb0, 0xed, 0x3d, 0x40, 0xa1, 0xf1, 0x06, 0x0a, 0xaf, 0x6b, 0xbd, 0x3d, 0x84, 0x27, 0x0b, 0xd5,
	0x57, 0x0d, 0x39, 0xf9, 0xb8, 0xda, 0x38, 0xe6, 0x7a, 0xe3, 0x98, 0x97, 0x1b, 0xc7, 0xfc, 0xb3,
	0x75, 0x8c, 0xf5, 0xd6, 0x31, 0xce, 0xb6, 0x8e, 0xf1, 0xc5, 0xdb, 0x9b, 0x45, 0xc4, 0x83, 0xcc,
	0xcf, 0xc4, 0x80, 0x33, 0xf9, 0x13, 0x16, 0xb3, 0xd2, 0xfa, 0xd7, 0x9e, 0xb9, 0x1a, 0x8d, 0xff,
	0x44, 0xfd, 0xe2, 0xf1, 0x75, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xaa, 0x9c, 0x7b, 0x57, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Balance queries the balance of all coins of x/community module.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// TotalBalance queries the balance of all coins, including x/distribution,
	// x/community, and supplied balances.
	TotalBalance(ctx context.Context, in *QueryTotalBalanceRequest, opts ...grpc.CallOption) (*QueryTotalBalanceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/nemo.community.v1beta1.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalBalance(ctx context.Context, in *QueryTotalBalanceRequest, opts ...grpc.CallOption) (*QueryTotalBalanceResponse, error) {
	out := new(QueryTotalBalanceResponse)
	err := c.cc.Invoke(ctx, "/nemo.community.v1beta1.Query/TotalBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Balance queries the balance of all coins of x/community module.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// TotalBalance queries the balance of all coins, including x/distribution,
	// x/community, and supplied balances.
	TotalBalance(context.Context, *QueryTotalBalanceRequest) (*QueryTotalBalanceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Balance(ctx context.Context, req *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (*UnimplementedQueryServer) TotalBalance(ctx context.Context, req *QueryTotalBalanceRequest) (*QueryTotalBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalBalance not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo.community.v1beta1.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo.community.v1beta1.Query/TotalBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalBalance(ctx, req.(*QueryTotalBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nemo.community.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "TotalBalance",
			Handler:    _Query_TotalBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nemo/community/v1beta1/query.proto",
}

func (m *QueryBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryTotalBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryTotalBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pool) > 0 {
		for iNdEx := len(m.Pool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryTotalBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryTotalBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pool) > 0 {
		for _, e := range m.Pool {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryBalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTotalBalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTotalBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTotalBalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTotalBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pool = append(m.Pool, types.DecCoin{})
			if err := m.Pool[len(m.Pool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
