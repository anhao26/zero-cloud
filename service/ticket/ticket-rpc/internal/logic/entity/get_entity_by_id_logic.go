package entity

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEntityByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEntityByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntityByIdLogic {
	return &GetEntityByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEntityByIdLogic) GetEntityById(in *ticket.IDReq) (*ticket.EntityInfo, error) {
	result, err := l.svcCtx.DB.Entity.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.EntityInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			EntityCode:	result.EntityCode,
			EntityClass:	result.EntityClass,
			EntityTable:	result.EntityTable,
			DefaultAttributeSetId:	result.DefaultAttributeSetID,
			AdditionalAttributeTable:	result.AdditionalAttributeTable,
			IsFlatEnabled:	result.IsFlatEnabled,
	}, nil
}

