// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/api/Jobs.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type JobStatusEnum int32

const (
	JobStatusEnum_INIT       JobStatusEnum = 0
	JobStatusEnum_RUNNING    JobStatusEnum = 1
	JobStatusEnum_SUCCESFULL JobStatusEnum = 2
	JobStatusEnum_ERROR      JobStatusEnum = 3
)

// Enum value maps for JobStatusEnum.
var (
	JobStatusEnum_name = map[int32]string{
		0: "INIT",
		1: "RUNNING",
		2: "SUCCESFULL",
		3: "ERROR",
	}
	JobStatusEnum_value = map[string]int32{
		"INIT":       0,
		"RUNNING":    1,
		"SUCCESFULL": 2,
		"ERROR":      3,
	}
)

func (x JobStatusEnum) Enum() *JobStatusEnum {
	p := new(JobStatusEnum)
	*p = x
	return p
}

func (x JobStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_api_Jobs_proto_enumTypes[0].Descriptor()
}

func (JobStatusEnum) Type() protoreflect.EnumType {
	return &file_proto_api_Jobs_proto_enumTypes[0]
}

func (x JobStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStatusEnum.Descriptor instead.
func (JobStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{0}
}

type InitJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasProdigalFile bool `protobuf:"varint,1,opt,name=hasProdigalFile,proto3" json:"hasProdigalFile,omitempty"`
}

func (x *InitJobRequest) Reset() {
	*x = InitJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitJobRequest) ProtoMessage() {}

func (x *InitJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitJobRequest.ProtoReflect.Descriptor instead.
func (*InitJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{0}
}

func (x *InitJobRequest) GetHasProdigalFile() bool {
	if x != nil {
		return x.HasProdigalFile
	}
	return false
}

type InitJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadLinkFasta    string   `protobuf:"bytes,1,opt,name=uploadLinkFasta,proto3" json:"uploadLinkFasta,omitempty"`
	UploadLinkProdigal string   `protobuf:"bytes,2,opt,name=uploadLinkProdigal,proto3" json:"uploadLinkProdigal,omitempty"`
	Secret             string   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	Job                *JobAuth `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *InitJobResponse) Reset() {
	*x = InitJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitJobResponse) ProtoMessage() {}

func (x *InitJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitJobResponse.ProtoReflect.Descriptor instead.
func (*InitJobResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{1}
}

func (x *InitJobResponse) GetUploadLinkFasta() string {
	if x != nil {
		return x.UploadLinkFasta
	}
	return ""
}

func (x *InitJobResponse) GetUploadLinkProdigal() string {
	if x != nil {
		return x.UploadLinkProdigal
	}
	return ""
}

func (x *InitJobResponse) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *InitJobResponse) GetJob() *JobAuth {
	if x != nil {
		return x.Job
	}
	return nil
}

type StartJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID   string     `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Secret  string     `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Confiig *JobConfig `protobuf:"bytes,3,opt,name=confiig,proto3" json:"confiig,omitempty"`
}

func (x *StartJobRequest) Reset() {
	*x = StartJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobRequest) ProtoMessage() {}

func (x *StartJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobRequest.ProtoReflect.Descriptor instead.
func (*StartJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{2}
}

func (x *StartJobRequest) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *StartJobRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *StartJobRequest) GetConfiig() *JobConfig {
	if x != nil {
		return x.Confiig
	}
	return nil
}

type JobConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JobConfig) Reset() {
	*x = JobConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobConfig) ProtoMessage() {}

func (x *JobConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobConfig.ProtoReflect.Descriptor instead.
func (*JobConfig) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{3}
}

type JobStatusReponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*JobStatusResponse `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *JobStatusReponseList) Reset() {
	*x = JobStatusReponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatusReponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusReponseList) ProtoMessage() {}

func (x *JobStatusReponseList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusReponseList.ProtoReflect.Descriptor instead.
func (*JobStatusReponseList) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{4}
}

func (x *JobStatusReponseList) GetJobs() []*JobStatusResponse {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type JobStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID     string        `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	JobStatus JobStatusEnum `protobuf:"varint,2,opt,name=jobStatus,proto3,enum=JobStatusEnum" json:"jobStatus,omitempty"`
}

func (x *JobStatusResponse) Reset() {
	*x = JobStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusResponse) ProtoMessage() {}

func (x *JobStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusResponse.ProtoReflect.Descriptor instead.
func (*JobStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{5}
}

func (x *JobStatusResponse) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *JobStatusResponse) GetJobStatus() JobStatusEnum {
	if x != nil {
		return x.JobStatus
	}
	return JobStatusEnum_INIT
}

type JobResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID      string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	ResultLink string `protobuf:"bytes,2,opt,name=resultLink,proto3" json:"resultLink,omitempty"`
}

func (x *JobResultResponse) Reset() {
	*x = JobResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResultResponse) ProtoMessage() {}

func (x *JobResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResultResponse.ProtoReflect.Descriptor instead.
func (*JobResultResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{6}
}

func (x *JobResultResponse) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *JobResultResponse) GetResultLink() string {
	if x != nil {
		return x.ResultLink
	}
	return ""
}

type JobStatusRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*JobAuth `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *JobStatusRequestList) Reset() {
	*x = JobStatusRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatusRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusRequestList) ProtoMessage() {}

func (x *JobStatusRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusRequestList.ProtoReflect.Descriptor instead.
func (*JobStatusRequestList) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{7}
}

func (x *JobStatusRequestList) GetJobs() []*JobAuth {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type JobAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	JobID  string `protobuf:"bytes,2,opt,name=jobID,proto3" json:"jobID,omitempty"`
}

func (x *JobAuth) Reset() {
	*x = JobAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobAuth) ProtoMessage() {}

func (x *JobAuth) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobAuth.ProtoReflect.Descriptor instead.
func (*JobAuth) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{8}
}

func (x *JobAuth) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *JobAuth) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_Jobs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_Jobs_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_api_Jobs_proto_rawDescGZIP(), []int{9}
}

var File_proto_api_Jobs_proto protoreflect.FileDescriptor

var file_proto_api_Jobs_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x4a, 0x6f, 0x62, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x50,
	0x72, 0x6f, 0x64, 0x69, 0x67, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x67, 0x61, 0x6c, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x61, 0x73, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x61, 0x73, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x12, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72,
	0x6f, 0x64, 0x69, 0x67, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x67, 0x61, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x03, 0x6a, 0x6f, 0x62, 0x22, 0x65, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x69, 0x67, 0x22, 0x0b, 0x0a, 0x09, 0x4a,
	0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3e, 0x0a, 0x14, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x57, 0x0a, 0x11, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f,
	0x62, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x49, 0x0a, 0x11, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x34, 0x0a, 0x14,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x6a, 0x6f,
	0x62, 0x73, 0x22, 0x37, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x41, 0x75, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x2a, 0x41, 0x0a, 0x0d, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x32, 0xcc, 0x01, 0x0a, 0x09, 0x42, 0x61, 0x6b, 0x74,
	0x61, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x4a, 0x6f, 0x62,
	0x12, 0x0f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f,
	0x62, 0x12, 0x08, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x08, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x75, 0x74, 0x68,
	0x1a, 0x12, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x62, 0x69, 0x6f, 0x2f, 0x62, 0x61, 0x6b, 0x74, 0x61, 0x2d,
	0x77, 0x65, 0x62, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_Jobs_proto_rawDescOnce sync.Once
	file_proto_api_Jobs_proto_rawDescData = file_proto_api_Jobs_proto_rawDesc
)

func file_proto_api_Jobs_proto_rawDescGZIP() []byte {
	file_proto_api_Jobs_proto_rawDescOnce.Do(func() {
		file_proto_api_Jobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_Jobs_proto_rawDescData)
	})
	return file_proto_api_Jobs_proto_rawDescData
}

var file_proto_api_Jobs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_api_Jobs_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_api_Jobs_proto_goTypes = []interface{}{
	(JobStatusEnum)(0),           // 0: JobStatusEnum
	(*InitJobRequest)(nil),       // 1: InitJobRequest
	(*InitJobResponse)(nil),      // 2: InitJobResponse
	(*StartJobRequest)(nil),      // 3: StartJobRequest
	(*JobConfig)(nil),            // 4: JobConfig
	(*JobStatusReponseList)(nil), // 5: JobStatusReponseList
	(*JobStatusResponse)(nil),    // 6: JobStatusResponse
	(*JobResultResponse)(nil),    // 7: JobResultResponse
	(*JobStatusRequestList)(nil), // 8: JobStatusRequestList
	(*JobAuth)(nil),              // 9: JobAuth
	(*Empty)(nil),                // 10: Empty
}
var file_proto_api_Jobs_proto_depIdxs = []int32{
	9,  // 0: InitJobResponse.job:type_name -> JobAuth
	4,  // 1: StartJobRequest.confiig:type_name -> JobConfig
	6,  // 2: JobStatusReponseList.jobs:type_name -> JobStatusResponse
	0,  // 3: JobStatusResponse.jobStatus:type_name -> JobStatusEnum
	9,  // 4: JobStatusRequestList.jobs:type_name -> JobAuth
	1,  // 5: BaktaJobs.InitJob:input_type -> InitJobRequest
	9,  // 6: BaktaJobs.StartJob:input_type -> JobAuth
	8,  // 7: BaktaJobs.GetJobsStatus:input_type -> JobStatusRequestList
	9,  // 8: BaktaJobs.GetJobResult:input_type -> JobAuth
	2,  // 9: BaktaJobs.InitJob:output_type -> InitJobResponse
	10, // 10: BaktaJobs.StartJob:output_type -> Empty
	5,  // 11: BaktaJobs.GetJobsStatus:output_type -> JobStatusReponseList
	7,  // 12: BaktaJobs.GetJobResult:output_type -> JobResultResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_api_Jobs_proto_init() }
func file_proto_api_Jobs_proto_init() {
	if File_proto_api_Jobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_Jobs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitJobRequest); i {
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
		file_proto_api_Jobs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitJobResponse); i {
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
		file_proto_api_Jobs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartJobRequest); i {
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
		file_proto_api_Jobs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobConfig); i {
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
		file_proto_api_Jobs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatusReponseList); i {
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
		file_proto_api_Jobs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatusResponse); i {
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
		file_proto_api_Jobs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResultResponse); i {
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
		file_proto_api_Jobs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatusRequestList); i {
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
		file_proto_api_Jobs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobAuth); i {
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
		file_proto_api_Jobs_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_api_Jobs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_Jobs_proto_goTypes,
		DependencyIndexes: file_proto_api_Jobs_proto_depIdxs,
		EnumInfos:         file_proto_api_Jobs_proto_enumTypes,
		MessageInfos:      file_proto_api_Jobs_proto_msgTypes,
	}.Build()
	File_proto_api_Jobs_proto = out.File
	file_proto_api_Jobs_proto_rawDesc = nil
	file_proto_api_Jobs_proto_goTypes = nil
	file_proto_api_Jobs_proto_depIdxs = nil
}
