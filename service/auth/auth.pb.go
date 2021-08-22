// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protobuf/auth.proto

package auth

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

type PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string `protobuf:"bytes,1,opt,name=X,proto3" json:"X,omitempty"`
	Y string `protobuf:"bytes,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (x *PrivateKey) Reset() {
	*x = PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKey) ProtoMessage() {}

func (x *PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKey.ProtoReflect.Descriptor instead.
func (*PrivateKey) Descriptor() ([]byte, []int) {
	return file_protobuf_auth_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateKey) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *PrivateKey) GetY() string {
	if x != nil {
		return x.Y
	}
	return ""
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string `protobuf:"bytes,1,opt,name=X,proto3" json:"X,omitempty"`
	Y string `protobuf:"bytes,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_protobuf_auth_proto_rawDescGZIP(), []int{1}
}

func (x *PublicKey) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *PublicKey) GetY() string {
	if x != nil {
		return x.Y
	}
	return ""
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R string `protobuf:"bytes,1,opt,name=R,proto3" json:"R,omitempty"`
	S string `protobuf:"bytes,2,opt,name=S,proto3" json:"S,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_protobuf_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Signature) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *Signature) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

var File_protobuf_auth_proto protoreflect.FileDescriptor

var file_protobuf_auth_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x28, 0x0a, 0x0a, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x59, 0x22, 0x27, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x58,
	0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x59, 0x22, 0x27,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x52,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x52, 0x12, 0x0c, 0x0a, 0x01, 0x53, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x53, 0x42, 0x0e, 0x5a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_auth_proto_rawDescOnce sync.Once
	file_protobuf_auth_proto_rawDescData = file_protobuf_auth_proto_rawDesc
)

func file_protobuf_auth_proto_rawDescGZIP() []byte {
	file_protobuf_auth_proto_rawDescOnce.Do(func() {
		file_protobuf_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_auth_proto_rawDescData)
	})
	return file_protobuf_auth_proto_rawDescData
}

var file_protobuf_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_auth_proto_goTypes = []interface{}{
	(*PrivateKey)(nil), // 0: auth.PrivateKey
	(*PublicKey)(nil),  // 1: auth.PublicKey
	(*Signature)(nil),  // 2: auth.Signature
}
var file_protobuf_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_auth_proto_init() }
func file_protobuf_auth_proto_init() {
	if File_protobuf_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateKey); i {
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
		file_protobuf_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
		file_protobuf_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
			RawDescriptor: file_protobuf_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_auth_proto_goTypes,
		DependencyIndexes: file_protobuf_auth_proto_depIdxs,
		MessageInfos:      file_protobuf_auth_proto_msgTypes,
	}.Build()
	File_protobuf_auth_proto = out.File
	file_protobuf_auth_proto_rawDesc = nil
	file_protobuf_auth_proto_goTypes = nil
	file_protobuf_auth_proto_depIdxs = nil
}
