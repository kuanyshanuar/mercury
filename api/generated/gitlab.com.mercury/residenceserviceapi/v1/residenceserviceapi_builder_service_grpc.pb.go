// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: residenceserviceapi_builder_service.proto

package v1

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

// BuilderServiceClient is the client API for BuilderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuilderServiceClient interface {
	// Create a new builder.
	CreateBuilder(ctx context.Context, in *CreateBuilderRequest, opts ...grpc.CallOption) (*CreateBuilderResponse, error)
	// Returns a list builder by criteria.
	//
	// If the builder does not exist, an error "Not found" will be returned.
	ListBuilders(ctx context.Context, in *ListBuildersRequest, opts ...grpc.CallOption) (*ListBuildersResponse, error)
	// Returns a builder by id.
	//
	// If the builder does not exist, an error "Not found" will be returned.
	GetBuilder(ctx context.Context, in *GetBuilderRequest, opts ...grpc.CallOption) (*GetBuilderResponse, error)
	// Updates a builder by id.
	//
	// If the builder does not exist, an error "Not found" will be returned.
	UpdateBuilder(ctx context.Context, in *UpdateBuilderRequest, opts ...grpc.CallOption) (*UpdateBuilderResponse, error)
	// Deletes a builder by id.
	//
	// If the builder does not exist, an error "Not found" will be returned.
	DeleteBuilder(ctx context.Context, in *DeleteBuilderRequest, opts ...grpc.CallOption) (*DeleteBuilderResponse, error)
}

type builderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuilderServiceClient(cc grpc.ClientConnInterface) BuilderServiceClient {
	return &builderServiceClient{cc}
}

func (c *builderServiceClient) CreateBuilder(ctx context.Context, in *CreateBuilderRequest, opts ...grpc.CallOption) (*CreateBuilderResponse, error) {
	out := new(CreateBuilderResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/CreateBuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderServiceClient) ListBuilders(ctx context.Context, in *ListBuildersRequest, opts ...grpc.CallOption) (*ListBuildersResponse, error) {
	out := new(ListBuildersResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/ListBuilders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderServiceClient) GetBuilder(ctx context.Context, in *GetBuilderRequest, opts ...grpc.CallOption) (*GetBuilderResponse, error) {
	out := new(GetBuilderResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/GetBuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderServiceClient) UpdateBuilder(ctx context.Context, in *UpdateBuilderRequest, opts ...grpc.CallOption) (*UpdateBuilderResponse, error) {
	out := new(UpdateBuilderResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/UpdateBuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderServiceClient) DeleteBuilder(ctx context.Context, in *DeleteBuilderRequest, opts ...grpc.CallOption) (*DeleteBuilderResponse, error) {
	out := new(DeleteBuilderResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/DeleteBuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuilderServiceServer is the server API for BuilderService service.
// All implementations should embed UnimplementedBuilderServiceServer
// for forward compatibility
type BuilderServiceServer interface {
	// Create a new builder.
	CreateBuilder(context.Context, *CreateBuilderRequest) (*CreateBuilderResponse, error)
	// Returns a list builder by criteria.
	//
	// If the builder does not exist, an error "Not found" will be returned.
	ListBuilders(context.Context, *ListBuildersRequest) (*ListBuildersResponse, error)
	// Returns a builder by id.
	//
	// If the builder does not exist, an error "Not found" will be returned.
	GetBuilder(context.Context, *GetBuilderRequest) (*GetBuilderResponse, error)
	// Updates a builder by id.
	//
	// If the builder does not exist, an error "Not found" will be returned.
	UpdateBuilder(context.Context, *UpdateBuilderRequest) (*UpdateBuilderResponse, error)
	// Deletes a builder by id.
	//
	// If the builder does not exist, an error "Not found" will be returned.
	DeleteBuilder(context.Context, *DeleteBuilderRequest) (*DeleteBuilderResponse, error)
}

// UnimplementedBuilderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBuilderServiceServer struct {
}

func (UnimplementedBuilderServiceServer) CreateBuilder(context.Context, *CreateBuilderRequest) (*CreateBuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBuilder not implemented")
}
func (UnimplementedBuilderServiceServer) ListBuilders(context.Context, *ListBuildersRequest) (*ListBuildersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuilders not implemented")
}
func (UnimplementedBuilderServiceServer) GetBuilder(context.Context, *GetBuilderRequest) (*GetBuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuilder not implemented")
}
func (UnimplementedBuilderServiceServer) UpdateBuilder(context.Context, *UpdateBuilderRequest) (*UpdateBuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBuilder not implemented")
}
func (UnimplementedBuilderServiceServer) DeleteBuilder(context.Context, *DeleteBuilderRequest) (*DeleteBuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBuilder not implemented")
}

// UnsafeBuilderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuilderServiceServer will
// result in compilation errors.
type UnsafeBuilderServiceServer interface {
	mustEmbedUnimplementedBuilderServiceServer()
}

func RegisterBuilderServiceServer(s grpc.ServiceRegistrar, srv BuilderServiceServer) {
	s.RegisterService(&BuilderService_ServiceDesc, srv)
}

func _BuilderService_CreateBuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).CreateBuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/CreateBuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).CreateBuilder(ctx, req.(*CreateBuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderService_ListBuilders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).ListBuilders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/ListBuilders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).ListBuilders(ctx, req.(*ListBuildersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderService_GetBuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).GetBuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/GetBuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).GetBuilder(ctx, req.(*GetBuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderService_UpdateBuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).UpdateBuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/UpdateBuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).UpdateBuilder(ctx, req.(*UpdateBuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderService_DeleteBuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).DeleteBuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService/DeleteBuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).DeleteBuilder(ctx, req.(*DeleteBuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BuilderService_ServiceDesc is the grpc.ServiceDesc for BuilderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuilderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.com.mercury.residenceservice.generated.model.v1.BuilderService",
	HandlerType: (*BuilderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBuilder",
			Handler:    _BuilderService_CreateBuilder_Handler,
		},
		{
			MethodName: "ListBuilders",
			Handler:    _BuilderService_ListBuilders_Handler,
		},
		{
			MethodName: "GetBuilder",
			Handler:    _BuilderService_GetBuilder_Handler,
		},
		{
			MethodName: "UpdateBuilder",
			Handler:    _BuilderService_UpdateBuilder_Handler,
		},
		{
			MethodName: "DeleteBuilder",
			Handler:    _BuilderService_DeleteBuilder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "residenceserviceapi_builder_service.proto",
}
