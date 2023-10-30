// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aserto/directory/exporter/v3/exporter.proto

package exporter

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
	Exporter_Export_FullMethodName = "/aserto.directory.exporter.v3.Exporter/Export"
)

// ExporterClient is the client API for Exporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExporterClient interface {
	Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (Exporter_ExportClient, error)
}

type exporterClient struct {
	cc grpc.ClientConnInterface
}

func NewExporterClient(cc grpc.ClientConnInterface) ExporterClient {
	return &exporterClient{cc}
}

func (c *exporterClient) Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (Exporter_ExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &Exporter_ServiceDesc.Streams[0], Exporter_Export_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &exporterExportClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Exporter_ExportClient interface {
	Recv() (*ExportResponse, error)
	grpc.ClientStream
}

type exporterExportClient struct {
	grpc.ClientStream
}

func (x *exporterExportClient) Recv() (*ExportResponse, error) {
	m := new(ExportResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExporterServer is the server API for Exporter service.
// All implementations should embed UnimplementedExporterServer
// for forward compatibility
type ExporterServer interface {
	Export(*ExportRequest, Exporter_ExportServer) error
}

// UnimplementedExporterServer should be embedded to have forward compatible implementations.
type UnimplementedExporterServer struct {
}

func (UnimplementedExporterServer) Export(*ExportRequest, Exporter_ExportServer) error {
	return status.Errorf(codes.Unimplemented, "method Export not implemented")
}

// UnsafeExporterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExporterServer will
// result in compilation errors.
type UnsafeExporterServer interface {
	mustEmbedUnimplementedExporterServer()
}

func RegisterExporterServer(s grpc.ServiceRegistrar, srv ExporterServer) {
	s.RegisterService(&Exporter_ServiceDesc, srv)
}

func _Exporter_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExportRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExporterServer).Export(m, &exporterExportServer{stream})
}

type Exporter_ExportServer interface {
	Send(*ExportResponse) error
	grpc.ServerStream
}

type exporterExportServer struct {
	grpc.ServerStream
}

func (x *exporterExportServer) Send(m *ExportResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Exporter_ServiceDesc is the grpc.ServiceDesc for Exporter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exporter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.directory.exporter.v3.Exporter",
	HandlerType: (*ExporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Export",
			Handler:       _Exporter_Export_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "aserto/directory/exporter/v3/exporter.proto",
}
