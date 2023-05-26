package entity

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEntityByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEntityByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntityByIdLogic {
	return &GetEntityByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEntityByIdLogic) GetEntityById(req *types.IDReq) (resp *types.EntityInfoResp, err error) {
	data, err := l.svcCtx.TicketRpc.GetEntityById(l.ctx, &ticket.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.EntityInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.EntityInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			EntityCode:               data.EntityCode,
			EntityClass:              data.EntityClass,
			EntityTable:              data.EntityTable,
			DefaultAttributeSetId:    data.DefaultAttributeSetId,
			AdditionalAttributeTable: data.AdditionalAttributeTable,
			IsFlatEnabled:            data.IsFlatEnabled,
		},
	}, nil
}
