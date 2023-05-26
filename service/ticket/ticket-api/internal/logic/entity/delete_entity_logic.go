package entity

import (
	"context"

    "github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEntityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEntityLogic {
	return &DeleteEntityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEntityLogic) DeleteEntity(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.TicketRpc.DeleteEntity(l.ctx, &ticket.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
