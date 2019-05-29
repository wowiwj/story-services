// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package go_common_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (UserService_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (UserService_PingPongService, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.common.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (UserService_StreamService, error) {
	req := c.c.NewRequest(c.name, "UserService.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &userServiceStream{stream}, nil
}

type UserService_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type userServiceStream struct {
	stream client.Stream
}

func (x *userServiceStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userService) PingPong(ctx context.Context, opts ...client.CallOption) (UserService_PingPongService, error) {
	req := c.c.NewRequest(c.name, "UserService.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &userServicePingPong{stream}, nil
}

type UserService_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type userServicePingPong struct {
	stream client.Stream
}

func (x *userServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *userServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *userServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, UserService_StreamStream) error
	PingPong(context.Context, UserService_PingPongStream) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.Call(ctx, in, out)
}

func (h *userServiceHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UserServiceHandler.Stream(ctx, m, &userServiceStreamStream{stream})
}

type UserService_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type userServiceStreamStream struct {
	stream server.Stream
}

func (x *userServiceStreamStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *userServiceHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.UserServiceHandler.PingPong(ctx, &userServicePingPongStream{stream})
}

type UserService_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type userServicePingPongStream struct {
	stream server.Stream
}

func (x *userServicePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *userServicePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServicePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServicePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *userServicePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
