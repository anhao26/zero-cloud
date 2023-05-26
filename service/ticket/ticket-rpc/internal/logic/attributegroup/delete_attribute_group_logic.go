package attributegroup

import (
	"context"

    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributegroup"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteAttributeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAttributeGroupLogic {
	return &DeleteAttributeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAttributeGroupLogic) DeleteAttributeGroup(in *ticket.IDsReq) (*ticket.BaseResp, error) {
	_, err := l.svcCtx.DB.AttributeGroup.Delete().Where(attributegroup.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
