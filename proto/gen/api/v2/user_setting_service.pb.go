// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v2/user_setting_service.proto

package apiv2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserSetting_Locale int32

const (
	UserSetting_LOCALE_UNSPECIFIED UserSetting_Locale = 0
	UserSetting_LOCALE_EN          UserSetting_Locale = 1
	UserSetting_LOCALE_ZH          UserSetting_Locale = 2
)

// Enum value maps for UserSetting_Locale.
var (
	UserSetting_Locale_name = map[int32]string{
		0: "LOCALE_UNSPECIFIED",
		1: "LOCALE_EN",
		2: "LOCALE_ZH",
	}
	UserSetting_Locale_value = map[string]int32{
		"LOCALE_UNSPECIFIED": 0,
		"LOCALE_EN":          1,
		"LOCALE_ZH":          2,
	}
)

func (x UserSetting_Locale) Enum() *UserSetting_Locale {
	p := new(UserSetting_Locale)
	*p = x
	return p
}

func (x UserSetting_Locale) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserSetting_Locale) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_user_setting_service_proto_enumTypes[0].Descriptor()
}

func (UserSetting_Locale) Type() protoreflect.EnumType {
	return &file_api_v2_user_setting_service_proto_enumTypes[0]
}

func (x UserSetting_Locale) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserSetting_Locale.Descriptor instead.
func (UserSetting_Locale) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_user_setting_service_proto_rawDescGZIP(), []int{0, 0}
}

type UserSetting_ColorTheme int32

const (
	UserSetting_COLOR_THEME_UNSPECIFIED UserSetting_ColorTheme = 0
	UserSetting_COLOR_THEME_SYSTEM      UserSetting_ColorTheme = 1
	UserSetting_COLOR_THEME_LIGHT       UserSetting_ColorTheme = 2
	UserSetting_COLOR_THEME_DARK        UserSetting_ColorTheme = 3
)

// Enum value maps for UserSetting_ColorTheme.
var (
	UserSetting_ColorTheme_name = map[int32]string{
		0: "COLOR_THEME_UNSPECIFIED",
		1: "COLOR_THEME_SYSTEM",
		2: "COLOR_THEME_LIGHT",
		3: "COLOR_THEME_DARK",
	}
	UserSetting_ColorTheme_value = map[string]int32{
		"COLOR_THEME_UNSPECIFIED": 0,
		"COLOR_THEME_SYSTEM":      1,
		"COLOR_THEME_LIGHT":       2,
		"COLOR_THEME_DARK":        3,
	}
)

func (x UserSetting_ColorTheme) Enum() *UserSetting_ColorTheme {
	p := new(UserSetting_ColorTheme)
	*p = x
	return p
}

func (x UserSetting_ColorTheme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserSetting_ColorTheme) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_user_setting_service_proto_enumTypes[1].Descriptor()
}

func (UserSetting_ColorTheme) Type() protoreflect.EnumType {
	return &file_api_v2_user_setting_service_proto_enumTypes[1]
}

func (x UserSetting_ColorTheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserSetting_ColorTheme.Descriptor instead.
func (UserSetting_ColorTheme) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_user_setting_service_proto_rawDescGZIP(), []int{0, 1}
}

type UserSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the user id.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// locale is the user locale.
	Locale UserSetting_Locale `protobuf:"varint,2,opt,name=locale,proto3,enum=slash.api.v2.UserSetting_Locale" json:"locale,omitempty"`
	// color_theme is the user color theme.
	ColorTheme UserSetting_ColorTheme `protobuf:"varint,3,opt,name=color_theme,json=colorTheme,proto3,enum=slash.api.v2.UserSetting_ColorTheme" json:"color_theme,omitempty"`
}

func (x *UserSetting) Reset() {
	*x = UserSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_user_setting_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSetting) ProtoMessage() {}

func (x *UserSetting) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_user_setting_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSetting.ProtoReflect.Descriptor instead.
func (*UserSetting) Descriptor() ([]byte, []int) {
	return file_api_v2_user_setting_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserSetting) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserSetting) GetLocale() UserSetting_Locale {
	if x != nil {
		return x.Locale
	}
	return UserSetting_LOCALE_UNSPECIFIED
}

func (x *UserSetting) GetColorTheme() UserSetting_ColorTheme {
	if x != nil {
		return x.ColorTheme
	}
	return UserSetting_COLOR_THEME_UNSPECIFIED
}

type GetUserSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the user id.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserSettingRequest) Reset() {
	*x = GetUserSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_user_setting_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSettingRequest) ProtoMessage() {}

func (x *GetUserSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_user_setting_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSettingRequest.ProtoReflect.Descriptor instead.
func (*GetUserSettingRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_user_setting_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserSettingRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSetting *UserSetting `protobuf:"bytes,1,opt,name=user_setting,json=userSetting,proto3" json:"user_setting,omitempty"`
}

func (x *GetUserSettingResponse) Reset() {
	*x = GetUserSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_user_setting_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSettingResponse) ProtoMessage() {}

func (x *GetUserSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_user_setting_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSettingResponse.ProtoReflect.Descriptor instead.
func (*GetUserSettingResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_user_setting_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserSettingResponse) GetUserSetting() *UserSetting {
	if x != nil {
		return x.UserSetting
	}
	return nil
}

type UpdateUserSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the user id.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user_setting is the user setting to update.
	UserSetting *UserSetting `protobuf:"bytes,2,opt,name=user_setting,json=userSetting,proto3" json:"user_setting,omitempty"`
	// update_mask is the field mask to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateUserSettingRequest) Reset() {
	*x = UpdateUserSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_user_setting_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserSettingRequest) ProtoMessage() {}

func (x *UpdateUserSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_user_setting_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserSettingRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserSettingRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_user_setting_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserSettingRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserSettingRequest) GetUserSetting() *UserSetting {
	if x != nil {
		return x.UserSetting
	}
	return nil
}

func (x *UpdateUserSettingRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateUserSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSetting *UserSetting `protobuf:"bytes,1,opt,name=user_setting,json=userSetting,proto3" json:"user_setting,omitempty"`
}

func (x *UpdateUserSettingResponse) Reset() {
	*x = UpdateUserSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_user_setting_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserSettingResponse) ProtoMessage() {}

func (x *UpdateUserSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_user_setting_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserSettingResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserSettingResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_user_setting_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserSettingResponse) GetUserSetting() *UserSetting {
	if x != nil {
		return x.UserSetting
	}
	return nil
}

var File_api_v2_user_setting_service_proto protoreflect.FileDescriptor

var file_api_v2_user_setting_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52,
	0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x5f, 0x5a, 0x48, 0x10, 0x02, 0x22, 0x6e, 0x0a, 0x0a, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4c,
	0x4f, 0x52, 0x5f, 0x54, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f,
	0x54, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x54, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x4c, 0x49,
	0x47, 0x48, 0x54, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x54,
	0x48, 0x45, 0x4d, 0x45, 0x5f, 0x44, 0x41, 0x52, 0x4b, 0x10, 0x03, 0x22, 0x27, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xa5, 0x01, 0x0a,
	0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x73, 0x6b, 0x22, 0x59, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x32,
	0xd1, 0x02, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x2e, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0xda, 0x41, 0x02, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0xb2,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73,
	0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0xda, 0x41, 0x18, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x32, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x42, 0xb5, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x17, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x6c, 0x66, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x2f, 0x73,
	0x6c, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x53, 0x41,
	0x58, 0xaa, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32,
	0xca, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0xe2,
	0x02, 0x18, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x53, 0x6c, 0x61,
	0x73, 0x68, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_v2_user_setting_service_proto_rawDescOnce sync.Once
	file_api_v2_user_setting_service_proto_rawDescData = file_api_v2_user_setting_service_proto_rawDesc
)

func file_api_v2_user_setting_service_proto_rawDescGZIP() []byte {
	file_api_v2_user_setting_service_proto_rawDescOnce.Do(func() {
		file_api_v2_user_setting_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_user_setting_service_proto_rawDescData)
	})
	return file_api_v2_user_setting_service_proto_rawDescData
}

var file_api_v2_user_setting_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v2_user_setting_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v2_user_setting_service_proto_goTypes = []interface{}{
	(UserSetting_Locale)(0),           // 0: slash.api.v2.UserSetting.Locale
	(UserSetting_ColorTheme)(0),       // 1: slash.api.v2.UserSetting.ColorTheme
	(*UserSetting)(nil),               // 2: slash.api.v2.UserSetting
	(*GetUserSettingRequest)(nil),     // 3: slash.api.v2.GetUserSettingRequest
	(*GetUserSettingResponse)(nil),    // 4: slash.api.v2.GetUserSettingResponse
	(*UpdateUserSettingRequest)(nil),  // 5: slash.api.v2.UpdateUserSettingRequest
	(*UpdateUserSettingResponse)(nil), // 6: slash.api.v2.UpdateUserSettingResponse
	(*fieldmaskpb.FieldMask)(nil),     // 7: google.protobuf.FieldMask
}
var file_api_v2_user_setting_service_proto_depIdxs = []int32{
	0, // 0: slash.api.v2.UserSetting.locale:type_name -> slash.api.v2.UserSetting.Locale
	1, // 1: slash.api.v2.UserSetting.color_theme:type_name -> slash.api.v2.UserSetting.ColorTheme
	2, // 2: slash.api.v2.GetUserSettingResponse.user_setting:type_name -> slash.api.v2.UserSetting
	2, // 3: slash.api.v2.UpdateUserSettingRequest.user_setting:type_name -> slash.api.v2.UserSetting
	7, // 4: slash.api.v2.UpdateUserSettingRequest.update_mask:type_name -> google.protobuf.FieldMask
	2, // 5: slash.api.v2.UpdateUserSettingResponse.user_setting:type_name -> slash.api.v2.UserSetting
	3, // 6: slash.api.v2.UserSettingService.GetUserSetting:input_type -> slash.api.v2.GetUserSettingRequest
	5, // 7: slash.api.v2.UserSettingService.UpdateUserSetting:input_type -> slash.api.v2.UpdateUserSettingRequest
	4, // 8: slash.api.v2.UserSettingService.GetUserSetting:output_type -> slash.api.v2.GetUserSettingResponse
	6, // 9: slash.api.v2.UserSettingService.UpdateUserSetting:output_type -> slash.api.v2.UpdateUserSettingResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v2_user_setting_service_proto_init() }
func file_api_v2_user_setting_service_proto_init() {
	if File_api_v2_user_setting_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v2_user_setting_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSetting); i {
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
		file_api_v2_user_setting_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSettingRequest); i {
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
		file_api_v2_user_setting_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSettingResponse); i {
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
		file_api_v2_user_setting_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserSettingRequest); i {
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
		file_api_v2_user_setting_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserSettingResponse); i {
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
			RawDescriptor: file_api_v2_user_setting_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_user_setting_service_proto_goTypes,
		DependencyIndexes: file_api_v2_user_setting_service_proto_depIdxs,
		EnumInfos:         file_api_v2_user_setting_service_proto_enumTypes,
		MessageInfos:      file_api_v2_user_setting_service_proto_msgTypes,
	}.Build()
	File_api_v2_user_setting_service_proto = out.File
	file_api_v2_user_setting_service_proto_rawDesc = nil
	file_api_v2_user_setting_service_proto_goTypes = nil
	file_api_v2_user_setting_service_proto_depIdxs = nil
}
