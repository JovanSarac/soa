// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: followings/follower_service.proto

package followings

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

type People struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ud         int64   `protobuf:"varint,1,opt,name=ud,proto3" json:"ud,omitempty"`
	UserId     int64   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name       string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string  `protobuf:"bytes,4,opt,name=surname,proto3" json:"surname,omitempty"`
	Email      string  `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	ProfilePic string  `protobuf:"bytes,6,opt,name=profile_pic,json=profilePic,proto3" json:"profile_pic,omitempty"`
	Biography  string  `protobuf:"bytes,7,opt,name=biography,proto3" json:"biography,omitempty"`
	Motto      string  `protobuf:"bytes,8,opt,name=motto,proto3" json:"motto,omitempty"`
	Latitude   float64 `protobuf:"fixed64,9,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude  float64 `protobuf:"fixed64,10,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *People) Reset() {
	*x = People{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followings_follower_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *People) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*People) ProtoMessage() {}

func (x *People) ProtoReflect() protoreflect.Message {
	mi := &file_followings_follower_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use People.ProtoReflect.Descriptor instead.
func (*People) Descriptor() ([]byte, []int) {
	return file_followings_follower_service_proto_rawDescGZIP(), []int{0}
}

func (x *People) GetUd() int64 {
	if x != nil {
		return x.Ud
	}
	return 0
}

func (x *People) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *People) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *People) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *People) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *People) GetProfilePic() string {
	if x != nil {
		return x.ProfilePic
	}
	return ""
}

func (x *People) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *People) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *People) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *People) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GetFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFollowRequest) Reset() {
	*x = GetFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followings_follower_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowRequest) ProtoMessage() {}

func (x *GetFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followings_follower_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowRequest.ProtoReflect.Descriptor instead.
func (*GetFollowRequest) Descriptor() ([]byte, []int) {
	return file_followings_follower_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetFollowRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*People `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *GetFollowResponse) Reset() {
	*x = GetFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followings_follower_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowResponse) ProtoMessage() {}

func (x *GetFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followings_follower_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowResponse.ProtoReflect.Descriptor instead.
func (*GetFollowResponse) Descriptor() ([]byte, []int) {
	return file_followings_follower_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetFollowResponse) GetPeople() []*People {
	if x != nil {
		return x.People
	}
	return nil
}

var File_followings_follower_service_proto protoreflect.FileDescriptor

var file_followings_follower_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x06, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x75, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x75, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x74,
	0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x06, 0x70, 0x65, 0x6f,
	0x70, 0x6c, 0x65, 0x32, 0x4b, 0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_followings_follower_service_proto_rawDescOnce sync.Once
	file_followings_follower_service_proto_rawDescData = file_followings_follower_service_proto_rawDesc
)

func file_followings_follower_service_proto_rawDescGZIP() []byte {
	file_followings_follower_service_proto_rawDescOnce.Do(func() {
		file_followings_follower_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_followings_follower_service_proto_rawDescData)
	})
	return file_followings_follower_service_proto_rawDescData
}

var file_followings_follower_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_followings_follower_service_proto_goTypes = []interface{}{
	(*People)(nil),            // 0: People
	(*GetFollowRequest)(nil),  // 1: GetFollowRequest
	(*GetFollowResponse)(nil), // 2: GetFollowResponse
}
var file_followings_follower_service_proto_depIdxs = []int32{
	0, // 0: GetFollowResponse.people:type_name -> People
	1, // 1: FollowerService.GetFollowings:input_type -> GetFollowRequest
	2, // 2: FollowerService.GetFollowings:output_type -> GetFollowResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_followings_follower_service_proto_init() }
func file_followings_follower_service_proto_init() {
	if File_followings_follower_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_followings_follower_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*People); i {
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
		file_followings_follower_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowRequest); i {
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
		file_followings_follower_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowResponse); i {
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
			RawDescriptor: file_followings_follower_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_followings_follower_service_proto_goTypes,
		DependencyIndexes: file_followings_follower_service_proto_depIdxs,
		MessageInfos:      file_followings_follower_service_proto_msgTypes,
	}.Build()
	File_followings_follower_service_proto = out.File
	file_followings_follower_service_proto_rawDesc = nil
	file_followings_follower_service_proto_goTypes = nil
	file_followings_follower_service_proto_depIdxs = nil
}