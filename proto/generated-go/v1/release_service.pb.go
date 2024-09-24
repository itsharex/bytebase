// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/release_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Release_File_Type int32

const (
	Release_File_TYPE_UNSPECIFIED Release_File_Type = 0
	Release_File_VERSIONED        Release_File_Type = 1
)

// Enum value maps for Release_File_Type.
var (
	Release_File_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "VERSIONED",
	}
	Release_File_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"VERSIONED":        1,
	}
)

func (x Release_File_Type) Enum() *Release_File_Type {
	p := new(Release_File_Type)
	*p = x
	return p
}

func (x Release_File_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Release_File_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_release_service_proto_enumTypes[0].Descriptor()
}

func (Release_File_Type) Type() protoreflect.EnumType {
	return &file_v1_release_service_proto_enumTypes[0]
}

func (x Release_File_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Release_File_Type.Descriptor instead.
func (Release_File_Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_release_service_proto_rawDescGZIP(), []int{4, 0, 0}
}

type GetReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: projects/{project}/releases/{release}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetReleaseRequest) Reset() {
	*x = GetReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_release_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReleaseRequest) ProtoMessage() {}

func (x *GetReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_release_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReleaseRequest.ProtoReflect.Descriptor instead.
func (*GetReleaseRequest) Descriptor() ([]byte, []int) {
	return file_v1_release_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetReleaseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListReleasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: projects/{project}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of change histories to return. The service may return fewer than this value.
	// If unspecified, at most 10 change histories will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Not used. A page token, received from a previous `ListReleasesRequest` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListReleasesRequest` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListReleasesRequest) Reset() {
	*x = ListReleasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_release_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReleasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReleasesRequest) ProtoMessage() {}

func (x *ListReleasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_release_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReleasesRequest.ProtoReflect.Descriptor instead.
func (*ListReleasesRequest) Descriptor() ([]byte, []int) {
	return file_v1_release_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListReleasesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListReleasesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListReleasesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListReleasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Releases []*Release `protobuf:"bytes,1,rep,name=releases,proto3" json:"releases,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListReleasesResponse) Reset() {
	*x = ListReleasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_release_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReleasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReleasesResponse) ProtoMessage() {}

func (x *ListReleasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_release_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReleasesResponse.ProtoReflect.Descriptor instead.
func (*ListReleasesResponse) Descriptor() ([]byte, []int) {
	return file_v1_release_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListReleasesResponse) GetReleases() []*Release {
	if x != nil {
		return x.Releases
	}
	return nil
}

func (x *ListReleasesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: projects/{project}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The release to create.
	Release *Release `protobuf:"bytes,2,opt,name=release,proto3" json:"release,omitempty"`
}

func (x *CreateReleaseRequest) Reset() {
	*x = CreateReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_release_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReleaseRequest) ProtoMessage() {}

func (x *CreateReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_release_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReleaseRequest.ProtoReflect.Descriptor instead.
func (*CreateReleaseRequest) Descriptor() ([]byte, []int) {
	return file_v1_release_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateReleaseRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateReleaseRequest) GetRelease() *Release {
	if x != nil {
		return x.Release
	}
	return nil
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: projects/{project}/releases/{release}
	Name      string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Files     []*Release_File    `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	VcsSource *Release_VCSSource `protobuf:"bytes,3,opt,name=vcs_source,json=vcsSource,proto3" json:"vcs_source,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_release_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_v1_release_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_v1_release_service_proto_rawDescGZIP(), []int{4}
}

func (x *Release) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Release) GetFiles() []*Release_File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Release) GetVcsSource() *Release_VCSSource {
	if x != nil {
		return x.VcsSource
	}
	return nil
}

type Release_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	// The sheet that holds the statement.
	// Format: projects/{project}/sheets/{sheet}
	Sheet string `protobuf:"bytes,2,opt,name=sheet,proto3" json:"sheet,omitempty"`
	// The SHA1 hash value of the sheet.
	SheetSha1 string            `protobuf:"bytes,3,opt,name=sheet_sha1,json=sheetSha1,proto3" json:"sheet_sha1,omitempty"`
	Type      Release_File_Type `protobuf:"varint,4,opt,name=type,proto3,enum=bytebase.v1.Release_File_Type" json:"type,omitempty"`
	Version   string            `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Release_File) Reset() {
	*x = Release_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_release_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release_File) ProtoMessage() {}

func (x *Release_File) ProtoReflect() protoreflect.Message {
	mi := &file_v1_release_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release_File.ProtoReflect.Descriptor instead.
func (*Release_File) Descriptor() ([]byte, []int) {
	return file_v1_release_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Release_File) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Release_File) GetSheet() string {
	if x != nil {
		return x.Sheet
	}
	return ""
}

func (x *Release_File) GetSheetSha1() string {
	if x != nil {
		return x.SheetSha1
	}
	return ""
}

func (x *Release_File) GetType() Release_File_Type {
	if x != nil {
		return x.Type
	}
	return Release_File_TYPE_UNSPECIFIED
}

func (x *Release_File) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Release_VCSSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VcsType        VCSType `protobuf:"varint,1,opt,name=vcs_type,json=vcsType,proto3,enum=bytebase.v1.VCSType" json:"vcs_type,omitempty"`
	PullRequestUrl string  `protobuf:"bytes,2,opt,name=pull_request_url,json=pullRequestUrl,proto3" json:"pull_request_url,omitempty"`
}

func (x *Release_VCSSource) Reset() {
	*x = Release_VCSSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_release_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release_VCSSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release_VCSSource) ProtoMessage() {}

func (x *Release_VCSSource) ProtoReflect() protoreflect.Message {
	mi := &file_v1_release_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release_VCSSource.ProtoReflect.Descriptor instead.
func (*Release_VCSSource) Descriptor() ([]byte, []int) {
	return file_v1_release_service_proto_rawDescGZIP(), []int{4, 1}
}

func (x *Release_VCSSource) GetVcsType() VCSType {
	if x != nil {
		return x.VcsType
	}
	return VCSType_VCS_TYPE_UNSPECIFIED
}

func (x *Release_VCSSource) GetPullRequestUrl() string {
	if x != nil {
		return x.PullRequestUrl
	}
	return ""
}

var File_v1_release_service_proto protoreflect.FileDescriptor

var file_v1_release_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d,
	0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x16, 0x0a, 0x14, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xe2, 0x41, 0x01,
	0x02, 0xfa, 0x41, 0x16, 0x0a, 0x14, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x70,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x83, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xe2, 0x41, 0x01, 0x02, 0xfa,
	0x41, 0x16, 0x0a, 0x14, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x07, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0xe3, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x76, 0x63, 0x73, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x2e, 0x56, 0x43, 0x53, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x76, 0x63, 0x73,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0xeb, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73,
	0x68, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xfa, 0x41, 0x14, 0x0a,
	0x12, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68,
	0x65, 0x65, 0x74, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x53, 0x68, 0x61, 0x31, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x45, 0x44, 0x10, 0x01, 0x1a, 0x66, 0x0a, 0x09, 0x56, 0x43, 0x53, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x2f, 0x0a, 0x08, 0x76, 0x63, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x43, 0x53, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x76, 0x63, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x75,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x32, 0x9d, 0x03, 0x0a,
	0x0e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x73, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x22, 0x2f, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x86, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0xda, 0x41, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x8c, 0x01,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0x42, 0xda, 0x41, 0x0e, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x2c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2b, 0x3a, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_release_service_proto_rawDescOnce sync.Once
	file_v1_release_service_proto_rawDescData = file_v1_release_service_proto_rawDesc
)

func file_v1_release_service_proto_rawDescGZIP() []byte {
	file_v1_release_service_proto_rawDescOnce.Do(func() {
		file_v1_release_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_release_service_proto_rawDescData)
	})
	return file_v1_release_service_proto_rawDescData
}

var file_v1_release_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_release_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_release_service_proto_goTypes = []any{
	(Release_File_Type)(0),       // 0: bytebase.v1.Release.File.Type
	(*GetReleaseRequest)(nil),    // 1: bytebase.v1.GetReleaseRequest
	(*ListReleasesRequest)(nil),  // 2: bytebase.v1.ListReleasesRequest
	(*ListReleasesResponse)(nil), // 3: bytebase.v1.ListReleasesResponse
	(*CreateReleaseRequest)(nil), // 4: bytebase.v1.CreateReleaseRequest
	(*Release)(nil),              // 5: bytebase.v1.Release
	(*Release_File)(nil),         // 6: bytebase.v1.Release.File
	(*Release_VCSSource)(nil),    // 7: bytebase.v1.Release.VCSSource
	(VCSType)(0),                 // 8: bytebase.v1.VCSType
}
var file_v1_release_service_proto_depIdxs = []int32{
	5, // 0: bytebase.v1.ListReleasesResponse.releases:type_name -> bytebase.v1.Release
	5, // 1: bytebase.v1.CreateReleaseRequest.release:type_name -> bytebase.v1.Release
	6, // 2: bytebase.v1.Release.files:type_name -> bytebase.v1.Release.File
	7, // 3: bytebase.v1.Release.vcs_source:type_name -> bytebase.v1.Release.VCSSource
	0, // 4: bytebase.v1.Release.File.type:type_name -> bytebase.v1.Release.File.Type
	8, // 5: bytebase.v1.Release.VCSSource.vcs_type:type_name -> bytebase.v1.VCSType
	1, // 6: bytebase.v1.ReleaseService.GetRelease:input_type -> bytebase.v1.GetReleaseRequest
	2, // 7: bytebase.v1.ReleaseService.ListReleases:input_type -> bytebase.v1.ListReleasesRequest
	4, // 8: bytebase.v1.ReleaseService.CreateRelease:input_type -> bytebase.v1.CreateReleaseRequest
	5, // 9: bytebase.v1.ReleaseService.GetRelease:output_type -> bytebase.v1.Release
	3, // 10: bytebase.v1.ReleaseService.ListReleases:output_type -> bytebase.v1.ListReleasesResponse
	5, // 11: bytebase.v1.ReleaseService.CreateRelease:output_type -> bytebase.v1.Release
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_release_service_proto_init() }
func file_v1_release_service_proto_init() {
	if File_v1_release_service_proto != nil {
		return
	}
	file_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_release_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetReleaseRequest); i {
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
		file_v1_release_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListReleasesRequest); i {
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
		file_v1_release_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListReleasesResponse); i {
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
		file_v1_release_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateReleaseRequest); i {
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
		file_v1_release_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Release); i {
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
		file_v1_release_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Release_File); i {
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
		file_v1_release_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Release_VCSSource); i {
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
			RawDescriptor: file_v1_release_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_release_service_proto_goTypes,
		DependencyIndexes: file_v1_release_service_proto_depIdxs,
		EnumInfos:         file_v1_release_service_proto_enumTypes,
		MessageInfos:      file_v1_release_service_proto_msgTypes,
	}.Build()
	File_v1_release_service_proto = out.File
	file_v1_release_service_proto_rawDesc = nil
	file_v1_release_service_proto_goTypes = nil
	file_v1_release_service_proto_depIdxs = nil
}