// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: identityserviceapi_permission_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPermissionsRequest) Reset() {
	*x = ListPermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identityserviceapi_permission_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsRequest) ProtoMessage() {}

func (x *ListPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identityserviceapi_permission_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsRequest.ProtoReflect.Descriptor instead.
func (*ListPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_identityserviceapi_permission_service_proto_rawDescGZIP(), []int{0}
}

type ListPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPermissionsResponse) Reset() {
	*x = ListPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identityserviceapi_permission_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsResponse) ProtoMessage() {}

func (x *ListPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identityserviceapi_permission_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsResponse.ProtoReflect.Descriptor instead.
func (*ListPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_identityserviceapi_permission_service_proto_rawDescGZIP(), []int{1}
}

type AllowPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PermissionKey string `protobuf:"bytes,1,opt,name=permission_key,json=permissionKey,proto3" json:"permission_key,omitempty"`
	UserId        int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleId        int64  `protobuf:"varint,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *AllowPermissionRequest) Reset() {
	*x = AllowPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identityserviceapi_permission_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowPermissionRequest) ProtoMessage() {}

func (x *AllowPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identityserviceapi_permission_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowPermissionRequest.ProtoReflect.Descriptor instead.
func (*AllowPermissionRequest) Descriptor() ([]byte, []int) {
	return file_identityserviceapi_permission_service_proto_rawDescGZIP(), []int{2}
}

func (x *AllowPermissionRequest) GetPermissionKey() string {
	if x != nil {
		return x.PermissionKey
	}
	return ""
}

func (x *AllowPermissionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AllowPermissionRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type AllowPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAllowed bool `protobuf:"varint,1,opt,name=is_allowed,json=isAllowed,proto3" json:"is_allowed,omitempty"`
}

func (x *AllowPermissionResponse) Reset() {
	*x = AllowPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identityserviceapi_permission_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowPermissionResponse) ProtoMessage() {}

func (x *AllowPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identityserviceapi_permission_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowPermissionResponse.ProtoReflect.Descriptor instead.
func (*AllowPermissionResponse) Descriptor() ([]byte, []int) {
	return file_identityserviceapi_permission_service_proto_rawDescGZIP(), []int{3}
}

func (x *AllowPermissionResponse) GetIsAllowed() bool {
	if x != nil {
		return x.IsAllowed
	}
	return false
}

var File_identityserviceapi_permission_service_proto protoreflect.FileDescriptor

var file_identityserviceapi_permission_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x19,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x71, 0x0a, 0x16, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x17,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x32, 0xe4, 0x02, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa5, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa6, 0x01, 0x0a, 0x05, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x4d,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x2e,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75,
	0x72, 0x79, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_identityserviceapi_permission_service_proto_rawDescOnce sync.Once
	file_identityserviceapi_permission_service_proto_rawDescData = file_identityserviceapi_permission_service_proto_rawDesc
)

func file_identityserviceapi_permission_service_proto_rawDescGZIP() []byte {
	file_identityserviceapi_permission_service_proto_rawDescOnce.Do(func() {
		file_identityserviceapi_permission_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_identityserviceapi_permission_service_proto_rawDescData)
	})
	return file_identityserviceapi_permission_service_proto_rawDescData
}

var file_identityserviceapi_permission_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_identityserviceapi_permission_service_proto_goTypes = []interface{}{
	(*ListPermissionsRequest)(nil),  // 0: gitlab.com.mercury.identityservice.generated.model.v1.ListPermissionsRequest
	(*ListPermissionsResponse)(nil), // 1: gitlab.com.mercury.identityservice.generated.model.v1.ListPermissionsResponse
	(*AllowPermissionRequest)(nil),  // 2: gitlab.com.mercury.identityservice.generated.model.v1.AllowPermissionRequest
	(*AllowPermissionResponse)(nil), // 3: gitlab.com.mercury.identityservice.generated.model.v1.AllowPermissionResponse
}
var file_identityserviceapi_permission_service_proto_depIdxs = []int32{
	0, // 0: gitlab.com.mercury.identityservice.generated.model.v1.PermissionService.List:input_type -> gitlab.com.mercury.identityservice.generated.model.v1.ListPermissionsRequest
	2, // 1: gitlab.com.mercury.identityservice.generated.model.v1.PermissionService.Allow:input_type -> gitlab.com.mercury.identityservice.generated.model.v1.AllowPermissionRequest
	1, // 2: gitlab.com.mercury.identityservice.generated.model.v1.PermissionService.List:output_type -> gitlab.com.mercury.identityservice.generated.model.v1.ListPermissionsResponse
	3, // 3: gitlab.com.mercury.identityservice.generated.model.v1.PermissionService.Allow:output_type -> gitlab.com.mercury.identityservice.generated.model.v1.AllowPermissionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_identityserviceapi_permission_service_proto_init() }
func file_identityserviceapi_permission_service_proto_init() {
	if File_identityserviceapi_permission_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identityserviceapi_permission_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identityserviceapi_permission_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identityserviceapi_permission_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowPermissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identityserviceapi_permission_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowPermissionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_identityserviceapi_permission_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identityserviceapi_permission_service_proto_goTypes,
		DependencyIndexes: file_identityserviceapi_permission_service_proto_depIdxs,
		MessageInfos:      file_identityserviceapi_permission_service_proto_msgTypes,
	}.Build()
	File_identityserviceapi_permission_service_proto = out.File
	file_identityserviceapi_permission_service_proto_rawDesc = nil
	file_identityserviceapi_permission_service_proto_goTypes = nil
	file_identityserviceapi_permission_service_proto_depIdxs = nil
}
