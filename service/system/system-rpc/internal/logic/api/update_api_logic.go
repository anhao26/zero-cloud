package api

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateApiLogic) UpdateApi(in *system.ApiInfo) (*system.BaseResp, error) {
	err := l.svcCtx.DB.API.UpdateOneID(in.Id).
		SetNotEmptyPath(in.Path).
		SetNotEmptyDescription(in.Description).
		SetNotEmptyAPIGroup(in.ApiGroup).
		SetNotEmptyMethod(in.Method).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
