package entity

import (
	"context"


	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEntityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEntityLogic {
	return &CreateEntityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateEntityLogic) CreateEntity(in *ticket.EntityInfo) (*ticket.BaseIDResp, error) {
    result, err := l.svcCtx.DB.Entity.Create().
			SetEntityCode(in.EntityCode).
			SetEntityClass(in.EntityClass).
			SetEntityTable(in.EntityTable).
			SetDefaultAttributeSetID(in.DefaultAttributeSetId).
			SetAdditionalAttributeTable(in.AdditionalAttributeTable).
			SetIsFlatEnabled(in.IsFlatEnabled).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
