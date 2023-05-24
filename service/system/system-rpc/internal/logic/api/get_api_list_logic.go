package api

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/api"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiListLogic) GetApiList(in *system.ApiListReq) (*system.ApiListResp, error) {
	var predicates []predicate.API
	if in.Path != "" {
		predicates = append(predicates, api.PathContains(in.Path))
	}
	if in.Description != "" {
		predicates = append(predicates, api.DescriptionContains(in.Description))
	}
	if in.ApiGroup != "" {
		predicates = append(predicates, api.APIGroupContains(in.ApiGroup))
	}
	if in.Method != "" {
		predicates = append(predicates, api.Method(in.Method))
	}
	result, err := l.svcCtx.DB.API.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.ApiListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.ApiInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			Path:        v.Path,
			Description: v.Description,
			ApiGroup:    v.APIGroup,
			Method:      v.Method,
		})
	}

	return resp, nil
}
