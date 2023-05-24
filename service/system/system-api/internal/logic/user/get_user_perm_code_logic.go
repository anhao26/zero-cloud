package user

import (
	"context"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"net/http"
	"strings"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermCodeLogic {
	return &GetUserPermCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermCodeLogic) GetUserPermCode() (resp *types.PermCodeResp, err error) {
	roleId := l.ctx.Value("roleId").(string)
	if roleId == "" {
		return nil, &errorx.ApiError{
			Code: http.StatusUnauthorized,
			Msg:  "login.requireLogin",
		}
	}

	return &types.PermCodeResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data:         strings.Split(roleId, ","),
	}, nil
}
