// Code generated by goctl. DO NOT EDIT.
// Source: ticket.proto

package ticketclient

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp     = ticket.BaseIDResp
	BaseResp       = ticket.BaseResp
	BaseUUIDResp   = ticket.BaseUUIDResp
	Empty          = ticket.Empty
	EntityInfo     = ticket.EntityInfo
	EntityListReq  = ticket.EntityListReq
	EntityListResp = ticket.EntityListResp
	IDReq          = ticket.IDReq
	IDsReq         = ticket.IDsReq
	PageInfoReq    = ticket.PageInfoReq
	UUIDReq        = ticket.UUIDReq
	UUIDsReq       = ticket.UUIDsReq

	Ticket interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Entity management
		CreateEntity(ctx context.Context, in *EntityInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateEntity(ctx context.Context, in *EntityInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetEntityList(ctx context.Context, in *EntityListReq, opts ...grpc.CallOption) (*EntityListResp, error)
		GetEntityById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*EntityInfo, error)
		DeleteEntity(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultTicket struct {
		cli zrpc.Client
	}
)

func NewTicket(cli zrpc.Client) Ticket {
	return &defaultTicket{
		cli: cli,
	}
}

func (m *defaultTicket) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := ticket.NewTicketClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Entity management
func (m *defaultTicket) CreateEntity(ctx context.Context, in *EntityInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := ticket.NewTicketClient(m.cli.Conn())
	return client.CreateEntity(ctx, in, opts...)
}

func (m *defaultTicket) UpdateEntity(ctx context.Context, in *EntityInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := ticket.NewTicketClient(m.cli.Conn())
	return client.UpdateEntity(ctx, in, opts...)
}

func (m *defaultTicket) GetEntityList(ctx context.Context, in *EntityListReq, opts ...grpc.CallOption) (*EntityListResp, error) {
	client := ticket.NewTicketClient(m.cli.Conn())
	return client.GetEntityList(ctx, in, opts...)
}

func (m *defaultTicket) GetEntityById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*EntityInfo, error) {
	client := ticket.NewTicketClient(m.cli.Conn())
	return client.GetEntityById(ctx, in, opts...)
}

func (m *defaultTicket) DeleteEntity(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := ticket.NewTicketClient(m.cli.Conn())
	return client.DeleteEntity(ctx, in, opts...)
}