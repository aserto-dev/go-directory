// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aserto/directory/reader/v3/reader.proto

package reader

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
	Reader_GetObject_FullMethodName       = "/aserto.directory.reader.v3.Reader/GetObject"
	Reader_GetObjectMany_FullMethodName   = "/aserto.directory.reader.v3.Reader/GetObjectMany"
	Reader_GetObjects_FullMethodName      = "/aserto.directory.reader.v3.Reader/GetObjects"
	Reader_GetRelation_FullMethodName     = "/aserto.directory.reader.v3.Reader/GetRelation"
	Reader_GetRelations_FullMethodName    = "/aserto.directory.reader.v3.Reader/GetRelations"
	Reader_CheckPermission_FullMethodName = "/aserto.directory.reader.v3.Reader/CheckPermission"
	Reader_CheckRelation_FullMethodName   = "/aserto.directory.reader.v3.Reader/CheckRelation"
	Reader_GetGraph_FullMethodName        = "/aserto.directory.reader.v3.Reader/GetGraph"
)

// ReaderClient is the client API for Reader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReaderClient interface {
	// object methods
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
	GetObjectMany(ctx context.Context, in *GetObjectManyRequest, opts ...grpc.CallOption) (*GetObjectManyResponse, error)
	GetObjects(ctx context.Context, in *GetObjectsRequest, opts ...grpc.CallOption) (*GetObjectsResponse, error)
	// relation methods
	GetRelation(ctx context.Context, in *GetRelationRequest, opts ...grpc.CallOption) (*GetRelationResponse, error)
	GetRelations(ctx context.Context, in *GetRelationsRequest, opts ...grpc.CallOption) (*GetRelationsResponse, error)
	// check permission method
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error)
	// check relation method
	CheckRelation(ctx context.Context, in *CheckRelationRequest, opts ...grpc.CallOption) (*CheckRelationResponse, error)
	// graph methods
	GetGraph(ctx context.Context, in *GetGraphRequest, opts ...grpc.CallOption) (*GetGraphResponse, error)
}

type readerClient struct {
	cc grpc.ClientConnInterface
}

func NewReaderClient(cc grpc.ClientConnInterface) ReaderClient {
	return &readerClient{cc}
}

func (c *readerClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, Reader_GetObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerClient) GetObjectMany(ctx context.Context, in *GetObjectManyRequest, opts ...grpc.CallOption) (*GetObjectManyResponse, error) {
	out := new(GetObjectManyResponse)
	err := c.cc.Invoke(ctx, Reader_GetObjectMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerClient) GetObjects(ctx context.Context, in *GetObjectsRequest, opts ...grpc.CallOption) (*GetObjectsResponse, error) {
	out := new(GetObjectsResponse)
	err := c.cc.Invoke(ctx, Reader_GetObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerClient) GetRelation(ctx context.Context, in *GetRelationRequest, opts ...grpc.CallOption) (*GetRelationResponse, error) {
	out := new(GetRelationResponse)
	err := c.cc.Invoke(ctx, Reader_GetRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerClient) GetRelations(ctx context.Context, in *GetRelationsRequest, opts ...grpc.CallOption) (*GetRelationsResponse, error) {
	out := new(GetRelationsResponse)
	err := c.cc.Invoke(ctx, Reader_GetRelations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error) {
	out := new(CheckPermissionResponse)
	err := c.cc.Invoke(ctx, Reader_CheckPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerClient) CheckRelation(ctx context.Context, in *CheckRelationRequest, opts ...grpc.CallOption) (*CheckRelationResponse, error) {
	out := new(CheckRelationResponse)
	err := c.cc.Invoke(ctx, Reader_CheckRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerClient) GetGraph(ctx context.Context, in *GetGraphRequest, opts ...grpc.CallOption) (*GetGraphResponse, error) {
	out := new(GetGraphResponse)
	err := c.cc.Invoke(ctx, Reader_GetGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReaderServer is the server API for Reader service.
// All implementations should embed UnimplementedReaderServer
// for forward compatibility
type ReaderServer interface {
	// object methods
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	GetObjectMany(context.Context, *GetObjectManyRequest) (*GetObjectManyResponse, error)
	GetObjects(context.Context, *GetObjectsRequest) (*GetObjectsResponse, error)
	// relation methods
	GetRelation(context.Context, *GetRelationRequest) (*GetRelationResponse, error)
	GetRelations(context.Context, *GetRelationsRequest) (*GetRelationsResponse, error)
	// check permission method
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error)
	// check relation method
	CheckRelation(context.Context, *CheckRelationRequest) (*CheckRelationResponse, error)
	// graph methods
	GetGraph(context.Context, *GetGraphRequest) (*GetGraphResponse, error)
}

// UnimplementedReaderServer should be embedded to have forward compatible implementations.
type UnimplementedReaderServer struct {
}

func (UnimplementedReaderServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedReaderServer) GetObjectMany(context.Context, *GetObjectManyRequest) (*GetObjectManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectMany not implemented")
}
func (UnimplementedReaderServer) GetObjects(context.Context, *GetObjectsRequest) (*GetObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjects not implemented")
}
func (UnimplementedReaderServer) GetRelation(context.Context, *GetRelationRequest) (*GetRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelation not implemented")
}
func (UnimplementedReaderServer) GetRelations(context.Context, *GetRelationsRequest) (*GetRelationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelations not implemented")
}
func (UnimplementedReaderServer) CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedReaderServer) CheckRelation(context.Context, *CheckRelationRequest) (*CheckRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRelation not implemented")
}
func (UnimplementedReaderServer) GetGraph(context.Context, *GetGraphRequest) (*GetGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraph not implemented")
}

// UnsafeReaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReaderServer will
// result in compilation errors.
type UnsafeReaderServer interface {
	mustEmbedUnimplementedReaderServer()
}

func RegisterReaderServer(s grpc.ServiceRegistrar, srv ReaderServer) {
	s.RegisterService(&Reader_ServiceDesc, srv)
}

func _Reader_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reader_GetObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reader_GetObjectMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).GetObjectMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reader_GetObjectMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).GetObjectMany(ctx, req.(*GetObjectManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reader_GetObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).GetObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reader_GetObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).GetObjects(ctx, req.(*GetObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reader_GetRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).GetRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reader_GetRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).GetRelation(ctx, req.(*GetRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reader_GetRelations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).GetRelations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reader_GetRelations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).GetRelations(ctx, req.(*GetRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reader_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reader_CheckPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reader_CheckRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).CheckRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reader_CheckRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).CheckRelation(ctx, req.(*CheckRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reader_GetGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).GetGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reader_GetGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).GetGraph(ctx, req.(*GetGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reader_ServiceDesc is the grpc.ServiceDesc for Reader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.directory.reader.v3.Reader",
	HandlerType: (*ReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetObject",
			Handler:    _Reader_GetObject_Handler,
		},
		{
			MethodName: "GetObjectMany",
			Handler:    _Reader_GetObjectMany_Handler,
		},
		{
			MethodName: "GetObjects",
			Handler:    _Reader_GetObjects_Handler,
		},
		{
			MethodName: "GetRelation",
			Handler:    _Reader_GetRelation_Handler,
		},
		{
			MethodName: "GetRelations",
			Handler:    _Reader_GetRelations_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _Reader_CheckPermission_Handler,
		},
		{
			MethodName: "CheckRelation",
			Handler:    _Reader_CheckRelation_Handler,
		},
		{
			MethodName: "GetGraph",
			Handler:    _Reader_GetGraph_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/directory/reader/v3/reader.proto",
}