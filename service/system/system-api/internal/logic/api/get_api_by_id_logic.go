package api

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiByIdLogic {
	return &GetApiByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiByIdLogic) GetApiById(req *types.IDReq) (resp *types.ApiInfoResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetApiById(l.ctx, &system.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.ApiInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ApiInfo{
			BaseIDInfo:  types.BaseIDInfo{Id: data.Id, CreatedAt: data.CreatedAt, UpdatedAt: data.UpdatedAt},
			Trans:       l.svcCtx.Trans.Trans(l.ctx, data.Description),
			Path:        data.Path,
			Description: data.Description,
			Group:       data.ApiGroup,
			Method:      data.Method,
		},
	}, nil
}
