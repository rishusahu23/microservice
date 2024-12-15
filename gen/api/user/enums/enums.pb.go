// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/user/enums/enums.proto

package enums

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

type UserType int32

const (
	UserType_USER_TYPE_UNSPECIFIED UserType = 0
	UserType_USER_TYPE_MERCHANT    UserType = 1
	UserType_USER_TYPE_CUSTOMER    UserType = 2
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "USER_TYPE_UNSPECIFIED",
		1: "USER_TYPE_MERCHANT",
		2: "USER_TYPE_CUSTOMER",
	}
	UserType_value = map[string]int32{
		"USER_TYPE_UNSPECIFIED": 0,
		"USER_TYPE_MERCHANT":    1,
		"USER_TYPE_CUSTOMER":    2,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_enums_enums_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_api_user_enums_enums_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_api_user_enums_enums_proto_rawDescGZIP(), []int{0}
}

var File_api_user_enums_enums_proto protoreflect.FileDescriptor

var file_api_user_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x55, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x52,
	0x43, 0x48, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x02, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69,
	0x73, 0x68, 0x75, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_enums_enums_proto_rawDescOnce sync.Once
	file_api_user_enums_enums_proto_rawDescData = file_api_user_enums_enums_proto_rawDesc
)

func file_api_user_enums_enums_proto_rawDescGZIP() []byte {
	file_api_user_enums_enums_proto_rawDescOnce.Do(func() {
		file_api_user_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_enums_enums_proto_rawDescData)
	})
	return file_api_user_enums_enums_proto_rawDescData
}

var file_api_user_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_user_enums_enums_proto_goTypes = []interface{}{
	(UserType)(0), // 0: user.enums.UserType
}
var file_api_user_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_enums_enums_proto_init() }
func file_api_user_enums_enums_proto_init() {
	if File_api_user_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_enums_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_enums_enums_proto_goTypes,
		DependencyIndexes: file_api_user_enums_enums_proto_depIdxs,
		EnumInfos:         file_api_user_enums_enums_proto_enumTypes,
	}.Build()
	File_api_user_enums_enums_proto = out.File
	file_api_user_enums_enums_proto_rawDesc = nil
	file_api_user_enums_enums_proto_goTypes = nil
	file_api_user_enums_enums_proto_depIdxs = nil
}
