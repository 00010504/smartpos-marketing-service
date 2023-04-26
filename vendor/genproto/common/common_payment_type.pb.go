// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: common_payment_type.proto

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

type CommonPaymentTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Logo    string   `protobuf:"bytes,3,opt,name=logo,proto3" json:"logo,omitempty"`
	Request *Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *CommonPaymentTypes) Reset() {
	*x = CommonPaymentTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_payment_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonPaymentTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonPaymentTypes) ProtoMessage() {}

func (x *CommonPaymentTypes) ProtoReflect() protoreflect.Message {
	mi := &file_common_payment_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonPaymentTypes.ProtoReflect.Descriptor instead.
func (*CommonPaymentTypes) Descriptor() ([]byte, []int) {
	return file_common_payment_type_proto_rawDescGZIP(), []int{0}
}

func (x *CommonPaymentTypes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommonPaymentTypes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommonPaymentTypes) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *CommonPaymentTypes) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

var File_common_payment_type_proto protoreflect.FileDescriptor

var file_common_payment_type_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x12, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x11, 0x5a, 0x0f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_payment_type_proto_rawDescOnce sync.Once
	file_common_payment_type_proto_rawDescData = file_common_payment_type_proto_rawDesc
)

func file_common_payment_type_proto_rawDescGZIP() []byte {
	file_common_payment_type_proto_rawDescOnce.Do(func() {
		file_common_payment_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_payment_type_proto_rawDescData)
	})
	return file_common_payment_type_proto_rawDescData
}

var file_common_payment_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_payment_type_proto_goTypes = []interface{}{
	(*CommonPaymentTypes)(nil), // 0: CommonPaymentTypes
	(*Request)(nil),            // 1: Request
}
var file_common_payment_type_proto_depIdxs = []int32{
	1, // 0: CommonPaymentTypes.request:type_name -> Request
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_payment_type_proto_init() }
func file_common_payment_type_proto_init() {
	if File_common_payment_type_proto != nil {
		return
	}
	file_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_payment_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonPaymentTypes); i {
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
			RawDescriptor: file_common_payment_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_payment_type_proto_goTypes,
		DependencyIndexes: file_common_payment_type_proto_depIdxs,
		MessageInfos:      file_common_payment_type_proto_msgTypes,
	}.Build()
	File_common_payment_type_proto = out.File
	file_common_payment_type_proto_rawDesc = nil
	file_common_payment_type_proto_goTypes = nil
	file_common_payment_type_proto_depIdxs = nil
}
