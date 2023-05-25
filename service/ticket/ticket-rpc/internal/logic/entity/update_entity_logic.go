package entity

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEntityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEntityLogic {
	return &UpdateEntityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEntityLogic) UpdateEntity(in *ticket.EntityInfo) (*ticket.BaseResp, error) {
	err := l.svcCtx.DB.Entity.UpdateOneID(in.Id).
		SetNotEmptyEntityCode(in.EntityCode).
		SetNotEmptyEntityClass(in.EntityClass).
		SetNotEmptyEntityTable(in.EntityTable).
		SetDefaultAttributeSetID(in.DefaultAttributeSetId).
		SetAdditionalAttributeTable(in.AdditionalAttributeTable).
		SetIsFlatEnabled(in.IsFlatEnabled).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.BaseResp{Msg: i18n.CreateSuccess}, nil
}
