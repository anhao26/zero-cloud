package user

import (
	"context"
	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/errorx"
	"strings"
	"time"

	"github.com/anhao26/zero-cloud/service/system/api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if ok := l.svcCtx.Captcha.Verify("CAPTCHA_"+req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.CoreRpc.GetUserByUsername(l.ctx,
			&core.UsernameReq{
				Username: req.Username,
			})
		if err != nil {
			return nil, err
		}

		if !encrypt.BcryptCheck(req.Password, user.Password) {
			return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
		}

		token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, jwt.WithOption("userId", user.Id), jwt.WithOption("roleId",
				strings.Join(user.RoleCodes, ",")))
		if err != nil {
			return nil, err
		}

		// add token into database
		expiredAt := time.Now().Add(time.Second * 259200).Unix()
		_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
			Id:        "",
			CreatedAt: 0,
			Uuid:      user.Id,
			Token:     token,
			Source:    "core_user",
			Status:    1,
			ExpiredAt: expiredAt,
		})

		if err != nil {
			return nil, err
		}

		resp = &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
			Data: types.LoginInfo{
				UserId: user.Id,
				Token:  token,
				Expire: uint64(expiredAt),
			},
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(errorcode.InvalidArgument, "login.wrongCaptcha")
	}
}
