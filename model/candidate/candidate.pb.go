// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/protobuf/candidate.proto

package candidate

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoterID   []byte `protobuf:"bytes,1,opt,name=VoterID,proto3" json:"VoterID,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
	KeyBytes  []byte `protobuf:"bytes,3,opt,name=KeyBytes,proto3" json:"KeyBytes,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Markdown  string `protobuf:"bytes,5,opt,name=Markdown,proto3" json:"Markdown,omitempty"`
	Hash      []byte `protobuf:"bytes,6,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_candidate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_candidate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_candidate_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetVoterID() []byte {
	if x != nil {
		return x.VoterID
	}
	return nil
}

func (x *Request) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Request) GetKeyBytes() []byte {
	if x != nil {
		return x.KeyBytes
	}
	return nil
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetMarkdown() string {
	if x != nil {
		return x.Markdown
	}
	return ""
}

func (x *Request) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK    bool   `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_candidate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_candidate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_candidate_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

func (x *Result) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Markdown string `protobuf:"bytes,2,opt,name=Markdown,proto3" json:"Markdown,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_candidate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_candidate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_candidate_proto_rawDescGZIP(), []int{2}
}

func (x *Candidate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Candidate) GetMarkdown() string {
	if x != nil {
		return x.Markdown
	}
	return ""
}

func (x *Candidate) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string   `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Names []string `protobuf:"bytes,2,rep,name=Names,proto3" json:"Names,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protobuf_candidate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protobuf_candidate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_proto_protobuf_candidate_proto_rawDescGZIP(), []int{3}
}

func (x *List) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *List) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_proto_protobuf_candidate_proto protoreflect.FileDescriptor

var file_proto_protobuf_candidate_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x6f, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x22,
	0x2e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x4b, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x57, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x32, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_protobuf_candidate_proto_rawDescOnce sync.Once
	file_proto_protobuf_candidate_proto_rawDescData = file_proto_protobuf_candidate_proto_rawDesc
)

func file_proto_protobuf_candidate_proto_rawDescGZIP() []byte {
	file_proto_protobuf_candidate_proto_rawDescOnce.Do(func() {
		file_proto_protobuf_candidate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protobuf_candidate_proto_rawDescData)
	})
	return file_proto_protobuf_candidate_proto_rawDescData
}

var file_proto_protobuf_candidate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_protobuf_candidate_proto_goTypes = []interface{}{
	(*Request)(nil),   // 0: candidate.Request
	(*Result)(nil),    // 1: candidate.Result
	(*Candidate)(nil), // 2: candidate.Candidate
	(*List)(nil),      // 3: candidate.List
}
var file_proto_protobuf_candidate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_protobuf_candidate_proto_init() }
func file_proto_protobuf_candidate_proto_init() {
	if File_proto_protobuf_candidate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protobuf_candidate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_protobuf_candidate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_proto_protobuf_candidate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
		file_proto_protobuf_candidate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
			RawDescriptor: file_proto_protobuf_candidate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_protobuf_candidate_proto_goTypes,
		DependencyIndexes: file_proto_protobuf_candidate_proto_depIdxs,
		MessageInfos:      file_proto_protobuf_candidate_proto_msgTypes,
	}.Build()
	File_proto_protobuf_candidate_proto = out.File
	file_proto_protobuf_candidate_proto_rawDesc = nil
	file_proto_protobuf_candidate_proto_goTypes = nil
	file_proto_protobuf_candidate_proto_depIdxs = nil
}
