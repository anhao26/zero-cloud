package entity

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEntityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEntityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntityListLogic {
	return &GetEntityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEntityListLogic) GetEntityList(req *types.EntityListReq) (resp *types.EntityListResp, err error) {
	data, err := l.svcCtx.TicketRpc.GetEntityList(l.ctx,
		&ticket.EntityListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			EntityCode:  req.EntityCode,
			EntityClass: req.EntityClass,
			EntityTable: req.EntityTable,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.EntityListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.EntityInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				EntityCode:               v.EntityCode,
				EntityClass:              v.EntityClass,
				EntityTable:              v.EntityTable,
				DefaultAttributeSetId:    v.DefaultAttributeSetId,
				AdditionalAttributeTable: v.AdditionalAttributeTable,
				IsFlatEnabled:            v.IsFlatEnabled,
			})
	}
	return resp, nil
}
