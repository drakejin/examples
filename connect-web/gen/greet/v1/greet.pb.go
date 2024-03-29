// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: greet/v1/greet.proto

package greetv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_v1_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_v1_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_greet_v1_greet_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	// Types that are assignable to HelloOneOf:
	//
	//	*GreetResponse_Foo
	//	*GreetResponse_Bar
	HelloOneOf isGreetResponse_HelloOneOf `protobuf_oneof:"HelloOneOf"`
	HelloAny   *anypb.Any                 `protobuf:"bytes,4,opt,name=helloAny,proto3" json:"helloAny,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_v1_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_v1_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_greet_v1_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (m *GreetResponse) GetHelloOneOf() isGreetResponse_HelloOneOf {
	if m != nil {
		return m.HelloOneOf
	}
	return nil
}

func (x *GreetResponse) GetFoo() *Foo {
	if x, ok := x.GetHelloOneOf().(*GreetResponse_Foo); ok {
		return x.Foo
	}
	return nil
}

func (x *GreetResponse) GetBar() *Bar {
	if x, ok := x.GetHelloOneOf().(*GreetResponse_Bar); ok {
		return x.Bar
	}
	return nil
}

func (x *GreetResponse) GetHelloAny() *anypb.Any {
	if x != nil {
		return x.HelloAny
	}
	return nil
}

type isGreetResponse_HelloOneOf interface {
	isGreetResponse_HelloOneOf()
}

type GreetResponse_Foo struct {
	Foo *Foo `protobuf:"bytes,2,opt,name=foo,proto3,oneof"`
}

type GreetResponse_Bar struct {
	Bar *Bar `protobuf:"bytes,3,opt,name=bar,proto3,oneof"`
}

func (*GreetResponse_Foo) isGreetResponse_HelloOneOf() {}

func (*GreetResponse_Bar) isGreetResponse_HelloOneOf() {}

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_v1_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_greet_v1_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_greet_v1_greet_proto_rawDescGZIP(), []int{2}
}

func (x *Foo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_v1_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_greet_v1_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_greet_v1_greet_proto_rawDescGZIP(), []int{3}
}

func (x *Bar) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_greet_v1_greet_proto protoreflect.FileDescriptor

var file_greet_v1_greet_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xb1, 0x01, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a,
	0x03, 0x66, 0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x48, 0x00, 0x52, 0x03, 0x66, 0x6f, 0x6f,
	0x12, 0x21, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x48, 0x00, 0x52, 0x03,
	0x62, 0x61, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x41, 0x6e, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x41, 0x6e, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4f, 0x6e,
	0x65, 0x4f, 0x66, 0x22, 0x19, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x17,
	0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32, 0x4a, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x16, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x72, 0x61, 0x6b, 0x65, 0x6a, 0x69, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2d, 0x77, 0x65, 0x62, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_v1_greet_proto_rawDescOnce sync.Once
	file_greet_v1_greet_proto_rawDescData = file_greet_v1_greet_proto_rawDesc
)

func file_greet_v1_greet_proto_rawDescGZIP() []byte {
	file_greet_v1_greet_proto_rawDescOnce.Do(func() {
		file_greet_v1_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_v1_greet_proto_rawDescData)
	})
	return file_greet_v1_greet_proto_rawDescData
}

var file_greet_v1_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_greet_v1_greet_proto_goTypes = []interface{}{
	(*GreetRequest)(nil),  // 0: greet.v1.GreetRequest
	(*GreetResponse)(nil), // 1: greet.v1.GreetResponse
	(*Foo)(nil),           // 2: greet.v1.Foo
	(*Bar)(nil),           // 3: greet.v1.Bar
	(*anypb.Any)(nil),     // 4: google.protobuf.Any
}
var file_greet_v1_greet_proto_depIdxs = []int32{
	2, // 0: greet.v1.GreetResponse.foo:type_name -> greet.v1.Foo
	3, // 1: greet.v1.GreetResponse.bar:type_name -> greet.v1.Bar
	4, // 2: greet.v1.GreetResponse.helloAny:type_name -> google.protobuf.Any
	0, // 3: greet.v1.GreetService.Greet:input_type -> greet.v1.GreetRequest
	1, // 4: greet.v1.GreetService.Greet:output_type -> greet.v1.GreetResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_greet_v1_greet_proto_init() }
func file_greet_v1_greet_proto_init() {
	if File_greet_v1_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_v1_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_greet_v1_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_greet_v1_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
		file_greet_v1_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
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
	file_greet_v1_greet_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GreetResponse_Foo)(nil),
		(*GreetResponse_Bar)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greet_v1_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greet_v1_greet_proto_goTypes,
		DependencyIndexes: file_greet_v1_greet_proto_depIdxs,
		MessageInfos:      file_greet_v1_greet_proto_msgTypes,
	}.Build()
	File_greet_v1_greet_proto = out.File
	file_greet_v1_greet_proto_rawDesc = nil
	file_greet_v1_greet_proto_goTypes = nil
	file_greet_v1_greet_proto_depIdxs = nil
}
