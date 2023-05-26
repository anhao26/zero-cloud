package attributeset

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAttributeSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAttributeSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAttributeSetLogic {
	return &UpdateAttributeSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAttributeSetLogic) UpdateAttributeSet(in *ticket.AttributeSetInfo) (*ticket.BaseResp, error) {
	err := l.svcCtx.DB.AttributeSet.UpdateOneID(in.Id).
		SetNotEmptyEntityID(in.EntityId).
		SetAttributeSetName(in.AttributeSetName).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.BaseResp{Msg: i18n.CreateSuccess}, nil
}
