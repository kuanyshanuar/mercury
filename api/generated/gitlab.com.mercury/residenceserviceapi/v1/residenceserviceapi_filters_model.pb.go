// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: residenceserviceapi_filters_model.proto

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

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_filters_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_filters_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_filters_model_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Filter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filters
	Filters []*Filter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *Filters) Reset() {
	*x = Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_filters_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_filters_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_filters_model_proto_rawDescGZIP(), []int{1}
}

func (x *Filters) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type FiltersV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Filters
	Filters []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *FiltersV2) Reset() {
	*x = FiltersV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_filters_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiltersV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiltersV2) ProtoMessage() {}

func (x *FiltersV2) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_filters_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiltersV2.ProtoReflect.Descriptor instead.
func (*FiltersV2) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_filters_model_proto_rawDescGZIP(), []int{2}
}

func (x *FiltersV2) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FiltersV2) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type FilterBuilder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *FilterBuilder) Reset() {
	*x = FilterBuilder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_filters_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterBuilder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterBuilder) ProtoMessage() {}

func (x *FilterBuilder) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_filters_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterBuilder.ProtoReflect.Descriptor instead.
func (*FilterBuilder) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_filters_model_proto_rawDescGZIP(), []int{3}
}

func (x *FilterBuilder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FilterBuilder) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

type CitySearchCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageRequest
	PageRequest *PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (x *CitySearchCriteria) Reset() {
	*x = CitySearchCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_filters_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitySearchCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitySearchCriteria) ProtoMessage() {}

func (x *CitySearchCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_filters_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitySearchCriteria.ProtoReflect.Descriptor instead.
func (*CitySearchCriteria) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_filters_model_proto_rawDescGZIP(), []int{4}
}

func (x *CitySearchCriteria) GetPageRequest() *PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

type DistrictSearchCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageRequest
	PageRequest *PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
	// City id
	CityId int64 `protobuf:"varint,2,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
}

func (x *DistrictSearchCriteria) Reset() {
	*x = DistrictSearchCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_filters_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistrictSearchCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistrictSearchCriteria) ProtoMessage() {}

func (x *DistrictSearchCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_filters_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistrictSearchCriteria.ProtoReflect.Descriptor instead.
func (*DistrictSearchCriteria) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_filters_model_proto_rawDescGZIP(), []int{5}
}

func (x *DistrictSearchCriteria) GetPageRequest() *PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

func (x *DistrictSearchCriteria) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

var File_residenceserviceapi_filters_model_proto protoreflect.FileDescriptor

var file_residenceserviceapi_filters_model_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x58,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x77, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x56, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x58, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x22, 0x3c, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x7c, 0x0a, 0x12, 0x43, 0x69, 0x74, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x66, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x99, 0x01,
	0x0a, 0x16, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x66, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2f,
	0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_residenceserviceapi_filters_model_proto_rawDescOnce sync.Once
	file_residenceserviceapi_filters_model_proto_rawDescData = file_residenceserviceapi_filters_model_proto_rawDesc
)

func file_residenceserviceapi_filters_model_proto_rawDescGZIP() []byte {
	file_residenceserviceapi_filters_model_proto_rawDescOnce.Do(func() {
		file_residenceserviceapi_filters_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_residenceserviceapi_filters_model_proto_rawDescData)
	})
	return file_residenceserviceapi_filters_model_proto_rawDescData
}

var file_residenceserviceapi_filters_model_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_residenceserviceapi_filters_model_proto_goTypes = []interface{}{
	(*Filter)(nil),                 // 0: gitlab.com.mercury.residenceservice.generated.model.v1.Filter
	(*Filters)(nil),                // 1: gitlab.com.mercury.residenceservice.generated.model.v1.Filters
	(*FiltersV2)(nil),              // 2: gitlab.com.mercury.residenceservice.generated.model.v1.FiltersV2
	(*FilterBuilder)(nil),          // 3: gitlab.com.mercury.residenceservice.generated.model.v1.FilterBuilder
	(*CitySearchCriteria)(nil),     // 4: gitlab.com.mercury.residenceservice.generated.model.v1.CitySearchCriteria
	(*DistrictSearchCriteria)(nil), // 5: gitlab.com.mercury.residenceservice.generated.model.v1.DistrictSearchCriteria
	(*PageRequest)(nil),            // 6: gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
}
var file_residenceserviceapi_filters_model_proto_depIdxs = []int32{
	0, // 0: gitlab.com.mercury.residenceservice.generated.model.v1.Filters.filters:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.Filter
	0, // 1: gitlab.com.mercury.residenceservice.generated.model.v1.FiltersV2.filters:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.Filter
	6, // 2: gitlab.com.mercury.residenceservice.generated.model.v1.CitySearchCriteria.page_request:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
	6, // 3: gitlab.com.mercury.residenceservice.generated.model.v1.DistrictSearchCriteria.page_request:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_residenceserviceapi_filters_model_proto_init() }
func file_residenceserviceapi_filters_model_proto_init() {
	if File_residenceserviceapi_filters_model_proto != nil {
		return
	}
	file_common_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_residenceserviceapi_filters_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_residenceserviceapi_filters_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filters); i {
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
		file_residenceserviceapi_filters_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiltersV2); i {
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
		file_residenceserviceapi_filters_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterBuilder); i {
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
		file_residenceserviceapi_filters_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CitySearchCriteria); i {
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
		file_residenceserviceapi_filters_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistrictSearchCriteria); i {
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
			RawDescriptor: file_residenceserviceapi_filters_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_residenceserviceapi_filters_model_proto_goTypes,
		DependencyIndexes: file_residenceserviceapi_filters_model_proto_depIdxs,
		MessageInfos:      file_residenceserviceapi_filters_model_proto_msgTypes,
	}.Build()
	File_residenceserviceapi_filters_model_proto = out.File
	file_residenceserviceapi_filters_model_proto_rawDesc = nil
	file_residenceserviceapi_filters_model_proto_goTypes = nil
	file_residenceserviceapi_filters_model_proto_depIdxs = nil
}