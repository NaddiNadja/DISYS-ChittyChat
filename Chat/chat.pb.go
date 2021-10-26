// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: chat/chat.proto

package chat

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

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SendersName string `protobuf:"bytes,2,opt,name=senders_name,json=sendersName,proto3" json:"senders_name,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Channel) GetSendersName() string {
	if x != nil {
		return x.SendersName
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender      string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Channel     *Channel `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Message     string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	LamportTime int64    `protobuf:"varint,4,opt,name=lamportTime,proto3" json:"lamportTime,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Message) GetChannel() *Channel {
	if x != nil {
		return x.Channel
	}
	return nil
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetLamportTime() int64 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

type MessageAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageAck string `protobuf:"bytes,1,opt,name=messageAck,proto3" json:"messageAck,omitempty"`
}

func (x *MessageAck) Reset() {
	*x = MessageAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAck) ProtoMessage() {}

func (x *MessageAck) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAck.ProtoReflect.Descriptor instead.
func (*MessageAck) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *MessageAck) GetMessageAck() string {
	if x != nil {
		return x.MessageAck
	}
	return ""
}

type LeaveAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaveAck string `protobuf:"bytes,1,opt,name=leaveAck,proto3" json:"leaveAck,omitempty"`
}

func (x *LeaveAck) Reset() {
	*x = LeaveAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveAck) ProtoMessage() {}

func (x *LeaveAck) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveAck.ProtoReflect.Descriptor instead.
func (*LeaveAck) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveAck) GetLeaveAck() string {
	if x != nil {
		return x.LeaveAck
	}
	return ""
}

var File_chat_chat_proto protoreflect.FileDescriptor

var file_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x40, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b,
	0x22, 0x26, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x65, 0x61, 0x76, 0x65, 0x41, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x65, 0x61, 0x76, 0x65, 0x41, 0x63, 0x6b, 0x32, 0xa9, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0d, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x0d, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x2f, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x0e,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x41, 0x63, 0x6b, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x61, 0x64, 0x64, 0x69,
	0x4e, 0x61, 0x64, 0x6a, 0x61, 0x2f, 0x44, 0x49, 0x53, 0x59, 0x53, 0x2d, 0x43, 0x68, 0x69, 0x74,
	0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chat_chat_proto_rawDescOnce sync.Once
	file_chat_chat_proto_rawDescData = file_chat_chat_proto_rawDesc
)

func file_chat_chat_proto_rawDescGZIP() []byte {
	file_chat_chat_proto_rawDescOnce.Do(func() {
		file_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_chat_proto_rawDescData)
	})
	return file_chat_chat_proto_rawDescData
}

var file_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chat_chat_proto_goTypes = []interface{}{
	(*Channel)(nil),    // 0: chat.Channel
	(*Message)(nil),    // 1: chat.Message
	(*MessageAck)(nil), // 2: chat.MessageAck
	(*LeaveAck)(nil),   // 3: chat.LeaveAck
}
var file_chat_chat_proto_depIdxs = []int32{
	0, // 0: chat.Message.channel:type_name -> chat.Channel
	0, // 1: chat.ChittyChatService.JoinChannel:input_type -> chat.Channel
	0, // 2: chat.ChittyChatService.LeaveChannel:input_type -> chat.Channel
	1, // 3: chat.ChittyChatService.SendMessage:input_type -> chat.Message
	1, // 4: chat.ChittyChatService.JoinChannel:output_type -> chat.Message
	3, // 5: chat.ChittyChatService.LeaveChannel:output_type -> chat.LeaveAck
	2, // 6: chat.ChittyChatService.SendMessage:output_type -> chat.MessageAck
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_chat_proto_init() }
func file_chat_chat_proto_init() {
	if File_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAck); i {
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
		file_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveAck); i {
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
			RawDescriptor: file_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_chat_proto_goTypes,
		DependencyIndexes: file_chat_chat_proto_depIdxs,
		MessageInfos:      file_chat_chat_proto_msgTypes,
	}.Build()
	File_chat_chat_proto = out.File
	file_chat_chat_proto_rawDesc = nil
	file_chat_chat_proto_goTypes = nil
	file_chat_chat_proto_depIdxs = nil
}
