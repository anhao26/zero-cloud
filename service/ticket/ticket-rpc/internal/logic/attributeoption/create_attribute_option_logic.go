package attributeoption

import (
	"context"


	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAttributeOptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAttributeOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAttributeOptionLogic {
	return &CreateAttributeOptionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAttributeOptionLogic) CreateAttributeOption(in *ticket.AttributeOptionInfo) (*ticket.BaseIDResp, error) {
    result, err := l.svcCtx.DB.AttributeOption.Create().
			SetAttributeOptionID(in.AttributeOptionId).
			SetLabel(in.Label).
			SetValue(in.Value).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
