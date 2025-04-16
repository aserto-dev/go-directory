// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: aserto/directory/writer/v3/writer.proto

package writer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Writer_SetObject_FullMethodName      = "/aserto.directory.writer.v3.Writer/SetObject"
	Writer_DeleteObject_FullMethodName   = "/aserto.directory.writer.v3.Writer/DeleteObject"
	Writer_SetRelation_FullMethodName    = "/aserto.directory.writer.v3.Writer/SetRelation"
	Writer_DeleteRelation_FullMethodName = "/aserto.directory.writer.v3.Writer/DeleteRelation"
	Writer_SetManifest_FullMethodName    = "/aserto.directory.writer.v3.Writer/SetManifest"
	Writer_DeleteManifest_FullMethodName = "/aserto.directory.writer.v3.Writer/DeleteManifest"
	Writer_Import_FullMethodName         = "/aserto.directory.writer.v3.Writer/Import"
)

// WriterClient is the client API for Writer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriterClient interface {
	// set object instance
	SetObject(ctx context.Context, in *SetObjectRequest, opts ...grpc.CallOption) (*SetObjectResponse, error)
	// delete object instance
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
	// set relation instance
	SetRelation(ctx context.Context, in *SetRelationRequest, opts ...grpc.CallOption) (*SetRelationResponse, error)
	// delete relation instance
	DeleteRelation(ctx context.Context, in *DeleteRelationRequest, opts ...grpc.CallOption) (*DeleteRelationResponse, error)
	// set manifest instance
	SetManifest(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SetManifestRequest, SetManifestResponse], error)
	// delete manifest instance
	DeleteManifest(ctx context.Context, in *DeleteManifestRequest, opts ...grpc.CallOption) (*DeleteManifestResponse, error)
	// import stream of objects and relations
	Import(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ImportRequest, ImportResponse], error)
}

type writerClient struct {
	cc grpc.ClientConnInterface
}

func NewWriterClient(cc grpc.ClientConnInterface) WriterClient {
	return &writerClient{cc}
}

func (c *writerClient) SetObject(ctx context.Context, in *SetObjectRequest, opts ...grpc.CallOption) (*SetObjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetObjectResponse)
	err := c.cc.Invoke(ctx, Writer_SetObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, Writer_DeleteObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) SetRelation(ctx context.Context, in *SetRelationRequest, opts ...grpc.CallOption) (*SetRelationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetRelationResponse)
	err := c.cc.Invoke(ctx, Writer_SetRelation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) DeleteRelation(ctx context.Context, in *DeleteRelationRequest, opts ...grpc.CallOption) (*DeleteRelationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRelationResponse)
	err := c.cc.Invoke(ctx, Writer_DeleteRelation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) SetManifest(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SetManifestRequest, SetManifestResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Writer_ServiceDesc.Streams[0], Writer_SetManifest_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SetManifestRequest, SetManifestResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Writer_SetManifestClient = grpc.ClientStreamingClient[SetManifestRequest, SetManifestResponse]

func (c *writerClient) DeleteManifest(ctx context.Context, in *DeleteManifestRequest, opts ...grpc.CallOption) (*DeleteManifestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteManifestResponse)
	err := c.cc.Invoke(ctx, Writer_DeleteManifest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) Import(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ImportRequest, ImportResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Writer_ServiceDesc.Streams[1], Writer_Import_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ImportRequest, ImportResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Writer_ImportClient = grpc.BidiStreamingClient[ImportRequest, ImportResponse]

// WriterServer is the server API for Writer service.
// All implementations should embed UnimplementedWriterServer
// for forward compatibility.
type WriterServer interface {
	// set object instance
	SetObject(context.Context, *SetObjectRequest) (*SetObjectResponse, error)
	// delete object instance
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
	// set relation instance
	SetRelation(context.Context, *SetRelationRequest) (*SetRelationResponse, error)
	// delete relation instance
	DeleteRelation(context.Context, *DeleteRelationRequest) (*DeleteRelationResponse, error)
	// set manifest instance
	SetManifest(grpc.ClientStreamingServer[SetManifestRequest, SetManifestResponse]) error
	// delete manifest instance
	DeleteManifest(context.Context, *DeleteManifestRequest) (*DeleteManifestResponse, error)
	// import stream of objects and relations
	Import(grpc.BidiStreamingServer[ImportRequest, ImportResponse]) error
}

// UnimplementedWriterServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWriterServer struct{}

func (UnimplementedWriterServer) SetObject(context.Context, *SetObjectRequest) (*SetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetObject not implemented")
}
func (UnimplementedWriterServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedWriterServer) SetRelation(context.Context, *SetRelationRequest) (*SetRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRelation not implemented")
}
func (UnimplementedWriterServer) DeleteRelation(context.Context, *DeleteRelationRequest) (*DeleteRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelation not implemented")
}
func (UnimplementedWriterServer) SetManifest(grpc.ClientStreamingServer[SetManifestRequest, SetManifestResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SetManifest not implemented")
}
func (UnimplementedWriterServer) DeleteManifest(context.Context, *DeleteManifestRequest) (*DeleteManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManifest not implemented")
}
func (UnimplementedWriterServer) Import(grpc.BidiStreamingServer[ImportRequest, ImportResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Import not implemented")
}
func (UnimplementedWriterServer) testEmbeddedByValue() {}

// UnsafeWriterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriterServer will
// result in compilation errors.
type UnsafeWriterServer interface {
	mustEmbedUnimplementedWriterServer()
}

func RegisterWriterServer(s grpc.ServiceRegistrar, srv WriterServer) {
	// If the following call pancis, it indicates UnimplementedWriterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Writer_ServiceDesc, srv)
}

func _Writer_SetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).SetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Writer_SetObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).SetObject(ctx, req.(*SetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Writer_DeleteObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_SetRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).SetRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Writer_SetRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).SetRelation(ctx, req.(*SetRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_DeleteRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).DeleteRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Writer_DeleteRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).DeleteRelation(ctx, req.(*DeleteRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_SetManifest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WriterServer).SetManifest(&grpc.GenericServerStream[SetManifestRequest, SetManifestResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Writer_SetManifestServer = grpc.ClientStreamingServer[SetManifestRequest, SetManifestResponse]

func _Writer_DeleteManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).DeleteManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Writer_DeleteManifest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).DeleteManifest(ctx, req.(*DeleteManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_Import_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WriterServer).Import(&grpc.GenericServerStream[ImportRequest, ImportResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Writer_ImportServer = grpc.BidiStreamingServer[ImportRequest, ImportResponse]

// Writer_ServiceDesc is the grpc.ServiceDesc for Writer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Writer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.directory.writer.v3.Writer",
	HandlerType: (*WriterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetObject",
			Handler:    _Writer_SetObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _Writer_DeleteObject_Handler,
		},
		{
			MethodName: "SetRelation",
			Handler:    _Writer_SetRelation_Handler,
		},
		{
			MethodName: "DeleteRelation",
			Handler:    _Writer_DeleteRelation_Handler,
		},
		{
			MethodName: "DeleteManifest",
			Handler:    _Writer_DeleteManifest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetManifest",
			Handler:       _Writer_SetManifest_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Import",
			Handler:       _Writer_Import_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "aserto/directory/writer/v3/writer.proto",
}
