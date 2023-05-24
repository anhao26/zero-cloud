package api

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiByIdLogic {
	return &GetApiByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiByIdLogic) GetApiById(in *system.IDReq) (*system.ApiInfo, error) {
	result, err := l.svcCtx.DB.API.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.ApiInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.CreatedAt.UnixMilli(),
		Path:        result.Path,
		Description: result.Description,
		ApiGroup:    result.APIGroup,
		Method:      result.Method,
	}, nil
}
