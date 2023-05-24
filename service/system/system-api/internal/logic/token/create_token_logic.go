package token

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTokenLogic) CreateToken(req *types.TokenInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.SystemRpc.CreateToken(l.ctx,
		&system.TokenInfo{
			Status:    req.Status,
			Uuid:      req.Uuid,
			Token:     req.Token,
			Source:    req.Source,
			ExpiredAt: req.ExpiredAt,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
