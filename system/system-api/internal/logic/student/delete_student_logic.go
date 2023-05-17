package student

import (
	"context"

    "github.com/anhao26/zero-cloud/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentLogic {
	return &DeleteStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStudentLogic) DeleteStudent(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.SystemRpc.DeleteStudent(l.ctx, &system.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
