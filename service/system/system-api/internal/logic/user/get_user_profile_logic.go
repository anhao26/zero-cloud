package user

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileLogic) GetUserProfile() (resp *types.ProfileResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetUserById(l.ctx, &system.UUIDReq{Id: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}

	return &types.ProfileResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ProfileInfo{
			Nickname: data.Nickname,
			Avatar:   data.Avatar,
			Mobile:   data.Mobile,
			Email:    data.Email,
		},
	}, nil
}
