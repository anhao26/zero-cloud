package attributeoption

import (
	"context"


	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAttributeOptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAttributeOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAttributeOptionLogic {
	return &UpdateAttributeOptionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAttributeOptionLogic) UpdateAttributeOption(in *ticket.AttributeOptionInfo) (*ticket.BaseResp, error) {
    err := l.svcCtx.DB.AttributeOption.UpdateOneID(in.Id).
			SetNotEmptyAttributeID(in.AttributeId).
			SetNotEmptyLabel(in.Label).
			SetNotEmptyValue(in.Value).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseResp{Msg: i18n.CreateSuccess }, nil
}
