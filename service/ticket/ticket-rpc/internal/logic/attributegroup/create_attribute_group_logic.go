package attributegroup

import (
	"context"


	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAttributeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAttributeGroupLogic {
	return &CreateAttributeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAttributeGroupLogic) CreateAttributeGroup(in *ticket.AttributeGroupInfo) (*ticket.BaseIDResp, error) {
    result, err := l.svcCtx.DB.AttributeGroup.Create().
			SetAttributeSetID(in.AttributeSetId).
			SetAttributeGroupName(in.AttributeGroupName).
			SetSequence(uint8(in.Sequence)).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
