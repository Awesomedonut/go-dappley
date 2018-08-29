// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/rpc/pb/rpc.proto

package rpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pb "github.com/dappley/go-dappley/network/pb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message
type CreateWalletRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWalletRequest) Reset()         { *m = CreateWalletRequest{} }
func (m *CreateWalletRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWalletRequest) ProtoMessage()    {}
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{0}
}
func (m *CreateWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWalletRequest.Unmarshal(m, b)
}
func (m *CreateWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWalletRequest.Marshal(b, m, deterministic)
}
func (dst *CreateWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWalletRequest.Merge(dst, src)
}
func (m *CreateWalletRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWalletRequest.Size(m)
}
func (m *CreateWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWalletRequest proto.InternalMessageInfo

func (m *CreateWalletRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetBalanceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceRequest) Reset()         { *m = GetBalanceRequest{} }
func (m *GetBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetBalanceRequest) ProtoMessage()    {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{1}
}
func (m *GetBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceRequest.Unmarshal(m, b)
}
func (m *GetBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceRequest.Marshal(b, m, deterministic)
}
func (dst *GetBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceRequest.Merge(dst, src)
}
func (m *GetBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_GetBalanceRequest.Size(m)
}
func (m *GetBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceRequest proto.InternalMessageInfo

func (m *GetBalanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetBalanceRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type SendRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Ammount              int64    `protobuf:"varint,4,opt,name=ammount,proto3" json:"ammount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{2}
}
func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (dst *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(dst, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SendRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendRequest) GetAmmount() int64 {
	if m != nil {
		return m.Ammount
	}
	return 0
}

type GetPeerInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPeerInfoRequest) Reset()         { *m = GetPeerInfoRequest{} }
func (m *GetPeerInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPeerInfoRequest) ProtoMessage()    {}
func (*GetPeerInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{3}
}
func (m *GetPeerInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeerInfoRequest.Unmarshal(m, b)
}
func (m *GetPeerInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeerInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetPeerInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeerInfoRequest.Merge(dst, src)
}
func (m *GetPeerInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPeerInfoRequest.Size(m)
}
func (m *GetPeerInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeerInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeerInfoRequest proto.InternalMessageInfo

type GetBlockchainInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockchainInfoRequest) Reset()         { *m = GetBlockchainInfoRequest{} }
func (m *GetBlockchainInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockchainInfoRequest) ProtoMessage()    {}
func (*GetBlockchainInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{4}
}
func (m *GetBlockchainInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockchainInfoRequest.Unmarshal(m, b)
}
func (m *GetBlockchainInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockchainInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetBlockchainInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockchainInfoRequest.Merge(dst, src)
}
func (m *GetBlockchainInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockchainInfoRequest.Size(m)
}
func (m *GetBlockchainInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockchainInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockchainInfoRequest proto.InternalMessageInfo

type CreateWalletResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWalletResponse) Reset()         { *m = CreateWalletResponse{} }
func (m *CreateWalletResponse) String() string { return proto.CompactTextString(m) }
func (*CreateWalletResponse) ProtoMessage()    {}
func (*CreateWalletResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{5}
}
func (m *CreateWalletResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWalletResponse.Unmarshal(m, b)
}
func (m *CreateWalletResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWalletResponse.Marshal(b, m, deterministic)
}
func (dst *CreateWalletResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWalletResponse.Merge(dst, src)
}
func (m *CreateWalletResponse) XXX_Size() int {
	return xxx_messageInfo_CreateWalletResponse.Size(m)
}
func (m *CreateWalletResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWalletResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWalletResponse proto.InternalMessageInfo

func (m *CreateWalletResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateWalletResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type GetBalanceResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Ammount              int64    `protobuf:"varint,2,opt,name=ammount,proto3" json:"ammount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceResponse) Reset()         { *m = GetBalanceResponse{} }
func (m *GetBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*GetBalanceResponse) ProtoMessage()    {}
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{6}
}
func (m *GetBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceResponse.Unmarshal(m, b)
}
func (m *GetBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceResponse.Marshal(b, m, deterministic)
}
func (dst *GetBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceResponse.Merge(dst, src)
}
func (m *GetBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_GetBalanceResponse.Size(m)
}
func (m *GetBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceResponse proto.InternalMessageInfo

func (m *GetBalanceResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetBalanceResponse) GetAmmount() int64 {
	if m != nil {
		return m.Ammount
	}
	return 0
}

type SendResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{7}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (dst *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(dst, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetPeerInfoResponse struct {
	PeerList             *pb.Peerlist `protobuf:"bytes,1,opt,name=peerList,proto3" json:"peerList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetPeerInfoResponse) Reset()         { *m = GetPeerInfoResponse{} }
func (m *GetPeerInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetPeerInfoResponse) ProtoMessage()    {}
func (*GetPeerInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{8}
}
func (m *GetPeerInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeerInfoResponse.Unmarshal(m, b)
}
func (m *GetPeerInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeerInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetPeerInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeerInfoResponse.Merge(dst, src)
}
func (m *GetPeerInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetPeerInfoResponse.Size(m)
}
func (m *GetPeerInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeerInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeerInfoResponse proto.InternalMessageInfo

func (m *GetPeerInfoResponse) GetPeerList() *pb.Peerlist {
	if m != nil {
		return m.PeerList
	}
	return nil
}

type GetBlockchainInfoResponse struct {
	TailBlockHash        []byte   `protobuf:"bytes,1,opt,name=tailBlockHash,proto3" json:"tailBlockHash,omitempty"`
	BlockHeight          uint64   `protobuf:"varint,2,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockchainInfoResponse) Reset()         { *m = GetBlockchainInfoResponse{} }
func (m *GetBlockchainInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockchainInfoResponse) ProtoMessage()    {}
func (*GetBlockchainInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_448b3d3d3b74811b, []int{9}
}
func (m *GetBlockchainInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockchainInfoResponse.Unmarshal(m, b)
}
func (m *GetBlockchainInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockchainInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetBlockchainInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockchainInfoResponse.Merge(dst, src)
}
func (m *GetBlockchainInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockchainInfoResponse.Size(m)
}
func (m *GetBlockchainInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockchainInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockchainInfoResponse proto.InternalMessageInfo

func (m *GetBlockchainInfoResponse) GetTailBlockHash() []byte {
	if m != nil {
		return m.TailBlockHash
	}
	return nil
}

func (m *GetBlockchainInfoResponse) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateWalletRequest)(nil), "rpcpb.CreateWalletRequest")
	proto.RegisterType((*GetBalanceRequest)(nil), "rpcpb.GetBalanceRequest")
	proto.RegisterType((*SendRequest)(nil), "rpcpb.SendRequest")
	proto.RegisterType((*GetPeerInfoRequest)(nil), "rpcpb.GetPeerInfoRequest")
	proto.RegisterType((*GetBlockchainInfoRequest)(nil), "rpcpb.GetBlockchainInfoRequest")
	proto.RegisterType((*CreateWalletResponse)(nil), "rpcpb.CreateWalletResponse")
	proto.RegisterType((*GetBalanceResponse)(nil), "rpcpb.GetBalanceResponse")
	proto.RegisterType((*SendResponse)(nil), "rpcpb.SendResponse")
	proto.RegisterType((*GetPeerInfoResponse)(nil), "rpcpb.GetPeerInfoResponse")
	proto.RegisterType((*GetBlockchainInfoResponse)(nil), "rpcpb.GetBlockchainInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConnectClient is the client API for Connect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectClient interface {
	RpcCreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	RpcGetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	RpcSend(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	RpcGetPeerInfo(ctx context.Context, in *GetPeerInfoRequest, opts ...grpc.CallOption) (*GetPeerInfoResponse, error)
	RpcGetBlockchainInfo(ctx context.Context, in *GetBlockchainInfoRequest, opts ...grpc.CallOption) (*GetBlockchainInfoResponse, error)
}

type connectClient struct {
	cc *grpc.ClientConn
}

func NewConnectClient(cc *grpc.ClientConn) ConnectClient {
	return &connectClient{cc}
}

func (c *connectClient) RpcCreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcCreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) RpcGetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcGetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) RpcSend(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) RpcGetPeerInfo(ctx context.Context, in *GetPeerInfoRequest, opts ...grpc.CallOption) (*GetPeerInfoResponse, error) {
	out := new(GetPeerInfoResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcGetPeerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) RpcGetBlockchainInfo(ctx context.Context, in *GetBlockchainInfoRequest, opts ...grpc.CallOption) (*GetBlockchainInfoResponse, error) {
	out := new(GetBlockchainInfoResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.Connect/RpcGetBlockchainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServer is the server API for Connect service.
type ConnectServer interface {
	RpcCreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	RpcGetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	RpcSend(context.Context, *SendRequest) (*SendResponse, error)
	RpcGetPeerInfo(context.Context, *GetPeerInfoRequest) (*GetPeerInfoResponse, error)
	RpcGetBlockchainInfo(context.Context, *GetBlockchainInfoRequest) (*GetBlockchainInfoResponse, error)
}

func RegisterConnectServer(s *grpc.Server, srv ConnectServer) {
	s.RegisterService(&_Connect_serviceDesc, srv)
}

func _Connect_RpcCreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcCreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcCreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcCreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_RpcGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcGetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcGetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_RpcSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcSend(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_RpcGetPeerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcGetPeerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcGetPeerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcGetPeerInfo(ctx, req.(*GetPeerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_RpcGetBlockchainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockchainInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).RpcGetBlockchainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Connect/RpcGetBlockchainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).RpcGetBlockchainInfo(ctx, req.(*GetBlockchainInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.Connect",
	HandlerType: (*ConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcCreateWallet",
			Handler:    _Connect_RpcCreateWallet_Handler,
		},
		{
			MethodName: "RpcGetBalance",
			Handler:    _Connect_RpcGetBalance_Handler,
		},
		{
			MethodName: "RpcSend",
			Handler:    _Connect_RpcSend_Handler,
		},
		{
			MethodName: "RpcGetPeerInfo",
			Handler:    _Connect_RpcGetPeerInfo_Handler,
		},
		{
			MethodName: "RpcGetBlockchainInfo",
			Handler:    _Connect_RpcGetBlockchainInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/dappley/go-dappley/rpc/pb/rpc.proto",
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/rpc/pb/rpc.proto", fileDescriptor_rpc_448b3d3d3b74811b)
}

var fileDescriptor_rpc_448b3d3d3b74811b = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0xe2, 0xe4, 0x6b, 0x60, 0xd2, 0x16, 0xb1, 0xc9, 0xc1, 0x35, 0x07, 0x22, 0x8b, 0x43,
	0x38, 0xe0, 0x48, 0x05, 0x89, 0x33, 0xad, 0xd4, 0xa6, 0x28, 0x87, 0x6a, 0x39, 0xf4, 0xc0, 0x01,
	0xad, 0x37, 0xd3, 0xc4, 0xaa, 0xbd, 0xbb, 0xec, 0x6e, 0x84, 0xf8, 0x3b, 0xfc, 0x4f, 0x24, 0xe4,
	0xf5, 0xda, 0xd8, 0x60, 0x02, 0x27, 0xef, 0xce, 0xcc, 0x9b, 0x99, 0xf7, 0xf6, 0x19, 0x92, 0x6d,
	0x66, 0x77, 0xfb, 0x34, 0xe1, 0xb2, 0x58, 0x6e, 0x98, 0x52, 0x39, 0x7e, 0x5d, 0x6e, 0xe5, 0xab,
	0xfa, 0xa8, 0x15, 0x5f, 0xaa, 0xb4, 0xfc, 0x24, 0x4a, 0x4b, 0x2b, 0xc9, 0xff, 0x5a, 0x71, 0x95,
	0x46, 0x6f, 0x0f, 0xc3, 0x04, 0xda, 0x2f, 0x52, 0x3f, 0x94, 0x50, 0x85, 0xa8, 0xf3, 0xcc, 0xd8,
	0x0a, 0x1f, 0xbf, 0x84, 0xe9, 0xa5, 0x46, 0x66, 0xf1, 0x8e, 0xe5, 0x39, 0x5a, 0x8a, 0x9f, 0xf7,
	0x68, 0x2c, 0x21, 0x30, 0x12, 0xac, 0xc0, 0x70, 0x30, 0x1f, 0x2c, 0x1e, 0x53, 0x77, 0x8e, 0xdf,
	0xc1, 0xd3, 0x6b, 0xb4, 0x17, 0x2c, 0x67, 0x82, 0xe3, 0x81, 0x42, 0x12, 0xc2, 0x98, 0x6d, 0x36,
	0x1a, 0x8d, 0x09, 0x03, 0x17, 0xae, 0xaf, 0xf1, 0x27, 0x98, 0x7c, 0x40, 0xb1, 0x39, 0x04, 0x26,
	0x30, 0xba, 0xd7, 0xb2, 0xf0, 0x48, 0x77, 0x26, 0xa7, 0x10, 0x58, 0x19, 0x0e, 0x5d, 0x24, 0xb0,
	0xd2, 0x0d, 0x28, 0x0a, 0xb9, 0x17, 0x36, 0x1c, 0xcd, 0x07, 0x8b, 0x21, 0xad, 0xaf, 0xf1, 0x0c,
	0xc8, 0x35, 0xda, 0x5b, 0x44, 0x7d, 0x23, 0xee, 0xa5, 0x9f, 0x13, 0x47, 0x10, 0x96, 0x9b, 0xe7,
	0x92, 0x3f, 0xf0, 0x1d, 0xcb, 0x44, 0x3b, 0xf7, 0x1e, 0x66, 0x5d, 0x01, 0x8c, 0x92, 0xc2, 0x38,
	0x12, 0x05, 0x1a, 0xc3, 0xb6, 0xf5, 0x7a, 0xf5, 0xf5, 0x00, 0xbd, 0x95, 0x9b, 0xde, 0x28, 0xf4,
	0x4f, 0x9d, 0x3c, 0x8f, 0xa0, 0xcb, 0x63, 0x01, 0xc7, 0x95, 0x50, 0x7f, 0xeb, 0x11, 0x5f, 0xc1,
	0xb4, 0xc3, 0xd8, 0x03, 0x96, 0xf0, 0xa8, 0x7c, 0xe9, 0x75, 0x66, 0xac, 0x43, 0x4c, 0xce, 0xa7,
	0x89, 0x77, 0x81, 0x4a, 0x93, 0x5b, 0x6f, 0x02, 0xda, 0x14, 0xc5, 0x1c, 0xce, 0x7a, 0x34, 0xf2,
	0xdd, 0x5e, 0xc0, 0x89, 0x65, 0x59, 0xee, 0xb2, 0x2b, 0x66, 0x76, 0xae, 0xe5, 0x31, 0xed, 0x06,
	0xc9, 0x1c, 0x26, 0xa9, 0xbb, 0x60, 0xb6, 0xdd, 0x55, 0x94, 0x46, 0xb4, 0x1d, 0x3a, 0xff, 0x1e,
	0xc0, 0xf8, 0x52, 0x0a, 0x81, 0xdc, 0x92, 0x35, 0x3c, 0xa1, 0x8a, 0xb7, 0xb5, 0x27, 0x51, 0xe2,
	0xdc, 0x9c, 0xf4, 0x38, 0x32, 0x7a, 0xd6, 0x9b, 0xab, 0xf6, 0x8b, 0xff, 0x23, 0x57, 0x70, 0x42,
	0x15, 0xff, 0xa9, 0x3e, 0x09, 0x7d, 0xfd, 0x6f, 0x96, 0x8d, 0xce, 0x7a, 0x32, 0x4d, 0x9f, 0x37,
	0x30, 0xa6, 0x8a, 0x97, 0xda, 0x13, 0xe2, 0xeb, 0x5a, 0x8e, 0x8d, 0xa6, 0x9d, 0x58, 0x83, 0xba,
	0x81, 0xd3, 0x6a, 0x7a, 0xfd, 0x0e, 0xa4, 0x35, 0xe4, 0x17, 0x37, 0x46, 0x51, 0x5f, 0xaa, 0x69,
	0xf5, 0x11, 0x66, 0x9e, 0x48, 0xe7, 0x29, 0xc8, 0xf3, 0xd6, 0xd6, 0x7d, 0x46, 0x8e, 0xe6, 0x7f,
	0x2e, 0xa8, 0x9b, 0x5f, 0x1c, 0x7d, 0x0b, 0x86, 0xab, 0xf5, 0x5d, 0x7a, 0xe4, 0x7e, 0xfe, 0xd7,
	0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xfc, 0x11, 0x17, 0x6e, 0x04, 0x00, 0x00,
}
