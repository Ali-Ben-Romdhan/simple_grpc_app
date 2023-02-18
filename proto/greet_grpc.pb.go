// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/greet.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	SayHelloFromServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_SayHelloFromServerStreamingClient, error)
	SayHelloFromClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloFromClientStreamingClient, error)
	SayHelloBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBidirectionalStreamingClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/greet_service.GreetService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) SayHelloFromServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_SayHelloFromServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet_service.GreetService/SayHelloFromServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloFromServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_SayHelloFromServerStreamingClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayHelloFromServerStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloFromServerStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SayHelloFromClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloFromClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet_service.GreetService/SayHelloFromClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloFromClientStreamingClient{stream}
	return x, nil
}

type GreetService_SayHelloFromClientStreamingClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*MessagesList, error)
	grpc.ClientStream
}

type greetServiceSayHelloFromClientStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloFromClientStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayHelloFromClientStreamingClient) CloseAndRecv() (*MessagesList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SayHelloBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/greet_service.GreetService/SayHelloBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloBidirectionalStreamingClient{stream}
	return x, nil
}

type GreetService_SayHelloBidirectionalStreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayHelloBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloBidirectionalStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayHelloBidirectionalStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	SayHello(context.Context, *NoParam) (*HelloResponse, error)
	SayHelloFromServerStreaming(*NamesList, GreetService_SayHelloFromServerStreamingServer) error
	SayHelloFromClientStreaming(GreetService_SayHelloFromClientStreamingServer) error
	SayHelloBidirectionalStreaming(GreetService_SayHelloBidirectionalStreamingServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) SayHello(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloFromServerStreaming(*NamesList, GreetService_SayHelloFromServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloFromServerStreaming not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloFromClientStreaming(GreetService_SayHelloFromClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloFromClientStreaming not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloBidirectionalStreaming(GreetService_SayHelloBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidirectionalStreaming not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet_service.GreetService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).SayHello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_SayHelloFromServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NamesList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).SayHelloFromServerStreaming(m, &greetServiceSayHelloFromServerStreamingServer{stream})
}

type GreetService_SayHelloFromServerStreamingServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetServiceSayHelloFromServerStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloFromServerStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_SayHelloFromClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloFromClientStreaming(&greetServiceSayHelloFromClientStreamingServer{stream})
}

type GreetService_SayHelloFromClientStreamingServer interface {
	SendAndClose(*MessagesList) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceSayHelloFromClientStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloFromClientStreamingServer) SendAndClose(m *MessagesList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayHelloFromClientStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_SayHelloBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloBidirectionalStreaming(&greetServiceSayHelloBidirectionalStreamingServer{stream})
}

type GreetService_SayHelloBidirectionalStreamingServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceSayHelloBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloBidirectionalStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayHelloBidirectionalStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreetService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloFromServerStreaming",
			Handler:       _GreetService_SayHelloFromServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloFromClientStreaming",
			Handler:       _GreetService_SayHelloFromClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloBidirectionalStreaming",
			Handler:       _GreetService_SayHelloBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}