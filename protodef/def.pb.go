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
	HabitStat
	Habit
	AddHabitRequest
	AddHabitResponse
	SessionSlice
	Session
	StartWorkSessionRequest
	StartWorkSessionResponse
	CancelWorkSessionRequest
	CancelWorkSessionResponse
	RefuseWorkSessionRequest
	RefuseWorkSessionResponse
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

type HabitStat struct {
	// @inject_tag: db:"forget"
	Forget uint32 `protobuf:"varint,1,opt,name=forget" json:"forget,omitempty" db:"forget"`
	// @inject_tag: db:"nMissed"
	NMissed uint32 `protobuf:"varint,2,opt,name=nMissed" json:"nMissed,omitempty" db:"nMissed"`
}

func (m *HabitStat) Reset()                    { *m = HabitStat{} }
func (m *HabitStat) String() string            { return proto.CompactTextString(m) }
func (*HabitStat) ProtoMessage()               {}
func (*HabitStat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *HabitStat) GetForget() uint32 {
	if m != nil {
		return m.Forget
	}
	return 0
}

func (m *HabitStat) GetNMissed() uint32 {
	if m != nil {
		return m.NMissed
	}
	return 0
}

type Habit struct {
	// @inject_tag: db:"id"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" db:"id"`
	// @inject_tag: db:"title"
	Title    string                     `protobuf:"bytes,2,opt,name=title" json:"title,omitempty" db:"title"`
	Duration string                     `protobuf:"bytes,3,opt,name=duration" json:"duration,omitempty"`
	Deadline *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=deadline" json:"deadline,omitempty"`
	// @inject_tag: db:"cron"
	Cron     string                     `protobuf:"bytes,6,opt,name=cron" json:"cron,omitempty" db:"cron"`
	LastDone *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=lastDone" json:"lastDone,omitempty"`
	Stats    *HabitStat                 `protobuf:"bytes,8,opt,name=stats" json:"stats,omitempty"`
}

func (m *Habit) Reset()                    { *m = Habit{} }
func (m *Habit) String() string            { return proto.CompactTextString(m) }
func (*Habit) ProtoMessage()               {}
func (*Habit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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

func (m *Habit) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Habit) GetDeadline() *google_protobuf.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
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

func (m *Habit) GetStats() *HabitStat {
	if m != nil {
		return m.Stats
	}
	return nil
}

type AddHabitRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Habit      *Habit `protobuf:"bytes,2,opt,name=habit" json:"habit,omitempty"`
}

func (m *AddHabitRequest) Reset()                    { *m = AddHabitRequest{} }
func (m *AddHabitRequest) String() string            { return proto.CompactTextString(m) }
func (*AddHabitRequest) ProtoMessage()               {}
func (*AddHabitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
func (*AddHabitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AddHabitResponse) GetHabit() *Habit {
	if m != nil {
		return m.Habit
	}
	return nil
}

type SessionSlice struct {
	// @inject_tag: db:"start"
	Start *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=start" json:"start,omitempty" db:"start"`
	// @inject_tag: db:"end"
	End *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=end" json:"end,omitempty" db:"end"`
	// @inject_tag: db:"habit"
	Habit *Habit `protobuf:"bytes,3,opt,name=habit" json:"habit,omitempty" db:"habit"`
}

func (m *SessionSlice) Reset()                    { *m = SessionSlice{} }
func (m *SessionSlice) String() string            { return proto.CompactTextString(m) }
func (*SessionSlice) ProtoMessage()               {}
func (*SessionSlice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SessionSlice) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *SessionSlice) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *SessionSlice) GetHabit() *Habit {
	if m != nil {
		return m.Habit
	}
	return nil
}

type Session struct {
	// @inject_tag: db:"start"
	Start *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=start" json:"start,omitempty" db:"start"`
	// @inject_tag: db:"end"
	End *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=end" json:"end,omitempty" db:"end"`
	// @inject_tag: db:"slices"
	Slices []*SessionSlice `protobuf:"bytes,3,rep,name=slices" json:"slices,omitempty" db:"slices"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Session) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Session) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Session) GetSlices() []*SessionSlice {
	if m != nil {
		return m.Slices
	}
	return nil
}

type StartWorkSessionRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Duration   uint32 `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
}

func (m *StartWorkSessionRequest) Reset()                    { *m = StartWorkSessionRequest{} }
func (m *StartWorkSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*StartWorkSessionRequest) ProtoMessage()               {}
func (*StartWorkSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *StartWorkSessionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *StartWorkSessionRequest) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type StartWorkSessionResponse struct {
	Session *Session `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
}

func (m *StartWorkSessionResponse) Reset()                    { *m = StartWorkSessionResponse{} }
func (m *StartWorkSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*StartWorkSessionResponse) ProtoMessage()               {}
func (*StartWorkSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *StartWorkSessionResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type CancelWorkSessionRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *CancelWorkSessionRequest) Reset()                    { *m = CancelWorkSessionRequest{} }
func (m *CancelWorkSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelWorkSessionRequest) ProtoMessage()               {}
func (*CancelWorkSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *CancelWorkSessionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type CancelWorkSessionResponse struct {
}

func (m *CancelWorkSessionResponse) Reset()                    { *m = CancelWorkSessionResponse{} }
func (m *CancelWorkSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*CancelWorkSessionResponse) ProtoMessage()               {}
func (*CancelWorkSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type RefuseWorkSessionRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Index      uint32 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
}

func (m *RefuseWorkSessionRequest) Reset()                    { *m = RefuseWorkSessionRequest{} }
func (m *RefuseWorkSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*RefuseWorkSessionRequest) ProtoMessage()               {}
func (*RefuseWorkSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *RefuseWorkSessionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *RefuseWorkSessionRequest) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type RefuseWorkSessionResponse struct {
}

func (m *RefuseWorkSessionResponse) Reset()                    { *m = RefuseWorkSessionResponse{} }
func (m *RefuseWorkSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*RefuseWorkSessionResponse) ProtoMessage()               {}
func (*RefuseWorkSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func init() {
	proto.RegisterType((*StartDarylRequest)(nil), "protodef.StartDarylRequest")
	proto.RegisterType((*HasDarylRequest)(nil), "protodef.HasDarylRequest")
	proto.RegisterType((*StartDarylResponse)(nil), "protodef.StartDarylResponse")
	proto.RegisterType((*HasDarylResponse)(nil), "protodef.HasDarylResponse")
	proto.RegisterType((*UserMessageRequest)(nil), "protodef.UserMessageRequest")
	proto.RegisterType((*MessageLink)(nil), "protodef.MessageLink")
	proto.RegisterType((*Message)(nil), "protodef.Message")
	proto.RegisterType((*UserMessageResponse)(nil), "protodef.UserMessageResponse")
	proto.RegisterType((*HabitStat)(nil), "protodef.HabitStat")
	proto.RegisterType((*Habit)(nil), "protodef.Habit")
	proto.RegisterType((*AddHabitRequest)(nil), "protodef.AddHabitRequest")
	proto.RegisterType((*AddHabitResponse)(nil), "protodef.AddHabitResponse")
	proto.RegisterType((*SessionSlice)(nil), "protodef.SessionSlice")
	proto.RegisterType((*Session)(nil), "protodef.Session")
	proto.RegisterType((*StartWorkSessionRequest)(nil), "protodef.StartWorkSessionRequest")
	proto.RegisterType((*StartWorkSessionResponse)(nil), "protodef.StartWorkSessionResponse")
	proto.RegisterType((*CancelWorkSessionRequest)(nil), "protodef.CancelWorkSessionRequest")
	proto.RegisterType((*CancelWorkSessionResponse)(nil), "protodef.CancelWorkSessionResponse")
	proto.RegisterType((*RefuseWorkSessionRequest)(nil), "protodef.RefuseWorkSessionRequest")
	proto.RegisterType((*RefuseWorkSessionResponse)(nil), "protodef.RefuseWorkSessionResponse")
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
	CancelWorkSession(ctx context.Context, in *CancelWorkSessionRequest, opts ...grpc.CallOption) (*CancelWorkSessionResponse, error)
	RefuseWorkSession(ctx context.Context, in *RefuseWorkSessionRequest, opts ...grpc.CallOption) (*RefuseWorkSessionResponse, error)
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

func (c *darylClient) CancelWorkSession(ctx context.Context, in *CancelWorkSessionRequest, opts ...grpc.CallOption) (*CancelWorkSessionResponse, error) {
	out := new(CancelWorkSessionResponse)
	err := grpc.Invoke(ctx, "/protodef.Daryl/CancelWorkSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darylClient) RefuseWorkSession(ctx context.Context, in *RefuseWorkSessionRequest, opts ...grpc.CallOption) (*RefuseWorkSessionResponse, error) {
	out := new(RefuseWorkSessionResponse)
	err := grpc.Invoke(ctx, "/protodef.Daryl/RefuseWorkSession", in, out, c.cc, opts...)
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
	CancelWorkSession(context.Context, *CancelWorkSessionRequest) (*CancelWorkSessionResponse, error)
	RefuseWorkSession(context.Context, *RefuseWorkSessionRequest) (*RefuseWorkSessionResponse, error)
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

func _Daryl_CancelWorkSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelWorkSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServer).CancelWorkSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.Daryl/CancelWorkSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServer).CancelWorkSession(ctx, req.(*CancelWorkSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daryl_RefuseWorkSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefuseWorkSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServer).RefuseWorkSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.Daryl/RefuseWorkSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServer).RefuseWorkSession(ctx, req.(*RefuseWorkSessionRequest))
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
		{
			MethodName: "CancelWorkSession",
			Handler:    _Daryl_CancelWorkSession_Handler,
		},
		{
			MethodName: "RefuseWorkSession",
			Handler:    _Daryl_RefuseWorkSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "def.proto",
}

func init() { proto.RegisterFile("def.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0x56, 0xda, 0xa5, 0x1f, 0xa7, 0xef, 0xde, 0x6d, 0xde, 0xde, 0xbd, 0x59, 0x06, 0xa8, 0x33,
	0x42, 0x1a, 0x0c, 0x05, 0xd8, 0x24, 0x24, 0x90, 0xb8, 0x18, 0x4c, 0x6c, 0x42, 0x4c, 0x42, 0x29,
	0x13, 0x20, 0x71, 0xe3, 0xcd, 0x6e, 0xb1, 0x96, 0xc6, 0x25, 0x76, 0xa5, 0xed, 0x6f, 0x20, 0x71,
	0xc1, 0x8f, 0x80, 0xdf, 0xc7, 0x25, 0x8a, 0xed, 0x34, 0x59, 0xd3, 0x96, 0x0e, 0x89, 0xab, 0xf8,
	0xe3, 0x39, 0xcf, 0x79, 0xce, 0xe3, 0xe3, 0x18, 0x9a, 0x94, 0x75, 0x83, 0x41, 0x22, 0x94, 0x40,
	0x0d, 0xfd, 0xa1, 0xac, 0xeb, 0xb7, 0x07, 0xea, 0x72, 0xc0, 0xe4, 0x03, 0xc5, 0xfb, 0x4c, 0x2a,
	0xd2, 0x1f, 0xe4, 0x23, 0x83, 0xc5, 0x7b, 0xb0, 0xd2, 0x51, 0x24, 0x51, 0x07, 0x24, 0xb9, 0x8c,
	0x42, 0xf6, 0x79, 0xc8, 0xa4, 0x42, 0xb7, 0x00, 0x38, 0x65, 0xb1, 0xe2, 0x5d, 0xce, 0x12, 0xcf,
	0x69, 0x3b, 0xdb, 0xcd, 0xb0, 0xb0, 0x82, 0x1f, 0xc1, 0xd2, 0x11, 0x91, 0xd7, 0x0a, 0x59, 0x03,
	0x54, 0xcc, 0x23, 0x07, 0x22, 0x96, 0x0c, 0x07, 0xb0, 0x9c, 0x13, 0x99, 0x35, 0xe4, 0x43, 0x23,
	0xb1, 0x63, 0xcd, 0xd3, 0x08, 0x47, 0x73, 0x4c, 0x00, 0x9d, 0x48, 0x96, 0x1c, 0x33, 0x29, 0x49,
	0x8f, 0xcd, 0x99, 0x1b, 0xed, 0x40, 0xbd, 0x6f, 0x22, 0xbc, 0x4a, 0xdb, 0xd9, 0x6e, 0xed, 0xae,
	0x04, 0x99, 0x43, 0x41, 0x46, 0x95, 0x21, 0xf0, 0x16, 0xb4, 0xec, 0xda, 0x6b, 0x1e, 0x9f, 0x23,
	0x04, 0x0b, 0x11, 0x8f, 0xcf, 0x2d, 0xab, 0x1e, 0xe3, 0x1f, 0x0e, 0xd4, 0x2d, 0x06, 0xfd, 0x0b,
	0x15, 0x4e, 0xed, 0x6e, 0x85, 0xd3, 0x14, 0xaf, 0xd8, 0x85, 0xd2, 0x89, 0x9a, 0xa1, 0x1e, 0xa3,
	0x7b, 0x50, 0x21, 0xca, 0xab, 0xea, 0xd4, 0x7e, 0xd0, 0x13, 0xa2, 0x17, 0x31, 0xa3, 0xe0, 0x74,
	0xd8, 0x0d, 0xde, 0x66, 0x27, 0x12, 0x56, 0x88, 0x4a, 0xe3, 0xa9, 0x88, 0x99, 0xb7, 0xa0, 0x2b,
	0xd7, 0x63, 0xcd, 0x29, 0xa8, 0xf0, 0x5c, 0xcb, 0x29, 0xa8, 0x40, 0x3b, 0xe0, 0xa6, 0x5a, 0xa4,
	0x57, 0x6b, 0x57, 0xb7, 0x5b, 0xbb, 0xff, 0x95, 0x2a, 0x4a, 0xd5, 0x87, 0x06, 0x83, 0x9f, 0xc3,
	0xea, 0x15, 0xdb, 0xac, 0xd3, 0xd7, 0xf2, 0xe5, 0x19, 0x34, 0x8f, 0xc8, 0x29, 0x57, 0x1d, 0x45,
	0x14, 0x5a, 0x87, 0x5a, 0x57, 0x24, 0x3d, 0xa6, 0x74, 0xe5, 0x8b, 0xa1, 0x9d, 0x21, 0x0f, 0xea,
	0xf1, 0x31, 0x97, 0x92, 0x51, 0xcd, 0xb8, 0x18, 0x66, 0x53, 0xfc, 0xd3, 0x01, 0x57, 0xc7, 0x97,
	0x1c, 0x5b, 0x03, 0x57, 0x71, 0x15, 0x31, 0x6b, 0x99, 0x99, 0xa4, 0x5d, 0x40, 0x87, 0x09, 0x51,
	0x5c, 0xc4, 0xda, 0xb9, 0x66, 0x38, 0x9a, 0xa3, 0xc7, 0xd0, 0xa0, 0x8c, 0xd0, 0x88, 0x5b, 0x9f,
	0x66, 0xbb, 0x3a, 0xc2, 0xa6, 0x3e, 0x9e, 0x25, 0x22, 0xf6, 0x6a, 0xc6, 0xc7, 0x74, 0x9c, 0x72,
	0x45, 0x44, 0xaa, 0x83, 0xd4, 0xf3, 0xfa, 0xef, 0xb9, 0x32, 0x2c, 0xba, 0x0b, 0xae, 0x54, 0x44,
	0x49, 0xaf, 0xa1, 0x83, 0x56, 0x73, 0xe7, 0x46, 0x2e, 0x85, 0x06, 0x81, 0xdf, 0xc3, 0xd2, 0x3e,
	0xa5, 0x7a, 0x79, 0xde, 0x8e, 0xbd, 0x03, 0xee, 0xa7, 0x14, 0x6f, 0xcf, 0x65, 0x69, 0x8c, 0x3d,
	0x34, 0xbb, 0xf8, 0x09, 0x2c, 0xe7, 0xcc, 0xf6, 0x50, 0xe7, 0x0c, 0xfd, 0xea, 0xc0, 0x3f, 0x1d,
	0x26, 0x25, 0x17, 0x71, 0x27, 0xe2, 0x67, 0x0c, 0x3d, 0xd4, 0x05, 0x25, 0xe6, 0x44, 0x67, 0xbb,
	0x60, 0x80, 0xe8, 0x3e, 0x54, 0x59, 0x4c, 0x6d, 0x9e, 0x59, 0xf8, 0x14, 0x96, 0xeb, 0xaa, 0xce,
	0xd4, 0xf5, 0xcd, 0x81, 0xba, 0xd5, 0xf5, 0xd7, 0x25, 0x05, 0x50, 0x93, 0x69, 0xed, 0xd2, 0xab,
	0xea, 0x4b, 0xb4, 0x9e, 0x6b, 0x2a, 0x5a, 0x13, 0x5a, 0x14, 0x3e, 0x81, 0xff, 0xf5, 0x3f, 0xec,
	0x9d, 0x48, 0xce, 0x2d, 0x60, 0xde, 0x03, 0x2d, 0xb6, 0xb3, 0xb9, 0x19, 0xa3, 0x39, 0x3e, 0x04,
	0xaf, 0x4c, 0x9b, 0x5f, 0x51, 0x69, 0x96, 0xca, 0x57, 0x34, 0xc3, 0x66, 0x08, 0xfc, 0x14, 0xbc,
	0x17, 0x24, 0x3e, 0x63, 0xd1, 0xf5, 0x05, 0xe2, 0x4d, 0xd8, 0x98, 0x10, 0x6b, 0x7f, 0xbb, 0x6f,
	0xc0, 0x0b, 0x59, 0x77, 0x28, 0xd9, 0x1f, 0x54, 0xbe, 0x06, 0x2e, 0x8f, 0x29, 0xbb, 0xb0, 0x65,
	0x9b, 0x49, 0x9a, 0x6e, 0x02, 0xa3, 0x49, 0xb7, 0xfb, 0xc5, 0x81, 0x85, 0x97, 0x24, 0xe9, 0xa3,
	0x43, 0x80, 0xfc, 0xd1, 0x40, 0x9b, 0x85, 0xd2, 0xc7, 0x9f, 0x2c, 0xff, 0xc6, 0xe4, 0x4d, 0x6b,
	0xe3, 0x3e, 0x34, 0xb2, 0x77, 0x06, 0x6d, 0x14, 0x3b, 0xef, 0xca, 0x23, 0xe6, 0xfb, 0x93, 0xb6,
	0xac, 0xa8, 0xef, 0x55, 0x70, 0x0d, 0xc1, 0x2b, 0x68, 0x15, 0xfe, 0xa6, 0xa8, 0x90, 0xb9, 0xfc,
	0x36, 0xf9, 0x37, 0xa7, 0xec, 0xe6, 0xc2, 0xb2, 0x1b, 0x5c, 0x14, 0x36, 0xf6, 0xbf, 0x28, 0x0a,
	0x2b, 0x5d, 0xf8, 0x0f, 0xb0, 0x3c, 0xde, 0x3e, 0x68, 0x6b, 0xcc, 0x8d, 0xf2, 0xb9, 0xf9, 0x78,
	0x16, 0xc4, 0x52, 0x7f, 0x84, 0x95, 0x52, 0x53, 0xa0, 0x42, 0xe0, 0xb4, 0x6e, 0xf3, 0x6f, 0xcf,
	0xc4, 0xe4, 0xec, 0xa5, 0x1e, 0x28, 0xb2, 0x4f, 0x6b, 0xb9, 0x22, 0xfb, 0xd4, 0x26, 0x3a, 0xad,
	0x69, 0xcc, 0xde, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xb3, 0x34, 0x73, 0x18, 0x09, 0x00,
	0x00,
}
