package user

import (
	"context"

    "github.com/anhao26/zero-cloud/service/system/system-rpc/ent/user"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *system.UUIDsReq) (*system.BaseResp, error) {
	_, err := l.svcCtx.DB.User.Delete().Where(user.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
