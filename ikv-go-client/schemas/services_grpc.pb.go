// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: services.proto

package schemas

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	InlineKVWriteService_HelloWorld_FullMethodName             = "/ikvschemas.InlineKVWriteService/helloWorld"
	InlineKVWriteService_UpsertFieldValues_FullMethodName      = "/ikvschemas.InlineKVWriteService/upsertFieldValues"
	InlineKVWriteService_BatchUpsertFieldValues_FullMethodName = "/ikvschemas.InlineKVWriteService/batchUpsertFieldValues"
	InlineKVWriteService_DeleteFieldValues_FullMethodName      = "/ikvschemas.InlineKVWriteService/deleteFieldValues"
	InlineKVWriteService_BatchDeleteFieldValues_FullMethodName = "/ikvschemas.InlineKVWriteService/batchDeleteFieldValues"
	InlineKVWriteService_DeleteDocument_FullMethodName         = "/ikvschemas.InlineKVWriteService/deleteDocument"
	InlineKVWriteService_BatchDeleteDocuments_FullMethodName   = "/ikvschemas.InlineKVWriteService/batchDeleteDocuments"
	InlineKVWriteService_GetUserStoreConfig_FullMethodName     = "/ikvschemas.InlineKVWriteService/getUserStoreConfig"
)

// InlineKVWriteServiceClient is the client API for InlineKVWriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InlineKVWriteServiceClient interface {
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
	// Write methods
	UpsertFieldValues(ctx context.Context, in *UpsertFieldValuesRequest, opts ...grpc.CallOption) (*Status, error)
	BatchUpsertFieldValues(ctx context.Context, in *BatchUpsertFieldValuesRequest, opts ...grpc.CallOption) (*Status, error)
	DeleteFieldValues(ctx context.Context, in *DeleteFieldValueRequest, opts ...grpc.CallOption) (*Status, error)
	BatchDeleteFieldValues(ctx context.Context, in *BatchDeleteFieldValuesRequest, opts ...grpc.CallOption) (*Status, error)
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*Status, error)
	BatchDeleteDocuments(ctx context.Context, in *BatchDeleteDocumentsRequest, opts ...grpc.CallOption) (*Status, error)
	// Gateway-specified configuration
	GetUserStoreConfig(ctx context.Context, in *GetUserStoreConfigRequest, opts ...grpc.CallOption) (*GetUserStoreConfigResponse, error)
}

type inlineKVWriteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInlineKVWriteServiceClient(cc grpc.ClientConnInterface) InlineKVWriteServiceClient {
	return &inlineKVWriteServiceClient{cc}
}

func (c *inlineKVWriteServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, InlineKVWriteService_HelloWorld_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inlineKVWriteServiceClient) UpsertFieldValues(ctx context.Context, in *UpsertFieldValuesRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, InlineKVWriteService_UpsertFieldValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inlineKVWriteServiceClient) BatchUpsertFieldValues(ctx context.Context, in *BatchUpsertFieldValuesRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, InlineKVWriteService_BatchUpsertFieldValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inlineKVWriteServiceClient) DeleteFieldValues(ctx context.Context, in *DeleteFieldValueRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, InlineKVWriteService_DeleteFieldValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inlineKVWriteServiceClient) BatchDeleteFieldValues(ctx context.Context, in *BatchDeleteFieldValuesRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, InlineKVWriteService_BatchDeleteFieldValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inlineKVWriteServiceClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, InlineKVWriteService_DeleteDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inlineKVWriteServiceClient) BatchDeleteDocuments(ctx context.Context, in *BatchDeleteDocumentsRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, InlineKVWriteService_BatchDeleteDocuments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inlineKVWriteServiceClient) GetUserStoreConfig(ctx context.Context, in *GetUserStoreConfigRequest, opts ...grpc.CallOption) (*GetUserStoreConfigResponse, error) {
	out := new(GetUserStoreConfigResponse)
	err := c.cc.Invoke(ctx, InlineKVWriteService_GetUserStoreConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InlineKVWriteServiceServer is the server API for InlineKVWriteService service.
// All implementations must embed UnimplementedInlineKVWriteServiceServer
// for forward compatibility
type InlineKVWriteServiceServer interface {
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	// Write methods
	UpsertFieldValues(context.Context, *UpsertFieldValuesRequest) (*Status, error)
	BatchUpsertFieldValues(context.Context, *BatchUpsertFieldValuesRequest) (*Status, error)
	DeleteFieldValues(context.Context, *DeleteFieldValueRequest) (*Status, error)
	BatchDeleteFieldValues(context.Context, *BatchDeleteFieldValuesRequest) (*Status, error)
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*Status, error)
	BatchDeleteDocuments(context.Context, *BatchDeleteDocumentsRequest) (*Status, error)
	// Gateway-specified configuration
	GetUserStoreConfig(context.Context, *GetUserStoreConfigRequest) (*GetUserStoreConfigResponse, error)
	mustEmbedUnimplementedInlineKVWriteServiceServer()
}

// UnimplementedInlineKVWriteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInlineKVWriteServiceServer struct {
}

func (UnimplementedInlineKVWriteServiceServer) HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedInlineKVWriteServiceServer) UpsertFieldValues(context.Context, *UpsertFieldValuesRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertFieldValues not implemented")
}
func (UnimplementedInlineKVWriteServiceServer) BatchUpsertFieldValues(context.Context, *BatchUpsertFieldValuesRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpsertFieldValues not implemented")
}
func (UnimplementedInlineKVWriteServiceServer) DeleteFieldValues(context.Context, *DeleteFieldValueRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFieldValues not implemented")
}
func (UnimplementedInlineKVWriteServiceServer) BatchDeleteFieldValues(context.Context, *BatchDeleteFieldValuesRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteFieldValues not implemented")
}
func (UnimplementedInlineKVWriteServiceServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedInlineKVWriteServiceServer) BatchDeleteDocuments(context.Context, *BatchDeleteDocumentsRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteDocuments not implemented")
}
func (UnimplementedInlineKVWriteServiceServer) GetUserStoreConfig(context.Context, *GetUserStoreConfigRequest) (*GetUserStoreConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStoreConfig not implemented")
}
func (UnimplementedInlineKVWriteServiceServer) mustEmbedUnimplementedInlineKVWriteServiceServer() {}

// UnsafeInlineKVWriteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InlineKVWriteServiceServer will
// result in compilation errors.
type UnsafeInlineKVWriteServiceServer interface {
	mustEmbedUnimplementedInlineKVWriteServiceServer()
}

func RegisterInlineKVWriteServiceServer(s grpc.ServiceRegistrar, srv InlineKVWriteServiceServer) {
	s.RegisterService(&InlineKVWriteService_ServiceDesc, srv)
}

func _InlineKVWriteService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InlineKVWriteServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InlineKVWriteService_HelloWorld_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InlineKVWriteServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InlineKVWriteService_UpsertFieldValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertFieldValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InlineKVWriteServiceServer).UpsertFieldValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InlineKVWriteService_UpsertFieldValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InlineKVWriteServiceServer).UpsertFieldValues(ctx, req.(*UpsertFieldValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InlineKVWriteService_BatchUpsertFieldValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpsertFieldValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InlineKVWriteServiceServer).BatchUpsertFieldValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InlineKVWriteService_BatchUpsertFieldValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InlineKVWriteServiceServer).BatchUpsertFieldValues(ctx, req.(*BatchUpsertFieldValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InlineKVWriteService_DeleteFieldValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFieldValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InlineKVWriteServiceServer).DeleteFieldValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InlineKVWriteService_DeleteFieldValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InlineKVWriteServiceServer).DeleteFieldValues(ctx, req.(*DeleteFieldValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InlineKVWriteService_BatchDeleteFieldValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteFieldValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InlineKVWriteServiceServer).BatchDeleteFieldValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InlineKVWriteService_BatchDeleteFieldValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InlineKVWriteServiceServer).BatchDeleteFieldValues(ctx, req.(*BatchDeleteFieldValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InlineKVWriteService_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InlineKVWriteServiceServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InlineKVWriteService_DeleteDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InlineKVWriteServiceServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InlineKVWriteService_BatchDeleteDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InlineKVWriteServiceServer).BatchDeleteDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InlineKVWriteService_BatchDeleteDocuments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InlineKVWriteServiceServer).BatchDeleteDocuments(ctx, req.(*BatchDeleteDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InlineKVWriteService_GetUserStoreConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStoreConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InlineKVWriteServiceServer).GetUserStoreConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InlineKVWriteService_GetUserStoreConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InlineKVWriteServiceServer).GetUserStoreConfig(ctx, req.(*GetUserStoreConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InlineKVWriteService_ServiceDesc is the grpc.ServiceDesc for InlineKVWriteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InlineKVWriteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ikvschemas.InlineKVWriteService",
	HandlerType: (*InlineKVWriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "helloWorld",
			Handler:    _InlineKVWriteService_HelloWorld_Handler,
		},
		{
			MethodName: "upsertFieldValues",
			Handler:    _InlineKVWriteService_UpsertFieldValues_Handler,
		},
		{
			MethodName: "batchUpsertFieldValues",
			Handler:    _InlineKVWriteService_BatchUpsertFieldValues_Handler,
		},
		{
			MethodName: "deleteFieldValues",
			Handler:    _InlineKVWriteService_DeleteFieldValues_Handler,
		},
		{
			MethodName: "batchDeleteFieldValues",
			Handler:    _InlineKVWriteService_BatchDeleteFieldValues_Handler,
		},
		{
			MethodName: "deleteDocument",
			Handler:    _InlineKVWriteService_DeleteDocument_Handler,
		},
		{
			MethodName: "batchDeleteDocuments",
			Handler:    _InlineKVWriteService_BatchDeleteDocuments_Handler,
		},
		{
			MethodName: "getUserStoreConfig",
			Handler:    _InlineKVWriteService_GetUserStoreConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

const (
	AdminService_HealthCheck_FullMethodName = "/ikvschemas.AdminService/healthCheck"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, AdminService_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ikvschemas.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "healthCheck",
			Handler:    _AdminService_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
