// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.10.0
// source: proto/common/danmaku.proto

package common

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

type Danmaku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string  `protobuf:"bytes,1,opt,name=Author,proto3" json:"Author,omitempty"`
	Time   float64 `protobuf:"fixed64,2,opt,name=Time,proto3" json:"Time,omitempty"`
	Text   string  `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
	Color  uint32  `protobuf:"varint,4,opt,name=Color,proto3" json:"Color,omitempty"`
	Type   uint32  `protobuf:"varint,5,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Danmaku) Reset() {
	*x = Danmaku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_danmaku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Danmaku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Danmaku) ProtoMessage() {}

func (x *Danmaku) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_danmaku_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Danmaku.ProtoReflect.Descriptor instead.
func (*Danmaku) Descriptor() ([]byte, []int) {
	return file_proto_common_danmaku_proto_rawDescGZIP(), []int{0}
}

func (x *Danmaku) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Danmaku) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Danmaku) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Danmaku) GetColor() uint32 {
	if x != nil {
		return x.Color
	}
	return 0
}

func (x *Danmaku) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_proto_common_danmaku_proto protoreflect.FileDescriptor

var file_proto_common_danmaku_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64,
	0x61, 0x6e, 0x6d, 0x61, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x07, 0x44, 0x61, 0x6e, 0x6d, 0x61, 0x6b, 0x75, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x74, 0x2d, 0x64, 0x61, 0x6e, 0x6d,
	0x61, 0x6b, 0x75, 0x2f, 0x64, 0x61, 0x6e, 0x6d, 0x61, 0x6b, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_danmaku_proto_rawDescOnce sync.Once
	file_proto_common_danmaku_proto_rawDescData = file_proto_common_danmaku_proto_rawDesc
)

func file_proto_common_danmaku_proto_rawDescGZIP() []byte {
	file_proto_common_danmaku_proto_rawDescOnce.Do(func() {
		file_proto_common_danmaku_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_danmaku_proto_rawDescData)
	})
	return file_proto_common_danmaku_proto_rawDescData
}

var file_proto_common_danmaku_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_common_danmaku_proto_goTypes = []interface{}{
	(*Danmaku)(nil), // 0: common.Danmaku
}
var file_proto_common_danmaku_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_common_danmaku_proto_init() }
func file_proto_common_danmaku_proto_init() {
	if File_proto_common_danmaku_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_danmaku_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Danmaku); i {
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
			RawDescriptor: file_proto_common_danmaku_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_danmaku_proto_goTypes,
		DependencyIndexes: file_proto_common_danmaku_proto_depIdxs,
		MessageInfos:      file_proto_common_danmaku_proto_msgTypes,
	}.Build()
	File_proto_common_danmaku_proto = out.File
	file_proto_common_danmaku_proto_rawDesc = nil
	file_proto_common_danmaku_proto_goTypes = nil
	file_proto_common_danmaku_proto_depIdxs = nil
}
