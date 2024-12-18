// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: eventstream.proto

package pb

import (
	v1 "github.com/jamieyoung5/kwikmedical-eventstream/pb/io/cloudevents/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types      []string `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"` // event types (e.g., "com.example.patients.created")
	Source     string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Subject    string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	ConsumerId string   `protobuf:"bytes,4,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	mi := &file_eventstream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventstream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_eventstream_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriptionRequest) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *SubscriptionRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *SubscriptionRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SubscriptionRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

type PublishEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success      bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`                              // indicates if the publish operation succeeded
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // optional
}

func (x *PublishEventResponse) Reset() {
	*x = PublishEventResponse{}
	mi := &file_eventstream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishEventResponse) ProtoMessage() {}

func (x *PublishEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventstream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishEventResponse.ProtoReflect.Descriptor instead.
func (*PublishEventResponse) Descriptor() ([]byte, []int) {
	return file_eventstream_proto_rawDescGZIP(), []int{1}
}

func (x *PublishEventResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublishEventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PublishEventResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	mi := &file_eventstream_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventstream_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_eventstream_proto_rawDescGZIP(), []int{2}
}

func (x *UnsubscribeRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

var File_eventstream_proto protoreflect.FileDescriptor

var file_eventstream_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x1a, 0x23, 0x69, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x14, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a, 0x12,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x32, 0x8a, 0x02, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x56, 0x31, 0x12, 0x53, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x11, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x23, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x61, 0x6d, 0x69, 0x65, 0x79, 0x6f, 0x75, 0x6e, 0x67, 0x35, 0x2f, 0x6b, 0x77, 0x69, 0x6b, 0x6d,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventstream_proto_rawDescOnce sync.Once
	file_eventstream_proto_rawDescData = file_eventstream_proto_rawDesc
)

func file_eventstream_proto_rawDescGZIP() []byte {
	file_eventstream_proto_rawDescOnce.Do(func() {
		file_eventstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventstream_proto_rawDescData)
	})
	return file_eventstream_proto_rawDescData
}

var file_eventstream_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eventstream_proto_goTypes = []any{
	(*SubscriptionRequest)(nil),  // 0: eventstream.v1.SubscriptionRequest
	(*PublishEventResponse)(nil), // 1: eventstream.v1.PublishEventResponse
	(*UnsubscribeRequest)(nil),   // 2: eventstream.v1.UnsubscribeRequest
	(*v1.CloudEvent)(nil),        // 3: io.cloudevents.v1.CloudEvent
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_eventstream_proto_depIdxs = []int32{
	3, // 0: eventstream.v1.EventStreamV1.PublishEvent:input_type -> io.cloudevents.v1.CloudEvent
	0, // 1: eventstream.v1.EventStreamV1.SubscribeToEvents:input_type -> eventstream.v1.SubscriptionRequest
	2, // 2: eventstream.v1.EventStreamV1.Unsubscribe:input_type -> eventstream.v1.UnsubscribeRequest
	1, // 3: eventstream.v1.EventStreamV1.PublishEvent:output_type -> eventstream.v1.PublishEventResponse
	3, // 4: eventstream.v1.EventStreamV1.SubscribeToEvents:output_type -> io.cloudevents.v1.CloudEvent
	4, // 5: eventstream.v1.EventStreamV1.Unsubscribe:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eventstream_proto_init() }
func file_eventstream_proto_init() {
	if File_eventstream_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eventstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eventstream_proto_goTypes,
		DependencyIndexes: file_eventstream_proto_depIdxs,
		MessageInfos:      file_eventstream_proto_msgTypes,
	}.Build()
	File_eventstream_proto = out.File
	file_eventstream_proto_rawDesc = nil
	file_eventstream_proto_goTypes = nil
	file_eventstream_proto_depIdxs = nil
}
