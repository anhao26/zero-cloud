package oauthprovider

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOauthProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderByIdLogic {
	return &GetOauthProviderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOauthProviderByIdLogic) GetOauthProviderById(req *types.IDReq) (resp *types.OauthProviderInfoResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetOauthProviderById(l.ctx, &system.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.OauthProviderInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.OauthProviderInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:         data.Name,
			ClientId:     data.ClientId,
			ClientSecret: data.ClientSecret,
			RedirectUrl:  data.RedirectUrl,
			Scopes:       data.Scopes,
			AuthUrl:      data.AuthUrl,
			TokenUrl:     data.TokenUrl,
			AuthStyle:    data.AuthStyle,
			InfoUrl:      data.InfoUrl,
		},
	}, nil
}
