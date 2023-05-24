package user

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *system.UserInfo) (*system.BaseResp, error) {
    err := l.svcCtx.DB.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
			SetNotEmptyStatus(uint8(in.Status)).
			SetNotEmptyUsername(in.Username).
			SetNotEmptyPassword(in.Password).
			SetNotEmptyNickname(in.Nickname).
			SetNotEmptyDescription(in.Description).
			SetNotEmptyHomePath(in.HomePath).
			SetNotEmptyMobile(in.Mobile).
			SetNotEmptyEmail(in.Email).
			SetNotEmptyAvatar(in.Avatar).
			SetNotEmptyDepartmentID(in.DepartmentId).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseResp{Msg: i18n.CreateSuccess }, nil
}
