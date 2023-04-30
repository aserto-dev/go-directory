// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aserto/directory/importer/v2/importer.proto

package importer

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
	Importer_Import_FullMethodName = "/aserto.directory.importer.v2.Importer/Import"
)

// ImporterClient is the client API for Importer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImporterClient interface {
	Import(ctx context.Context, opts ...grpc.CallOption) (Importer_ImportClient, error)
}

type importerClient struct {
	cc grpc.ClientConnInterface
}

func NewImporterClient(cc grpc.ClientConnInterface) ImporterClient {
	return &importerClient{cc}
}

func (c *importerClient) Import(ctx context.Context, opts ...grpc.CallOption) (Importer_ImportClient, error) {
	stream, err := c.cc.NewStream(ctx, &Importer_ServiceDesc.Streams[0], Importer_Import_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &importerImportClient{stream}
	return x, nil
}

type Importer_ImportClient interface {
	Send(*ImportRequest) error
	Recv() (*ImportResponse, error)
	grpc.ClientStream
}

type importerImportClient struct {
	grpc.ClientStream
}

func (x *importerImportClient) Send(m *ImportRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *importerImportClient) Recv() (*ImportResponse, error) {
	m := new(ImportResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImporterServer is the server API for Importer service.
// All implementations should embed UnimplementedImporterServer
// for forward compatibility
type ImporterServer interface {
	Import(Importer_ImportServer) error
}

// UnimplementedImporterServer should be embedded to have forward compatible implementations.
type UnimplementedImporterServer struct {
}

func (UnimplementedImporterServer) Import(Importer_ImportServer) error {
	return status.Errorf(codes.Unimplemented, "method Import not implemented")
}

// UnsafeImporterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImporterServer will
// result in compilation errors.
type UnsafeImporterServer interface {
	mustEmbedUnimplementedImporterServer()
}

func RegisterImporterServer(s grpc.ServiceRegistrar, srv ImporterServer) {
	s.RegisterService(&Importer_ServiceDesc, srv)
}

func _Importer_Import_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImporterServer).Import(&importerImportServer{stream})
}

type Importer_ImportServer interface {
	Send(*ImportResponse) error
	Recv() (*ImportRequest, error)
	grpc.ServerStream
}

type importerImportServer struct {
	grpc.ServerStream
}

func (x *importerImportServer) Send(m *ImportResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *importerImportServer) Recv() (*ImportRequest, error) {
	m := new(ImportRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Importer_ServiceDesc is the grpc.ServiceDesc for Importer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Importer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.directory.importer.v2.Importer",
	HandlerType: (*ImporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Import",
			Handler:       _Importer_Import_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "aserto/directory/importer/v2/importer.proto",
}
