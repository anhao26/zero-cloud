package oauthprovider

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderByIdLogic {
	return &GetOauthProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderByIdLogic) GetOauthProviderById(in *system.IDReq) (*system.OauthProviderInfo, error) {
	result, err := l.svcCtx.DB.OauthProvider.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.OauthProviderInfo{
		Id:           result.ID,
		CreatedAt:    result.CreatedAt.UnixMilli(),
		UpdatedAt:    result.UpdatedAt.UnixMilli(),
		Name:         result.Name,
		ClientId:     result.ClientID,
		ClientSecret: result.ClientSecret,
		RedirectUrl:  result.RedirectURL,
		Scopes:       result.Scopes,
		AuthUrl:      result.AuthURL,
		TokenUrl:     result.TokenURL,
		AuthStyle:    result.AuthStyle,
		InfoUrl:      result.InfoURL,
	}, nil
}
