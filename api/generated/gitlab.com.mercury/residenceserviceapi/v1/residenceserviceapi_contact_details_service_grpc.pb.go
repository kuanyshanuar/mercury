// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: residenceserviceapi_contact_details_service.proto

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

// ContactDetailsServiceClient is the client API for ContactDetailsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactDetailsServiceClient interface {
	CreateContactDetails(ctx context.Context, in *CreateContactDetailsRequest, opts ...grpc.CallOption) (*CreateContactDetailsResponse, error)
	CreateResidenceContactDetails(ctx context.Context, in *CreateResidenceContactDetailsRequest, opts ...grpc.CallOption) (*CreateResidenceContactDetailsResponse, error)
	// *
	// Returns a list cities.
	ListContactDetails(ctx context.Context, in *ListContactDetailsRequest, opts ...grpc.CallOption) (*ListContactDetailsResponse, error)
	// *
	// Returns a list district by city id.
	//
	// If the city does not exist, an error "Not found" will be returned.
	ListResidenceContactDetails(ctx context.Context, in *ListResidenceContactDetailsRequest, opts ...grpc.CallOption) (*ListResidenceContactDetailsResponse, error)
}

type contactDetailsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactDetailsServiceClient(cc grpc.ClientConnInterface) ContactDetailsServiceClient {
	return &contactDetailsServiceClient{cc}
}

func (c *contactDetailsServiceClient) CreateContactDetails(ctx context.Context, in *CreateContactDetailsRequest, opts ...grpc.CallOption) (*CreateContactDetailsResponse, error) {
	out := new(CreateContactDetailsResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService/CreateContactDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactDetailsServiceClient) CreateResidenceContactDetails(ctx context.Context, in *CreateResidenceContactDetailsRequest, opts ...grpc.CallOption) (*CreateResidenceContactDetailsResponse, error) {
	out := new(CreateResidenceContactDetailsResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService/CreateResidenceContactDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactDetailsServiceClient) ListContactDetails(ctx context.Context, in *ListContactDetailsRequest, opts ...grpc.CallOption) (*ListContactDetailsResponse, error) {
	out := new(ListContactDetailsResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService/ListContactDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactDetailsServiceClient) ListResidenceContactDetails(ctx context.Context, in *ListResidenceContactDetailsRequest, opts ...grpc.CallOption) (*ListResidenceContactDetailsResponse, error) {
	out := new(ListResidenceContactDetailsResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService/ListResidenceContactDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactDetailsServiceServer is the server API for ContactDetailsService service.
// All implementations should embed UnimplementedContactDetailsServiceServer
// for forward compatibility
type ContactDetailsServiceServer interface {
	CreateContactDetails(context.Context, *CreateContactDetailsRequest) (*CreateContactDetailsResponse, error)
	CreateResidenceContactDetails(context.Context, *CreateResidenceContactDetailsRequest) (*CreateResidenceContactDetailsResponse, error)
	// *
	// Returns a list cities.
	ListContactDetails(context.Context, *ListContactDetailsRequest) (*ListContactDetailsResponse, error)
	// *
	// Returns a list district by city id.
	//
	// If the city does not exist, an error "Not found" will be returned.
	ListResidenceContactDetails(context.Context, *ListResidenceContactDetailsRequest) (*ListResidenceContactDetailsResponse, error)
}

// UnimplementedContactDetailsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedContactDetailsServiceServer struct {
}

func (UnimplementedContactDetailsServiceServer) CreateContactDetails(context.Context, *CreateContactDetailsRequest) (*CreateContactDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContactDetails not implemented")
}
func (UnimplementedContactDetailsServiceServer) CreateResidenceContactDetails(context.Context, *CreateResidenceContactDetailsRequest) (*CreateResidenceContactDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResidenceContactDetails not implemented")
}
func (UnimplementedContactDetailsServiceServer) ListContactDetails(context.Context, *ListContactDetailsRequest) (*ListContactDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContactDetails not implemented")
}
func (UnimplementedContactDetailsServiceServer) ListResidenceContactDetails(context.Context, *ListResidenceContactDetailsRequest) (*ListResidenceContactDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResidenceContactDetails not implemented")
}

// UnsafeContactDetailsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactDetailsServiceServer will
// result in compilation errors.
type UnsafeContactDetailsServiceServer interface {
	mustEmbedUnimplementedContactDetailsServiceServer()
}

func RegisterContactDetailsServiceServer(s grpc.ServiceRegistrar, srv ContactDetailsServiceServer) {
	s.RegisterService(&ContactDetailsService_ServiceDesc, srv)
}

func _ContactDetailsService_CreateContactDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactDetailsServiceServer).CreateContactDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService/CreateContactDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactDetailsServiceServer).CreateContactDetails(ctx, req.(*CreateContactDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactDetailsService_CreateResidenceContactDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResidenceContactDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactDetailsServiceServer).CreateResidenceContactDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService/CreateResidenceContactDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactDetailsServiceServer).CreateResidenceContactDetails(ctx, req.(*CreateResidenceContactDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactDetailsService_ListContactDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContactDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactDetailsServiceServer).ListContactDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService/ListContactDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactDetailsServiceServer).ListContactDetails(ctx, req.(*ListContactDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactDetailsService_ListResidenceContactDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResidenceContactDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactDetailsServiceServer).ListResidenceContactDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService/ListResidenceContactDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactDetailsServiceServer).ListResidenceContactDetails(ctx, req.(*ListResidenceContactDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactDetailsService_ServiceDesc is the grpc.ServiceDesc for ContactDetailsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactDetailsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.com.mercury.residenceservice.generated.model.v1.ContactDetailsService",
	HandlerType: (*ContactDetailsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContactDetails",
			Handler:    _ContactDetailsService_CreateContactDetails_Handler,
		},
		{
			MethodName: "CreateResidenceContactDetails",
			Handler:    _ContactDetailsService_CreateResidenceContactDetails_Handler,
		},
		{
			MethodName: "ListContactDetails",
			Handler:    _ContactDetailsService_ListContactDetails_Handler,
		},
		{
			MethodName: "ListResidenceContactDetails",
			Handler:    _ContactDetailsService_ListResidenceContactDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "residenceserviceapi_contact_details_service.proto",
}