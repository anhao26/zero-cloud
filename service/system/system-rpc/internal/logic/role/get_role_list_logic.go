package role

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/role"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleListLogic) GetRoleList(in *system.RoleListReq) (*system.RoleListResp, error) {
	var predicates []predicate.Role
	if in.Name != "" {
		predicates = append(predicates, role.NameContains(in.Name))
	}
	if in.Code != "" {
		predicates = append(predicates, role.CodeContains(in.Code))
	}
	if in.DefaultRouter != "" {
		predicates = append(predicates, role.DefaultRouterContains(in.DefaultRouter))
	}
	result, err := l.svcCtx.DB.Role.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.RoleListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.RoleInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Status:	uint32(v.Status),
			Name:	v.Name,
			Code:	v.Code,
			DefaultRouter:	v.DefaultRouter,
			Remark:	v.Remark,
			Sort:	v.Sort,
		})
	}

	return resp, nil
}
