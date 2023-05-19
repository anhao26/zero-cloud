package base

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *rpc.Empty) (*rpc.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &rpc.BaseResp{}, nil
}
