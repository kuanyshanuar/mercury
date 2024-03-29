// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: identityserviceapi_authentication_model.proto

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

type GenderType int32

const (
	GenderType_GENDER_TYPE_UNKNOWN GenderType = 0
	GenderType_GENDER_TYPE_FEMALE  GenderType = 1
	GenderType_GENDER_TYPE_MALE    GenderType = 2
)

// Enum value maps for GenderType.
var (
	GenderType_name = map[int32]string{
		0: "GENDER_TYPE_UNKNOWN",
		1: "GENDER_TYPE_FEMALE",
		2: "GENDER_TYPE_MALE",
	}
	GenderType_value = map[string]int32{
		"GENDER_TYPE_UNKNOWN": 0,
		"GENDER_TYPE_FEMALE":  1,
		"GENDER_TYPE_MALE":    2,
	}
)

func (x GenderType) Enum() *GenderType {
	p := new(GenderType)
	*p = x
	return p
}

func (x GenderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenderType) Descriptor() protoreflect.EnumDescriptor {
	return file_identityserviceapi_authentication_model_proto_enumTypes[0].Descriptor()
}

func (GenderType) Type() protoreflect.EnumType {
	return &file_identityserviceapi_authentication_model_proto_enumTypes[0]
}

func (x GenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenderType.Descriptor instead.
func (GenderType) EnumDescriptor() ([]byte, []int) {
	return file_identityserviceapi_authentication_model_proto_rawDescGZIP(), []int{0}
}

// *
// Represents the user write model.
type UserWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Role id
	RoleId int64 `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// First name
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// Last name
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// City
	City string `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	// Phone
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	// Consultation phone number
	// only for builders
	ConsultationPhoneNumber string `protobuf:"bytes,10,opt,name=consultation_phone_number,json=consultationPhoneNumber,proto3" json:"consultation_phone_number,omitempty"`
	// Email
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// Gender
	//
	// Deprecated: Do not use.
	Gender string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	// Birthdate
	//
	// Deprecated: Do not use.
	BirthDate string `protobuf:"bytes,7,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	// Password
	Password string `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserWrite) Reset() {
	*x = UserWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identityserviceapi_authentication_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWrite) ProtoMessage() {}

func (x *UserWrite) ProtoReflect() protoreflect.Message {
	mi := &file_identityserviceapi_authentication_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWrite.ProtoReflect.Descriptor instead.
func (*UserWrite) Descriptor() ([]byte, []int) {
	return file_identityserviceapi_authentication_model_proto_rawDescGZIP(), []int{0}
}

func (x *UserWrite) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *UserWrite) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserWrite) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserWrite) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserWrite) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserWrite) GetConsultationPhoneNumber() string {
	if x != nil {
		return x.ConsultationPhoneNumber
	}
	return ""
}

func (x *UserWrite) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Deprecated: Do not use.
func (x *UserWrite) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

// Deprecated: Do not use.
func (x *UserWrite) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *UserWrite) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// *
// Represents the user read model.
type UserRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Role id
	RoleId int64 `protobuf:"varint,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// First name
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// Last name
	LastName string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// City
	City string `protobuf:"bytes,12,opt,name=city,proto3" json:"city,omitempty"`
	// Phone
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	// Consultation phone number
	// only for builders
	ConsultationPhoneNumber string `protobuf:"bytes,13,opt,name=consultation_phone_number,json=consultationPhoneNumber,proto3" json:"consultation_phone_number,omitempty"`
	// Email
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	// Gender
	//
	// Deprecated: Do not use.
	Gender string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	// Birthdate
	//
	// Deprecated: Do not use.
	BirthDate string `protobuf:"bytes,8,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	// Created time
	CreatedAt int64 `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Update time
	UpdatedAt int64 `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Deleted time
	DeletedAt int64 `protobuf:"varint,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *UserRead) Reset() {
	*x = UserRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identityserviceapi_authentication_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRead) ProtoMessage() {}

func (x *UserRead) ProtoReflect() protoreflect.Message {
	mi := &file_identityserviceapi_authentication_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRead.ProtoReflect.Descriptor instead.
func (*UserRead) Descriptor() ([]byte, []int) {
	return file_identityserviceapi_authentication_model_proto_rawDescGZIP(), []int{1}
}

func (x *UserRead) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRead) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *UserRead) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserRead) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserRead) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserRead) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserRead) GetConsultationPhoneNumber() string {
	if x != nil {
		return x.ConsultationPhoneNumber
	}
	return ""
}

func (x *UserRead) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Deprecated: Do not use.
func (x *UserRead) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

// Deprecated: Do not use.
func (x *UserRead) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *UserRead) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserRead) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *UserRead) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_identityserviceapi_authentication_model_proto protoreflect.FileDescriptor

var file_identityserviceapi_authentication_model_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x61, 0x70, 0x69, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x35, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0xb7, 0x02, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x87, 0x03, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x3a, 0x0a, 0x19,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0a, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x53, 0x0a, 0x0a, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x45, 0x4e, 0x44,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x45, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x75, 0x72, 0x79, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_identityserviceapi_authentication_model_proto_rawDescOnce sync.Once
	file_identityserviceapi_authentication_model_proto_rawDescData = file_identityserviceapi_authentication_model_proto_rawDesc
)

func file_identityserviceapi_authentication_model_proto_rawDescGZIP() []byte {
	file_identityserviceapi_authentication_model_proto_rawDescOnce.Do(func() {
		file_identityserviceapi_authentication_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_identityserviceapi_authentication_model_proto_rawDescData)
	})
	return file_identityserviceapi_authentication_model_proto_rawDescData
}

var file_identityserviceapi_authentication_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_identityserviceapi_authentication_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_identityserviceapi_authentication_model_proto_goTypes = []interface{}{
	(GenderType)(0),   // 0: gitlab.com.mercury.identityservice.generated.model.v1.GenderType
	(*UserWrite)(nil), // 1: gitlab.com.mercury.identityservice.generated.model.v1.UserWrite
	(*UserRead)(nil),  // 2: gitlab.com.mercury.identityservice.generated.model.v1.UserRead
}
var file_identityserviceapi_authentication_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_identityserviceapi_authentication_model_proto_init() }
func file_identityserviceapi_authentication_model_proto_init() {
	if File_identityserviceapi_authentication_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identityserviceapi_authentication_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserWrite); i {
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
		file_identityserviceapi_authentication_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRead); i {
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
			RawDescriptor: file_identityserviceapi_authentication_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_identityserviceapi_authentication_model_proto_goTypes,
		DependencyIndexes: file_identityserviceapi_authentication_model_proto_depIdxs,
		EnumInfos:         file_identityserviceapi_authentication_model_proto_enumTypes,
		MessageInfos:      file_identityserviceapi_authentication_model_proto_msgTypes,
	}.Build()
	File_identityserviceapi_authentication_model_proto = out.File
	file_identityserviceapi_authentication_model_proto_rawDesc = nil
	file_identityserviceapi_authentication_model_proto_goTypes = nil
	file_identityserviceapi_authentication_model_proto_depIdxs = nil
}
