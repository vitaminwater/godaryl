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
	Daryl
	UserMessageRequest
	Message
	UserMessageResponse
	Trigger
	Habit
	AddHabitRequest
	AddHabitResponse
	SessionSlice
	Session
	SessionConfig
	StartWorkSessionRequest
	StartWorkSessionResponse
	CancelWorkSessionRequest
	CancelWorkSessionResponse
	SessionSliceIndex
	RefuseSessionSliceRequest
	RefuseSessionSliceResponse
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
	Daryl *Daryl `protobuf:"bytes,1,opt,name=daryl" json:"daryl,omitempty"`
}

func (m *StartDarylRequest) Reset()                    { *m = StartDarylRequest{} }
func (m *StartDarylRequest) String() string            { return proto.CompactTextString(m) }
func (*StartDarylRequest) ProtoMessage()               {}
func (*StartDarylRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartDarylRequest) GetDaryl() *Daryl {
	if m != nil {
		return m.Daryl
	}
	return nil
}

type HasDarylRequest struct {
	DarylIdentifier string `protobuf:"bytes,1,opt,name=daryl_identifier,json=darylIdentifier" json:"daryl_identifier,omitempty"`
}

func (m *HasDarylRequest) Reset()                    { *m = HasDarylRequest{} }
func (m *HasDarylRequest) String() string            { return proto.CompactTextString(m) }
func (*HasDarylRequest) ProtoMessage()               {}
func (*HasDarylRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HasDarylRequest) GetDarylIdentifier() string {
	if m != nil {
		return m.DarylIdentifier
	}
	return ""
}

type StartDarylResponse struct {
	Daryl *Daryl `protobuf:"bytes,1,opt,name=daryl" json:"daryl,omitempty"`
}

func (m *StartDarylResponse) Reset()                    { *m = StartDarylResponse{} }
func (m *StartDarylResponse) String() string            { return proto.CompactTextString(m) }
func (*StartDarylResponse) ProtoMessage()               {}
func (*StartDarylResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StartDarylResponse) GetDaryl() *Daryl {
	if m != nil {
		return m.Daryl
	}
	return nil
}

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

type Daryl struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *Daryl) Reset()                    { *m = Daryl{} }
func (m *Daryl) String() string            { return proto.CompactTextString(m) }
func (*Daryl) ProtoMessage()               {}
func (*Daryl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Daryl) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Daryl) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Daryl) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserMessageRequest struct {
	DarylIdentifier string   `protobuf:"bytes,1,opt,name=daryl_identifier,json=darylIdentifier" json:"daryl_identifier,omitempty"`
	Message         *Message `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *UserMessageRequest) Reset()                    { *m = UserMessageRequest{} }
func (m *UserMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*UserMessageRequest) ProtoMessage()               {}
func (*UserMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserMessageRequest) GetDarylIdentifier() string {
	if m != nil {
		return m.DarylIdentifier
	}
	return ""
}

func (m *UserMessageRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type Message struct {
	Id              string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Text            string                     `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	HabitIdentifier string                     `protobuf:"bytes,3,opt,name=habit_identifier,json=habitIdentifier" json:"habit_identifier,omitempty"`
	At              *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=at" json:"at,omitempty"`
	Attrs           []byte                     `protobuf:"bytes,5,opt,name=attrs,proto3" json:"attrs,omitempty"`
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

func (m *Message) GetHabitIdentifier() string {
	if m != nil {
		return m.HabitIdentifier
	}
	return ""
}

func (m *Message) GetAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.At
	}
	return nil
}

func (m *Message) GetAttrs() []byte {
	if m != nil {
		return m.Attrs
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

type Trigger struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Engine string `protobuf:"bytes,3,opt,name=engine" json:"engine,omitempty"`
	Params []byte `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *Trigger) Reset()                    { *m = Trigger{} }
func (m *Trigger) String() string            { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()               {}
func (*Trigger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Trigger) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trigger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Trigger) GetEngine() string {
	if m != nil {
		return m.Engine
	}
	return ""
}

func (m *Trigger) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

type Habit struct {
	Id       string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title    string     `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Duration string     `protobuf:"bytes,3,opt,name=duration" json:"duration,omitempty"`
	Triggers []*Trigger `protobuf:"bytes,4,rep,name=triggers" json:"triggers,omitempty"`
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

func (m *Habit) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type AddHabitRequest struct {
	DarylIdentifier string `protobuf:"bytes,1,opt,name=daryl_identifier,json=darylIdentifier" json:"daryl_identifier,omitempty"`
	Habit           *Habit `protobuf:"bytes,2,opt,name=habit" json:"habit,omitempty"`
}

func (m *AddHabitRequest) Reset()                    { *m = AddHabitRequest{} }
func (m *AddHabitRequest) String() string            { return proto.CompactTextString(m) }
func (*AddHabitRequest) ProtoMessage()               {}
func (*AddHabitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AddHabitRequest) GetDarylIdentifier() string {
	if m != nil {
		return m.DarylIdentifier
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
	Start *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	Habit *Habit                     `protobuf:"bytes,3,opt,name=habit" json:"habit,omitempty"`
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
	Start  *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End    *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	Slices []*SessionSlice            `protobuf:"bytes,3,rep,name=slices" json:"slices,omitempty"`
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

type SessionConfig struct {
	Duration string `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
}

func (m *SessionConfig) Reset()                    { *m = SessionConfig{} }
func (m *SessionConfig) String() string            { return proto.CompactTextString(m) }
func (*SessionConfig) ProtoMessage()               {}
func (*SessionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SessionConfig) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

type StartWorkSessionRequest struct {
	DarylIdentifier string         `protobuf:"bytes,1,opt,name=daryl_identifier,json=darylIdentifier" json:"daryl_identifier,omitempty"`
	Config          *SessionConfig `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *StartWorkSessionRequest) Reset()                    { *m = StartWorkSessionRequest{} }
func (m *StartWorkSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*StartWorkSessionRequest) ProtoMessage()               {}
func (*StartWorkSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *StartWorkSessionRequest) GetDarylIdentifier() string {
	if m != nil {
		return m.DarylIdentifier
	}
	return ""
}

func (m *StartWorkSessionRequest) GetConfig() *SessionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type StartWorkSessionResponse struct {
	Session *Session `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
}

func (m *StartWorkSessionResponse) Reset()                    { *m = StartWorkSessionResponse{} }
func (m *StartWorkSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*StartWorkSessionResponse) ProtoMessage()               {}
func (*StartWorkSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *StartWorkSessionResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type CancelWorkSessionRequest struct {
	DarylIdentifier string `protobuf:"bytes,1,opt,name=daryl_identifier,json=darylIdentifier" json:"daryl_identifier,omitempty"`
}

func (m *CancelWorkSessionRequest) Reset()                    { *m = CancelWorkSessionRequest{} }
func (m *CancelWorkSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelWorkSessionRequest) ProtoMessage()               {}
func (*CancelWorkSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *CancelWorkSessionRequest) GetDarylIdentifier() string {
	if m != nil {
		return m.DarylIdentifier
	}
	return ""
}

type CancelWorkSessionResponse struct {
}

func (m *CancelWorkSessionResponse) Reset()                    { *m = CancelWorkSessionResponse{} }
func (m *CancelWorkSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*CancelWorkSessionResponse) ProtoMessage()               {}
func (*CancelWorkSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type SessionSliceIndex struct {
	Index uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (m *SessionSliceIndex) Reset()                    { *m = SessionSliceIndex{} }
func (m *SessionSliceIndex) String() string            { return proto.CompactTextString(m) }
func (*SessionSliceIndex) ProtoMessage()               {}
func (*SessionSliceIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *SessionSliceIndex) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type RefuseSessionSliceRequest struct {
	DarylIdentifier string             `protobuf:"bytes,1,opt,name=daryl_identifier,json=darylIdentifier" json:"daryl_identifier,omitempty"`
	Index           *SessionSliceIndex `protobuf:"bytes,2,opt,name=index" json:"index,omitempty"`
}

func (m *RefuseSessionSliceRequest) Reset()                    { *m = RefuseSessionSliceRequest{} }
func (m *RefuseSessionSliceRequest) String() string            { return proto.CompactTextString(m) }
func (*RefuseSessionSliceRequest) ProtoMessage()               {}
func (*RefuseSessionSliceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *RefuseSessionSliceRequest) GetDarylIdentifier() string {
	if m != nil {
		return m.DarylIdentifier
	}
	return ""
}

func (m *RefuseSessionSliceRequest) GetIndex() *SessionSliceIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

type RefuseSessionSliceResponse struct {
}

func (m *RefuseSessionSliceResponse) Reset()                    { *m = RefuseSessionSliceResponse{} }
func (m *RefuseSessionSliceResponse) String() string            { return proto.CompactTextString(m) }
func (*RefuseSessionSliceResponse) ProtoMessage()               {}
func (*RefuseSessionSliceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func init() {
	proto.RegisterType((*StartDarylRequest)(nil), "protodef.StartDarylRequest")
	proto.RegisterType((*HasDarylRequest)(nil), "protodef.HasDarylRequest")
	proto.RegisterType((*StartDarylResponse)(nil), "protodef.StartDarylResponse")
	proto.RegisterType((*HasDarylResponse)(nil), "protodef.HasDarylResponse")
	proto.RegisterType((*Daryl)(nil), "protodef.Daryl")
	proto.RegisterType((*UserMessageRequest)(nil), "protodef.UserMessageRequest")
	proto.RegisterType((*Message)(nil), "protodef.Message")
	proto.RegisterType((*UserMessageResponse)(nil), "protodef.UserMessageResponse")
	proto.RegisterType((*Trigger)(nil), "protodef.Trigger")
	proto.RegisterType((*Habit)(nil), "protodef.Habit")
	proto.RegisterType((*AddHabitRequest)(nil), "protodef.AddHabitRequest")
	proto.RegisterType((*AddHabitResponse)(nil), "protodef.AddHabitResponse")
	proto.RegisterType((*SessionSlice)(nil), "protodef.SessionSlice")
	proto.RegisterType((*Session)(nil), "protodef.Session")
	proto.RegisterType((*SessionConfig)(nil), "protodef.SessionConfig")
	proto.RegisterType((*StartWorkSessionRequest)(nil), "protodef.StartWorkSessionRequest")
	proto.RegisterType((*StartWorkSessionResponse)(nil), "protodef.StartWorkSessionResponse")
	proto.RegisterType((*CancelWorkSessionRequest)(nil), "protodef.CancelWorkSessionRequest")
	proto.RegisterType((*CancelWorkSessionResponse)(nil), "protodef.CancelWorkSessionResponse")
	proto.RegisterType((*SessionSliceIndex)(nil), "protodef.SessionSliceIndex")
	proto.RegisterType((*RefuseSessionSliceRequest)(nil), "protodef.RefuseSessionSliceRequest")
	proto.RegisterType((*RefuseSessionSliceResponse)(nil), "protodef.RefuseSessionSliceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FarmService service

type FarmServiceClient interface {
	StartDaryl(ctx context.Context, in *StartDarylRequest, opts ...grpc.CallOption) (*StartDarylResponse, error)
	HasDaryl(ctx context.Context, in *HasDarylRequest, opts ...grpc.CallOption) (*HasDarylResponse, error)
}

type farmServiceClient struct {
	cc *grpc.ClientConn
}

func NewFarmServiceClient(cc *grpc.ClientConn) FarmServiceClient {
	return &farmServiceClient{cc}
}

func (c *farmServiceClient) StartDaryl(ctx context.Context, in *StartDarylRequest, opts ...grpc.CallOption) (*StartDarylResponse, error) {
	out := new(StartDarylResponse)
	err := grpc.Invoke(ctx, "/protodef.FarmService/StartDaryl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmServiceClient) HasDaryl(ctx context.Context, in *HasDarylRequest, opts ...grpc.CallOption) (*HasDarylResponse, error) {
	out := new(HasDarylResponse)
	err := grpc.Invoke(ctx, "/protodef.FarmService/HasDaryl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FarmService service

type FarmServiceServer interface {
	StartDaryl(context.Context, *StartDarylRequest) (*StartDarylResponse, error)
	HasDaryl(context.Context, *HasDarylRequest) (*HasDarylResponse, error)
}

func RegisterFarmServiceServer(s *grpc.Server, srv FarmServiceServer) {
	s.RegisterService(&_FarmService_serviceDesc, srv)
}

func _FarmService_StartDaryl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDarylRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServiceServer).StartDaryl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.FarmService/StartDaryl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServiceServer).StartDaryl(ctx, req.(*StartDarylRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FarmService_HasDaryl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasDarylRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServiceServer).HasDaryl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.FarmService/HasDaryl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServiceServer).HasDaryl(ctx, req.(*HasDarylRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FarmService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protodef.FarmService",
	HandlerType: (*FarmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartDaryl",
			Handler:    _FarmService_StartDaryl_Handler,
		},
		{
			MethodName: "HasDaryl",
			Handler:    _FarmService_HasDaryl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "def.proto",
}

// Client API for DarylService service

type DarylServiceClient interface {
	UserMessage(ctx context.Context, in *UserMessageRequest, opts ...grpc.CallOption) (*UserMessageResponse, error)
	AddHabit(ctx context.Context, in *AddHabitRequest, opts ...grpc.CallOption) (*AddHabitResponse, error)
	StartWorkSession(ctx context.Context, in *StartWorkSessionRequest, opts ...grpc.CallOption) (*StartWorkSessionResponse, error)
	CancelWorkSession(ctx context.Context, in *CancelWorkSessionRequest, opts ...grpc.CallOption) (*CancelWorkSessionResponse, error)
	RefuseSessionSlice(ctx context.Context, in *RefuseSessionSliceRequest, opts ...grpc.CallOption) (*RefuseSessionSliceResponse, error)
}

type darylServiceClient struct {
	cc *grpc.ClientConn
}

func NewDarylServiceClient(cc *grpc.ClientConn) DarylServiceClient {
	return &darylServiceClient{cc}
}

func (c *darylServiceClient) UserMessage(ctx context.Context, in *UserMessageRequest, opts ...grpc.CallOption) (*UserMessageResponse, error) {
	out := new(UserMessageResponse)
	err := grpc.Invoke(ctx, "/protodef.DarylService/UserMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darylServiceClient) AddHabit(ctx context.Context, in *AddHabitRequest, opts ...grpc.CallOption) (*AddHabitResponse, error) {
	out := new(AddHabitResponse)
	err := grpc.Invoke(ctx, "/protodef.DarylService/AddHabit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darylServiceClient) StartWorkSession(ctx context.Context, in *StartWorkSessionRequest, opts ...grpc.CallOption) (*StartWorkSessionResponse, error) {
	out := new(StartWorkSessionResponse)
	err := grpc.Invoke(ctx, "/protodef.DarylService/StartWorkSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darylServiceClient) CancelWorkSession(ctx context.Context, in *CancelWorkSessionRequest, opts ...grpc.CallOption) (*CancelWorkSessionResponse, error) {
	out := new(CancelWorkSessionResponse)
	err := grpc.Invoke(ctx, "/protodef.DarylService/CancelWorkSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *darylServiceClient) RefuseSessionSlice(ctx context.Context, in *RefuseSessionSliceRequest, opts ...grpc.CallOption) (*RefuseSessionSliceResponse, error) {
	out := new(RefuseSessionSliceResponse)
	err := grpc.Invoke(ctx, "/protodef.DarylService/RefuseSessionSlice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DarylService service

type DarylServiceServer interface {
	UserMessage(context.Context, *UserMessageRequest) (*UserMessageResponse, error)
	AddHabit(context.Context, *AddHabitRequest) (*AddHabitResponse, error)
	StartWorkSession(context.Context, *StartWorkSessionRequest) (*StartWorkSessionResponse, error)
	CancelWorkSession(context.Context, *CancelWorkSessionRequest) (*CancelWorkSessionResponse, error)
	RefuseSessionSlice(context.Context, *RefuseSessionSliceRequest) (*RefuseSessionSliceResponse, error)
}

func RegisterDarylServiceServer(s *grpc.Server, srv DarylServiceServer) {
	s.RegisterService(&_DarylService_serviceDesc, srv)
}

func _DarylService_UserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServiceServer).UserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.DarylService/UserMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServiceServer).UserMessage(ctx, req.(*UserMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DarylService_AddHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHabitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServiceServer).AddHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.DarylService/AddHabit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServiceServer).AddHabit(ctx, req.(*AddHabitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DarylService_StartWorkSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartWorkSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServiceServer).StartWorkSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.DarylService/StartWorkSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServiceServer).StartWorkSession(ctx, req.(*StartWorkSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DarylService_CancelWorkSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelWorkSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServiceServer).CancelWorkSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.DarylService/CancelWorkSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServiceServer).CancelWorkSession(ctx, req.(*CancelWorkSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DarylService_RefuseSessionSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefuseSessionSliceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarylServiceServer).RefuseSessionSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodef.DarylService/RefuseSessionSlice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarylServiceServer).RefuseSessionSlice(ctx, req.(*RefuseSessionSliceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DarylService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protodef.DarylService",
	HandlerType: (*DarylServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserMessage",
			Handler:    _DarylService_UserMessage_Handler,
		},
		{
			MethodName: "AddHabit",
			Handler:    _DarylService_AddHabit_Handler,
		},
		{
			MethodName: "StartWorkSession",
			Handler:    _DarylService_StartWorkSession_Handler,
		},
		{
			MethodName: "CancelWorkSession",
			Handler:    _DarylService_CancelWorkSession_Handler,
		},
		{
			MethodName: "RefuseSessionSlice",
			Handler:    _DarylService_RefuseSessionSlice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "def.proto",
}

func init() { proto.RegisterFile("def.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 786 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0x96, 0x13, 0x9c, 0x84, 0x93, 0xb0, 0x49, 0x66, 0x11, 0x18, 0xc3, 0x4a, 0xec, 0xec, 0x22,
	0xc1, 0xb2, 0x0d, 0x2d, 0xbd, 0xea, 0xcf, 0x0d, 0xa5, 0x2d, 0x50, 0xa9, 0x37, 0x0e, 0x55, 0x55,
	0xa9, 0x15, 0x1a, 0xe2, 0x49, 0x3a, 0x6a, 0x62, 0xa7, 0x33, 0x93, 0x16, 0x1e, 0xa4, 0x17, 0xf4,
	0x45, 0xfa, 0x7a, 0x95, 0xe7, 0x27, 0x76, 0xe2, 0x10, 0x01, 0x55, 0xaf, 0xec, 0x73, 0xe6, 0x3b,
	0xe7, 0x7c, 0xe7, 0x6f, 0x06, 0x16, 0x43, 0xda, 0x6d, 0x0d, 0x79, 0x2c, 0x63, 0x54, 0x51, 0x9f,
	0x90, 0x76, 0xfd, 0xcd, 0xa1, 0xbc, 0x1c, 0x52, 0xb1, 0x27, 0xd9, 0x80, 0x0a, 0x49, 0x06, 0xc3,
	0xf4, 0x4f, 0x63, 0xf1, 0x63, 0x68, 0xb6, 0x25, 0xe1, 0xf2, 0x39, 0xe1, 0x97, 0xfd, 0x80, 0x7e,
	0x1e, 0x51, 0x21, 0xd1, 0x16, 0xb8, 0x61, 0x22, 0x7b, 0xce, 0xa6, 0xb3, 0x5d, 0xdd, 0xaf, 0xb7,
	0xac, 0xc3, 0x96, 0x86, 0xe9, 0x53, 0xfc, 0x14, 0xea, 0xc7, 0x44, 0x4c, 0x58, 0xee, 0x40, 0x43,
	0x9d, 0x9d, 0xb1, 0x90, 0x46, 0x92, 0x75, 0x19, 0xe5, 0xca, 0xc9, 0x62, 0x50, 0x57, 0xfa, 0x93,
	0xb1, 0x1a, 0x3f, 0x01, 0x94, 0x8d, 0x2c, 0x86, 0x71, 0x24, 0xe8, 0x4d, 0x43, 0xb7, 0xa0, 0x91,
	0x86, 0x36, 0xa6, 0x3e, 0x54, 0xb8, 0xf9, 0x57, 0xd6, 0x95, 0x60, 0x2c, 0xe3, 0x23, 0x70, 0x15,
	0x18, 0xfd, 0x01, 0x05, 0x16, 0x1a, 0x4a, 0x05, 0x16, 0x22, 0x04, 0x0b, 0x11, 0x19, 0x50, 0xaf,
	0xa0, 0x34, 0xea, 0x3f, 0x71, 0x34, 0x24, 0x42, 0x7c, 0x8d, 0x79, 0xe8, 0x15, 0x95, 0x7e, 0x2c,
	0xe3, 0x3e, 0xa0, 0x37, 0x82, 0xf2, 0xd7, 0x54, 0x08, 0xd2, 0xa3, 0xb7, 0x4f, 0x1b, 0xed, 0x42,
	0x79, 0xa0, 0x8d, 0x55, 0xcc, 0xea, 0x7e, 0x33, 0x4d, 0xd1, 0x7a, 0xb5, 0x08, 0xfc, 0xdd, 0x81,
	0xb2, 0x51, 0xce, 0x62, 0x2e, 0xe9, 0x85, 0xb4, 0xcc, 0x93, 0xff, 0x84, 0xc7, 0x47, 0x72, 0xce,
	0x64, 0x96, 0x87, 0xce, 0xa0, 0xae, 0xf4, 0x19, 0x1e, 0xff, 0x41, 0x81, 0x48, 0x6f, 0x41, 0x51,
	0xf0, 0x5b, 0xbd, 0x38, 0xee, 0xf5, 0xa9, 0x66, 0x72, 0x3e, 0xea, 0xb6, 0x4e, 0xed, 0x98, 0x04,
	0x05, 0x22, 0xd1, 0x32, 0xb8, 0x44, 0x4a, 0x2e, 0x3c, 0x77, 0xd3, 0xd9, 0xae, 0x05, 0x5a, 0xc0,
	0xcf, 0xe0, 0xcf, 0x89, 0x52, 0x98, 0x36, 0xdc, 0x2a, 0xc1, 0x0f, 0x50, 0x3e, 0xe5, 0xac, 0xd7,
	0xa3, 0xfc, 0x46, 0x9d, 0x59, 0x81, 0x12, 0x8d, 0x7a, 0x2c, 0xa2, 0x26, 0x2b, 0x23, 0x25, 0xfa,
	0x21, 0xe1, 0x64, 0x20, 0x54, 0x42, 0xb5, 0xc0, 0x48, 0xf8, 0x02, 0xdc, 0xe3, 0x24, 0xef, 0x9c,
	0xf3, 0x65, 0x70, 0x25, 0x93, 0x7d, 0xeb, 0x5d, 0x0b, 0x49, 0xe3, 0xc3, 0x11, 0x27, 0x92, 0xc5,
	0x91, 0x6d, 0xbc, 0x95, 0xd1, 0x3d, 0xa8, 0x48, 0xcd, 0x34, 0x09, 0x52, 0x9c, 0xcc, 0xcb, 0xe4,
	0x10, 0x8c, 0x21, 0xb8, 0x03, 0xf5, 0x83, 0x30, 0x54, 0xc1, 0xef, 0x30, 0x24, 0x5b, 0xe0, 0xaa,
	0x7e, 0x99, 0x0a, 0x66, 0xb6, 0x40, 0x7b, 0xd4, 0xa7, 0xf8, 0x11, 0x34, 0xd2, 0x20, 0xe9, 0x02,
	0xdd, 0xc4, 0xf4, 0x9b, 0x03, 0xb5, 0x36, 0x15, 0x82, 0xc5, 0x51, 0xbb, 0xcf, 0x3a, 0x14, 0xdd,
	0x07, 0x57, 0x24, 0xeb, 0x68, 0x16, 0x6f, 0xde, 0x48, 0x68, 0x20, 0xfa, 0x1f, 0x8a, 0x34, 0x0a,
	0x4d, 0x9c, 0x79, 0xf8, 0x04, 0x96, 0xf2, 0x2a, 0xce, 0xe5, 0x75, 0xe5, 0x40, 0xd9, 0xf0, 0xfa,
	0xed, 0x94, 0x5a, 0x50, 0x12, 0x49, 0xee, 0xc2, 0x2b, 0xaa, 0x86, 0xae, 0xa4, 0x9c, 0xb2, 0xa5,
	0x09, 0x0c, 0x0a, 0xef, 0xc2, 0x92, 0xd1, 0x1f, 0xc6, 0x51, 0x97, 0xf5, 0x26, 0xe6, 0xc5, 0x99,
	0x9c, 0x17, 0x3c, 0x82, 0x55, 0x75, 0xbd, 0xbd, 0x8d, 0xf9, 0x27, 0x63, 0x75, 0x87, 0x41, 0xd8,
	0x83, 0x52, 0x47, 0xc5, 0x32, 0x39, 0xad, 0xe6, 0x28, 0x6a, 0x2a, 0x81, 0x81, 0xe1, 0x23, 0xf0,
	0xf2, 0x61, 0xd3, 0xcd, 0x14, 0x5a, 0x95, 0xdf, 0x4c, 0x8b, 0xb5, 0x08, 0xfc, 0x02, 0xbc, 0x43,
	0x12, 0x75, 0x68, 0xff, 0x97, 0x12, 0xc0, 0xeb, 0xb0, 0x36, 0xc3, 0x8d, 0xb9, 0x95, 0x77, 0xa0,
	0x99, 0x2d, 0xf4, 0x49, 0x14, 0xd2, 0x8b, 0x64, 0x35, 0x59, 0xf2, 0xa3, 0x3c, 0x2e, 0x05, 0x5a,
	0xc0, 0x97, 0xb0, 0x16, 0xd0, 0xee, 0x48, 0xd0, 0x89, 0xce, 0xdc, 0xbe, 0xa0, 0x0f, 0xac, 0x77,
	0x5d, 0x81, 0xf5, 0xd9, 0x2d, 0x57, 0x4c, 0x6c, 0xe8, 0x0d, 0xf0, 0x67, 0x85, 0xd6, 0x39, 0xec,
	0x5f, 0x39, 0x50, 0x7d, 0x49, 0xf8, 0xa0, 0x4d, 0xf9, 0x97, 0x64, 0x8f, 0x8e, 0x00, 0xd2, 0x67,
	0x0d, 0x65, 0xfd, 0x4f, 0x3f, 0xb3, 0xfe, 0xc6, 0xec, 0x43, 0xd3, 0xad, 0x03, 0xa8, 0xd8, 0x27,
	0x0e, 0xad, 0x65, 0xb7, 0x65, 0xe2, 0xc5, 0xf5, 0xfd, 0x59, 0x47, 0x86, 0xdb, 0x8f, 0x22, 0xd4,
	0x94, 0xc6, 0x92, 0x7b, 0x05, 0xd5, 0xcc, 0x95, 0x8d, 0x32, 0x04, 0xf2, 0x8f, 0x9a, 0xff, 0xd7,
	0x35, 0xa7, 0x29, 0x3f, 0x7b, 0xf9, 0x64, 0xf9, 0x4d, 0xdd, 0x7a, 0x59, 0x7e, 0xb9, 0xbb, 0xea,
	0x1d, 0x34, 0xa6, 0x87, 0x15, 0xfd, 0x3d, 0x55, 0x94, 0xfc, 0xf8, 0xf9, 0x78, 0x1e, 0xc4, 0xb8,
	0x7e, 0x0f, 0xcd, 0xdc, 0xdc, 0xa1, 0x8c, 0xe1, 0x75, 0xb3, 0xed, 0xff, 0x33, 0x17, 0x63, 0xbc,
	0x9f, 0x01, 0xca, 0x8f, 0x04, 0xca, 0x98, 0x5e, 0x3b, 0xab, 0xfe, 0xbf, 0xf3, 0x41, 0x3a, 0xc0,
	0x79, 0x49, 0x81, 0x1e, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x26, 0xcc, 0x5e, 0xaa, 0xd6, 0x09,
	0x00, 0x00,
}
