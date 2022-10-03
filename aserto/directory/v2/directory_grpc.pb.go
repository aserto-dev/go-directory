// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: aserto/directory/v2/directory.proto

package directory

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

// DirectoryClient is the client API for Directory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryClient interface {
	// object type metadata methods
	GetObjectType(ctx context.Context, in *GetObjectTypeRequest, opts ...grpc.CallOption) (*GetObjectTypeResponse, error)
	GetObjectTypes(ctx context.Context, in *GetObjectTypesRequest, opts ...grpc.CallOption) (*GetObjectTypesResponse, error)
	// relation type metadata methods
	GetRelationType(ctx context.Context, in *GetRelationTypeRequest, opts ...grpc.CallOption) (*GetRelationTypeResponse, error)
	GetRelationTypes(ctx context.Context, in *GetRelationTypesRequest, opts ...grpc.CallOption) (*GetRelationTypesResponse, error)
	// permission metadata methods
	GetPermission(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*GetPermissionResponse, error)
	GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error)
	// object methods
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
	GetObjectMany(ctx context.Context, in *GetObjectManyRequest, opts ...grpc.CallOption) (*GetObjectManyResponse, error)
	GetObjects(ctx context.Context, in *GetObjectsRequest, opts ...grpc.CallOption) (*GetObjectsResponse, error)
	// relation methods
	GetRelation(ctx context.Context, in *GetRelationRequest, opts ...grpc.CallOption) (*GetRelationResponse, error)
	GetRelations(ctx context.Context, in *GetRelationsRequest, opts ...grpc.CallOption) (*GetRelationsResponse, error)
	// check methods
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error)
	CheckRelation(ctx context.Context, in *CheckRelationRequest, opts ...grpc.CallOption) (*CheckRelationResponse, error)
	// graph methods
	GetGraph(ctx context.Context, in *GetGraphRequest, opts ...grpc.CallOption) (*GetGraphResponse, error)
}

type directoryClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryClient(cc grpc.ClientConnInterface) DirectoryClient {
	return &directoryClient{cc}
}

func (c *directoryClient) GetObjectType(ctx context.Context, in *GetObjectTypeRequest, opts ...grpc.CallOption) (*GetObjectTypeResponse, error) {
	out := new(GetObjectTypeResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetObjectType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetObjectTypes(ctx context.Context, in *GetObjectTypesRequest, opts ...grpc.CallOption) (*GetObjectTypesResponse, error) {
	out := new(GetObjectTypesResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetObjectTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetRelationType(ctx context.Context, in *GetRelationTypeRequest, opts ...grpc.CallOption) (*GetRelationTypeResponse, error) {
	out := new(GetRelationTypeResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetRelationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetRelationTypes(ctx context.Context, in *GetRelationTypesRequest, opts ...grpc.CallOption) (*GetRelationTypesResponse, error) {
	out := new(GetRelationTypesResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetRelationTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetPermission(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*GetPermissionResponse, error) {
	out := new(GetPermissionResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error) {
	out := new(GetPermissionsResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetObjectMany(ctx context.Context, in *GetObjectManyRequest, opts ...grpc.CallOption) (*GetObjectManyResponse, error) {
	out := new(GetObjectManyResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetObjectMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetObjects(ctx context.Context, in *GetObjectsRequest, opts ...grpc.CallOption) (*GetObjectsResponse, error) {
	out := new(GetObjectsResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetRelation(ctx context.Context, in *GetRelationRequest, opts ...grpc.CallOption) (*GetRelationResponse, error) {
	out := new(GetRelationResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetRelations(ctx context.Context, in *GetRelationsRequest, opts ...grpc.CallOption) (*GetRelationsResponse, error) {
	out := new(GetRelationsResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetRelations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error) {
	out := new(CheckPermissionResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) CheckRelation(ctx context.Context, in *CheckRelationRequest, opts ...grpc.CallOption) (*CheckRelationResponse, error) {
	out := new(CheckRelationResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/CheckRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetGraph(ctx context.Context, in *GetGraphRequest, opts ...grpc.CallOption) (*GetGraphResponse, error) {
	out := new(GetGraphResponse)
	err := c.cc.Invoke(ctx, "/aserto.directory.v2.Directory/GetGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryServer is the server API for Directory service.
// All implementations should embed UnimplementedDirectoryServer
// for forward compatibility
type DirectoryServer interface {
	// object type metadata methods
	GetObjectType(context.Context, *GetObjectTypeRequest) (*GetObjectTypeResponse, error)
	GetObjectTypes(context.Context, *GetObjectTypesRequest) (*GetObjectTypesResponse, error)
	// relation type metadata methods
	GetRelationType(context.Context, *GetRelationTypeRequest) (*GetRelationTypeResponse, error)
	GetRelationTypes(context.Context, *GetRelationTypesRequest) (*GetRelationTypesResponse, error)
	// permission metadata methods
	GetPermission(context.Context, *GetPermissionRequest) (*GetPermissionResponse, error)
	GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error)
	// object methods
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	GetObjectMany(context.Context, *GetObjectManyRequest) (*GetObjectManyResponse, error)
	GetObjects(context.Context, *GetObjectsRequest) (*GetObjectsResponse, error)
	// relation methods
	GetRelation(context.Context, *GetRelationRequest) (*GetRelationResponse, error)
	GetRelations(context.Context, *GetRelationsRequest) (*GetRelationsResponse, error)
	// check methods
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error)
	CheckRelation(context.Context, *CheckRelationRequest) (*CheckRelationResponse, error)
	// graph methods
	GetGraph(context.Context, *GetGraphRequest) (*GetGraphResponse, error)
}

// UnimplementedDirectoryServer should be embedded to have forward compatible implementations.
type UnimplementedDirectoryServer struct {
}

func (UnimplementedDirectoryServer) GetObjectType(context.Context, *GetObjectTypeRequest) (*GetObjectTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectType not implemented")
}
func (UnimplementedDirectoryServer) GetObjectTypes(context.Context, *GetObjectTypesRequest) (*GetObjectTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectTypes not implemented")
}
func (UnimplementedDirectoryServer) GetRelationType(context.Context, *GetRelationTypeRequest) (*GetRelationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelationType not implemented")
}
func (UnimplementedDirectoryServer) GetRelationTypes(context.Context, *GetRelationTypesRequest) (*GetRelationTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelationTypes not implemented")
}
func (UnimplementedDirectoryServer) GetPermission(context.Context, *GetPermissionRequest) (*GetPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermission not implemented")
}
func (UnimplementedDirectoryServer) GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedDirectoryServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedDirectoryServer) GetObjectMany(context.Context, *GetObjectManyRequest) (*GetObjectManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectMany not implemented")
}
func (UnimplementedDirectoryServer) GetObjects(context.Context, *GetObjectsRequest) (*GetObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjects not implemented")
}
func (UnimplementedDirectoryServer) GetRelation(context.Context, *GetRelationRequest) (*GetRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelation not implemented")
}
func (UnimplementedDirectoryServer) GetRelations(context.Context, *GetRelationsRequest) (*GetRelationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelations not implemented")
}
func (UnimplementedDirectoryServer) CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedDirectoryServer) CheckRelation(context.Context, *CheckRelationRequest) (*CheckRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRelation not implemented")
}
func (UnimplementedDirectoryServer) GetGraph(context.Context, *GetGraphRequest) (*GetGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraph not implemented")
}

// UnsafeDirectoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryServer will
// result in compilation errors.
type UnsafeDirectoryServer interface {
	mustEmbedUnimplementedDirectoryServer()
}

func RegisterDirectoryServer(s grpc.ServiceRegistrar, srv DirectoryServer) {
	s.RegisterService(&Directory_ServiceDesc, srv)
}

func _Directory_GetObjectType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetObjectType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetObjectType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetObjectType(ctx, req.(*GetObjectTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetObjectTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetObjectTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetObjectTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetObjectTypes(ctx, req.(*GetObjectTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetRelationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetRelationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetRelationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetRelationType(ctx, req.(*GetRelationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetRelationTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetRelationTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetRelationTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetRelationTypes(ctx, req.(*GetRelationTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetPermission(ctx, req.(*GetPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetPermissions(ctx, req.(*GetPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetObjectMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetObjectMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetObjectMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetObjectMany(ctx, req.(*GetObjectManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetObjects(ctx, req.(*GetObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetRelation(ctx, req.(*GetRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetRelations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetRelations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetRelations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetRelations(ctx, req.(*GetRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_CheckRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).CheckRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/CheckRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).CheckRelation(ctx, req.(*CheckRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.directory.v2.Directory/GetGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetGraph(ctx, req.(*GetGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Directory_ServiceDesc is the grpc.ServiceDesc for Directory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Directory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.directory.v2.Directory",
	HandlerType: (*DirectoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetObjectType",
			Handler:    _Directory_GetObjectType_Handler,
		},
		{
			MethodName: "GetObjectTypes",
			Handler:    _Directory_GetObjectTypes_Handler,
		},
		{
			MethodName: "GetRelationType",
			Handler:    _Directory_GetRelationType_Handler,
		},
		{
			MethodName: "GetRelationTypes",
			Handler:    _Directory_GetRelationTypes_Handler,
		},
		{
			MethodName: "GetPermission",
			Handler:    _Directory_GetPermission_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _Directory_GetPermissions_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _Directory_GetObject_Handler,
		},
		{
			MethodName: "GetObjectMany",
			Handler:    _Directory_GetObjectMany_Handler,
		},
		{
			MethodName: "GetObjects",
			Handler:    _Directory_GetObjects_Handler,
		},
		{
			MethodName: "GetRelation",
			Handler:    _Directory_GetRelation_Handler,
		},
		{
			MethodName: "GetRelations",
			Handler:    _Directory_GetRelations_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _Directory_CheckPermission_Handler,
		},
		{
			MethodName: "CheckRelation",
			Handler:    _Directory_CheckRelation_Handler,
		},
		{
			MethodName: "GetGraph",
			Handler:    _Directory_GetGraph_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/directory/v2/directory.proto",
}
