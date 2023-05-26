package entity

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEntityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEntityLogic {
	return &CreateEntityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEntityLogic) CreateEntity(req *types.EntityInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.TicketRpc.CreateEntity(l.ctx,
		&ticket.EntityInfo{ 
        	EntityCode: req.EntityCode,
        	EntityClass: req.EntityClass,
        	EntityTable: req.EntityTable,
        	DefaultAttributeSetId: req.DefaultAttributeSetId,
        	AdditionalAttributeTable: req.AdditionalAttributeTable,
        	IsFlatEnabled: req.IsFlatEnabled,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
