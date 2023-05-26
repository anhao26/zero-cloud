package entity

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEntityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEntityLogic {
	return &UpdateEntityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEntityLogic) UpdateEntity(req *types.EntityInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.TicketRpc.UpdateEntity(l.ctx,
		&ticket.EntityInfo{
			Id:          req.Id,
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
