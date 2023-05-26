package entityattribute

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEntityAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEntityAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEntityAttributeLogic {
	return &UpdateEntityAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEntityAttributeLogic) UpdateEntityAttribute(in *ticket.EntityAttributeInfo) (*ticket.BaseResp, error) {
	err := l.svcCtx.DB.EntityAttribute.UpdateOneID(in.Id).
		SetNotEmptyAttributeID(in.AttributeId).
		SetNotEmptyEntityID(in.EntityId).
		SetAttributeSetID(in.AttributeSetId).
		SetAttributeGroupID(in.AttributeGroupId).
		SetSequence(uint8(in.Sequence)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.BaseResp{Msg: i18n.CreateSuccess}, nil
}
