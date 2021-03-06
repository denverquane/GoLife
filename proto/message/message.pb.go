// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: message.proto

package message

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MessageType int32

const (
	//sent by the client when they want to register, and echoed back by the server on a successful register
	MessageType_REGISTER MessageType = 0
	//data about the server
	MessageType_SERVER_DATA MessageType = 1
	//raw world data, alongside any additional information about the world (dimensions, tick, etc)
	MessageType_WORLD_DATA MessageType = 2
	//A command from the client to the server
	MessageType_COMMAND MessageType = 3
	//A response in regards to the client issuing a command
	MessageType_RESPONSE MessageType = 4
	//A chat message that has gone out
	MessageType_CHAT_LOG MessageType = 5
	//All available RLEs that the server knows about
	MessageType_RLE_OPTIONS MessageType = 6
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "REGISTER",
		1: "SERVER_DATA",
		2: "WORLD_DATA",
		3: "COMMAND",
		4: "RESPONSE",
		5: "CHAT_LOG",
		6: "RLE_OPTIONS",
	}
	MessageType_value = map[string]int32{
		"REGISTER":    0,
		"SERVER_DATA": 1,
		"WORLD_DATA":  2,
		"COMMAND":     3,
		"RESPONSE":    4,
		"CHAT_LOG":    5,
		"RLE_OPTIONS": 6,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type CommandType int32

const (
	CommandType_MARK_CELL    CommandType = 0
	CommandType_PLACE_RLE    CommandType = 1
	CommandType_TOGGLE_PAUSE CommandType = 2
	CommandType_POST_CHAT    CommandType = 3
	CommandType_CLEAR_BOARD  CommandType = 4
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "MARK_CELL",
		1: "PLACE_RLE",
		2: "TOGGLE_PAUSE",
		3: "POST_CHAT",
		4: "CLEAR_BOARD",
	}
	CommandType_value = map[string]int32{
		"MARK_CELL":    0,
		"PLACE_RLE":    1,
		"TOGGLE_PAUSE": 2,
		"POST_CHAT":    3,
		"CLEAR_BOARD":  4,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

type ResponseCode int32

const (
	ResponseCode_GENERIC_SUCCESS ResponseCode = 0
	ResponseCode_GENERIC_FAILURE ResponseCode = 1
)

// Enum value maps for ResponseCode.
var (
	ResponseCode_name = map[int32]string{
		0: "GENERIC_SUCCESS",
		1: "GENERIC_FAILURE",
	}
	ResponseCode_value = map[string]int32{
		"GENERIC_SUCCESS": 0,
		"GENERIC_FAILURE": 1,
	}
)

func (x ResponseCode) Enum() *ResponseCode {
	p := new(ResponseCode)
	*p = x
	return p
}

func (x ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[2].Descriptor()
}

func (ResponseCode) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[2]
}

func (x ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseCode.Descriptor instead.
func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=message.MessageType" json:"type,omitempty"`
	Content []byte      `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
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
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_REGISTER
}

func (x *Message) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//24bit color
	Color uint32 `protobuf:"fixed32,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetColor() uint32 {
	if x != nil {
		return x.Color
	}
	return 0
}

type WorldData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []uint32 `protobuf:"fixed32,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	Tick   uint64   `protobuf:"varint,2,opt,name=tick,proto3" json:"tick,omitempty"`
	Width  uint32   `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32   `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Paused bool     `protobuf:"varint,5,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (x *WorldData) Reset() {
	*x = WorldData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldData) ProtoMessage() {}

func (x *WorldData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldData.ProtoReflect.Descriptor instead.
func (*WorldData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *WorldData) GetData() []uint32 {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *WorldData) GetTick() uint64 {
	if x != nil {
		return x.Tick
	}
	return 0
}

func (x *WorldData) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *WorldData) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *WorldData) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

type ServerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *ServerData) Reset() {
	*x = ServerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerData) ProtoMessage() {}

func (x *ServerData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerData.ProtoReflect.Descriptor instead.
func (*ServerData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *ServerData) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type CommandType `protobuf:"varint,1,opt,name=type,proto3,enum=message.CommandType" json:"type,omitempty"`
	X    uint32      `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y    uint32      `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Text string      `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *Command) GetType() CommandType {
	if x != nil {
		return x.Type
	}
	return CommandType_MARK_CELL
}

func (x *Command) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Command) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Command) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=message.ResponseCode" json:"code,omitempty"`
	Text string       `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetCode() ResponseCode {
	if x != nil {
		return x.Code
	}
	return ResponseCode_GENERIC_SUCCESS
}

func (x *Response) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Text   string  `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *Chat) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *Chat) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type RLEs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rles []*RLE `protobuf:"bytes,1,rep,name=rles,proto3" json:"rles,omitempty"`
}

func (x *RLEs) Reset() {
	*x = RLEs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RLEs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RLEs) ProtoMessage() {}

func (x *RLEs) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RLEs.ProtoReflect.Descriptor instead.
func (*RLEs) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *RLEs) GetRles() []*RLE {
	if x != nil {
		return x.Rles
	}
	return nil
}

type RLE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Width  uint32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RLE) Reset() {
	*x = RLE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RLE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RLE) ProtoMessage() {}

func (x *RLE) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RLE.ProtoReflect.Descriptor instead.
func (*RLE) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{8}
}

func (x *RLE) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RLE) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *RLE) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *RLE) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x07, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x79, 0x0a, 0x09, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x07, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x22, 0x37, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22,
	0x63, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x49, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x43, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x28, 0x0a, 0x04, 0x52, 0x4c, 0x45, 0x73, 0x12, 0x20, 0x0a, 0x04,
	0x72, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x4c, 0x45, 0x52, 0x04, 0x72, 0x6c, 0x65, 0x73, 0x22, 0x5b,
	0x0a, 0x03, 0x52, 0x4c, 0x45, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x76, 0x0a, 0x0b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x4f, 0x52,
	0x4c, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d,
	0x4d, 0x41, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x4c, 0x4f, 0x47,
	0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x53, 0x10, 0x06, 0x2a, 0x5d, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x52, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x4f, 0x47, 0x47, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x10,
	0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x5f, 0x42, 0x4f, 0x41, 0x52, 0x44,
	0x10, 0x04, 0x2a, 0x38, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x45, 0x4e, 0x45, 0x52,
	0x49, 0x43, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_message_proto_goTypes = []interface{}{
	(MessageType)(0),   // 0: message.MessageType
	(CommandType)(0),   // 1: message.CommandType
	(ResponseCode)(0),  // 2: message.ResponseCode
	(*Message)(nil),    // 3: message.Message
	(*Player)(nil),     // 4: message.Player
	(*WorldData)(nil),  // 5: message.WorldData
	(*ServerData)(nil), // 6: message.ServerData
	(*Command)(nil),    // 7: message.Command
	(*Response)(nil),   // 8: message.Response
	(*Chat)(nil),       // 9: message.Chat
	(*RLEs)(nil),       // 10: message.RLEs
	(*RLE)(nil),        // 11: message.RLE
}
var file_message_proto_depIdxs = []int32{
	0,  // 0: message.Message.type:type_name -> message.MessageType
	4,  // 1: message.ServerData.players:type_name -> message.Player
	1,  // 2: message.Command.type:type_name -> message.CommandType
	2,  // 3: message.Response.code:type_name -> message.ResponseCode
	4,  // 4: message.Chat.player:type_name -> message.Player
	11, // 5: message.RLEs.rles:type_name -> message.RLE
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldData); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerData); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RLEs); i {
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
		file_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RLE); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
