package entityattribute

import (
	"context"


	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEntityAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEntityAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEntityAttributeLogic {
	return &CreateEntityAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateEntityAttributeLogic) CreateEntityAttribute(in *ticket.EntityAttributeInfo) (*ticket.BaseIDResp, error) {
    result, err := l.svcCtx.DB.EntityAttribute.Create().
			SetAttributeID(in.AttributeId).
			SetEntityID(in.EntityId).
			SetAttributeSetID(in.AttributeSetId).
			SetAttributeGroupID(in.AttributeGroupId).
			SetSequence(uint8(in.Sequence)).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
