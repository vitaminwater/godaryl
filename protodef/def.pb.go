// Code generated by protoc-gen-go. DO NOT EDIT.
// source: def.proto

/*
Package protodef is a generated protocol buffer package.

It is generated from these files:
	def.proto

It has these top-level messages:
	StartDarylRequest
	HasDarylRequest
	StartDarylResponse
	HasDarylResponse
	UserMessageRequest
	MessageLink
	Message
	UserMessageResponse
	Habit
	AddHabitRequest
	AddHabitResponse
	SessionSlice
	Session
	StartWorkSessionRequest
	StartWorkSessionResponse
*/
package protodef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StartDarylRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *StartDarylRequest) Reset()                    { *m = StartDarylRequest{} }
func (m *StartDarylRequest) String() string            { return proto.CompactTextString(m) }
func (*StartDarylRequest) ProtoMessage()               {}
func (*StartDarylRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartDarylRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type HasDarylRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *HasDarylRequest) Reset()                    { *m = HasDarylRequest{} }
func (m *HasDarylRequest) String() string            { return proto.CompactTextString(m) }
func (*HasDarylRequest) ProtoMessage()               {}
func (*HasDarylRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HasDarylRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type StartDarylResponse struct {
}

func (m *StartDarylResponse) Reset()                    { *m = StartDarylResponse{} }
func (m *StartDarylResponse) String() string            { return proto.CompactTextString(m) }
func (*StartDarylResponse) ProtoMessage()               {}
func (*StartDarylResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type HasDarylResponse struct {
	Response bool `protobuf:"varint,1,opt,name=response" json:"response,omitempty"`
}

func (m *HasDarylResponse) Reset()                    { *m = HasDarylResponse{} }
func (m *HasDarylResponse) String() string            { return proto.CompactTextString(m) }
func (*HasDarylResponse) ProtoMessage()               {}
func (*HasDarylResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HasDarylResponse) GetResponse() bool {
	if m != nil {
		return m.Response
	}
	return false
}

type UserMessageRequest struct {
	Identifier string   `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Message    *Message `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *UserMessageRequest) Reset()                    { *m = UserMessageRequest{} }
func (m *UserMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*UserMessageRequest) ProtoMessage()               {}
func (*UserMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserMessageRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *UserMessageRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageLink struct {
	// @inject_tag: db:"link"
	Link string `protobuf:"bytes,1,opt,name=link" json:"link,omitempty" db:"link"`
}

func (m *MessageLink) Reset()                    { *m = MessageLink{} }
func (m *MessageLink) String() string            { return proto.CompactTextString(m) }
func (*MessageLink) ProtoMessage()               {}
func (*MessageLink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MessageLink) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type Message struct {
	// @inject_tag: db:"id"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" db:"id"`
	// @inject_tag: db:"text"
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty" db:"text"`
	// @inject_tag: db:"at"
	At *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=at" json:"at,omitempty" db:"at"`
	// @inject_tag: db:"link"
	Done bool `protobuf:"varint,4,opt,name=done" json:"done,omitempty" db:"link"`
	// @inject_tag: db:"todo"
	Todo string `protobuf:"bytes,5,opt,name=todo" json:"todo,omitempty" db:"todo"`
	// @inject_tag: db:"links"
	Links []*MessageLink `protobuf:"bytes,6,rep,name=links" json:"links,omitempty" db:"links"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.At
	}
	return nil
}

func (m *Message) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Message) GetTodo() string {
	if m != nil {
		return m.Todo
	}
	return ""
}

func (m *Message) GetLinks() []*MessageLink {
	if m != nil {
		return m.Links
	}
	return nil
}

type UserMessageResponse struct {
	Message *Message `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *UserMessageResponse) Reset()                    { *m = UserMessageResponse{} }
func (m *UserMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*UserMessageResponse) ProtoMessage()               {}
func (*UserMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserMessageResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type Habit struct {
	// @inject_tag: db:"id"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" db:"id"`
	// @inject_tag: db:"title"
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty" db:"title"`
	// @inject_tag: db:"avgDuration"
	Duration uint32                     `protobuf:"varint,3,opt,name=duration" json:"duration,omitempty" db:"avgDuration"`
	Deadline *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=deadline" json:"deadline,omitempty"`
	// @inject_tag: db:"during"
	During uint32 `protobuf:"varint,5,opt,name=during" json:"during,omitempty" db:"during"`
	// @inject_tag: db:"cron"
	Cron     string                     `protobuf:"bytes,6,opt,name=cron" json:"cron,omitempty" db:"cron"`
	LastDone *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=lastDone" json:"lastDone,omitempty"`
	// @inject_tag: db:"nMissed"
	NMissed uint32 `protobuf:"varint,7,opt,name=nMissed" json:"nMissed,omitempty" db:"nMissed"`
}

func (m *Habit) Reset()                    { *m = Habit{} }
func (m *Habit) String() string            { return proto.CompactTextString(m) }
func (*Habit) ProtoMessage()               {}
func (*Habit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Habit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Habit) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Habit) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Habit) GetDeadline() *google_protobuf.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *Habit) GetDuring() uint32 {
	if m != nil {
		return m.During
	}
	return 0
}

func (m *Habit) GetCron() string {
	if m != nil {
		return m.Cron
	}
	return ""
}

func (m *Habit) GetLastDone() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastDone
	}
	return nil
}

func (m *Habit) GetNMissed() uint32 {
	if m != nil {
		return m.NMissed
	}
	return 0
}

type AddHabitRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Habit      *Habit `protobuf:"bytes,2,opt,name=habit" json:"habit,omitempty"`
}

func (m *AddHabitRequest) Reset()                    { *m = AddHabitRequest{} }
func (m *AddHabitRequest) String() string            { return proto.CompactTextString(m) }
func (*AddHabitRequest) ProtoMessage()               {}
func (*AddHabitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AddHabitRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *AddHabitRequest) GetHabit() *Habit {
	if m != nil {
		return m.Habit
	}
	return nil
}

type AddHabitResponse struct {
	Habit *Habit `protobuf:"bytes,2,opt,name=habit" json:"habit,omitempty"`
}

func (m *AddHabitResponse) Reset()                    { *m = AddHabitResponse{} }
func (m *AddHabitResponse) String() string            { return proto.CompactTextString(m) }
func (*AddHabitResponse) ProtoMessage()               {}
func (*AddHabitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AddHabitResponse) GetHabit() *Habit {
	if m != nil {
		return m.Habit
	}
	return nil
}

type SessionSlice struct {
	// @inject_tag: db:"start"
	Start string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty" db:"start"`
	// @inject_tag: db:"end"
	End string `protobuf:"bytes,2,opt,name=end" json:"end,omitempty" db:"end"`
	// @inject_tag: db:"habit"
	Habit *Habit `protobuf:"bytes,3,opt,name=habit" json:"habit,omitempty" db:"habit"`
}

func (m *SessionSlice) Reset()                    { *m = SessionSlice{} }
func (m *SessionSlice) String() string            { return proto.CompactTextString(m) }
func (*SessionSlice) ProtoMessage()               {}
func (*SessionSlice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SessionSlice) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *SessionSlice) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *SessionSlice) GetHabit() *Habit {
	if m != nil {
		return m.Habit
	}
	return nil
}

type Session struct {
	// @inject_tag: db:"start"
	Start string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty" db:"start"`
	// @inject_tag: db:"end"
	End string `protobuf:"bytes,2,opt,name=end" json:"end,omitempty" db:"end"`
	// @inject_tag: db:"slices"
	Slices []*SessionSlice `protobuf:"bytes,3,rep,name=slices" json:"slices,omitempty" db:"slices"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Session) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Session) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *Session) GetSlices() []*SessionSlice {
	if m != nil {
		return m.Slices
	}
	return nil
}

type StartWorkSessionRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *StartWorkSessionRequest) Reset()                    { *m = StartWorkSessionRequest{} }
func (m *StartWorkSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*StartWorkSessionRequest) ProtoMessage()               {}
func (*StartWorkSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *StartWorkSessionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type StartWorkSessionResponse struct {
	Session *Session `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
}

func (m *StartWorkSessionResponse) Reset()                    { *m = StartWorkSessionResponse{} }
func (m *StartWorkSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*StartWorkSessionResponse) ProtoMessage()               {}
func (*StartWorkSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *StartWorkSessionResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func init() {
	proto.RegisterType((*StartDarylRequest)(nil), "protodef.StartDarylRequest")
	proto.RegisterType((*HasDarylRequest)(nil), "protodef.HasDarylRequest")
	proto.RegisterType((*StartDarylResponse)(nil), "protodef.StartDarylResponse")
	proto.RegisterType((*HasDarylResponse)(nil), "protodef.HasDarylResponse")
	proto.RegisterType((*UserMessageRequest)(nil), "protodef.UserMessageRequest")
	proto.RegisterType((*MessageLink)(nil), "protodef.MessageLink")
	proto.RegisterType((*Message)(nil), "protodef.Message")
	proto.RegisterType((*UserMessageResponse)(nil), "protodef.UserMessageResponse")
	proto.RegisterType((*Habit)(nil), "protodef.Habit")
	proto.RegisterType((*AddHabitRequest)(nil), "protodef.AddHabitRequest")
	proto.RegisterType((*AddHabitResponse)(nil), "protodef.AddHabitResponse")
	proto.RegisterType((*SessionSlice)(nil), "protodef.SessionSlice")
	proto.RegisterType((*Session)(nil), "protodef.Session")
	proto.RegisterType((*StartWorkSessionRequest)(nil), "protodef.StartWorkSessionRequest")
	proto.RegisterType((*StartWorkSessionResponse)(nil), "protodef.StartWorkSessionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Farm service

type FarmClient interface {
	StartDaryl(ctx context.Context, in *StartDarylRequest, opts ...grpc.CallOption) (*StartDarylResponse, error)
	HasDaryl(ctx context.Context, in *HasDarylRequest, opts ...grpc.CallOption) (*HasDarylResponse, error)
}

type farmClient struct {
	cc *grpc.ClientConn
}

func NewFarmClient(cc *grpc.ClientConn) FarmClient {
	return &farmClient{cc}
}

func (c *farmClient) StartDaryl(ctx context.Context, in *StartDarylRequest, opts ...grpc.CallOption) (*StartDarylResponse, error) {
	out := new(StartDarylResponse)
	err := grpc.Invoke(ctx, "/protodef.Farm/StartDaryl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) HasDaryl(ctx context.Context, in *HasDarylRequest, opts ...grpc.CallOption) (*HasDarylResponse, error) {
	out := new(HasDarylResponse)
	err := grpc.Invoke(ctx, "/protodef.Farm/HasDaryl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Farm service

type FarmServer interface {
	StartDaryl(context.Context, *StartDarylRequest) (*StartDarylResponse, error)
	HasDaryl(context.Context, *HasDarylRequest) (*HasDarylResponse, error)
}

func RegisterFarmServer(s *grpc.Server, srv FarmServer) {
	s.RegisterService(&_Farm_serviceDesc, srv)
}

func _Farm_StartDaryl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDarylRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).StartDaryl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.Farm/StartDaryl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).StartDaryl(ctx, req.(*StartDarylRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_HasDaryl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasDarylRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).HasDaryl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.Farm/HasDaryl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).HasDaryl(ctx, req.(*HasDarylRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Farm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protodef.Farm",
	HandlerType: (*FarmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartDaryl",
			Handler:    _Farm_StartDaryl_Handler,
		},
		{
			MethodName: "HasDaryl",
			Handler:    _Farm_HasDaryl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "def.proto",
}

// Client API for Daryl service

type DarylClient interface {
	UserMessage(ctx context.Context, in *UserMessageRequest, opts ...grpc.CallOption) (*UserMessageResponse, error)
	AddHabit(ctx context.Context, in *AddHabitRequest, opts ...grpc.CallOption) (*AddHabitResponse, error)
	StartWorkSession(ctx context.Context, in *StartWorkSessionRequest, opts ...grpc.CallOption) (*StartWorkSessionResponse, error)
}

type darylClient struct {
	cc *grpc.ClientConn
}

func NewDarylClient(cc *grpc.ClientConn) DarylClient {
	return &darylClient{cc}
}

func (c *darylClient) UserMessage(ctx context.Context, in *UserMessageRequest, opts ...grpc.CallOption) (*UserMessageResponse, error) {
	out := new(UserMessageResponse)
	err := grpc.Invoke(ctx, "/protodef.Daryl/UserMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darylClient) AddHabit(ctx context.Context, in *AddHabitRequest, opts ...grpc.CallOption) (*AddHabitResponse, error) {
	out := new(AddHabitResponse)
	err := grpc.Invoke(ctx, "/protodef.Daryl/AddHabit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darylClient) StartWorkSession(ctx context.Context, in *StartWorkSessionRequest, opts ...grpc.CallOption) (*StartWorkSessionResponse, error) {
	out := new(StartWorkSessionResponse)
	err := grpc.Invoke(ctx, "/protodef.Daryl/StartWorkSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Daryl service

type DarylServer interface {
	UserMessage(context.Context, *UserMessageRequest) (*UserMessageResponse, error)
	AddHabit(context.Context, *AddHabitRequest) (*AddHabitResponse, error)
	StartWorkSession(context.Context, *StartWorkSessionRequest) (*StartWorkSessionResponse, error)
}

func RegisterDarylServer(s *grpc.Server, srv DarylServer) {
	s.RegisterService(&_Daryl_serviceDesc, srv)
}

func _Daryl_UserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServer).UserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.Daryl/UserMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServer).UserMessage(ctx, req.(*UserMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daryl_AddHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHabitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServer).AddHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.Daryl/AddHabit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServer).AddHabit(ctx, req.(*AddHabitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daryl_StartWorkSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartWorkSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServer).StartWorkSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.Daryl/StartWorkSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServer).StartWorkSession(ctx, req.(*StartWorkSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Daryl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protodef.Daryl",
	HandlerType: (*DarylServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserMessage",
			Handler:    _Daryl_UserMessage_Handler,
		},
		{
			MethodName: "AddHabit",
			Handler:    _Daryl_AddHabit_Handler,
		},
		{
			MethodName: "StartWorkSession",
			Handler:    _Daryl_StartWorkSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "def.proto",
}

func init() { proto.RegisterFile("def.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x93, 0x38, 0x71, 0x26, 0x94, 0xa4, 0x4b, 0x29, 0xc6, 0xfc, 0x28, 0xb5, 0x84, 0x14,
	0x51, 0xc9, 0x88, 0x54, 0x42, 0xea, 0xb1, 0xa8, 0xa2, 0x15, 0xa2, 0x17, 0x07, 0x04, 0x1c, 0x38,
	0x6c, 0xba, 0x9b, 0xb0, 0x8a, 0xe3, 0x0d, 0xde, 0x8d, 0x44, 0x5f, 0x83, 0x07, 0xe1, 0xdd, 0x78,
	0x02, 0xd0, 0xfe, 0x38, 0x76, 0xe3, 0x92, 0xa6, 0x27, 0xcf, 0xee, 0x7e, 0xf3, 0xcd, 0xcc, 0x37,
	0xe3, 0x81, 0x36, 0xa1, 0x93, 0x68, 0x91, 0x71, 0xc9, 0x91, 0xa7, 0x3f, 0x84, 0x4e, 0x82, 0xfe,
	0x42, 0x5e, 0x2d, 0xa8, 0x78, 0x25, 0xd9, 0x9c, 0x0a, 0x89, 0xe7, 0x8b, 0xc2, 0x32, 0xd8, 0xf0,
	0x08, 0x76, 0x47, 0x12, 0x67, 0xf2, 0x14, 0x67, 0x57, 0x49, 0x4c, 0x7f, 0x2c, 0xa9, 0x90, 0xe8,
	0x39, 0x00, 0x23, 0x34, 0x95, 0x6c, 0xc2, 0x68, 0xe6, 0x3b, 0x7d, 0x67, 0xd0, 0x8e, 0x4b, 0x37,
	0xe1, 0x6b, 0xe8, 0x9e, 0x63, 0x71, 0x27, 0x97, 0x3d, 0x40, 0xe5, 0x38, 0x62, 0xc1, 0x53, 0x41,
	0xc3, 0x08, 0x7a, 0x05, 0x91, 0xb9, 0x43, 0x01, 0x78, 0x99, 0xb5, 0x35, 0x8f, 0x17, 0xaf, 0xce,
	0x21, 0x06, 0xf4, 0x49, 0xd0, 0xec, 0x82, 0x0a, 0x81, 0xa7, 0x74, 0xcb, 0xd8, 0xe8, 0x10, 0x5a,
	0x73, 0xe3, 0xe1, 0xd7, 0xfa, 0xce, 0xa0, 0x33, 0xdc, 0x8d, 0x72, 0x85, 0xa2, 0x9c, 0x2a, 0x47,
	0x84, 0x07, 0xd0, 0xb1, 0x77, 0x1f, 0x58, 0x3a, 0x43, 0x08, 0x1a, 0x09, 0x4b, 0x67, 0x96, 0x55,
	0xdb, 0xe1, 0x6f, 0x07, 0x5a, 0x16, 0x83, 0xee, 0x43, 0x8d, 0x11, 0xfb, 0x5a, 0x63, 0x44, 0xe1,
	0x25, 0xfd, 0x29, 0x75, 0xa0, 0x76, 0xac, 0x6d, 0xf4, 0x12, 0x6a, 0x58, 0xfa, 0x75, 0x1d, 0x3a,
	0x88, 0xa6, 0x9c, 0x4f, 0x13, 0x6a, 0x32, 0x18, 0x2f, 0x27, 0xd1, 0xc7, 0xbc, 0x23, 0x71, 0x0d,
	0x4b, 0xe5, 0x4f, 0x78, 0x4a, 0xfd, 0x86, 0xae, 0x5c, 0xdb, 0x9a, 0x93, 0x13, 0xee, 0xbb, 0x96,
	0x93, 0x13, 0x8e, 0x0e, 0xc1, 0x55, 0xb9, 0x08, 0xbf, 0xd9, 0xaf, 0x0f, 0x3a, 0xc3, 0x87, 0x95,
	0x8a, 0x54, 0xf6, 0xb1, 0xc1, 0x84, 0x6f, 0xe1, 0xc1, 0x35, 0xd9, 0xac, 0xd2, 0x77, 0xd2, 0xe5,
	0xaf, 0x03, 0xee, 0x39, 0x1e, 0x33, 0x59, 0x29, 0x79, 0x0f, 0x5c, 0xc9, 0x64, 0x42, 0x6d, 0xcd,
	0xe6, 0xa0, 0xda, 0x48, 0x96, 0x19, 0x96, 0x8c, 0xa7, 0xba, 0xf4, 0x9d, 0x78, 0x75, 0x46, 0x6f,
	0xc0, 0x23, 0x14, 0x93, 0x84, 0xd9, 0x42, 0x37, 0xcb, 0xb2, 0xc2, 0xa2, 0x7d, 0x68, 0x92, 0x65,
	0xc6, 0xd2, 0xa9, 0x96, 0x62, 0x27, 0xb6, 0x27, 0x25, 0xd0, 0x65, 0xc6, 0x53, 0xbf, 0x69, 0x04,
	0x52, 0xb6, 0x8a, 0x91, 0x60, 0x21, 0x4f, 0x95, 0x98, 0xde, 0xed, 0x31, 0x72, 0x2c, 0xf2, 0xa1,
	0x95, 0x5e, 0x30, 0x21, 0x28, 0xf1, 0x5b, 0x3a, 0x48, 0x7e, 0x0c, 0xbf, 0x40, 0xf7, 0x84, 0x10,
	0xad, 0xc1, 0xb6, 0x93, 0xf7, 0x02, 0xdc, 0xef, 0x0a, 0x6f, 0xf5, 0xed, 0x16, 0xfa, 0x1a, 0x1a,
	0xf3, 0x1a, 0x1e, 0x43, 0xaf, 0x60, 0xb6, 0xcd, 0xd9, 0xd2, 0xf5, 0x1b, 0xdc, 0x1b, 0x51, 0x21,
	0x18, 0x4f, 0x47, 0x09, 0xbb, 0xa4, 0xaa, 0x19, 0x42, 0xfd, 0x67, 0x36, 0x19, 0x73, 0x40, 0x3d,
	0xa8, 0xd3, 0x94, 0xd8, 0x06, 0x29, 0xb3, 0xa0, 0xaf, 0x6f, 0xa4, 0xc7, 0xd0, 0xb2, 0xf4, 0x5b,
	0x33, 0x47, 0xd0, 0x14, 0x2a, 0x15, 0xe1, 0xd7, 0xf5, 0x68, 0xee, 0x17, 0xd4, 0xe5, 0x4c, 0x63,
	0x8b, 0x0a, 0x8f, 0xe1, 0x91, 0xde, 0x0c, 0x9f, 0x79, 0x36, 0xb3, 0x80, 0x6d, 0x97, 0xca, 0x19,
	0xf8, 0x55, 0xd7, 0x62, 0xb8, 0x85, 0xb9, 0xaa, 0x0e, 0x77, 0x8e, 0xcd, 0x11, 0xc3, 0x5f, 0x0e,
	0x34, 0xde, 0xe1, 0x6c, 0x8e, 0xce, 0x00, 0x8a, 0x35, 0x85, 0x9e, 0x94, 0x5c, 0xd6, 0x97, 0x64,
	0xf0, 0xf4, 0xe6, 0x47, 0x1b, 0xfe, 0x04, 0xbc, 0x7c, 0xb3, 0xa1, 0xc7, 0x65, 0x71, 0xaf, 0xad,
	0xcd, 0x20, 0xb8, 0xe9, 0xc9, 0x50, 0x0c, 0xff, 0x38, 0xe0, 0x1a, 0x82, 0xf7, 0xd0, 0x29, 0xfd,
	0xbf, 0xa8, 0x14, 0xb9, 0xba, 0x0d, 0x83, 0x67, 0xff, 0x79, 0x2d, 0x12, 0xcb, 0x67, 0xad, 0x9c,
	0xd8, 0xda, 0x64, 0x97, 0x13, 0xab, 0x8c, 0xe6, 0x57, 0xe8, 0xad, 0xcb, 0x8e, 0x0e, 0xd6, 0xd4,
	0xa8, 0x76, 0x33, 0x08, 0x37, 0x41, 0x0c, 0xf5, 0xb8, 0xa9, 0x21, 0x47, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xe6, 0xb0, 0xba, 0xf1, 0xce, 0x06, 0x00, 0x00,
}
