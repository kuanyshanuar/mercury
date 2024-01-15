// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: residenceserviceapi_cottage_service.proto

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

// CottageServiceClient is the client API for CottageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CottageServiceClient interface {
	CreateCottage(ctx context.Context, in *CreateCottageRequest, opts ...grpc.CallOption) (*CreateCottageResponse, error)
	UpdateCottage(ctx context.Context, in *UpdateCottageRequest, opts ...grpc.CallOption) (*UpdateCottageResponse, error)
	DeleteCottage(ctx context.Context, in *DeleteCottageRequest, opts ...grpc.CallOption) (*DeleteCottageResponse, error)
	GetCottage(ctx context.Context, in *GetCottageRequest, opts ...grpc.CallOption) (*GetCottageResponse, error)
	ListCottages(ctx context.Context, in *ListCottagesRequest, opts ...grpc.CallOption) (*ListCottagesResponse, error)
	ListCottageByID(ctx context.Context, in *ListCottageByIDRequest, opts ...grpc.CallOption) (*ListCottageByIDResponse, error)
	ListPopularCottages(ctx context.Context, in *ListPopularCottagesRequest, opts ...grpc.CallOption) (*ListPopularCottagesResponse, error)
	CreateCottagePlan(ctx context.Context, in *CreateCottagePlanRequest, opts ...grpc.CallOption) (*CreateCottagePlanResponse, error)
	UpdateCottagePlan(ctx context.Context, in *UpdateCottagePlanRequest, opts ...grpc.CallOption) (*UpdateCottagePlanResponse, error)
	DeleteCottagePlan(ctx context.Context, in *DeleteCottagePlanRequest, opts ...grpc.CallOption) (*DeleteCottagePlanResponse, error)
}

type cottageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCottageServiceClient(cc grpc.ClientConnInterface) CottageServiceClient {
	return &cottageServiceClient{cc}
}

func (c *cottageServiceClient) CreateCottage(ctx context.Context, in *CreateCottageRequest, opts ...grpc.CallOption) (*CreateCottageResponse, error) {
	out := new(CreateCottageResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/CreateCottage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) UpdateCottage(ctx context.Context, in *UpdateCottageRequest, opts ...grpc.CallOption) (*UpdateCottageResponse, error) {
	out := new(UpdateCottageResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/UpdateCottage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) DeleteCottage(ctx context.Context, in *DeleteCottageRequest, opts ...grpc.CallOption) (*DeleteCottageResponse, error) {
	out := new(DeleteCottageResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/DeleteCottage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) GetCottage(ctx context.Context, in *GetCottageRequest, opts ...grpc.CallOption) (*GetCottageResponse, error) {
	out := new(GetCottageResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/GetCottage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) ListCottages(ctx context.Context, in *ListCottagesRequest, opts ...grpc.CallOption) (*ListCottagesResponse, error) {
	out := new(ListCottagesResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/ListCottages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) ListCottageByID(ctx context.Context, in *ListCottageByIDRequest, opts ...grpc.CallOption) (*ListCottageByIDResponse, error) {
	out := new(ListCottageByIDResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/ListCottageByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) ListPopularCottages(ctx context.Context, in *ListPopularCottagesRequest, opts ...grpc.CallOption) (*ListPopularCottagesResponse, error) {
	out := new(ListPopularCottagesResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/ListPopularCottages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) CreateCottagePlan(ctx context.Context, in *CreateCottagePlanRequest, opts ...grpc.CallOption) (*CreateCottagePlanResponse, error) {
	out := new(CreateCottagePlanResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/CreateCottagePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) UpdateCottagePlan(ctx context.Context, in *UpdateCottagePlanRequest, opts ...grpc.CallOption) (*UpdateCottagePlanResponse, error) {
	out := new(UpdateCottagePlanResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/UpdateCottagePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cottageServiceClient) DeleteCottagePlan(ctx context.Context, in *DeleteCottagePlanRequest, opts ...grpc.CallOption) (*DeleteCottagePlanResponse, error) {
	out := new(DeleteCottagePlanResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/DeleteCottagePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CottageServiceServer is the server API for CottageService service.
// All implementations should embed UnimplementedCottageServiceServer
// for forward compatibility
type CottageServiceServer interface {
	CreateCottage(context.Context, *CreateCottageRequest) (*CreateCottageResponse, error)
	UpdateCottage(context.Context, *UpdateCottageRequest) (*UpdateCottageResponse, error)
	DeleteCottage(context.Context, *DeleteCottageRequest) (*DeleteCottageResponse, error)
	GetCottage(context.Context, *GetCottageRequest) (*GetCottageResponse, error)
	ListCottages(context.Context, *ListCottagesRequest) (*ListCottagesResponse, error)
	ListCottageByID(context.Context, *ListCottageByIDRequest) (*ListCottageByIDResponse, error)
	ListPopularCottages(context.Context, *ListPopularCottagesRequest) (*ListPopularCottagesResponse, error)
	CreateCottagePlan(context.Context, *CreateCottagePlanRequest) (*CreateCottagePlanResponse, error)
	UpdateCottagePlan(context.Context, *UpdateCottagePlanRequest) (*UpdateCottagePlanResponse, error)
	DeleteCottagePlan(context.Context, *DeleteCottagePlanRequest) (*DeleteCottagePlanResponse, error)
}

// UnimplementedCottageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCottageServiceServer struct {
}

func (UnimplementedCottageServiceServer) CreateCottage(context.Context, *CreateCottageRequest) (*CreateCottageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCottage not implemented")
}
func (UnimplementedCottageServiceServer) UpdateCottage(context.Context, *UpdateCottageRequest) (*UpdateCottageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCottage not implemented")
}
func (UnimplementedCottageServiceServer) DeleteCottage(context.Context, *DeleteCottageRequest) (*DeleteCottageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCottage not implemented")
}
func (UnimplementedCottageServiceServer) GetCottage(context.Context, *GetCottageRequest) (*GetCottageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCottage not implemented")
}
func (UnimplementedCottageServiceServer) ListCottages(context.Context, *ListCottagesRequest) (*ListCottagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCottages not implemented")
}
func (UnimplementedCottageServiceServer) ListCottageByID(context.Context, *ListCottageByIDRequest) (*ListCottageByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCottageByID not implemented")
}
func (UnimplementedCottageServiceServer) ListPopularCottages(context.Context, *ListPopularCottagesRequest) (*ListPopularCottagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPopularCottages not implemented")
}
func (UnimplementedCottageServiceServer) CreateCottagePlan(context.Context, *CreateCottagePlanRequest) (*CreateCottagePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCottagePlan not implemented")
}
func (UnimplementedCottageServiceServer) UpdateCottagePlan(context.Context, *UpdateCottagePlanRequest) (*UpdateCottagePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCottagePlan not implemented")
}
func (UnimplementedCottageServiceServer) DeleteCottagePlan(context.Context, *DeleteCottagePlanRequest) (*DeleteCottagePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCottagePlan not implemented")
}

// UnsafeCottageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CottageServiceServer will
// result in compilation errors.
type UnsafeCottageServiceServer interface {
	mustEmbedUnimplementedCottageServiceServer()
}

func RegisterCottageServiceServer(s grpc.ServiceRegistrar, srv CottageServiceServer) {
	s.RegisterService(&CottageService_ServiceDesc, srv)
}

func _CottageService_CreateCottage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCottageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).CreateCottage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/CreateCottage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).CreateCottage(ctx, req.(*CreateCottageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_UpdateCottage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCottageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).UpdateCottage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/UpdateCottage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).UpdateCottage(ctx, req.(*UpdateCottageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_DeleteCottage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCottageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).DeleteCottage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/DeleteCottage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).DeleteCottage(ctx, req.(*DeleteCottageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_GetCottage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCottageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).GetCottage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/GetCottage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).GetCottage(ctx, req.(*GetCottageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_ListCottages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCottagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).ListCottages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/ListCottages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).ListCottages(ctx, req.(*ListCottagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_ListCottageByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCottageByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).ListCottageByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/ListCottageByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).ListCottageByID(ctx, req.(*ListCottageByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_ListPopularCottages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPopularCottagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).ListPopularCottages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/ListPopularCottages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).ListPopularCottages(ctx, req.(*ListPopularCottagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_CreateCottagePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCottagePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).CreateCottagePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/CreateCottagePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).CreateCottagePlan(ctx, req.(*CreateCottagePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_UpdateCottagePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCottagePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).UpdateCottagePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/UpdateCottagePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).UpdateCottagePlan(ctx, req.(*UpdateCottagePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CottageService_DeleteCottagePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCottagePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CottageServiceServer).DeleteCottagePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.CottageService/DeleteCottagePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CottageServiceServer).DeleteCottagePlan(ctx, req.(*DeleteCottagePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CottageService_ServiceDesc is the grpc.ServiceDesc for CottageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CottageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.com.mercury.residenceservice.generated.model.v1.CottageService",
	HandlerType: (*CottageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCottage",
			Handler:    _CottageService_CreateCottage_Handler,
		},
		{
			MethodName: "UpdateCottage",
			Handler:    _CottageService_UpdateCottage_Handler,
		},
		{
			MethodName: "DeleteCottage",
			Handler:    _CottageService_DeleteCottage_Handler,
		},
		{
			MethodName: "GetCottage",
			Handler:    _CottageService_GetCottage_Handler,
		},
		{
			MethodName: "ListCottages",
			Handler:    _CottageService_ListCottages_Handler,
		},
		{
			MethodName: "ListCottageByID",
			Handler:    _CottageService_ListCottageByID_Handler,
		},
		{
			MethodName: "ListPopularCottages",
			Handler:    _CottageService_ListPopularCottages_Handler,
		},
		{
			MethodName: "CreateCottagePlan",
			Handler:    _CottageService_CreateCottagePlan_Handler,
		},
		{
			MethodName: "UpdateCottagePlan",
			Handler:    _CottageService_UpdateCottagePlan_Handler,
		},
		{
			MethodName: "DeleteCottagePlan",
			Handler:    _CottageService_DeleteCottagePlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "residenceserviceapi_cottage_service.proto",
}
