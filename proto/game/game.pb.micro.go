// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: game/game.proto

package game

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GameService service

func NewGameServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GameService service

type GameService interface {
	GetPlayerInfoByID(ctx context.Context, in *GetPlayerInfoByIDRequest, opts ...client.CallOption) (*GetPlayerInfoByIDReply, error)
	GetGuildInfoByID(ctx context.Context, in *GetGuildInfoByIDRequest, opts ...client.CallOption) (*GetGuildInfoByIDReply, error)
}

type gameService struct {
	c    client.Client
	name string
}

func NewGameService(name string, c client.Client) GameService {
	return &gameService{
		c:    c,
		name: name,
	}
}

func (c *gameService) GetPlayerInfoByID(ctx context.Context, in *GetPlayerInfoByIDRequest, opts ...client.CallOption) (*GetPlayerInfoByIDReply, error) {
	req := c.c.NewRequest(c.name, "GameService.GetPlayerInfoByID", in)
	out := new(GetPlayerInfoByIDReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameService) GetGuildInfoByID(ctx context.Context, in *GetGuildInfoByIDRequest, opts ...client.CallOption) (*GetGuildInfoByIDReply, error) {
	req := c.c.NewRequest(c.name, "GameService.GetGuildInfoByID", in)
	out := new(GetGuildInfoByIDReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GameService service

type GameServiceHandler interface {
	GetPlayerInfoByID(context.Context, *GetPlayerInfoByIDRequest, *GetPlayerInfoByIDReply) error
	GetGuildInfoByID(context.Context, *GetGuildInfoByIDRequest, *GetGuildInfoByIDReply) error
}

func RegisterGameServiceHandler(s server.Server, hdlr GameServiceHandler, opts ...server.HandlerOption) error {
	type gameService interface {
		GetPlayerInfoByID(ctx context.Context, in *GetPlayerInfoByIDRequest, out *GetPlayerInfoByIDReply) error
		GetGuildInfoByID(ctx context.Context, in *GetGuildInfoByIDRequest, out *GetGuildInfoByIDReply) error
	}
	type GameService struct {
		gameService
	}
	h := &gameServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GameService{h}, opts...))
}

type gameServiceHandler struct {
	GameServiceHandler
}

func (h *gameServiceHandler) GetPlayerInfoByID(ctx context.Context, in *GetPlayerInfoByIDRequest, out *GetPlayerInfoByIDReply) error {
	return h.GameServiceHandler.GetPlayerInfoByID(ctx, in, out)
}

func (h *gameServiceHandler) GetGuildInfoByID(ctx context.Context, in *GetGuildInfoByIDRequest, out *GetGuildInfoByIDReply) error {
	return h.GameServiceHandler.GetGuildInfoByID(ctx, in, out)
}