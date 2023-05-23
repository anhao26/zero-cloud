package user

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *system.UserInfo) (*system.BaseUUIDResp, error) {
    result, err := l.svcCtx.DB.User.Create().
			SetStatus(uint8(in.Status)).
			SetUsername(in.Username).
			SetPassword(in.Password).
			SetNickname(in.Nickname).
			SetDescription(in.Description).
			SetHomePath(in.HomePath).
			SetMobile(in.Mobile).
			SetEmail(in.Email).
			SetAvatar(in.Avatar).
			SetDepartmentID(in.DepartmentId).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess }, nil
}
