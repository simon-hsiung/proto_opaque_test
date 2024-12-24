// Protocol Buffers - Google's data interchange format
// Copyright 2023 Google Inc.  All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: google/protobuf/go_features.proto

package gofeaturespb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GoFeatures_APILevel int32

const (
	// API_LEVEL_UNSPECIFIED results in selecting the OPEN API,
	// but needs to be a separate value to distinguish between
	// an explicitly set api level or a missing api level.
	GoFeatures_API_LEVEL_UNSPECIFIED GoFeatures_APILevel = 0
	GoFeatures_API_OPEN              GoFeatures_APILevel = 1
	GoFeatures_API_HYBRID            GoFeatures_APILevel = 2
	GoFeatures_API_OPAQUE            GoFeatures_APILevel = 3
)

// Enum value maps for GoFeatures_APILevel.
var (
	GoFeatures_APILevel_name = map[int32]string{
		0: "API_LEVEL_UNSPECIFIED",
		1: "API_OPEN",
		2: "API_HYBRID",
		3: "API_OPAQUE",
	}
	GoFeatures_APILevel_value = map[string]int32{
		"API_LEVEL_UNSPECIFIED": 0,
		"API_OPEN":              1,
		"API_HYBRID":            2,
		"API_OPAQUE":            3,
	}
)

func (x GoFeatures_APILevel) Enum() *GoFeatures_APILevel {
	p := new(GoFeatures_APILevel)
	*p = x
	return p
}

func (x GoFeatures_APILevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoFeatures_APILevel) Descriptor() protoreflect.EnumDescriptor {
	return file_google_protobuf_go_features_proto_enumTypes[0].Descriptor()
}

func (GoFeatures_APILevel) Type() protoreflect.EnumType {
	return &file_google_protobuf_go_features_proto_enumTypes[0]
}

func (x GoFeatures_APILevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GoFeatures_APILevel) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GoFeatures_APILevel(num)
	return nil
}

// Deprecated: Use GoFeatures_APILevel.Descriptor instead.
func (GoFeatures_APILevel) EnumDescriptor() ([]byte, []int) {
	return file_google_protobuf_go_features_proto_rawDescGZIP(), []int{0, 0}
}

type GoFeatures_StripEnumPrefix int32

const (
	GoFeatures_STRIP_ENUM_PREFIX_UNSPECIFIED   GoFeatures_StripEnumPrefix = 0
	GoFeatures_STRIP_ENUM_PREFIX_KEEP          GoFeatures_StripEnumPrefix = 1
	GoFeatures_STRIP_ENUM_PREFIX_GENERATE_BOTH GoFeatures_StripEnumPrefix = 2
	GoFeatures_STRIP_ENUM_PREFIX_STRIP         GoFeatures_StripEnumPrefix = 3
)

// Enum value maps for GoFeatures_StripEnumPrefix.
var (
	GoFeatures_StripEnumPrefix_name = map[int32]string{
		0: "STRIP_ENUM_PREFIX_UNSPECIFIED",
		1: "STRIP_ENUM_PREFIX_KEEP",
		2: "STRIP_ENUM_PREFIX_GENERATE_BOTH",
		3: "STRIP_ENUM_PREFIX_STRIP",
	}
	GoFeatures_StripEnumPrefix_value = map[string]int32{
		"STRIP_ENUM_PREFIX_UNSPECIFIED":   0,
		"STRIP_ENUM_PREFIX_KEEP":          1,
		"STRIP_ENUM_PREFIX_GENERATE_BOTH": 2,
		"STRIP_ENUM_PREFIX_STRIP":         3,
	}
)

func (x GoFeatures_StripEnumPrefix) Enum() *GoFeatures_StripEnumPrefix {
	p := new(GoFeatures_StripEnumPrefix)
	*p = x
	return p
}

func (x GoFeatures_StripEnumPrefix) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoFeatures_StripEnumPrefix) Descriptor() protoreflect.EnumDescriptor {
	return file_google_protobuf_go_features_proto_enumTypes[1].Descriptor()
}

func (GoFeatures_StripEnumPrefix) Type() protoreflect.EnumType {
	return &file_google_protobuf_go_features_proto_enumTypes[1]
}

func (x GoFeatures_StripEnumPrefix) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GoFeatures_StripEnumPrefix) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GoFeatures_StripEnumPrefix(num)
	return nil
}

// Deprecated: Use GoFeatures_StripEnumPrefix.Descriptor instead.
func (GoFeatures_StripEnumPrefix) EnumDescriptor() ([]byte, []int) {
	return file_google_protobuf_go_features_proto_rawDescGZIP(), []int{0, 1}
}

type GoFeatures struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Whether or not to generate the deprecated UnmarshalJSON method for enums.
	// Can only be true for proto using the Open Struct api.
	LegacyUnmarshalJsonEnum *bool `protobuf:"varint,1,opt,name=legacy_unmarshal_json_enum,json=legacyUnmarshalJsonEnum" json:"legacy_unmarshal_json_enum,omitempty"`
	// One of OPEN, HYBRID or OPAQUE.
	ApiLevel        *GoFeatures_APILevel        `protobuf:"varint,2,opt,name=api_level,json=apiLevel,enum=pb.GoFeatures_APILevel" json:"api_level,omitempty"`
	StripEnumPrefix *GoFeatures_StripEnumPrefix `protobuf:"varint,3,opt,name=strip_enum_prefix,json=stripEnumPrefix,enum=pb.GoFeatures_StripEnumPrefix" json:"strip_enum_prefix,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GoFeatures) Reset() {
	*x = GoFeatures{}
	mi := &file_google_protobuf_go_features_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoFeatures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoFeatures) ProtoMessage() {}

func (x *GoFeatures) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_go_features_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoFeatures.ProtoReflect.Descriptor instead.
func (*GoFeatures) Descriptor() ([]byte, []int) {
	return file_google_protobuf_go_features_proto_rawDescGZIP(), []int{0}
}

func (x *GoFeatures) GetLegacyUnmarshalJsonEnum() bool {
	if x != nil && x.LegacyUnmarshalJsonEnum != nil {
		return *x.LegacyUnmarshalJsonEnum
	}
	return false
}

func (x *GoFeatures) GetApiLevel() GoFeatures_APILevel {
	if x != nil && x.ApiLevel != nil {
		return *x.ApiLevel
	}
	return GoFeatures_API_LEVEL_UNSPECIFIED
}

func (x *GoFeatures) GetStripEnumPrefix() GoFeatures_StripEnumPrefix {
	if x != nil && x.StripEnumPrefix != nil {
		return *x.StripEnumPrefix
	}
	return GoFeatures_STRIP_ENUM_PREFIX_UNSPECIFIED
}

var file_google_protobuf_go_features_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FeatureSet)(nil),
		ExtensionType: (*GoFeatures)(nil),
		Field:         1002,
		Name:          "pb.go",
		Tag:           "bytes,1002,opt,name=go",
		Filename:      "google/protobuf/go_features.proto",
	},
}

// Extension fields to descriptorpb.FeatureSet.
var (
	// optional pb.GoFeatures go = 1002;
	E_Go = &file_google_protobuf_go_features_proto_extTypes[0]
)

var File_google_protobuf_go_features_proto protoreflect.FileDescriptor

var file_google_protobuf_go_features_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x05, 0x0a, 0x0a, 0x47, 0x6f,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0xbe, 0x01, 0x0a, 0x1a, 0x6c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x5f, 0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x5f, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x80, 0x01,
	0x88, 0x01, 0x01, 0x98, 0x01, 0x06, 0x98, 0x01, 0x01, 0xa2, 0x01, 0x09, 0x12, 0x04, 0x74, 0x72,
	0x75, 0x65, 0x18, 0x84, 0x07, 0xa2, 0x01, 0x0a, 0x12, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x18,
	0xe7, 0x07, 0xb2, 0x01, 0x5b, 0x08, 0xe8, 0x07, 0x10, 0xe8, 0x07, 0x1a, 0x53, 0x54, 0x68, 0x65,
	0x20, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x20, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61,
	0x6c, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x6c, 0x6c,
	0x20, 0x62, 0x65, 0x20, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x61,
	0x20, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x20, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x17, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61,
	0x6c, 0x4a, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x74, 0x0a, 0x09, 0x61, 0x70, 0x69,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x6f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x41, 0x50, 0x49,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x3e, 0x88, 0x01, 0x01, 0x98, 0x01, 0x03, 0x98, 0x01, 0x01,
	0xa2, 0x01, 0x1a, 0x12, 0x15, 0x41, 0x50, 0x49, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x18, 0x84, 0x07, 0xa2, 0x01, 0x0f,
	0x12, 0x0a, 0x41, 0x50, 0x49, 0x5f, 0x4f, 0x50, 0x41, 0x51, 0x55, 0x45, 0x18, 0xe9, 0x07, 0xb2,
	0x01, 0x03, 0x08, 0xe8, 0x07, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x7c, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x6f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x70,
	0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x30, 0x88, 0x01, 0x01, 0x98,
	0x01, 0x06, 0x98, 0x01, 0x07, 0x98, 0x01, 0x01, 0xa2, 0x01, 0x1b, 0x12, 0x16, 0x53, 0x54, 0x52,
	0x49, 0x50, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x5f, 0x4b,
	0x45, 0x45, 0x50, 0x18, 0x84, 0x07, 0xb2, 0x01, 0x03, 0x08, 0xe9, 0x07, 0x52, 0x0f, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x53, 0x0a,
	0x08, 0x41, 0x50, 0x49, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x50, 0x49,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x49, 0x5f, 0x4f, 0x50, 0x45, 0x4e,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x5f, 0x48, 0x59, 0x42, 0x52, 0x49, 0x44,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x5f, 0x4f, 0x50, 0x41, 0x51, 0x55, 0x45,
	0x10, 0x03, 0x22, 0x92, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x69, 0x70, 0x45, 0x6e, 0x75, 0x6d,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x54, 0x52, 0x49, 0x50, 0x5f,
	0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x52,
	0x49, 0x50, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x5f, 0x4b,
	0x45, 0x45, 0x50, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x54, 0x52, 0x49, 0x50, 0x5f, 0x45,
	0x4e, 0x55, 0x4d, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52,
	0x41, 0x54, 0x45, 0x5f, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54,
	0x52, 0x49, 0x50, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x5f,
	0x53, 0x54, 0x52, 0x49, 0x50, 0x10, 0x03, 0x3a, 0x3c, 0x0a, 0x02, 0x67, 0x6f, 0x12, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x52, 0x02, 0x67, 0x6f, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x70, 0x62,
}

var (
	file_google_protobuf_go_features_proto_rawDescOnce sync.Once
	file_google_protobuf_go_features_proto_rawDescData = file_google_protobuf_go_features_proto_rawDesc
)

func file_google_protobuf_go_features_proto_rawDescGZIP() []byte {
	file_google_protobuf_go_features_proto_rawDescOnce.Do(func() {
		file_google_protobuf_go_features_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_go_features_proto_rawDescData)
	})
	return file_google_protobuf_go_features_proto_rawDescData
}

var file_google_protobuf_go_features_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_protobuf_go_features_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_go_features_proto_goTypes = []any{
	(GoFeatures_APILevel)(0),        // 0: pb.GoFeatures.APILevel
	(GoFeatures_StripEnumPrefix)(0), // 1: pb.GoFeatures.StripEnumPrefix
	(*GoFeatures)(nil),              // 2: pb.GoFeatures
	(*descriptorpb.FeatureSet)(nil), // 3: google.protobuf.FeatureSet
}
var file_google_protobuf_go_features_proto_depIdxs = []int32{
	0, // 0: pb.GoFeatures.api_level:type_name -> pb.GoFeatures.APILevel
	1, // 1: pb.GoFeatures.strip_enum_prefix:type_name -> pb.GoFeatures.StripEnumPrefix
	3, // 2: pb.go:extendee -> google.protobuf.FeatureSet
	2, // 3: pb.go:type_name -> pb.GoFeatures
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_protobuf_go_features_proto_init() }
func file_google_protobuf_go_features_proto_init() {
	if File_google_protobuf_go_features_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_protobuf_go_features_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_go_features_proto_goTypes,
		DependencyIndexes: file_google_protobuf_go_features_proto_depIdxs,
		EnumInfos:         file_google_protobuf_go_features_proto_enumTypes,
		MessageInfos:      file_google_protobuf_go_features_proto_msgTypes,
		ExtensionInfos:    file_google_protobuf_go_features_proto_extTypes,
	}.Build()
	File_google_protobuf_go_features_proto = out.File
	file_google_protobuf_go_features_proto_rawDesc = nil
	file_google_protobuf_go_features_proto_goTypes = nil
	file_google_protobuf_go_features_proto_depIdxs = nil
}
