// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: residenceserviceapi_manager_model.proto

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

type ManagerWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// First name
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// Last name
	LastName string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// Email
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// City
	City string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	// Phone
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	// Image
	Image string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	// Password
	Password string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	// Is banned
	IsBanned bool `protobuf:"varint,8,opt,name=is_banned,json=isBanned,proto3" json:"is_banned,omitempty"`
}

func (x *ManagerWrite) Reset() {
	*x = ManagerWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerWrite) ProtoMessage() {}

func (x *ManagerWrite) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerWrite.ProtoReflect.Descriptor instead.
func (*ManagerWrite) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_model_proto_rawDescGZIP(), []int{0}
}

func (x *ManagerWrite) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ManagerWrite) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ManagerWrite) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ManagerWrite) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ManagerWrite) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ManagerWrite) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ManagerWrite) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ManagerWrite) GetIsBanned() bool {
	if x != nil {
		return x.IsBanned
	}
	return false
}

type ManagerRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID - id of the manager
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// First name
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// Last name
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// Email
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// Phone
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	// Image
	Image string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	// Is banned
	IsBanned bool `protobuf:"varint,7,opt,name=is_banned,json=isBanned,proto3" json:"is_banned,omitempty"`
}

func (x *ManagerRead) Reset() {
	*x = ManagerRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerRead) ProtoMessage() {}

func (x *ManagerRead) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerRead.ProtoReflect.Descriptor instead.
func (*ManagerRead) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_model_proto_rawDescGZIP(), []int{1}
}

func (x *ManagerRead) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ManagerRead) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ManagerRead) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ManagerRead) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ManagerRead) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ManagerRead) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ManagerRead) GetIsBanned() bool {
	if x != nil {
		return x.IsBanned
	}
	return false
}

// ManagerSearchCriteria - manager search criteria
type ManagerSearchCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageRequest
	PageRequest *PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
	// Filter by id
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Filter by name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Filter by email
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// Filter by phone
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *ManagerSearchCriteria) Reset() {
	*x = ManagerSearchCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_manager_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerSearchCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerSearchCriteria) ProtoMessage() {}

func (x *ManagerSearchCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_manager_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerSearchCriteria.ProtoReflect.Descriptor instead.
func (*ManagerSearchCriteria) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_manager_model_proto_rawDescGZIP(), []int{2}
}

func (x *ManagerSearchCriteria) GetPageRequest() *PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

func (x *ManagerSearchCriteria) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ManagerSearchCriteria) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ManagerSearchCriteria) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ManagerSearchCriteria) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_residenceserviceapi_manager_model_proto protoreflect.FileDescriptor

var file_residenceserviceapi_manager_model_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x64, 0x22, 0xb8, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x22, 0xcf, 0x01, 0x0a,
	0x15, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x66, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_residenceserviceapi_manager_model_proto_rawDescOnce sync.Once
	file_residenceserviceapi_manager_model_proto_rawDescData = file_residenceserviceapi_manager_model_proto_rawDesc
)

func file_residenceserviceapi_manager_model_proto_rawDescGZIP() []byte {
	file_residenceserviceapi_manager_model_proto_rawDescOnce.Do(func() {
		file_residenceserviceapi_manager_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_residenceserviceapi_manager_model_proto_rawDescData)
	})
	return file_residenceserviceapi_manager_model_proto_rawDescData
}

var file_residenceserviceapi_manager_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_residenceserviceapi_manager_model_proto_goTypes = []interface{}{
	(*ManagerWrite)(nil),          // 0: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerWrite
	(*ManagerRead)(nil),           // 1: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerRead
	(*ManagerSearchCriteria)(nil), // 2: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerSearchCriteria
	(*PageRequest)(nil),           // 3: gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
}
var file_residenceserviceapi_manager_model_proto_depIdxs = []int32{
	3, // 0: gitlab.com.mercury.residenceservice.generated.model.v1.ManagerSearchCriteria.page_request:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_residenceserviceapi_manager_model_proto_init() }
func file_residenceserviceapi_manager_model_proto_init() {
	if File_residenceserviceapi_manager_model_proto != nil {
		return
	}
	file_common_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_residenceserviceapi_manager_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerWrite); i {
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
		file_residenceserviceapi_manager_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerRead); i {
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
		file_residenceserviceapi_manager_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerSearchCriteria); i {
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
			RawDescriptor: file_residenceserviceapi_manager_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_residenceserviceapi_manager_model_proto_goTypes,
		DependencyIndexes: file_residenceserviceapi_manager_model_proto_depIdxs,
		MessageInfos:      file_residenceserviceapi_manager_model_proto_msgTypes,
	}.Build()
	File_residenceserviceapi_manager_model_proto = out.File
	file_residenceserviceapi_manager_model_proto_rawDesc = nil
	file_residenceserviceapi_manager_model_proto_goTypes = nil
	file_residenceserviceapi_manager_model_proto_depIdxs = nil
}
