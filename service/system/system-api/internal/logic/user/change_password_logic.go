package user

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.BaseMsgResp, err error) {
	userData, err := l.svcCtx.SystemRpc.GetUserById(l.ctx, &system.UUIDReq{Id: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}

	if encrypt.BcryptCheck(req.OldPassword, userData.Password) {
		result, err := l.svcCtx.SystemRpc.UpdateUser(l.ctx, &system.UserInfo{
			Id:       l.ctx.Value("userId").(string),
			Password: req.NewPassword,
		})
		if err != nil {
			return nil, err
		}

		return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
	}

	return nil, errorx.NewCodeInvalidArgumentError("login.wrongPassword")
}
