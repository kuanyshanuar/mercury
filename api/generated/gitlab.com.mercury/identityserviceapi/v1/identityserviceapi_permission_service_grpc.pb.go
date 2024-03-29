// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: identityserviceapi_permission_service.proto

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

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	// *
	// Returns a list of permission.
	List(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error)
	// *
	// Allows user.
	Allow(ctx context.Context, in *AllowPermissionRequest, opts ...grpc.CallOption) (*AllowPermissionResponse, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) List(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error) {
	out := new(ListPermissionsResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.identityservice.generated.model.v1.PermissionService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) Allow(ctx context.Context, in *AllowPermissionRequest, opts ...grpc.CallOption) (*AllowPermissionResponse, error) {
	out := new(AllowPermissionResponse)
	err := c.cc.Invoke(ctx, "/gitlab.com.mercury.identityservice.generated.model.v1.PermissionService/Allow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations should embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	// *
	// Returns a list of permission.
	List(context.Context, *ListPermissionsRequest) (*ListPermissionsResponse, error)
	// *
	// Allows user.
	Allow(context.Context, *AllowPermissionRequest) (*AllowPermissionResponse, error)
}

// UnimplementedPermissionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) List(context.Context, *ListPermissionsRequest) (*ListPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPermissionServiceServer) Allow(context.Context, *AllowPermissionRequest) (*AllowPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allow not implemented")
}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.identityservice.generated.model.v1.PermissionService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).List(ctx, req.(*ListPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_Allow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllowPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).Allow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitlab.com.mercury.identityservice.generated.model.v1.PermissionService/Allow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).Allow(ctx, req.(*AllowPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.com.mercury.identityservice.generated.model.v1.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PermissionService_List_Handler,
		},
		{
			MethodName: "Allow",
			Handler:    _PermissionService_Allow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identityserviceapi_permission_service.proto",
}
