package entity

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/entity"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetEntityListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEntityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntityListLogic {
	return &GetEntityListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEntityListLogic) GetEntityList(in *ticket.EntityListReq) (*ticket.EntityListResp, error) {
	var predicates []predicate.Entity
	if in.EntityCode != "" {
		predicates = append(predicates, entity.EntityCodeContains(in.EntityCode))
	}
	if in.EntityClass != "" {
		predicates = append(predicates, entity.EntityClassContains(in.EntityClass))
	}
	if in.EntityTable != "" {
		predicates = append(predicates, entity.EntityTableContains(in.EntityTable))
	}
	result, err := l.svcCtx.DB.Entity.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &ticket.EntityListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &ticket.EntityInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			EntityCode:	v.EntityCode,
			EntityClass:	v.EntityClass,
			EntityTable:	v.EntityTable,
			DefaultAttributeSetId:	v.DefaultAttributeSetID,
			AdditionalAttributeTable:	v.AdditionalAttributeTable,
			IsFlatEnabled:	v.IsFlatEnabled,
		})
	}

	return resp, nil
}
