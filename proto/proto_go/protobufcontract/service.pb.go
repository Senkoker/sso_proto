// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: protobufcontract/service.proto

package sso_v1_ssov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Registrequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Registrequest) Reset() {
	*x = Registrequest{}
	mi := &file_protobufcontract_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registrequest) ProtoMessage() {}

func (x *Registrequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registrequest.ProtoReflect.Descriptor instead.
func (*Registrequest) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{0}
}

func (x *Registrequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Registrequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Registresponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Userid        int64                  `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Registresponse) Reset() {
	*x = Registresponse{}
	mi := &file_protobufcontract_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registresponse) ProtoMessage() {}

func (x *Registresponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registresponse.ProtoReflect.Descriptor instead.
func (*Registresponse) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{1}
}

func (x *Registresponse) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type Loginrequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Appid         string                 `protobuf:"bytes,3,opt,name=appid,proto3" json:"appid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Loginrequest) Reset() {
	*x = Loginrequest{}
	mi := &file_protobufcontract_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Loginrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loginrequest) ProtoMessage() {}

func (x *Loginrequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loginrequest.ProtoReflect.Descriptor instead.
func (*Loginrequest) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{2}
}

func (x *Loginrequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Loginrequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Loginrequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

type Loginresponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Loginresponse) Reset() {
	*x = Loginresponse{}
	mi := &file_protobufcontract_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Loginresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loginresponse) ProtoMessage() {}

func (x *Loginresponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loginresponse.ProtoReflect.Descriptor instead.
func (*Loginresponse) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{3}
}

func (x *Loginresponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Acceptrequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Usercode      string                 `protobuf:"bytes,1,opt,name=usercode,proto3" json:"usercode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Acceptrequest) Reset() {
	*x = Acceptrequest{}
	mi := &file_protobufcontract_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Acceptrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acceptrequest) ProtoMessage() {}

func (x *Acceptrequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acceptrequest.ProtoReflect.Descriptor instead.
func (*Acceptrequest) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{4}
}

func (x *Acceptrequest) GetUsercode() string {
	if x != nil {
		return x.Usercode
	}
	return ""
}

type Acceptresponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Accresp       bool                   `protobuf:"varint,1,opt,name=accresp,proto3" json:"accresp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Acceptresponse) Reset() {
	*x = Acceptresponse{}
	mi := &file_protobufcontract_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Acceptresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acceptresponse) ProtoMessage() {}

func (x *Acceptresponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acceptresponse.ProtoReflect.Descriptor instead.
func (*Acceptresponse) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{5}
}

func (x *Acceptresponse) GetAccresp() bool {
	if x != nil {
		return x.Accresp
	}
	return false
}

type Retryrequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Retryrequest) Reset() {
	*x = Retryrequest{}
	mi := &file_protobufcontract_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Retryrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retryrequest) ProtoMessage() {}

func (x *Retryrequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retryrequest.ProtoReflect.Descriptor instead.
func (*Retryrequest) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{6}
}

func (x *Retryrequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Retryresponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Retryresp     bool                   `protobuf:"varint,1,opt,name=retryresp,proto3" json:"retryresp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Retryresponse) Reset() {
	*x = Retryresponse{}
	mi := &file_protobufcontract_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Retryresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retryresponse) ProtoMessage() {}

func (x *Retryresponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retryresponse.ProtoReflect.Descriptor instead.
func (*Retryresponse) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{7}
}

func (x *Retryresponse) GetRetryresp() bool {
	if x != nil {
		return x.Retryresp
	}
	return false
}

type IsAdminrequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsAdminrequest) Reset() {
	*x = IsAdminrequest{}
	mi := &file_protobufcontract_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAdminrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminrequest) ProtoMessage() {}

func (x *IsAdminrequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminrequest.ProtoReflect.Descriptor instead.
func (*IsAdminrequest) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{8}
}

func (x *IsAdminrequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type IsAdminresponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Adminresp     bool                   `protobuf:"varint,1,opt,name=adminresp,proto3" json:"adminresp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsAdminresponse) Reset() {
	*x = IsAdminresponse{}
	mi := &file_protobufcontract_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAdminresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminresponse) ProtoMessage() {}

func (x *IsAdminresponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobufcontract_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminresponse.ProtoReflect.Descriptor instead.
func (*IsAdminresponse) Descriptor() ([]byte, []int) {
	return file_protobufcontract_service_proto_rawDescGZIP(), []int{9}
}

func (x *IsAdminresponse) GetAdminresp() bool {
	if x != nil {
		return x.Adminresp
	}
	return false
}

var File_protobufcontract_service_proto protoreflect.FileDescriptor

var file_protobufcontract_service_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x2b, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x2a, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x61, 0x63, 0x63, 0x72, 0x65, 0x73, 0x70, 0x22, 0x1e, 0x0a, 0x0c, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0d, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x72, 0x65, 0x73, 0x70, 0x22, 0x26, 0x0a, 0x0e, 0x49, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x2f, 0x0a, 0x0f, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x72, 0x65,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x72,
	0x65, 0x73, 0x70, 0x32, 0x8e, 0x02, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x52, 0x65, 0x74, 0x72, 0x79, 0x12, 0x12,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07,
	0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x3a, 0x73,
	0x73, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protobufcontract_service_proto_rawDescOnce sync.Once
	file_protobufcontract_service_proto_rawDescData []byte
)

func file_protobufcontract_service_proto_rawDescGZIP() []byte {
	file_protobufcontract_service_proto_rawDescOnce.Do(func() {
		file_protobufcontract_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protobufcontract_service_proto_rawDesc), len(file_protobufcontract_service_proto_rawDesc)))
	})
	return file_protobufcontract_service_proto_rawDescData
}

var file_protobufcontract_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protobufcontract_service_proto_goTypes = []any{
	(*Registrequest)(nil),   // 0: auth.Registrequest
	(*Registresponse)(nil),  // 1: auth.Registresponse
	(*Loginrequest)(nil),    // 2: auth.Loginrequest
	(*Loginresponse)(nil),   // 3: auth.Loginresponse
	(*Acceptrequest)(nil),   // 4: auth.Acceptrequest
	(*Acceptresponse)(nil),  // 5: auth.Acceptresponse
	(*Retryrequest)(nil),    // 6: auth.Retryrequest
	(*Retryresponse)(nil),   // 7: auth.Retryresponse
	(*IsAdminrequest)(nil),  // 8: auth.IsAdminrequest
	(*IsAdminresponse)(nil), // 9: auth.IsAdminresponse
}
var file_protobufcontract_service_proto_depIdxs = []int32{
	0, // 0: auth.Auth.Register:input_type -> auth.Registrequest
	2, // 1: auth.Auth.Login:input_type -> auth.Loginrequest
	6, // 2: auth.Auth.Retry:input_type -> auth.Retryrequest
	4, // 3: auth.Auth.Accept:input_type -> auth.Acceptrequest
	8, // 4: auth.Auth.IsAdmin:input_type -> auth.IsAdminrequest
	1, // 5: auth.Auth.Register:output_type -> auth.Registresponse
	3, // 6: auth.Auth.Login:output_type -> auth.Loginresponse
	7, // 7: auth.Auth.Retry:output_type -> auth.Retryresponse
	5, // 8: auth.Auth.Accept:output_type -> auth.Acceptresponse
	9, // 9: auth.Auth.IsAdmin:output_type -> auth.IsAdminresponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobufcontract_service_proto_init() }
func file_protobufcontract_service_proto_init() {
	if File_protobufcontract_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protobufcontract_service_proto_rawDesc), len(file_protobufcontract_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobufcontract_service_proto_goTypes,
		DependencyIndexes: file_protobufcontract_service_proto_depIdxs,
		MessageInfos:      file_protobufcontract_service_proto_msgTypes,
	}.Build()
	File_protobufcontract_service_proto = out.File
	file_protobufcontract_service_proto_goTypes = nil
	file_protobufcontract_service_proto_depIdxs = nil
}
