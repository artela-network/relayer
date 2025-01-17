// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: penumbra/cnidarium/v1alpha1/cnidarium.proto

package cnidariumv1alpha1

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v8/modules/core/23-commitment/types"
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

// Performs a key-value query, either by key or by key hash.
//
// Proofs are only supported by key.
type KeyValueRequest struct {
	// The expected chain id (empty string if no expectation).
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// If set, the key to fetch from storage.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// whether to return a proof
	Proof bool `protobuf:"varint,3,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *KeyValueRequest) Reset()         { *m = KeyValueRequest{} }
func (m *KeyValueRequest) String() string { return proto.CompactTextString(m) }
func (*KeyValueRequest) ProtoMessage()    {}
func (*KeyValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6006139f070ea05, []int{0}
}
func (m *KeyValueRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyValueRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValueRequest.Merge(m, src)
}
func (m *KeyValueRequest) XXX_Size() int {
	return m.Size()
}
func (m *KeyValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValueRequest proto.InternalMessageInfo

func (m *KeyValueRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *KeyValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValueRequest) GetProof() bool {
	if m != nil {
		return m.Proof
	}
	return false
}

type KeyValueResponse struct {
	// The value corresponding to the specified key, if it was found.
	Value *KeyValueResponse_Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// A proof of existence or non-existence.
	Proof *types.MerkleProof `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *KeyValueResponse) Reset()         { *m = KeyValueResponse{} }
func (m *KeyValueResponse) String() string { return proto.CompactTextString(m) }
func (*KeyValueResponse) ProtoMessage()    {}
func (*KeyValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6006139f070ea05, []int{1}
}
func (m *KeyValueResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyValueResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValueResponse.Merge(m, src)
}
func (m *KeyValueResponse) XXX_Size() int {
	return m.Size()
}
func (m *KeyValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValueResponse proto.InternalMessageInfo

func (m *KeyValueResponse) GetValue() *KeyValueResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyValueResponse) GetProof() *types.MerkleProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type KeyValueResponse_Value struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValueResponse_Value) Reset()         { *m = KeyValueResponse_Value{} }
func (m *KeyValueResponse_Value) String() string { return proto.CompactTextString(m) }
func (*KeyValueResponse_Value) ProtoMessage()    {}
func (*KeyValueResponse_Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6006139f070ea05, []int{1, 0}
}
func (m *KeyValueResponse_Value) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyValueResponse_Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyValueResponse_Value.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyValueResponse_Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValueResponse_Value.Merge(m, src)
}
func (m *KeyValueResponse_Value) XXX_Size() int {
	return m.Size()
}
func (m *KeyValueResponse_Value) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValueResponse_Value.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValueResponse_Value proto.InternalMessageInfo

func (m *KeyValueResponse_Value) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Performs a prefixed key-value query, by string prefix.
type PrefixValueRequest struct {
	// The expected chain id (empty string if no expectation).
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// The prefix to fetch subkeys from storage.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (m *PrefixValueRequest) Reset()         { *m = PrefixValueRequest{} }
func (m *PrefixValueRequest) String() string { return proto.CompactTextString(m) }
func (*PrefixValueRequest) ProtoMessage()    {}
func (*PrefixValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6006139f070ea05, []int{2}
}
func (m *PrefixValueRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrefixValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrefixValueRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrefixValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrefixValueRequest.Merge(m, src)
}
func (m *PrefixValueRequest) XXX_Size() int {
	return m.Size()
}
func (m *PrefixValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrefixValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrefixValueRequest proto.InternalMessageInfo

func (m *PrefixValueRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *PrefixValueRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type PrefixValueResponse struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PrefixValueResponse) Reset()         { *m = PrefixValueResponse{} }
func (m *PrefixValueResponse) String() string { return proto.CompactTextString(m) }
func (*PrefixValueResponse) ProtoMessage()    {}
func (*PrefixValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6006139f070ea05, []int{3}
}
func (m *PrefixValueResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrefixValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrefixValueResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrefixValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrefixValueResponse.Merge(m, src)
}
func (m *PrefixValueResponse) XXX_Size() int {
	return m.Size()
}
func (m *PrefixValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrefixValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrefixValueResponse proto.InternalMessageInfo

func (m *PrefixValueResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PrefixValueResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyValueRequest)(nil), "penumbra.cnidarium.v1alpha1.KeyValueRequest")
	proto.RegisterType((*KeyValueResponse)(nil), "penumbra.cnidarium.v1alpha1.KeyValueResponse")
	proto.RegisterType((*KeyValueResponse_Value)(nil), "penumbra.cnidarium.v1alpha1.KeyValueResponse.Value")
	proto.RegisterType((*PrefixValueRequest)(nil), "penumbra.cnidarium.v1alpha1.PrefixValueRequest")
	proto.RegisterType((*PrefixValueResponse)(nil), "penumbra.cnidarium.v1alpha1.PrefixValueResponse")
}

func init() {
	proto.RegisterFile("penumbra/cnidarium/v1alpha1/cnidarium.proto", fileDescriptor_a6006139f070ea05)
}

var fileDescriptor_a6006139f070ea05 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xcd, 0x38, 0x4a, 0x09, 0xd3, 0x0a, 0xaa, 0x01, 0xa1, 0x90, 0xaa, 0xa6, 0x0a, 0x8b, 0x56,
	0x02, 0x66, 0x9a, 0x74, 0x45, 0x10, 0x9b, 0x74, 0x51, 0x55, 0xa8, 0x92, 0x09, 0xa8, 0x02, 0x14,
	0x09, 0x8d, 0x9d, 0xdb, 0x66, 0xd4, 0xd8, 0x63, 0xc6, 0x0f, 0x91, 0xbf, 0xe0, 0x1b, 0x58, 0xb2,
	0x64, 0xc3, 0x2f, 0x20, 0x56, 0x5d, 0xb2, 0x44, 0xc9, 0x0e, 0x7e, 0x02, 0xf9, 0x31, 0x76, 0x0a,
	0x92, 0x95, 0xee, 0x7c, 0xcf, 0x3d, 0xf7, 0xcc, 0xb9, 0xc7, 0x33, 0xf8, 0x91, 0x0f, 0x5e, 0xe4,
	0xda, 0x8a, 0x33, 0xc7, 0x13, 0x63, 0xae, 0x44, 0xe4, 0xb2, 0xb8, 0xcb, 0xa7, 0xfe, 0x84, 0x77,
	0x4b, 0x88, 0xfa, 0x4a, 0x86, 0x92, 0x6c, 0x69, 0x32, 0x2d, 0x3b, 0x9a, 0xdc, 0xde, 0x15, 0xb6,
	0xc3, 0x1c, 0xa9, 0x80, 0x39, 0xd2, 0x75, 0x45, 0xe8, 0x82, 0x17, 0xb2, 0xb8, 0xbb, 0x54, 0x65,
	0x2a, 0x9d, 0xd7, 0xf8, 0xf6, 0x0b, 0x98, 0x9d, 0xf2, 0x69, 0x04, 0x43, 0xf8, 0x10, 0x41, 0x10,
	0x92, 0xfb, 0xb8, 0xe9, 0x4c, 0xb8, 0xf0, 0xde, 0x8b, 0x71, 0x0b, 0xed, 0xa0, 0xbd, 0x9b, 0xc3,
	0x1b, 0x69, 0x7d, 0x3c, 0x26, 0x9b, 0xb8, 0x7e, 0x01, 0xb3, 0x96, 0x91, 0xa2, 0xc9, 0x27, 0xb9,
	0x8b, 0x1b, 0xbe, 0x92, 0xf2, 0xac, 0x55, 0xdf, 0x41, 0x7b, 0xcd, 0x61, 0x56, 0x74, 0xbe, 0x21,
	0xbc, 0x59, 0xca, 0x06, 0xbe, 0xf4, 0x02, 0x20, 0xc7, 0xb8, 0x11, 0x27, 0x40, 0x2a, 0xba, 0xde,
	0x3b, 0xa0, 0x15, 0x0b, 0xd0, 0x7f, 0xa7, 0x69, 0x56, 0x65, 0x0a, 0xe4, 0xa9, 0x3e, 0xd5, 0x48,
	0xa5, 0x1e, 0x52, 0x61, 0x3b, 0x34, 0x59, 0x97, 0x2e, 0x2d, 0x18, 0x77, 0xe9, 0x09, 0xa8, 0x8b,
	0x29, 0x58, 0x09, 0x35, 0xb7, 0xd6, 0xde, 0xc6, 0x8d, 0x54, 0x2a, 0x71, 0x5e, 0xda, 0xd9, 0xc8,
	0x95, 0x3b, 0x47, 0x98, 0x58, 0x0a, 0xce, 0xc4, 0xc7, 0x55, 0x23, 0xb9, 0x87, 0xd7, 0xfc, 0x74,
	0x20, 0x4f, 0x25, 0xaf, 0x3a, 0xcf, 0xf1, 0x9d, 0x2b, 0x42, 0x79, 0x08, 0x79, 0x82, 0xe8, 0x4a,
	0x82, 0x99, 0x0f, 0x63, 0xc9, 0x47, 0xef, 0x0f, 0xc2, 0x1b, 0x2f, 0x23, 0x50, 0xb3, 0x57, 0xa0,
	0x62, 0xe1, 0x00, 0x39, 0xc7, 0x4d, 0x9d, 0x09, 0x79, 0xbc, 0x62, 0x74, 0xa9, 0xf9, 0xf6, 0x93,
	0x6b, 0x05, 0x4d, 0x14, 0x5e, 0x5f, 0x32, 0x4e, 0x58, 0xe5, 0xf4, 0xff, 0x59, 0xb5, 0xf7, 0x57,
	0x1f, 0xc8, 0x4e, 0xdc, 0x47, 0x83, 0xaf, 0xc6, 0xf7, 0xb9, 0x89, 0x2e, 0xe7, 0x26, 0xfa, 0x35,
	0x37, 0xd1, 0xa7, 0x85, 0x59, 0xbb, 0x5c, 0x98, 0xb5, 0x9f, 0x0b, 0xb3, 0x86, 0x1f, 0x38, 0xd2,
	0xad, 0x52, 0x1c, 0xdc, 0x3a, 0xd4, 0x98, 0x95, 0xdc, 0x68, 0x0b, 0xbd, 0x7b, 0x7b, 0x2e, 0xc2,
	0x49, 0x64, 0x27, 0x77, 0x81, 0x39, 0x32, 0x70, 0x65, 0xc0, 0x14, 0x4c, 0xf9, 0x0c, 0x14, 0x8b,
	0x7b, 0xc5, 0x67, 0xfa, 0x03, 0x03, 0x56, 0xf1, 0xe4, 0x9e, 0x15, 0x90, 0x46, 0x3e, 0x1b, 0x75,
	0xeb, 0xf0, 0xcd, 0x17, 0x63, 0xcb, 0xd2, 0x86, 0x8a, 0xc3, 0xe9, 0x69, 0xce, 0xf9, 0x51, 0x76,
	0x47, 0x45, 0x77, 0xa4, 0xbb, 0x73, 0x63, 0xb7, 0xa2, 0x3b, 0x3a, 0xb2, 0x06, 0x27, 0x10, 0xf2,
	0x31, 0x0f, 0xf9, 0x6f, 0x63, 0x5b, 0x33, 0xfb, 0xfd, 0x82, 0xda, 0xef, 0x6b, 0xae, 0xbd, 0x96,
	0xbe, 0xe0, 0x83, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x61, 0x0d, 0xb7, 0x36, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// General-purpose key-value state query API, that can be used to query
	// arbitrary keys in the JMT storage.
	KeyValue(ctx context.Context, in *KeyValueRequest, opts ...grpc.CallOption) (*KeyValueResponse, error)
	// General-purpose prefixed key-value state query API, that can be used to query
	// arbitrary prefixes in the JMT storage.
	// Returns a stream of `PrefixValueResponse`s.
	PrefixValue(ctx context.Context, in *PrefixValueRequest, opts ...grpc.CallOption) (QueryService_PrefixValueClient, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) KeyValue(ctx context.Context, in *KeyValueRequest, opts ...grpc.CallOption) (*KeyValueResponse, error) {
	out := new(KeyValueResponse)
	err := c.cc.Invoke(ctx, "/penumbra.cnidarium.v1alpha1.QueryService/KeyValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) PrefixValue(ctx context.Context, in *PrefixValueRequest, opts ...grpc.CallOption) (QueryService_PrefixValueClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[0], "/penumbra.cnidarium.v1alpha1.QueryService/PrefixValue", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServicePrefixValueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_PrefixValueClient interface {
	Recv() (*PrefixValueResponse, error)
	grpc.ClientStream
}

type queryServicePrefixValueClient struct {
	grpc.ClientStream
}

func (x *queryServicePrefixValueClient) Recv() (*PrefixValueResponse, error) {
	m := new(PrefixValueResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// General-purpose key-value state query API, that can be used to query
	// arbitrary keys in the JMT storage.
	KeyValue(context.Context, *KeyValueRequest) (*KeyValueResponse, error)
	// General-purpose prefixed key-value state query API, that can be used to query
	// arbitrary prefixes in the JMT storage.
	// Returns a stream of `PrefixValueResponse`s.
	PrefixValue(*PrefixValueRequest, QueryService_PrefixValueServer) error
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) KeyValue(ctx context.Context, req *KeyValueRequest) (*KeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyValue not implemented")
}
func (*UnimplementedQueryServiceServer) PrefixValue(req *PrefixValueRequest, srv QueryService_PrefixValueServer) error {
	return status.Errorf(codes.Unimplemented, "method PrefixValue not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_KeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).KeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/penumbra.cnidarium.v1alpha1.QueryService/KeyValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).KeyValue(ctx, req.(*KeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_PrefixValue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrefixValueRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).PrefixValue(m, &queryServicePrefixValueServer{stream})
}

type QueryService_PrefixValueServer interface {
	Send(*PrefixValueResponse) error
	grpc.ServerStream
}

type queryServicePrefixValueServer struct {
	grpc.ServerStream
}

func (x *queryServicePrefixValueServer) Send(m *PrefixValueResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "penumbra.cnidarium.v1alpha1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeyValue",
			Handler:    _QueryService_KeyValue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrefixValue",
			Handler:       _QueryService_PrefixValue_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "penumbra/cnidarium/v1alpha1/cnidarium.proto",
}

func (m *KeyValueRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyValueRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyValueRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof {
		i--
		if m.Proof {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintCnidarium(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintCnidarium(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeyValueResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyValueResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyValueResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCnidarium(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCnidarium(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeyValueResponse_Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyValueResponse_Value) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyValueResponse_Value) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCnidarium(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PrefixValueRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrefixValueRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrefixValueRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Prefix) > 0 {
		i -= len(m.Prefix)
		copy(dAtA[i:], m.Prefix)
		i = encodeVarintCnidarium(dAtA, i, uint64(len(m.Prefix)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintCnidarium(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PrefixValueResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrefixValueResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrefixValueResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCnidarium(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintCnidarium(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCnidarium(dAtA []byte, offset int, v uint64) int {
	offset -= sovCnidarium(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyValueRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovCnidarium(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCnidarium(uint64(l))
	}
	if m.Proof {
		n += 2
	}
	return n
}

func (m *KeyValueResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovCnidarium(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovCnidarium(uint64(l))
	}
	return n
}

func (m *KeyValueResponse_Value) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCnidarium(uint64(l))
	}
	return n
}

func (m *PrefixValueRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovCnidarium(uint64(l))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovCnidarium(uint64(l))
	}
	return n
}

func (m *PrefixValueResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCnidarium(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCnidarium(uint64(l))
	}
	return n
}

func sovCnidarium(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCnidarium(x uint64) (n int) {
	return sovCnidarium(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyValueRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCnidarium
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
			return fmt.Errorf("proto: KeyValueRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValueRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
			m.Proof = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCnidarium(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCnidarium
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
func (m *KeyValueResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCnidarium
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
			return fmt.Errorf("proto: KeyValueResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValueResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &KeyValueResponse_Value{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &types.MerkleProof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCnidarium(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCnidarium
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
func (m *KeyValueResponse_Value) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCnidarium
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
			return fmt.Errorf("proto: Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCnidarium(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCnidarium
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
func (m *PrefixValueRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCnidarium
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
			return fmt.Errorf("proto: PrefixValueRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrefixValueRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCnidarium(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCnidarium
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
func (m *PrefixValueResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCnidarium
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
			return fmt.Errorf("proto: PrefixValueResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrefixValueResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCnidarium
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
				return ErrInvalidLengthCnidarium
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCnidarium
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCnidarium(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCnidarium
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
func skipCnidarium(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCnidarium
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
					return 0, ErrIntOverflowCnidarium
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
					return 0, ErrIntOverflowCnidarium
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
				return 0, ErrInvalidLengthCnidarium
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCnidarium
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCnidarium
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCnidarium        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCnidarium          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCnidarium = fmt.Errorf("proto: unexpected end of group")
)
