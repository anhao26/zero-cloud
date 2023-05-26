package attributegroup

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAttributeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAttributeGroupLogic {
	return &UpdateAttributeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAttributeGroupLogic) UpdateAttributeGroup(in *ticket.AttributeGroupInfo) (*ticket.BaseResp, error) {
	err := l.svcCtx.DB.AttributeGroup.UpdateOneID(in.Id).
		SetAttributeSetID(in.AttributeSetId).
		SetAttributeGroupName(in.AttributeGroupName).
		SetSequence(uint8(in.Sequence)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.BaseResp{Msg: i18n.CreateSuccess}, nil
}
