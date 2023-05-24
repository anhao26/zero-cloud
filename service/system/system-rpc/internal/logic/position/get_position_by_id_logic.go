package position

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionByIdLogic {
	return &GetPositionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionByIdLogic) GetPositionById(in *system.IDReq) (*system.PositionInfo, error) {
	result, err := l.svcCtx.DB.Position.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.PositionInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			Status:	uint32(result.Status),
			Sort:	result.Sort,
			Name:	result.Name,
			Code:	result.Code,
			Remark:	result.Remark,
	}, nil
}

