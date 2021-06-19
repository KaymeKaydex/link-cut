// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: link-cut-api.proto

package link_cut_api

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

// The request message containing the user's name.
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_cut_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_cut_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_link_cut_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// The response message containing the greetings
type CreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateReply) Reset() {
	*x = CreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_cut_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReply) ProtoMessage() {}

func (x *CreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_link_cut_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReply.ProtoReflect.Descriptor instead.
func (*CreateReply) Descriptor() ([]byte, []int) {
	return file_link_cut_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// The request message containing the user's name.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_cut_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_cut_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_link_cut_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// The response message containing the greetings
type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_cut_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_link_cut_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_link_cut_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_link_cut_api_proto protoreflect.FileDescriptor

var file_link_cut_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x75, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x75, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x22, 0x21, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x1f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x1c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x32, 0x84, 0x01, 0x0a, 0x07, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x75, 0x74,
	0x12, 0x40, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x5f, 0x63, 0x75, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63,
	0x75, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x5f, 0x63, 0x75, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x75, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x2f, 0x3b, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x75, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_link_cut_api_proto_rawDescOnce sync.Once
	file_link_cut_api_proto_rawDescData = file_link_cut_api_proto_rawDesc
)

func file_link_cut_api_proto_rawDescGZIP() []byte {
	file_link_cut_api_proto_rawDescOnce.Do(func() {
		file_link_cut_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_link_cut_api_proto_rawDescData)
	})
	return file_link_cut_api_proto_rawDescData
}

var file_link_cut_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_link_cut_api_proto_goTypes = []interface{}{
	(*CreateRequest)(nil), // 0: link_cut_api.CreateRequest
	(*CreateReply)(nil),   // 1: link_cut_api.CreateReply
	(*GetRequest)(nil),    // 2: link_cut_api.GetRequest
	(*GetReply)(nil),      // 3: link_cut_api.GetReply
}
var file_link_cut_api_proto_depIdxs = []int32{
	0, // 0: link_cut_api.LinkCut.Create:input_type -> link_cut_api.CreateRequest
	2, // 1: link_cut_api.LinkCut.Get:input_type -> link_cut_api.GetRequest
	1, // 2: link_cut_api.LinkCut.Create:output_type -> link_cut_api.CreateReply
	3, // 3: link_cut_api.LinkCut.Get:output_type -> link_cut_api.GetReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_link_cut_api_proto_init() }
func file_link_cut_api_proto_init() {
	if File_link_cut_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_link_cut_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_link_cut_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReply); i {
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
		file_link_cut_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_link_cut_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
			RawDescriptor: file_link_cut_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_link_cut_api_proto_goTypes,
		DependencyIndexes: file_link_cut_api_proto_depIdxs,
		MessageInfos:      file_link_cut_api_proto_msgTypes,
	}.Build()
	File_link_cut_api_proto = out.File
	file_link_cut_api_proto_rawDesc = nil
	file_link_cut_api_proto_goTypes = nil
	file_link_cut_api_proto_depIdxs = nil
}