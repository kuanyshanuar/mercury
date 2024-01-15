// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: residenceserviceapi_lead_cottage_model.proto

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

type LeadCottageRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CottageId   int64       `protobuf:"varint,2,opt,name=cottage_id,json=cottageId,proto3" json:"cottage_id,omitempty"`
	CottageName string      `protobuf:"bytes,3,opt,name=cottage_name,json=cottageName,proto3" json:"cottage_name,omitempty"`
	StatusId    int64       `protobuf:"varint,4,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	Status      *LeadStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	DateStart   int64       `protobuf:"varint,6,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd     int64       `protobuf:"varint,7,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
}

func (x *LeadCottageRead) Reset() {
	*x = LeadCottageRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_lead_cottage_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadCottageRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadCottageRead) ProtoMessage() {}

func (x *LeadCottageRead) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_lead_cottage_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadCottageRead.ProtoReflect.Descriptor instead.
func (*LeadCottageRead) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_lead_cottage_model_proto_rawDescGZIP(), []int{0}
}

func (x *LeadCottageRead) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeadCottageRead) GetCottageId() int64 {
	if x != nil {
		return x.CottageId
	}
	return 0
}

func (x *LeadCottageRead) GetCottageName() string {
	if x != nil {
		return x.CottageName
	}
	return ""
}

func (x *LeadCottageRead) GetStatusId() int64 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

func (x *LeadCottageRead) GetStatus() *LeadStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *LeadCottageRead) GetDateStart() int64 {
	if x != nil {
		return x.DateStart
	}
	return 0
}

func (x *LeadCottageRead) GetDateEnd() int64 {
	if x != nil {
		return x.DateEnd
	}
	return 0
}

type LeadCottageWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CottageId  int64 `protobuf:"varint,2,opt,name=cottage_id,json=cottageId,proto3" json:"cottage_id,omitempty"`
	StatusId   int64 `protobuf:"varint,3,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	IssueDate  int64 `protobuf:"varint,4,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	ExpireDate int64 `protobuf:"varint,5,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
}

func (x *LeadCottageWrite) Reset() {
	*x = LeadCottageWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_lead_cottage_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadCottageWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadCottageWrite) ProtoMessage() {}

func (x *LeadCottageWrite) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_lead_cottage_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadCottageWrite.ProtoReflect.Descriptor instead.
func (*LeadCottageWrite) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_lead_cottage_model_proto_rawDescGZIP(), []int{1}
}

func (x *LeadCottageWrite) GetCottageId() int64 {
	if x != nil {
		return x.CottageId
	}
	return 0
}

func (x *LeadCottageWrite) GetStatusId() int64 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

func (x *LeadCottageWrite) GetIssueDate() int64 {
	if x != nil {
		return x.IssueDate
	}
	return 0
}

func (x *LeadCottageWrite) GetExpireDate() int64 {
	if x != nil {
		return x.ExpireDate
	}
	return 0
}

type LeadCottageSearchCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageRequest
	PageRequest *PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
	// Filter by status
	StatusId int64 `protobuf:"varint,2,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	// Filter by name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LeadCottageSearchCriteria) Reset() {
	*x = LeadCottageSearchCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceserviceapi_lead_cottage_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadCottageSearchCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadCottageSearchCriteria) ProtoMessage() {}

func (x *LeadCottageSearchCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_residenceserviceapi_lead_cottage_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadCottageSearchCriteria.ProtoReflect.Descriptor instead.
func (*LeadCottageSearchCriteria) Descriptor() ([]byte, []int) {
	return file_residenceserviceapi_lead_cottage_model_proto_rawDescGZIP(), []int{2}
}

func (x *LeadCottageSearchCriteria) GetPageRequest() *PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

func (x *LeadCottageSearchCriteria) GetStatusId() int64 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

func (x *LeadCottageSearchCriteria) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_residenceserviceapi_lead_cottage_model_proto protoreflect.FileDescriptor

var file_residenceserviceapi_lead_cottage_model_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x74, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75,
	0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x0f, 0x4c,
	0x65, 0x61, 0x64, 0x43, 0x6f, 0x74, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x74, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x74, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x74, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x74, 0x74, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x5a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75,
	0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x10, 0x4c, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x74, 0x74,
	0x61, 0x67, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x74, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x74, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x19, 0x4c, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x74,
	0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x12, 0x66, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_residenceserviceapi_lead_cottage_model_proto_rawDescOnce sync.Once
	file_residenceserviceapi_lead_cottage_model_proto_rawDescData = file_residenceserviceapi_lead_cottage_model_proto_rawDesc
)

func file_residenceserviceapi_lead_cottage_model_proto_rawDescGZIP() []byte {
	file_residenceserviceapi_lead_cottage_model_proto_rawDescOnce.Do(func() {
		file_residenceserviceapi_lead_cottage_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_residenceserviceapi_lead_cottage_model_proto_rawDescData)
	})
	return file_residenceserviceapi_lead_cottage_model_proto_rawDescData
}

var file_residenceserviceapi_lead_cottage_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_residenceserviceapi_lead_cottage_model_proto_goTypes = []interface{}{
	(*LeadCottageRead)(nil),           // 0: gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageRead
	(*LeadCottageWrite)(nil),          // 1: gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageWrite
	(*LeadCottageSearchCriteria)(nil), // 2: gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageSearchCriteria
	(*LeadStatus)(nil),                // 3: gitlab.com.mercury.residenceservice.generated.model.v1.LeadStatus
	(*PageRequest)(nil),               // 4: gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
}
var file_residenceserviceapi_lead_cottage_model_proto_depIdxs = []int32{
	3, // 0: gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageRead.status:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.LeadStatus
	4, // 1: gitlab.com.mercury.residenceservice.generated.model.v1.LeadCottageSearchCriteria.page_request:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_residenceserviceapi_lead_cottage_model_proto_init() }
func file_residenceserviceapi_lead_cottage_model_proto_init() {
	if File_residenceserviceapi_lead_cottage_model_proto != nil {
		return
	}
	file_common_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_residenceserviceapi_lead_cottage_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadCottageRead); i {
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
		file_residenceserviceapi_lead_cottage_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadCottageWrite); i {
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
		file_residenceserviceapi_lead_cottage_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadCottageSearchCriteria); i {
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
			RawDescriptor: file_residenceserviceapi_lead_cottage_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_residenceserviceapi_lead_cottage_model_proto_goTypes,
		DependencyIndexes: file_residenceserviceapi_lead_cottage_model_proto_depIdxs,
		MessageInfos:      file_residenceserviceapi_lead_cottage_model_proto_msgTypes,
	}.Build()
	File_residenceserviceapi_lead_cottage_model_proto = out.File
	file_residenceserviceapi_lead_cottage_model_proto_rawDesc = nil
	file_residenceserviceapi_lead_cottage_model_proto_goTypes = nil
	file_residenceserviceapi_lead_cottage_model_proto_depIdxs = nil
}