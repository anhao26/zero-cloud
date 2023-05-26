package attributeset

import (
	"context"

    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeset"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteAttributeSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAttributeSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAttributeSetLogic {
	return &DeleteAttributeSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAttributeSetLogic) DeleteAttributeSet(in *ticket.IDsReq) (*ticket.BaseResp, error) {
	_, err := l.svcCtx.DB.AttributeSet.Delete().Where(attributeset.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
