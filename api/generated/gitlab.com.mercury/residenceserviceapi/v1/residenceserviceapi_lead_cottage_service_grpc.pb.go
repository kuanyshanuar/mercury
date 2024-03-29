// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: residenceserviceapi_lead_cottage_service.proto

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

// LeadCottageServiceClient is the client API for LeadCottageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeadCottageServiceClient interface {
	CreateLeadCottage(ctx context.Context, in *CreateLeadCottageRequest, opts ...grpc.CallOption) (*CreateLeadCottageResponse, error)
	UpdateLeadCottage(ctx context.Context, in *UpdateLeadCottageRequest, opts ...grpc.CallOption) (*UpdateLeadCottageResponse, error)
	GetLeadCottage(ctx context.Context, in *GetLeadCottageRequest, opts ...grpc.CallOption) (*GetLeadCottageResponse, error)
	ListLeadCottages(ctx context.Context, in *ListLeadCottagesRequest, opts ...grpc.CallOption) (*ListLeadCottagesResponse, error)
	DeleteLeadCottage(ctx context.Context, in *DeleteLeadCottageRequest, opts ...grpc.CallOption) (*DeleteLeadCottageResponse, error)
}

type leadCottageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeadCottageServiceClient(cc grpc.ClientConnInterface) LeadCottageServiceClient {
	return &leadCottageServiceClient{cc}
}

func (c *leadCottageServiceClient) CreateLeadCottage(ctx context.Context, in *CreateLeadCottageRequest, opts ...grpc.CallOption) (*CreateLeadCottageResponse, error) {
	out := new(CreateLeadCottageResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/CreateLeadCottage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadCottageServiceClient) UpdateLeadCottage(ctx context.Context, in *UpdateLeadCottageRequest, opts ...grpc.CallOption) (*UpdateLeadCottageResponse, error) {
	out := new(UpdateLeadCottageResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/UpdateLeadCottage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadCottageServiceClient) GetLeadCottage(ctx context.Context, in *GetLeadCottageRequest, opts ...grpc.CallOption) (*GetLeadCottageResponse, error) {
	out := new(GetLeadCottageResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/GetLeadCottage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadCottageServiceClient) ListLeadCottages(ctx context.Context, in *ListLeadCottagesRequest, opts ...grpc.CallOption) (*ListLeadCottagesResponse, error) {
	out := new(ListLeadCottagesResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/ListLeadCottages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leadCottageServiceClient) DeleteLeadCottage(ctx context.Context, in *DeleteLeadCottageRequest, opts ...grpc.CallOption) (*DeleteLeadCottageResponse, error) {
	out := new(DeleteLeadCottageResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/DeleteLeadCottage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeadCottageServiceServer is the server API for LeadCottageService service.
// All implementations should embed UnimplementedLeadCottageServiceServer
// for forward compatibility
type LeadCottageServiceServer interface {
	CreateLeadCottage(context.Context, *CreateLeadCottageRequest) (*CreateLeadCottageResponse, error)
	UpdateLeadCottage(context.Context, *UpdateLeadCottageRequest) (*UpdateLeadCottageResponse, error)
	GetLeadCottage(context.Context, *GetLeadCottageRequest) (*GetLeadCottageResponse, error)
	ListLeadCottages(context.Context, *ListLeadCottagesRequest) (*ListLeadCottagesResponse, error)
	DeleteLeadCottage(context.Context, *DeleteLeadCottageRequest) (*DeleteLeadCottageResponse, error)
}

// UnimplementedLeadCottageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLeadCottageServiceServer struct {
}

func (UnimplementedLeadCottageServiceServer) CreateLeadCottage(context.Context, *CreateLeadCottageRequest) (*CreateLeadCottageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLeadCottage not implemented")
}
func (UnimplementedLeadCottageServiceServer) UpdateLeadCottage(context.Context, *UpdateLeadCottageRequest) (*UpdateLeadCottageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLeadCottage not implemented")
}
func (UnimplementedLeadCottageServiceServer) GetLeadCottage(context.Context, *GetLeadCottageRequest) (*GetLeadCottageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeadCottage not implemented")
}
func (UnimplementedLeadCottageServiceServer) ListLeadCottages(context.Context, *ListLeadCottagesRequest) (*ListLeadCottagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLeadCottages not implemented")
}
func (UnimplementedLeadCottageServiceServer) DeleteLeadCottage(context.Context, *DeleteLeadCottageRequest) (*DeleteLeadCottageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLeadCottage not implemented")
}

// UnsafeLeadCottageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeadCottageServiceServer will
// result in compilation errors.
type UnsafeLeadCottageServiceServer interface {
	mustEmbedUnimplementedLeadCottageServiceServer()
}

func RegisterLeadCottageServiceServer(s grpc.ServiceRegistrar, srv LeadCottageServiceServer) {
	s.RegisterService(&LeadCottageService_ServiceDesc, srv)
}

func _LeadCottageService_CreateLeadCottage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLeadCottageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadCottageServiceServer).CreateLeadCottage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/CreateLeadCottage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadCottageServiceServer).CreateLeadCottage(ctx, req.(*CreateLeadCottageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadCottageService_UpdateLeadCottage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLeadCottageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadCottageServiceServer).UpdateLeadCottage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/UpdateLeadCottage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadCottageServiceServer).UpdateLeadCottage(ctx, req.(*UpdateLeadCottageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadCottageService_GetLeadCottage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeadCottageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadCottageServiceServer).GetLeadCottage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/GetLeadCottage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadCottageServiceServer).GetLeadCottage(ctx, req.(*GetLeadCottageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadCottageService_ListLeadCottages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLeadCottagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadCottageServiceServer).ListLeadCottages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/ListLeadCottages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadCottageServiceServer).ListLeadCottages(ctx, req.(*ListLeadCottagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeadCottageService_DeleteLeadCottage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLeadCottageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadCottageServiceServer).DeleteLeadCottage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService/DeleteLeadCottage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadCottageServiceServer).DeleteLeadCottage(ctx, req.(*DeleteLeadCottageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LeadCottageService_ServiceDesc is the grpc.ServiceDesc for LeadCottageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeadCottageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageService",
	HandlerType: (*LeadCottageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLeadCottage",
			Handler:    _LeadCottageService_CreateLeadCottage_Handler,
		},
		{
			MethodName: "UpdateLeadCottage",
			Handler:    _LeadCottageService_UpdateLeadCottage_Handler,
		},
		{
			MethodName: "GetLeadCottage",
			Handler:    _LeadCottageService_GetLeadCottage_Handler,
		},
		{
			MethodName: "ListLeadCottages",
			Handler:    _LeadCottageService_ListLeadCottages_Handler,
		},
		{
			MethodName: "DeleteLeadCottage",
			Handler:    _LeadCottageService_DeleteLeadCottage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "residenceserviceapi_lead_cottage_service.proto",
}
