// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: aserto/directory/exporter/v2/exporter.proto

package exporter

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
	Exporter_Export_FullMethodName = "/aserto.directory.exporter.v2.Exporter/Export"
)

// ExporterClient is the client API for Exporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExporterClient interface {
	// Deprecated: Do not use.
	// export objects and relations as a stream
	Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ExportResponse], error)
}

type exporterClient struct {
	cc grpc.ClientConnInterface
}

func NewExporterClient(cc grpc.ClientConnInterface) ExporterClient {
	return &exporterClient{cc}
}

// Deprecated: Do not use.
func (c *exporterClient) Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ExportResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Exporter_ServiceDesc.Streams[0], Exporter_Export_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ExportRequest, ExportResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Exporter_ExportClient = grpc.ServerStreamingClient[ExportResponse]

// ExporterServer is the server API for Exporter service.
// All implementations should embed UnimplementedExporterServer
// for forward compatibility.
type ExporterServer interface {
	// Deprecated: Do not use.
	// export objects and relations as a stream
	Export(*ExportRequest, grpc.ServerStreamingServer[ExportResponse]) error
}

// UnimplementedExporterServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExporterServer struct{}

func (UnimplementedExporterServer) Export(*ExportRequest, grpc.ServerStreamingServer[ExportResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedExporterServer) testEmbeddedByValue() {}

// UnsafeExporterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExporterServer will
// result in compilation errors.
type UnsafeExporterServer interface {
	mustEmbedUnimplementedExporterServer()
}

func RegisterExporterServer(s grpc.ServiceRegistrar, srv ExporterServer) {
	// If the following call pancis, it indicates UnimplementedExporterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Exporter_ServiceDesc, srv)
}

func _Exporter_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExportRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExporterServer).Export(m, &grpc.GenericServerStream[ExportRequest, ExportResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Exporter_ExportServer = grpc.ServerStreamingServer[ExportResponse]

// Exporter_ServiceDesc is the grpc.ServiceDesc for Exporter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exporter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.directory.exporter.v2.Exporter",
	HandlerType: (*ExporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Export",
			Handler:       _Exporter_Export_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "aserto/directory/exporter/v2/exporter.proto",
}
