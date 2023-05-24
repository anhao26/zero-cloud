package token

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenByIdLogic) GetTokenById(req *types.UUIDReq) (resp *types.TokenInfoResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetTokenById(l.ctx, &system.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.TokenInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.TokenInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:    data.Status,
			Uuid:      data.Uuid,
			Token:     data.Token,
			Source:    data.Source,
			ExpiredAt: data.ExpiredAt,
		},
	}, nil
}
