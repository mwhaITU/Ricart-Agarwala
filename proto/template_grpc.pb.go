// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateClient interface {
	//one message is sent and one is recieved
	Increment(ctx context.Context, in *Amount, opts ...grpc.CallOption) (*Ack, error)
	// many messages are sent and one is recieved
	SayHi(ctx context.Context, opts ...grpc.CallOption) (Template_SayHiClient, error)
}

type templateClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateClient(cc grpc.ClientConnInterface) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) Increment(ctx context.Context, in *Amount, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/proto.Template/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) SayHi(ctx context.Context, opts ...grpc.CallOption) (Template_SayHiClient, error) {
	stream, err := c.cc.NewStream(ctx, &Template_ServiceDesc.Streams[0], "/proto.Template/SayHi", opts...)
	if err != nil {
		return nil, err
	}
	x := &templateSayHiClient{stream}
	return x, nil
}

type Template_SayHiClient interface {
	Send(*Greeding) error
	CloseAndRecv() (*Farewell, error)
	grpc.ClientStream
}

type templateSayHiClient struct {
	grpc.ClientStream
}

func (x *templateSayHiClient) Send(m *Greeding) error {
	return x.ClientStream.SendMsg(m)
}

func (x *templateSayHiClient) CloseAndRecv() (*Farewell, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Farewell)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TemplateServer is the server API for Template service.
// All implementations must embed UnimplementedTemplateServer
// for forward compatibility
type TemplateServer interface {
	//one message is sent and one is recieved
	Increment(context.Context, *Amount) (*Ack, error)
	// many messages are sent and one is recieved
	SayHi(Template_SayHiServer) error
	mustEmbedUnimplementedTemplateServer()
}

// UnimplementedTemplateServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateServer struct {
}

func (UnimplementedTemplateServer) Increment(context.Context, *Amount) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedTemplateServer) SayHi(Template_SayHiServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedTemplateServer) mustEmbedUnimplementedTemplateServer() {}

// UnsafeTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServer will
// result in compilation errors.
type UnsafeTemplateServer interface {
	mustEmbedUnimplementedTemplateServer()
}

func RegisterTemplateServer(s grpc.ServiceRegistrar, srv TemplateServer) {
	s.RegisterService(&Template_ServiceDesc, srv)
}

func _Template_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Amount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Increment(ctx, req.(*Amount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_SayHi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TemplateServer).SayHi(&templateSayHiServer{stream})
}

type Template_SayHiServer interface {
	SendAndClose(*Farewell) error
	Recv() (*Greeding, error)
	grpc.ServerStream
}

type templateSayHiServer struct {
	grpc.ServerStream
}

func (x *templateSayHiServer) SendAndClose(m *Farewell) error {
	return x.ServerStream.SendMsg(m)
}

func (x *templateSayHiServer) Recv() (*Greeding, error) {
	m := new(Greeding)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Template_ServiceDesc is the grpc.ServiceDesc for Template service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Template_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _Template_Increment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHi",
			Handler:       _Template_SayHi_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/template.proto",
}
