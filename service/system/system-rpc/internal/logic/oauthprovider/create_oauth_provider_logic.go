package oauthprovider

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOauthProviderLogic {
	return &CreateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOauthProviderLogic) CreateOauthProvider(in *system.OauthProviderInfo) (*system.BaseIDResp, error) {
	result, err := l.svcCtx.DB.OauthProvider.Create().
		SetName(in.Name).
		SetClientID(in.ClientId).
		SetClientSecret(in.ClientSecret).
		SetRedirectURL(in.RedirectUrl).
		SetScopes(in.Scopes).
		SetAuthURL(in.AuthUrl).
		SetTokenURL(in.TokenUrl).
		SetAuthStyle(in.AuthStyle).
		SetInfoURL(in.InfoUrl).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
