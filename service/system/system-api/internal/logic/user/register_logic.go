package user

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseMsgResp, err error) {
	if ok := l.svcCtx.Captcha.Verify("CAPTCHA_"+req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.SystemRpc.CreateUser(l.ctx,
			&system.UserInfo{
				Id:           "",
				Username:     req.Username,
				Password:     req.Password,
				Email:        req.Email,
				Nickname:     req.Username,
				Status:       1,
				HomePath:     "/dashboard",
				RoleIds:      []uint64{1},
				DepartmentId: 1,
				PositionIds:  []uint64{1},
			})
		if err != nil {
			return nil, err
		}
		resp = &types.BaseMsgResp{
			Msg: l.svcCtx.Trans.Trans(l.ctx, user.Msg),
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(errorcode.InvalidArgument,
			l.svcCtx.Trans.Trans(l.ctx, "login.wrongCaptcha"))
	}
}
