// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: stakeholders_service.proto

package auth

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

type RequestLogIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RequestLogIn) Reset() {
	*x = RequestLogIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLogIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogIn) ProtoMessage() {}

func (x *RequestLogIn) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogIn.ProtoReflect.Descriptor instead.
func (*RequestLogIn) Descriptor() ([]byte, []int) {
	return file_stakeholders_service_proto_rawDescGZIP(), []int{0}
}

func (x *RequestLogIn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RequestLogIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ResponseLogIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	AccessToken string `protobuf:"bytes,4,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
}

func (x *ResponseLogIn) Reset() {
	*x = ResponseLogIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLogIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLogIn) ProtoMessage() {}

func (x *ResponseLogIn) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLogIn.ProtoReflect.Descriptor instead.
func (*ResponseLogIn) Descriptor() ([]byte, []int) {
	return file_stakeholders_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseLogIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResponseLogIn) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RequestRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Email    string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Name     string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Surname  string `protobuf:"bytes,9,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *RequestRegister) Reset() {
	*x = RequestRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRegister) ProtoMessage() {}

func (x *RequestRegister) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRegister.ProtoReflect.Descriptor instead.
func (*RequestRegister) Descriptor() ([]byte, []int) {
	return file_stakeholders_service_proto_rawDescGZIP(), []int{2}
}

func (x *RequestRegister) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RequestRegister) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RequestRegister) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestRegister) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestRegister) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

type RequestActivateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,10,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RequestActivateUser) Reset() {
	*x = RequestActivateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestActivateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestActivateUser) ProtoMessage() {}

func (x *RequestActivateUser) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestActivateUser.ProtoReflect.Descriptor instead.
func (*RequestActivateUser) Descriptor() ([]byte, []int) {
	return file_stakeholders_service_proto_rawDescGZIP(), []int{3}
}

func (x *RequestActivateUser) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RequestChangePassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPassword        string `protobuf:"bytes,11,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	NewPasswordConfirm string `protobuf:"bytes,12,opt,name=newPasswordConfirm,proto3" json:"newPasswordConfirm,omitempty"`
	Token              string `protobuf:"bytes,13,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RequestChangePassword) Reset() {
	*x = RequestChangePassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChangePassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChangePassword) ProtoMessage() {}

func (x *RequestChangePassword) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChangePassword.ProtoReflect.Descriptor instead.
func (*RequestChangePassword) Descriptor() ([]byte, []int) {
	return file_stakeholders_service_proto_rawDescGZIP(), []int{4}
}

func (x *RequestChangePassword) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *RequestChangePassword) GetNewPasswordConfirm() string {
	if x != nil {
		return x.NewPasswordConfirm
	}
	return ""
}

func (x *RequestChangePassword) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RequestChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,14,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RequestChangePasswordRequest) Reset() {
	*x = RequestChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChangePasswordRequest) ProtoMessage() {}

func (x *RequestChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*RequestChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_stakeholders_service_proto_rawDescGZIP(), []int{5}
}

func (x *RequestChangePasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_stakeholders_service_proto protoreflect.FileDescriptor

var file_stakeholders_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x7f, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2e, 0x0a,
	0x12, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x65, 0x77, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x34, 0x0a, 0x1c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xfb, 0x03, 0x0a, 0x12, 0x53, 0x74,
	0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x12, 0x0d, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x4a, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x58, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a,
	0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x71, 0x0a, 0x0e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0x31, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x86,
	0x01, 0x0a, 0x15, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0x38, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stakeholders_service_proto_rawDescOnce sync.Once
	file_stakeholders_service_proto_rawDescData = file_stakeholders_service_proto_rawDesc
)

func file_stakeholders_service_proto_rawDescGZIP() []byte {
	file_stakeholders_service_proto_rawDescOnce.Do(func() {
		file_stakeholders_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stakeholders_service_proto_rawDescData)
	})
	return file_stakeholders_service_proto_rawDescData
}

var file_stakeholders_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stakeholders_service_proto_goTypes = []interface{}{
	(*RequestLogIn)(nil),                 // 0: RequestLogIn
	(*ResponseLogIn)(nil),                // 1: ResponseLogIn
	(*RequestRegister)(nil),              // 2: RequestRegister
	(*RequestActivateUser)(nil),          // 3: RequestActivateUser
	(*RequestChangePassword)(nil),        // 4: RequestChangePassword
	(*RequestChangePasswordRequest)(nil), // 5: RequestChangePasswordRequest
}
var file_stakeholders_service_proto_depIdxs = []int32{
	0, // 0: StakeholderService.LogIn:input_type -> RequestLogIn
	2, // 1: StakeholderService.RegisterTourist:input_type -> RequestRegister
	3, // 2: StakeholderService.ActivateUser:input_type -> RequestActivateUser
	4, // 3: StakeholderService.ChangePassword:input_type -> RequestChangePassword
	5, // 4: StakeholderService.changePasswordRequest:input_type -> RequestChangePasswordRequest
	1, // 5: StakeholderService.LogIn:output_type -> ResponseLogIn
	1, // 6: StakeholderService.RegisterTourist:output_type -> ResponseLogIn
	1, // 7: StakeholderService.ActivateUser:output_type -> ResponseLogIn
	3, // 8: StakeholderService.ChangePassword:output_type -> RequestActivateUser
	3, // 9: StakeholderService.changePasswordRequest:output_type -> RequestActivateUser
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stakeholders_service_proto_init() }
func file_stakeholders_service_proto_init() {
	if File_stakeholders_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stakeholders_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLogIn); i {
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
		file_stakeholders_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLogIn); i {
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
		file_stakeholders_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRegister); i {
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
		file_stakeholders_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestActivateUser); i {
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
		file_stakeholders_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChangePassword); i {
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
		file_stakeholders_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChangePasswordRequest); i {
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
			RawDescriptor: file_stakeholders_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stakeholders_service_proto_goTypes,
		DependencyIndexes: file_stakeholders_service_proto_depIdxs,
		MessageInfos:      file_stakeholders_service_proto_msgTypes,
	}.Build()
	File_stakeholders_service_proto = out.File
	file_stakeholders_service_proto_rawDesc = nil
	file_stakeholders_service_proto_goTypes = nil
	file_stakeholders_service_proto_depIdxs = nil
}
