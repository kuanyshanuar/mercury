// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: residenceapiservice_news_model.proto

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

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID - id of the article
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Title is a title of the article
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// ShortDescription is a short description at the beginning
	ShortDescription string `protobuf:"bytes,3,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
	// Slug is a slug of the article on the URL
	Slug string `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	// Content is a content of the article
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	// ViewsCount is value of the viewers
	ViewsCount int64 `protobuf:"varint,12,opt,name=views_count,json=viewsCount,proto3" json:"views_count,omitempty"`
	// SourceUrl is for the case when the article was taken from other website
	SourceUrl string `protobuf:"bytes,13,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// AuthorName is the name of the article's author
	AuthorName string `protobuf:"bytes,14,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	// Images array of the URL of images
	Images []string `protobuf:"bytes,15,rep,name=images,proto3" json:"images,omitempty"`
	// CreatedAt stores the date of creation of article
	CreatedAt int64 `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// UpdatedAt stores update date of the article
	UpdatedAt int64 `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// DeletedAt stores the deletion date
	DeletedAt int64 `protobuf:"varint,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	// CreatedBy stores the ID of the moderator
	CreatedBy int64 `protobuf:"varint,9,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// DeletedBy stores the ID of the moderator
	DeletedBy int64 `protobuf:"varint,10,opt,name=deleted_by,json=deletedBy,proto3" json:"deleted_by,omitempty"`
	// UpdatedBy stores the ID of the moderator
	UpdatedBy int64 `protobuf:"varint,11,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	// HasLikedByMe points whether user has liked article or not
	HasLikedByMe bool `protobuf:"varint,16,opt,name=has_liked_by_me,json=hasLikedByMe,proto3" json:"has_liked_by_me,omitempty"`
	// HasDislikedByMe points whether user has liked article or not
	HasDislikedByMe bool `protobuf:"varint,17,opt,name=has_disliked_by_me,json=hasDislikedByMe,proto3" json:"has_disliked_by_me,omitempty"`
	// Likes count
	Likes int64 `protobuf:"varint,18,opt,name=likes,proto3" json:"likes,omitempty"`
	// Dislikes count
	Dislikes int64 `protobuf:"varint,19,opt,name=dislikes,proto3" json:"dislikes,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceapiservice_news_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_residenceapiservice_news_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_residenceapiservice_news_model_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *Article) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetViewsCount() int64 {
	if x != nil {
		return x.ViewsCount
	}
	return 0
}

func (x *Article) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *Article) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *Article) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Article) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Article) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Article) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *Article) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *Article) GetDeletedBy() int64 {
	if x != nil {
		return x.DeletedBy
	}
	return 0
}

func (x *Article) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *Article) GetHasLikedByMe() bool {
	if x != nil {
		return x.HasLikedByMe
	}
	return false
}

func (x *Article) GetHasDislikedByMe() bool {
	if x != nil {
		return x.HasDislikedByMe
	}
	return false
}

func (x *Article) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *Article) GetDislikes() int64 {
	if x != nil {
		return x.Dislikes
	}
	return 0
}

type NewsSearchCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageRequest *PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"` // page request
	Id          int64        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`                                     // filter by id
	Title       string       `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                // filter by title
	Short       string       `protobuf:"bytes,4,opt,name=short,proto3" json:"short,omitempty"`                                // filter by short description
	Slug        string       `protobuf:"bytes,5,opt,name=slug,proto3" json:"slug,omitempty"`                                  // filter by slug
	Sorts       []*Sort      `protobuf:"bytes,6,rep,name=sorts,proto3" json:"sorts,omitempty"`                                // sorting
}

func (x *NewsSearchCriteria) Reset() {
	*x = NewsSearchCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residenceapiservice_news_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsSearchCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsSearchCriteria) ProtoMessage() {}

func (x *NewsSearchCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_residenceapiservice_news_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsSearchCriteria.ProtoReflect.Descriptor instead.
func (*NewsSearchCriteria) Descriptor() ([]byte, []int) {
	return file_residenceapiservice_news_model_proto_rawDescGZIP(), []int{1}
}

func (x *NewsSearchCriteria) GetPageRequest() *PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

func (x *NewsSearchCriteria) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewsSearchCriteria) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsSearchCriteria) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *NewsSearchCriteria) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *NewsSearchCriteria) GetSorts() []*Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

var File_residenceapiservice_news_model_proto protoreflect.FileDescriptor

var file_residenceapiservice_news_model_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x12,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc3, 0x04, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x65, 0x77, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x25, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x5f, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x4d, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x68, 0x61, 0x73,
	0x5f, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x6d, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b,
	0x65, 0x64, 0x42, 0x79, 0x4d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x22, 0xa0, 0x02, 0x0a, 0x12, 0x4e, 0x65, 0x77,
	0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12,
	0x66, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x52, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_residenceapiservice_news_model_proto_rawDescOnce sync.Once
	file_residenceapiservice_news_model_proto_rawDescData = file_residenceapiservice_news_model_proto_rawDesc
)

func file_residenceapiservice_news_model_proto_rawDescGZIP() []byte {
	file_residenceapiservice_news_model_proto_rawDescOnce.Do(func() {
		file_residenceapiservice_news_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_residenceapiservice_news_model_proto_rawDescData)
	})
	return file_residenceapiservice_news_model_proto_rawDescData
}

var file_residenceapiservice_news_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_residenceapiservice_news_model_proto_goTypes = []interface{}{
	(*Article)(nil),            // 0: gitlab.com.mercury.residenceservice.generated.model.v1.Article
	(*NewsSearchCriteria)(nil), // 1: gitlab.com.mercury.residenceservice.generated.model.v1.NewsSearchCriteria
	(*PageRequest)(nil),        // 2: gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
	(*Sort)(nil),               // 3: gitlab.com.mercury.residenceservice.generated.model.v1.Sort
}
var file_residenceapiservice_news_model_proto_depIdxs = []int32{
	2, // 0: gitlab.com.mercury.residenceservice.generated.model.v1.NewsSearchCriteria.page_request:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.PageRequest
	3, // 1: gitlab.com.mercury.residenceservice.generated.model.v1.NewsSearchCriteria.sorts:type_name -> gitlab.com.mercury.residenceservice.generated.model.v1.Sort
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_residenceapiservice_news_model_proto_init() }
func file_residenceapiservice_news_model_proto_init() {
	if File_residenceapiservice_news_model_proto != nil {
		return
	}
	file_common_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_residenceapiservice_news_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_residenceapiservice_news_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsSearchCriteria); i {
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
			RawDescriptor: file_residenceapiservice_news_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_residenceapiservice_news_model_proto_goTypes,
		DependencyIndexes: file_residenceapiservice_news_model_proto_depIdxs,
		MessageInfos:      file_residenceapiservice_news_model_proto_msgTypes,
	}.Build()
	File_residenceapiservice_news_model_proto = out.File
	file_residenceapiservice_news_model_proto_rawDesc = nil
	file_residenceapiservice_news_model_proto_goTypes = nil
	file_residenceapiservice_news_model_proto_depIdxs = nil
}
