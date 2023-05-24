package position

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePositionLogic {
	return &CreatePositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePositionLogic) CreatePosition(req *types.PositionInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.SystemRpc.CreatePosition(l.ctx,
		&system.PositionInfo{
			Status: req.Status,
			Sort:   req.Sort,
			Name:   req.Name,
			Code:   req.Code,
			Remark: req.Remark,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
