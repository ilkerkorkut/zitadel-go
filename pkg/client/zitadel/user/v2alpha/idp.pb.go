// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: zitadel/user/v2alpha/idp.proto

package user

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IDPInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Access:
	//
	//	*IDPInformation_Oauth
	Access         isIDPInformation_Access `protobuf_oneof:"access"`
	IdpId          string                  `protobuf:"bytes,2,opt,name=idp_id,json=idpId,proto3" json:"idp_id,omitempty"`
	UserId         string                  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName       string                  `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RawInformation *structpb.Struct        `protobuf:"bytes,5,opt,name=raw_information,json=rawInformation,proto3" json:"raw_information,omitempty"`
}

func (x *IDPInformation) Reset() {
	*x = IDPInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2alpha_idp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPInformation) ProtoMessage() {}

func (x *IDPInformation) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2alpha_idp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPInformation.ProtoReflect.Descriptor instead.
func (*IDPInformation) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2alpha_idp_proto_rawDescGZIP(), []int{0}
}

func (m *IDPInformation) GetAccess() isIDPInformation_Access {
	if m != nil {
		return m.Access
	}
	return nil
}

func (x *IDPInformation) GetOauth() *IDPOAuthAccessInformation {
	if x, ok := x.GetAccess().(*IDPInformation_Oauth); ok {
		return x.Oauth
	}
	return nil
}

func (x *IDPInformation) GetIdpId() string {
	if x != nil {
		return x.IdpId
	}
	return ""
}

func (x *IDPInformation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IDPInformation) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *IDPInformation) GetRawInformation() *structpb.Struct {
	if x != nil {
		return x.RawInformation
	}
	return nil
}

type isIDPInformation_Access interface {
	isIDPInformation_Access()
}

type IDPInformation_Oauth struct {
	Oauth *IDPOAuthAccessInformation `protobuf:"bytes,1,opt,name=oauth,proto3,oneof"`
}

func (*IDPInformation_Oauth) isIDPInformation_Access() {}

type IDPOAuthAccessInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string  `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	IdToken     *string `protobuf:"bytes,2,opt,name=id_token,json=idToken,proto3,oneof" json:"id_token,omitempty"`
}

func (x *IDPOAuthAccessInformation) Reset() {
	*x = IDPOAuthAccessInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2alpha_idp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPOAuthAccessInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPOAuthAccessInformation) ProtoMessage() {}

func (x *IDPOAuthAccessInformation) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2alpha_idp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPOAuthAccessInformation.ProtoReflect.Descriptor instead.
func (*IDPOAuthAccessInformation) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2alpha_idp_proto_rawDescGZIP(), []int{1}
}

func (x *IDPOAuthAccessInformation) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *IDPOAuthAccessInformation) GetIdToken() string {
	if x != nil && x.IdToken != nil {
		return *x.IdToken
	}
	return ""
}

type IDPLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdpId    string `protobuf:"bytes,1,opt,name=idp_id,json=idpId,proto3" json:"idp_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *IDPLink) Reset() {
	*x = IDPLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2alpha_idp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPLink) ProtoMessage() {}

func (x *IDPLink) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2alpha_idp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPLink.ProtoReflect.Descriptor instead.
func (*IDPLink) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2alpha_idp_proto_rawDescGZIP(), []int{2}
}

func (x *IDPLink) GetIdpId() string {
	if x != nil {
		return x.IdpId
	}
	return ""
}

func (x *IDPLink) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IDPLink) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

var File_zitadel_user_v2alpha_idp_proto protoreflect.FileDescriptor

var file_zitadel_user_v2alpha_idp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda,
	0x04, 0x0a, 0x0e, 0x49, 0x44, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x90, 0x01, 0x0a, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49, 0x44, 0x50, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x47, 0x92, 0x41, 0x44, 0x32, 0x42, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x2f, 0x4f,
	0x49, 0x44, 0x43, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x28, 0x61, 0x6e, 0x64, 0x20,
	0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x29, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x12, 0x5f, 0x0a, 0x06, 0x69, 0x64, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x48, 0x92, 0x41, 0x45, 0x32, 0x1b, 0x49, 0x44, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4a, 0x26, 0x22, 0x64, 0x36, 0x35, 0x34, 0x65, 0x36, 0x62,
	0x61, 0x2d, 0x37, 0x30, 0x61, 0x33, 0x2d, 0x34, 0x38, 0x65, 0x66, 0x2d, 0x61, 0x39, 0x35, 0x64,
	0x2d, 0x33, 0x37, 0x63, 0x38, 0x64, 0x38, 0x61, 0x37, 0x39, 0x30, 0x31, 0x61, 0x22, 0x52, 0x05,
	0x69, 0x64, 0x70, 0x49, 0x64, 0x12, 0x65, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4c, 0x92, 0x41, 0x49, 0x32, 0x27, 0x49, 0x44, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x4a, 0x1e, 0x22, 0x36, 0x35, 0x31, 0x36, 0x38, 0x34, 0x39, 0x38, 0x30,
	0x34, 0x38, 0x39, 0x30, 0x34, 0x36, 0x38, 0x30, 0x34, 0x38, 0x34, 0x36, 0x31, 0x34, 0x30, 0x33,
	0x35, 0x31, 0x38, 0x22, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x64, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x47, 0x92, 0x41, 0x44, 0x32, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x4a, 0x13, 0x22, 0x75, 0x73, 0x65, 0x72, 0x40, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x7d, 0x0a, 0x0f, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x42, 0x3b, 0x92, 0x41, 0x38, 0x32, 0x36, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x0e, 0x72, 0x61, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x6b, 0x0a, 0x19, 0x49,
	0x44, 0x50, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x08, 0x69,
	0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xe7, 0x02, 0x0a, 0x07, 0x49, 0x44, 0x50,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x6f, 0x0a, 0x06, 0x69, 0x64, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x58, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01,
	0x92, 0x41, 0x4b, 0x32, 0x1b, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x4a, 0x26, 0x22, 0x64, 0x36, 0x35, 0x34, 0x65, 0x36, 0x62, 0x61, 0x2d, 0x37, 0x30, 0x61, 0x33,
	0x2d, 0x34, 0x38, 0x65, 0x66, 0x2d, 0x61, 0x39, 0x35, 0x64, 0x2d, 0x33, 0x37, 0x63, 0x38, 0x64,
	0x38, 0x61, 0x37, 0x39, 0x30, 0x31, 0x61, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0x52, 0x05,
	0x69, 0x64, 0x70, 0x49, 0x64, 0x12, 0x75, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x5c, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0xc8, 0x01, 0x92, 0x41, 0x4f, 0x32, 0x27, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4a, 0x1e,
	0x22, 0x36, 0x35, 0x31, 0x36, 0x38, 0x34, 0x39, 0x38, 0x30, 0x34, 0x38, 0x39, 0x30, 0x34, 0x36,
	0x38, 0x30, 0x34, 0x38, 0x34, 0x36, 0x31, 0x34, 0x30, 0x33, 0x35, 0x31, 0x38, 0x22, 0x78, 0xc8,
	0x01, 0x80, 0x01, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x74, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x57, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x92, 0x41, 0x4a, 0x32, 0x2d,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4a, 0x13, 0x22,
	0x75, 0x73, 0x65, 0x72, 0x40, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f,
	0x6d, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_zitadel_user_v2alpha_idp_proto_rawDescOnce sync.Once
	file_zitadel_user_v2alpha_idp_proto_rawDescData = file_zitadel_user_v2alpha_idp_proto_rawDesc
)

func file_zitadel_user_v2alpha_idp_proto_rawDescGZIP() []byte {
	file_zitadel_user_v2alpha_idp_proto_rawDescOnce.Do(func() {
		file_zitadel_user_v2alpha_idp_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_user_v2alpha_idp_proto_rawDescData)
	})
	return file_zitadel_user_v2alpha_idp_proto_rawDescData
}

var file_zitadel_user_v2alpha_idp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_zitadel_user_v2alpha_idp_proto_goTypes = []interface{}{
	(*IDPInformation)(nil),            // 0: zitadel.user.v2alpha.IDPInformation
	(*IDPOAuthAccessInformation)(nil), // 1: zitadel.user.v2alpha.IDPOAuthAccessInformation
	(*IDPLink)(nil),                   // 2: zitadel.user.v2alpha.IDPLink
	(*structpb.Struct)(nil),           // 3: google.protobuf.Struct
}
var file_zitadel_user_v2alpha_idp_proto_depIdxs = []int32{
	1, // 0: zitadel.user.v2alpha.IDPInformation.oauth:type_name -> zitadel.user.v2alpha.IDPOAuthAccessInformation
	3, // 1: zitadel.user.v2alpha.IDPInformation.raw_information:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_zitadel_user_v2alpha_idp_proto_init() }
func file_zitadel_user_v2alpha_idp_proto_init() {
	if File_zitadel_user_v2alpha_idp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_user_v2alpha_idp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPInformation); i {
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
		file_zitadel_user_v2alpha_idp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPOAuthAccessInformation); i {
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
		file_zitadel_user_v2alpha_idp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPLink); i {
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
	file_zitadel_user_v2alpha_idp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*IDPInformation_Oauth)(nil),
	}
	file_zitadel_user_v2alpha_idp_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_user_v2alpha_idp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_user_v2alpha_idp_proto_goTypes,
		DependencyIndexes: file_zitadel_user_v2alpha_idp_proto_depIdxs,
		MessageInfos:      file_zitadel_user_v2alpha_idp_proto_msgTypes,
	}.Build()
	File_zitadel_user_v2alpha_idp_proto = out.File
	file_zitadel_user_v2alpha_idp_proto_rawDesc = nil
	file_zitadel_user_v2alpha_idp_proto_goTypes = nil
	file_zitadel_user_v2alpha_idp_proto_depIdxs = nil
}
