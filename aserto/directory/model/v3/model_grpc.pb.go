// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aserto/directory/model/v3/model.proto

package model

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

const (
	Model_GetManifest_FullMethodName    = "/aserto.directory.model.v3.Model/GetManifest"
	Model_SetManifest_FullMethodName    = "/aserto.directory.model.v3.Model/SetManifest"
	Model_DeleteManifest_FullMethodName = "/aserto.directory.model.v3.Model/DeleteManifest"
)

// ModelClient is the client API for Model service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelClient interface {
	GetManifest(ctx context.Context, in *GetManifestRequest, opts ...grpc.CallOption) (Model_GetManifestClient, error)
	SetManifest(ctx context.Context, opts ...grpc.CallOption) (Model_SetManifestClient, error)
	DeleteManifest(ctx context.Context, in *DeleteManifestRequest, opts ...grpc.CallOption) (*DeleteManifestResponse, error)
}

type modelClient struct {
	cc grpc.ClientConnInterface
}

func NewModelClient(cc grpc.ClientConnInterface) ModelClient {
	return &modelClient{cc}
}

func (c *modelClient) GetManifest(ctx context.Context, in *GetManifestRequest, opts ...grpc.CallOption) (Model_GetManifestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Model_ServiceDesc.Streams[0], Model_GetManifest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &modelGetManifestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Model_GetManifestClient interface {
	Recv() (*GetManifestResponse, error)
	grpc.ClientStream
}

type modelGetManifestClient struct {
	grpc.ClientStream
}

func (x *modelGetManifestClient) Recv() (*GetManifestResponse, error) {
	m := new(GetManifestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *modelClient) SetManifest(ctx context.Context, opts ...grpc.CallOption) (Model_SetManifestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Model_ServiceDesc.Streams[1], Model_SetManifest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &modelSetManifestClient{stream}
	return x, nil
}

type Model_SetManifestClient interface {
	Send(*SetManifestRequest) error
	CloseAndRecv() (*SetManifestResponse, error)
	grpc.ClientStream
}

type modelSetManifestClient struct {
	grpc.ClientStream
}

func (x *modelSetManifestClient) Send(m *SetManifestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *modelSetManifestClient) CloseAndRecv() (*SetManifestResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SetManifestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *modelClient) DeleteManifest(ctx context.Context, in *DeleteManifestRequest, opts ...grpc.CallOption) (*DeleteManifestResponse, error) {
	out := new(DeleteManifestResponse)
	err := c.cc.Invoke(ctx, Model_DeleteManifest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelServer is the server API for Model service.
// All implementations should embed UnimplementedModelServer
// for forward compatibility
type ModelServer interface {
	GetManifest(*GetManifestRequest, Model_GetManifestServer) error
	SetManifest(Model_SetManifestServer) error
	DeleteManifest(context.Context, *DeleteManifestRequest) (*DeleteManifestResponse, error)
}

// UnimplementedModelServer should be embedded to have forward compatible implementations.
type UnimplementedModelServer struct {
}

func (UnimplementedModelServer) GetManifest(*GetManifestRequest, Model_GetManifestServer) error {
	return status.Errorf(codes.Unimplemented, "method GetManifest not implemented")
}
func (UnimplementedModelServer) SetManifest(Model_SetManifestServer) error {
	return status.Errorf(codes.Unimplemented, "method SetManifest not implemented")
}
func (UnimplementedModelServer) DeleteManifest(context.Context, *DeleteManifestRequest) (*DeleteManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManifest not implemented")
}

// UnsafeModelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServer will
// result in compilation errors.
type UnsafeModelServer interface {
	mustEmbedUnimplementedModelServer()
}

func RegisterModelServer(s grpc.ServiceRegistrar, srv ModelServer) {
	s.RegisterService(&Model_ServiceDesc, srv)
}

func _Model_GetManifest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetManifestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModelServer).GetManifest(m, &modelGetManifestServer{stream})
}

type Model_GetManifestServer interface {
	Send(*GetManifestResponse) error
	grpc.ServerStream
}

type modelGetManifestServer struct {
	grpc.ServerStream
}

func (x *modelGetManifestServer) Send(m *GetManifestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Model_SetManifest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModelServer).SetManifest(&modelSetManifestServer{stream})
}

type Model_SetManifestServer interface {
	SendAndClose(*SetManifestResponse) error
	Recv() (*SetManifestRequest, error)
	grpc.ServerStream
}

type modelSetManifestServer struct {
	grpc.ServerStream
}

func (x *modelSetManifestServer) SendAndClose(m *SetManifestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *modelSetManifestServer) Recv() (*SetManifestRequest, error) {
	m := new(SetManifestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Model_DeleteManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServer).DeleteManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Model_DeleteManifest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServer).DeleteManifest(ctx, req.(*DeleteManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Model_ServiceDesc is the grpc.ServiceDesc for Model service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Model_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.directory.model.v3.Model",
	HandlerType: (*ModelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteManifest",
			Handler:    _Model_DeleteManifest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetManifest",
			Handler:       _Model_GetManifest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SetManifest",
			Handler:       _Model_SetManifest_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "aserto/directory/model/v3/model.proto",
}
