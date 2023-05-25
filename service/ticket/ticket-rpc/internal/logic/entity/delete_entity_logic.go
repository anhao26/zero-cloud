package entity

import (
	"context"

    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/entity"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteEntityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEntityLogic {
	return &DeleteEntityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEntityLogic) DeleteEntity(in *ticket.IDsReq) (*ticket.BaseResp, error) {
	_, err := l.svcCtx.DB.Entity.Delete().Where(entity.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
