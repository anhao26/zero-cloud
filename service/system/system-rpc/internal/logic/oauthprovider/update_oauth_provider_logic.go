package oauthprovider

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOauthProviderLogic {
	return &UpdateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOauthProviderLogic) UpdateOauthProvider(in *system.OauthProviderInfo) (*system.BaseResp, error) {
	err := l.svcCtx.DB.OauthProvider.UpdateOneID(in.Id).
		SetNotEmptyName(in.Name).
		SetNotEmptyClientID(in.ClientId).
		SetNotEmptyClientSecret(in.ClientSecret).
		SetNotEmptyRedirectURL(in.RedirectUrl).
		SetNotEmptyScopes(in.Scopes).
		SetNotEmptyAuthURL(in.AuthUrl).
		SetNotEmptyTokenURL(in.TokenUrl).
		SetNotEmptyAuthStyle(in.AuthStyle).
		SetNotEmptyInfoURL(in.InfoUrl).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	if _, ok := providerConfig[in.Name]; ok {
		delete(providerConfig, in.Name)
	}

	return &system.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
