// Usage:
//    make proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/v1/delivery.proto

// Package api.v1 defines various messages that can be used by these
// microservices. You may either write protobuf definitions here for easier
// inter-language usage, or write the model structs manually like in
// internal/domain/email/request.go
//
// If you choose to use protobuf, run `make proto` and it generates both Go
// structs and php classes in the same directory: `delivery.pb.go`.
//
// You may consider moving the whole api folder into its own repository so that
// each project can import the generated files for consistency.

package delivery_v1

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

// Request is a common message used by all incoming messages. Typically, you
// want to embed this in your own custom message. See RefereeRequest.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SentByUserID keeps track of who is sending the message
	SentByUserID uint64 `protobuf:"varint,1,opt,name=SentByUserID,proto3" json:"SentByUserID,omitempty"`
	// Type tells what kind of UoW it is. They are defined by constants in
	// task/tasks.go and must be followed by all services using this Producer
	// api.
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	// UUID defines a unique ID sent by client. We can send it by header
	// request, or as part oj JSON payload.
	// This field is important to ensure deduplication in the event of client
	// retries.
	UUID string `protobuf:"bytes,3,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delivery_proto_msgTypes[0]
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
	return file_api_v1_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetSentByUserID() uint64 {
	if x != nil {
		return x.SentByUserID
	}
	return 0
}

func (x *Request) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Request) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

// Email is a custom message concerned with just Email parameters.
type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// To is the ID of the user to be sent to
	To int64 `protobuf:"varint,1,opt,name=To,proto3" json:"To,omitempty"`
	// Subject is the Email subject
	Subject string `protobuf:"bytes,2,opt,name=Subject,proto3" json:"Subject,omitempty"`
	// Content is the text to be sent in the body of an email
	Content string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_delivery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delivery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_api_v1_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *Email) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *Email) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Email) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// RefereeRequest user request is parsed using this message.
// Other Unit of Work (UoW) will also embed Request message.
// Other parameters are then defined is their own field, in this case an Email
// message
type RefereeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The compulsory common Request message
	Request *Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Specific parameters that RefereeRequest needs, in this case, Email message
	Parameters *Email `protobuf:"bytes,2,opt,name=Parameters,proto3" json:"Parameters,omitempty"`
}

func (x *RefereeRequest) Reset() {
	*x = RefereeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_delivery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefereeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefereeRequest) ProtoMessage() {}

func (x *RefereeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delivery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefereeRequest.ProtoReflect.Descriptor instead.
func (*RefereeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_delivery_proto_rawDescGZIP(), []int{2}
}

func (x *RefereeRequest) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *RefereeRequest) GetParameters() *Email {
	if x != nil {
		return x.Parameters
	}
	return nil
}

var File_api_v1_delivery_proto protoreflect.FileDescriptor

var file_api_v1_delivery_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22,
	0x55, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x4b, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x54, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x6a, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x42,
	0x13, 0x5a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_delivery_proto_rawDescOnce sync.Once
	file_api_v1_delivery_proto_rawDescData = file_api_v1_delivery_proto_rawDesc
)

func file_api_v1_delivery_proto_rawDescGZIP() []byte {
	file_api_v1_delivery_proto_rawDescOnce.Do(func() {
		file_api_v1_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_delivery_proto_rawDescData)
	})
	return file_api_v1_delivery_proto_rawDescData
}

var file_api_v1_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_delivery_proto_goTypes = []interface{}{
	(*Request)(nil),        // 0: api.v1.Request
	(*Email)(nil),          // 1: api.v1.Email
	(*RefereeRequest)(nil), // 2: api.v1.RefereeRequest
}
var file_api_v1_delivery_proto_depIdxs = []int32{
	0, // 0: api.v1.RefereeRequest.request:type_name -> api.v1.Request
	1, // 1: api.v1.RefereeRequest.Parameters:type_name -> api.v1.Email
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_delivery_proto_init() }
func file_api_v1_delivery_proto_init() {
	if File_api_v1_delivery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_delivery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_v1_delivery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_api_v1_delivery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefereeRequest); i {
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
			RawDescriptor: file_api_v1_delivery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_delivery_proto_goTypes,
		DependencyIndexes: file_api_v1_delivery_proto_depIdxs,
		MessageInfos:      file_api_v1_delivery_proto_msgTypes,
	}.Build()
	File_api_v1_delivery_proto = out.File
	file_api_v1_delivery_proto_rawDesc = nil
	file_api_v1_delivery_proto_goTypes = nil
	file_api_v1_delivery_proto_depIdxs = nil
}
