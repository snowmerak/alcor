// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protobuf/transaction.proto

package transaction

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_transaction_proto_msgTypes[0]
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
	return file_protobuf_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type EnrollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message []byte `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	From    []byte `protobuf:"bytes,4,opt,name=From,proto3" json:"From,omitempty"`
}

func (x *EnrollRequest) Reset() {
	*x = EnrollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollRequest) ProtoMessage() {}

func (x *EnrollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollRequest.ProtoReflect.Descriptor instead.
func (*EnrollRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *EnrollRequest) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *EnrollRequest) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

type EnrollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash      []byte   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Message   []byte   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Timestamp []byte   `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	From      []byte   `protobuf:"bytes,4,opt,name=From,proto3" json:"From,omitempty"`
	Befores   [][]byte `protobuf:"bytes,5,rep,name=Befores,proto3" json:"Befores,omitempty"`
}

func (x *EnrollResponse) Reset() {
	*x = EnrollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollResponse) ProtoMessage() {}

func (x *EnrollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollResponse.ProtoReflect.Descriptor instead.
func (*EnrollResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *EnrollResponse) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *EnrollResponse) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *EnrollResponse) GetTimestamp() []byte {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *EnrollResponse) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *EnrollResponse) GetBefores() [][]byte {
	if x != nil {
		return x.Befores
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash      []byte   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Message   []byte   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Timestamp []byte   `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	From      []byte   `protobuf:"bytes,4,opt,name=From,proto3" json:"From,omitempty"`
	Befores   [][]byte `protobuf:"bytes,5,rep,name=Befores,proto3" json:"Befores,omitempty"`
	Signature []byte   `protobuf:"bytes,6,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_protobuf_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Transaction) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Transaction) GetTimestamp() []byte {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Transaction) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Transaction) GetBefores() [][]byte {
	if x != nil {
		return x.Befores
	}
	return nil
}

func (x *Transaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_protobuf_transaction_proto protoreflect.FileDescriptor

var file_protobuf_transaction_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3d, 0x0a, 0x0d, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07,
	0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_transaction_proto_rawDescOnce sync.Once
	file_protobuf_transaction_proto_rawDescData = file_protobuf_transaction_proto_rawDesc
)

func file_protobuf_transaction_proto_rawDescGZIP() []byte {
	file_protobuf_transaction_proto_rawDescOnce.Do(func() {
		file_protobuf_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_transaction_proto_rawDescData)
	})
	return file_protobuf_transaction_proto_rawDescData
}

var file_protobuf_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_transaction_proto_goTypes = []interface{}{
	(*GetRequest)(nil),     // 0: transaction.GetRequest
	(*EnrollRequest)(nil),  // 1: transaction.EnrollRequest
	(*EnrollResponse)(nil), // 2: transaction.EnrollResponse
	(*Transaction)(nil),    // 3: transaction.Transaction
}
var file_protobuf_transaction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_transaction_proto_init() }
func file_protobuf_transaction_proto_init() {
	if File_protobuf_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protobuf_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollRequest); i {
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
		file_protobuf_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollResponse); i {
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
		file_protobuf_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_protobuf_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_transaction_proto_goTypes,
		DependencyIndexes: file_protobuf_transaction_proto_depIdxs,
		MessageInfos:      file_protobuf_transaction_proto_msgTypes,
	}.Build()
	File_protobuf_transaction_proto = out.File
	file_protobuf_transaction_proto_rawDesc = nil
	file_protobuf_transaction_proto_goTypes = nil
	file_protobuf_transaction_proto_depIdxs = nil
}