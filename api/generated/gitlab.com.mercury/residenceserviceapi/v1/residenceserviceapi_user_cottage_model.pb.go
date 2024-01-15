// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: residenceserviceapi_user_cottage_model.proto

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

// Represents the User favourite cottages model
type UserCottages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CottageId int64 `protobuf:"varint,2,opt,name=cottage_id,json=cottageId,proto3" json:"cottage_id,omitempty"`
	CreatedAt int64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UserCottages) Reset() {
	*x = UserCottages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_user_cottage_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCottages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCottages) ProtoMessage() {}

func (x *UserCottages) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_user_cottage_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCottages.ProtoReflect.Descriptor instead.
func (*UserCottages) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_user_cottage_model_proto_rawDescGZIP(), []int{0}
}

func (x *UserCottages) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserCottages) GetCottageId() int64 {
	if x != nil {
		return x.CottageId
	}
	return 0
}

func (x *UserCottages) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// *
// Represents the favourite cottage search criteria model.
type FavouriteCottageSearchCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageRequest
	PageRequest *PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
	// Sorting
	Sorts []*Sort `protobuf:"bytes,2,rep,name=sorts,proto3" json:"sorts,omitempty"`
}

func (x *FavouriteCottageSearchCriteria) Reset() {
	*x = FavouriteCottageSearchCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_user_cottage_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavouriteCottageSearchCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavouriteCottageSearchCriteria) ProtoMessage() {}

func (x *FavouriteCottageSearchCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_user_cottage_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavouriteCottageSearchCriteria.ProtoReflect.Descriptor instead.
func (*FavouriteCottageSearchCriteria) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_user_cottage_model_proto_rawDescGZIP(), []int{1}
}

func (x *FavouriteCottageSearchCriteria) GetPageRequest() *PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

func (x *FavouriteCottageSearchCriteria) GetSorts() []*Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

var File_residenceserviceapi_user_cottage_model_proto protoreflect.FileDescriptor

var file_residenceserviceapi_user_cottage_model_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x74, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75,
	0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x6f, 0x74, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x74, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x74, 0x74, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xdc, 0x01, 0x0a, 0x1e, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x74, 0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x12, 0x66, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x05,
	0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_residenceserviceapi_user_cottage_model_proto_rawDescOnce sync.Once
	file_residenceserviceapi_user_cottage_model_proto_rawDescData = file_residenceserviceapi_user_cottage_model_proto_rawDesc
)

func file_residenceserviceapi_user_cottage_model_proto_rawDescGZIP() []byte {
	file_residenceserviceapi_user_cottage_model_proto_rawDescOnce.Do(func() {
		file_residenceserviceapi_user_cottage_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_residenceserviceapi_user_cottage_model_proto_rawDescData)
	})
	return file_residenceserviceapi_user_cottage_model_proto_rawDescData
}

var file_residenceserviceapi_user_cottage_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_residenceserviceapi_user_cottage_model_proto_goTypes = []interface{}{
	(*UserCottages)(nil),                   // 0: gitlab.com.mercury.residenceservice.generated.model.v1.UserCottages
	(*FavouriteCottageSearchCriteria)(nil), // 1: gitlab.com.mercury.residenceservice.generated.model.v1.FavouriteCottageSearchCriteria
	(*PageRequest)(nil),                    // 2: gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
	(*Sort)(nil),                           // 3: gitlab.com.mercury.residenceservice.generated.model.v1.Sort
}
var file_residenceserviceapi_user_cottage_model_proto_depIdxs = []int32{
	2, // 0: gitlab.com.mercury.residenceservice.generated.model.v1.FavouriteCottageSearchCriteria.page_request:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
	3, // 1: gitlab.com.mercury.residenceservice.generated.model.v1.FavouriteCottageSearchCriteria.sorts:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.Sort
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_residenceserviceapi_user_cottage_model_proto_init() }
func file_residenceserviceapi_user_cottage_model_proto_init() {
	if File_residenceserviceapi_user_cottage_model_proto != nil {
		return
	}
	file_common_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_residenceserviceapi_user_cottage_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCottages); i {
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
		file_residenceserviceapi_user_cottage_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavouriteCottageSearchCriteria); i {
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
			RawDescriptor: file_residenceserviceapi_user_cottage_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_residenceserviceapi_user_cottage_model_proto_goTypes,
		DependencyIndexes: file_residenceserviceapi_user_cottage_model_proto_depIdxs,
		MessageInfos:      file_residenceserviceapi_user_cottage_model_proto_msgTypes,
	}.Build()
	File_residenceserviceapi_user_cottage_model_proto = out.File
	file_residenceserviceapi_user_cottage_model_proto_rawDesc = nil
	file_residenceserviceapi_user_cottage_model_proto_goTypes = nil
	file_residenceserviceapi_user_cottage_model_proto_depIdxs = nil
}
