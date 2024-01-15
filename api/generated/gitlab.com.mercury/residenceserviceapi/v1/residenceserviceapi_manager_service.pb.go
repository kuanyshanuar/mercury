// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: residenceserviceapi_manager_service.proto

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

type CreateManagerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manager *ManagerWrite `protobuf:"bytes,1,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (x *CreateManagerRequest) Reset() {
	*x = CreateManagerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManagerRequest) ProtoMessage() {}

func (x *CreateManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManagerRequest.ProtoReflect.Descriptor instead.
func (*CreateManagerRequest) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateManagerRequest) GetManager() *ManagerWrite {
	if x != nil {
		return x.Manager
	}
	return nil
}

type CreateManagerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateManagerResponse) Reset() {
	*x = CreateManagerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManagerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManagerResponse) ProtoMessage() {}

func (x *CreateManagerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManagerResponse.ProtoReflect.Descriptor instead.
func (*CreateManagerResponse) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{1}
}

type ListManagersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Criteria *ManagerSearchCriteria `protobuf:"bytes,1,opt,name=criteria,proto3" json:"criteria,omitempty"`
}

func (x *ListManagersRequest) Reset() {
	*x = ListManagersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListManagersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListManagersRequest) ProtoMessage() {}

func (x *ListManagersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListManagersRequest.ProtoReflect.Descriptor instead.
func (*ListManagersRequest) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListManagersRequest) GetCriteria() *ManagerSearchCriteria {
	if x != nil {
		return x.Criteria
	}
	return nil
}

type ListManagersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Managers []*ManagerRead `protobuf:"bytes,1,rep,name=managers,proto3" json:"managers,omitempty"`
	Total    int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListManagersResponse) Reset() {
	*x = ListManagersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListManagersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListManagersResponse) ProtoMessage() {}

func (x *ListManagersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListManagersResponse.ProtoReflect.Descriptor instead.
func (*ListManagersResponse) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListManagersResponse) GetManagers() []*ManagerRead {
	if x != nil {
		return x.Managers
	}
	return nil
}

func (x *ListManagersResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetManagerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetManagerRequest) Reset() {
	*x = GetManagerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManagerRequest) ProtoMessage() {}

func (x *GetManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManagerRequest.ProtoReflect.Descriptor instead.
func (*GetManagerRequest) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetManagerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetManagerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manager *ManagerRead `protobuf:"bytes,1,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (x *GetManagerResponse) Reset() {
	*x = GetManagerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManagerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManagerResponse) ProtoMessage() {}

func (x *GetManagerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManagerResponse.ProtoReflect.Descriptor instead.
func (*GetManagerResponse) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetManagerResponse) GetManager() *ManagerRead {
	if x != nil {
		return x.Manager
	}
	return nil
}

type UpdateManagerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Manager *ManagerWrite `protobuf:"bytes,2,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (x *UpdateManagerRequest) Reset() {
	*x = UpdateManagerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateManagerRequest) ProtoMessage() {}

func (x *UpdateManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateManagerRequest.ProtoReflect.Descriptor instead.
func (*UpdateManagerRequest) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateManagerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateManagerRequest) GetManager() *ManagerWrite {
	if x != nil {
		return x.Manager
	}
	return nil
}

type UpdateManagerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateManagerResponse) Reset() {
	*x = UpdateManagerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateManagerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateManagerResponse) ProtoMessage() {}

func (x *UpdateManagerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateManagerResponse.ProtoReflect.Descriptor instead.
func (*UpdateManagerResponse) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{7}
}

type DeleteManagerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteManagerRequest) Reset() {
	*x = DeleteManagerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteManagerRequest) ProtoMessage() {}

func (x *DeleteManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteManagerRequest.ProtoReflect.Descriptor instead.
func (*DeleteManagerRequest) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteManagerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteManagerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteManagerResponse) Reset() {
	*x = DeleteManagerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteManagerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteManagerResponse) ProtoMessage() {}

func (x *DeleteManagerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteManagerResponse.ProtoReflect.Descriptor instead.
func (*DeleteManagerResponse) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_service_proto_rawDescGZIP(), []int{9}
}

var File_residenceserviceapi_manager_service_proto protoreflect.FileDescriptor

var file_residenceserviceapi_manager_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x1a, 0x27, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x80, 0x01,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x22, 0x8d, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x08, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x07, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x5e, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf9, 0x06,
	0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xae, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x4c, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x4d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0xab, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x12, 0x4b, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x4c, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0xa5, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x49,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xae, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4c, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xae, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4c, 0x2e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2f,
	0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_residenceserviceapi_manager_service_proto_rawDescOnce sync.Once
	file_residenceserviceapi_manager_service_proto_rawDescData = file_residenceserviceapi_manager_service_proto_rawDesc
)

func file_residenceserviceapi_manager_service_proto_rawDescGZIP() []byte {
	file_residenceserviceapi_manager_service_proto_rawDescOnce.Do(func() {
		file_residenceserviceapi_manager_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_residenceserviceapi_manager_service_proto_rawDescData)
	})
	return file_residenceserviceapi_manager_service_proto_rawDescData
}

var file_residenceserviceapi_manager_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_residenceserviceapi_manager_service_proto_goTypes = []interface{}{
	(*CreateManagerRequest)(nil),  // 0: gitlab.com.mercury.residenceservice.generated.model.v1.CreateManagerRequest
	(*CreateManagerResponse)(nil), // 1: gitlab.com.mercury.residenceservice.generated.model.v1.CreateManagerResponse
	(*ListManagersRequest)(nil),   // 2: gitlab.com.mercury.residenceservice.generated.model.v1.ListManagersRequest
	(*ListManagersResponse)(nil),  // 3: gitlab.com.mercury.residenceservice.generated.model.v1.ListManagersResponse
	(*GetManagerRequest)(nil),     // 4: gitlab.com.mercury.residenceservice.generated.model.v1.GetManagerRequest
	(*GetManagerResponse)(nil),    // 5: gitlab.com.mercury.residenceservice.generated.model.v1.GetManagerResponse
	(*UpdateManagerRequest)(nil),  // 6: gitlab.com.mercury.residenceservice.generated.model.v1.UpdateManagerRequest
	(*UpdateManagerResponse)(nil), // 7: gitlab.com.mercury.residenceservice.generated.model.v1.UpdateManagerResponse
	(*DeleteManagerRequest)(nil),  // 8: gitlab.com.mercury.residenceservice.generated.model.v1.DeleteManagerRequest
	(*DeleteManagerResponse)(nil), // 9: gitlab.com.mercury.residenceservice.generated.model.v1.DeleteManagerResponse
	(*ManagerWrite)(nil),          // 10: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerWrite
	(*ManagerSearchCriteria)(nil), // 11: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerSearchCriteria
	(*ManagerRead)(nil),           // 12: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerRead
}
var file_residenceserviceapi_manager_service_proto_depIdxs = []int32{
	10, // 0: gitlab.com.mercury.residenceservice.generated.model.v1.CreateManagerRequest.manager:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.ManagerWrite
	11, // 1: gitlab.com.mercury.residenceservice.generated.model.v1.ListManagersRequest.criteria:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.ManagerSearchCriteria
	12, // 2: gitlab.com.mercury.residenceservice.generated.model.v1.ListManagersResponse.managers:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.ManagerRead
	12, // 3: gitlab.com.mercury.residenceservice.generated.model.v1.GetManagerResponse.manager:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.ManagerRead
	10, // 4: gitlab.com.mercury.residenceservice.generated.model.v1.UpdateManagerRequest.manager:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.ManagerWrite
	0,  // 5: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.CreateManager:input_type -> gitlab.com.mercury.residenceservice.generated.model.v1.CreateManagerRequest
	2,  // 6: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.ListManagers:input_type -> gitlab.com.mercury.residenceservice.generated.model.v1.ListManagersRequest
	4,  // 7: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.GetManager:input_type -> gitlab.com.mercury.residenceservice.generated.model.v1.GetManagerRequest
	6,  // 8: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.UpdateManager:input_type -> gitlab.com.mercury.residenceservice.generated.model.v1.UpdateManagerRequest
	8,  // 9: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.DeleteManager:input_type -> gitlab.com.mercury.residenceservice.generated.model.v1.DeleteManagerRequest
	1,  // 10: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.CreateManager:output_type -> gitlab.com.mercury.residenceservice.generated.model.v1.CreateManagerResponse
	3,  // 11: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.ListManagers:output_type -> gitlab.com.mercury.residenceservice.generated.model.v1.ListManagersResponse
	5,  // 12: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.GetManager:output_type -> gitlab.com.mercury.residenceservice.generated.model.v1.GetManagerResponse
	7,  // 13: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.UpdateManager:output_type -> gitlab.com.mercury.residenceservice.generated.model.v1.UpdateManagerResponse
	9,  // 14: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerService.DeleteManager:output_type -> gitlab.com.mercury.residenceservice.generated.model.v1.DeleteManagerResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_residenceserviceapi_manager_service_proto_init() }
func file_residenceserviceapi_manager_service_proto_init() {
	if File_residenceserviceapi_manager_service_proto != nil {
		return
	}
	file_residenceserviceapi_manager_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_residenceserviceapi_manager_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManagerRequest); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManagerResponse); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListManagersRequest); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListManagersResponse); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManagerRequest); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManagerResponse); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateManagerRequest); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateManagerResponse); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteManagerRequest); i {
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
		file_residenceserviceapi_manager_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteManagerResponse); i {
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
			RawDescriptor: file_residenceserviceapi_manager_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_residenceserviceapi_manager_service_proto_goTypes,
		DependencyIndexes: file_residenceserviceapi_manager_service_proto_depIdxs,
		MessageInfos:      file_residenceserviceapi_manager_service_proto_msgTypes,
	}.Build()
	File_residenceserviceapi_manager_service_proto = out.File
	file_residenceserviceapi_manager_service_proto_rawDesc = nil
	file_residenceserviceapi_manager_service_proto_goTypes = nil
	file_residenceserviceapi_manager_service_proto_depIdxs = nil
}