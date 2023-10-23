// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: index_platform.proto

package index_platform

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

type BuildIndexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag:form:"file_path" uri:"file_path"
	FilePath []string `protobuf:"bytes,1,rep,name=file_path,json=filePath,proto3" json:"file_path,omitempty" form:"file_path" uri:"file_path"`
}

func (x *BuildIndexReq) Reset() {
	*x = BuildIndexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildIndexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildIndexReq) ProtoMessage() {}

func (x *BuildIndexReq) ProtoReflect() protoreflect.Message {
	mi := &file_index_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildIndexReq.ProtoReflect.Descriptor instead.
func (*BuildIndexReq) Descriptor() ([]byte, []int) {
	return file_index_platform_proto_rawDescGZIP(), []int{0}
}

func (x *BuildIndexReq) GetFilePath() []string {
	if x != nil {
		return x.FilePath
	}
	return nil
}

type BuildIndexResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag:form:"code" uri:"code"
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" uri:"code"`
	// @inject_tag:form:"message" uri:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" form:"message" uri:"message"`
}

func (x *BuildIndexResp) Reset() {
	*x = BuildIndexResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildIndexResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildIndexResp) ProtoMessage() {}

func (x *BuildIndexResp) ProtoReflect() protoreflect.Message {
	mi := &file_index_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildIndexResp.ProtoReflect.Descriptor instead.
func (*BuildIndexResp) Descriptor() ([]byte, []int) {
	return file_index_platform_proto_rawDescGZIP(), []int{1}
}

func (x *BuildIndexResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BuildIndexResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_index_platform_proto protoreflect.FileDescriptor

var file_index_platform_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x22, 0x3e, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x4c, 0x0a, 0x14, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x11,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x0e, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_index_platform_proto_rawDescOnce sync.Once
	file_index_platform_proto_rawDescData = file_index_platform_proto_rawDesc
)

func file_index_platform_proto_rawDescGZIP() []byte {
	file_index_platform_proto_rawDescOnce.Do(func() {
		file_index_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_index_platform_proto_rawDescData)
	})
	return file_index_platform_proto_rawDescData
}

var file_index_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_index_platform_proto_goTypes = []interface{}{
	(*BuildIndexReq)(nil),  // 0: BuildIndexReq
	(*BuildIndexResp)(nil), // 1: BuildIndexResp
}
var file_index_platform_proto_depIdxs = []int32{
	0, // 0: IndexPlatformService.BuildIndexService:input_type -> BuildIndexReq
	1, // 1: IndexPlatformService.BuildIndexService:output_type -> BuildIndexResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_index_platform_proto_init() }
func file_index_platform_proto_init() {
	if File_index_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_index_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildIndexReq); i {
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
		file_index_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildIndexResp); i {
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
			RawDescriptor: file_index_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_index_platform_proto_goTypes,
		DependencyIndexes: file_index_platform_proto_depIdxs,
		MessageInfos:      file_index_platform_proto_msgTypes,
	}.Build()
	File_index_platform_proto = out.File
	file_index_platform_proto_rawDesc = nil
	file_index_platform_proto_goTypes = nil
	file_index_platform_proto_depIdxs = nil
}
