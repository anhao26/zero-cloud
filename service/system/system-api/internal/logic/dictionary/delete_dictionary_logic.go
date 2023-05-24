package dictionary

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryLogic {
	return &DeleteDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDictionaryLogic) DeleteDictionary(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.SystemRpc.DeleteDictionary(l.ctx, &system.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
