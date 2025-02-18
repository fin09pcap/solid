// Licensed to SolID under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. SolID licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: oidc/core/v1/session.proto

package corev1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeviceCodeStatus int32

const (
	DeviceCodeStatus_DEVICE_CODE_STATUS_INVALID               DeviceCodeStatus = 0
	DeviceCodeStatus_DEVICE_CODE_STATUS_UNKNOWN               DeviceCodeStatus = 1
	DeviceCodeStatus_DEVICE_CODE_STATUS_AUTHORIZATION_PENDING DeviceCodeStatus = 2
	DeviceCodeStatus_DEVICE_CODE_STATUS_VALIDATED             DeviceCodeStatus = 3
)

// Enum value maps for DeviceCodeStatus.
var (
	DeviceCodeStatus_name = map[int32]string{
		0: "DEVICE_CODE_STATUS_INVALID",
		1: "DEVICE_CODE_STATUS_UNKNOWN",
		2: "DEVICE_CODE_STATUS_AUTHORIZATION_PENDING",
		3: "DEVICE_CODE_STATUS_VALIDATED",
	}
	DeviceCodeStatus_value = map[string]int32{
		"DEVICE_CODE_STATUS_INVALID":               0,
		"DEVICE_CODE_STATUS_UNKNOWN":               1,
		"DEVICE_CODE_STATUS_AUTHORIZATION_PENDING": 2,
		"DEVICE_CODE_STATUS_VALIDATED":             3,
	}
)

func (x DeviceCodeStatus) Enum() *DeviceCodeStatus {
	p := new(DeviceCodeStatus)
	*p = x
	return p
}

func (x DeviceCodeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceCodeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_oidc_core_v1_session_proto_enumTypes[0].Descriptor()
}

func (DeviceCodeStatus) Type() protoreflect.EnumType {
	return &file_oidc_core_v1_session_proto_enumTypes[0]
}

func (x DeviceCodeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceCodeStatus.Descriptor instead.
func (DeviceCodeStatus) EnumDescriptor() ([]byte, []int) {
	return file_oidc_core_v1_session_proto_rawDescGZIP(), []int{0}
}

type AuthorizationCodeSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client  *Client               `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Issuer  string                `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Request *AuthorizationRequest `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Subject string                `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *AuthorizationCodeSession) Reset() {
	*x = AuthorizationCodeSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_core_v1_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationCodeSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationCodeSession) ProtoMessage() {}

func (x *AuthorizationCodeSession) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_core_v1_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationCodeSession.ProtoReflect.Descriptor instead.
func (*AuthorizationCodeSession) Descriptor() ([]byte, []int) {
	return file_oidc_core_v1_session_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizationCodeSession) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *AuthorizationCodeSession) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *AuthorizationCodeSession) GetRequest() *AuthorizationRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *AuthorizationCodeSession) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type DeviceCodeSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client     *Client                     `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Issuer     string                      `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Request    *DeviceAuthorizationRequest `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	DeviceCode string                      `protobuf:"bytes,4,opt,name=device_code,json=deviceCode,proto3" json:"device_code,omitempty"`
	UserCode   string                      `protobuf:"bytes,5,opt,name=user_code,json=userCode,proto3" json:"user_code,omitempty"`
	ExpiresAt  uint64                      `protobuf:"fixed64,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Status     DeviceCodeStatus            `protobuf:"varint,7,opt,name=status,proto3,enum=oidc.core.v1.DeviceCodeStatus" json:"status,omitempty"`
	Scope      string                      `protobuf:"bytes,8,opt,name=scope,proto3" json:"scope,omitempty"`
	Audience   string                      `protobuf:"bytes,9,opt,name=audience,proto3" json:"audience,omitempty"`
	Subject    string                      `protobuf:"bytes,10,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *DeviceCodeSession) Reset() {
	*x = DeviceCodeSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_core_v1_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCodeSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCodeSession) ProtoMessage() {}

func (x *DeviceCodeSession) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_core_v1_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCodeSession.ProtoReflect.Descriptor instead.
func (*DeviceCodeSession) Descriptor() ([]byte, []int) {
	return file_oidc_core_v1_session_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceCodeSession) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *DeviceCodeSession) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *DeviceCodeSession) GetRequest() *DeviceAuthorizationRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *DeviceCodeSession) GetDeviceCode() string {
	if x != nil {
		return x.DeviceCode
	}
	return ""
}

func (x *DeviceCodeSession) GetUserCode() string {
	if x != nil {
		return x.UserCode
	}
	return ""
}

func (x *DeviceCodeSession) GetExpiresAt() uint64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *DeviceCodeSession) GetStatus() DeviceCodeStatus {
	if x != nil {
		return x.Status
	}
	return DeviceCodeStatus_DEVICE_CODE_STATUS_INVALID
}

func (x *DeviceCodeSession) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *DeviceCodeSession) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *DeviceCodeSession) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

var File_oidc_core_v1_session_proto protoreflect.FileDescriptor

var file_oidc_core_v1_session_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x69,
	0x64, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x6f, 0x69, 0x64, 0x63,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x18,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x3c,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xfe, 0x02, 0x0a, 0x11, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f,
	0x69, 0x64, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x12, 0x42, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x06, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2a, 0xa2, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x2c, 0x0a, 0x28,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x45,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x15, 0x5a, 0x13,
	0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72,
	0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oidc_core_v1_session_proto_rawDescOnce sync.Once
	file_oidc_core_v1_session_proto_rawDescData = file_oidc_core_v1_session_proto_rawDesc
)

func file_oidc_core_v1_session_proto_rawDescGZIP() []byte {
	file_oidc_core_v1_session_proto_rawDescOnce.Do(func() {
		file_oidc_core_v1_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidc_core_v1_session_proto_rawDescData)
	})
	return file_oidc_core_v1_session_proto_rawDescData
}

var (
	file_oidc_core_v1_session_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_oidc_core_v1_session_proto_msgTypes  = make([]protoimpl.MessageInfo, 2)
	file_oidc_core_v1_session_proto_goTypes   = []interface{}{
		(DeviceCodeStatus)(0),              // 0: oidc.core.v1.DeviceCodeStatus
		(*AuthorizationCodeSession)(nil),   // 1: oidc.core.v1.AuthorizationCodeSession
		(*DeviceCodeSession)(nil),          // 2: oidc.core.v1.DeviceCodeSession
		(*Client)(nil),                     // 3: oidc.core.v1.Client
		(*AuthorizationRequest)(nil),       // 4: oidc.core.v1.AuthorizationRequest
		(*DeviceAuthorizationRequest)(nil), // 5: oidc.core.v1.DeviceAuthorizationRequest
	}
)

var file_oidc_core_v1_session_proto_depIdxs = []int32{
	3, // 0: oidc.core.v1.AuthorizationCodeSession.client:type_name -> oidc.core.v1.Client
	4, // 1: oidc.core.v1.AuthorizationCodeSession.request:type_name -> oidc.core.v1.AuthorizationRequest
	3, // 2: oidc.core.v1.DeviceCodeSession.client:type_name -> oidc.core.v1.Client
	5, // 3: oidc.core.v1.DeviceCodeSession.request:type_name -> oidc.core.v1.DeviceAuthorizationRequest
	0, // 4: oidc.core.v1.DeviceCodeSession.status:type_name -> oidc.core.v1.DeviceCodeStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_oidc_core_v1_session_proto_init() }
func file_oidc_core_v1_session_proto_init() {
	if File_oidc_core_v1_session_proto != nil {
		return
	}
	file_oidc_core_v1_core_proto_init()
	file_oidc_core_v1_core_api_proto_init()
	file_oidc_core_v1_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_oidc_core_v1_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationCodeSession); i {
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
		file_oidc_core_v1_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCodeSession); i {
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
			RawDescriptor: file_oidc_core_v1_session_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidc_core_v1_session_proto_goTypes,
		DependencyIndexes: file_oidc_core_v1_session_proto_depIdxs,
		EnumInfos:         file_oidc_core_v1_session_proto_enumTypes,
		MessageInfos:      file_oidc_core_v1_session_proto_msgTypes,
	}.Build()
	File_oidc_core_v1_session_proto = out.File
	file_oidc_core_v1_session_proto_rawDesc = nil
	file_oidc_core_v1_session_proto_goTypes = nil
	file_oidc_core_v1_session_proto_depIdxs = nil
}
