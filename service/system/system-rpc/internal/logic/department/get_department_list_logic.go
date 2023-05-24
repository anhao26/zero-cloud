package department

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/department"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepartmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentListLogic {
	return &GetDepartmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepartmentListLogic) GetDepartmentList(in *system.DepartmentListReq) (*system.DepartmentListResp, error) {
	var predicates []predicate.Department
	if in.Name != "" {
		predicates = append(predicates, department.NameContains(in.Name))
	}
	if in.Ancestors != "" {
		predicates = append(predicates, department.AncestorsContains(in.Ancestors))
	}
	if in.Leader != "" {
		predicates = append(predicates, department.LeaderContains(in.Leader))
	}
	result, err := l.svcCtx.DB.Department.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.DepartmentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.DepartmentInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Status:	uint32(v.Status),
			Sort:	v.Sort,
			Name:	v.Name,
			Ancestors:	v.Ancestors,
			Leader:	v.Leader,
			Phone:	v.Phone,
			Email:	v.Email,
			Remark:	v.Remark,
			ParentId:	v.ParentID,
		})
	}

	return resp, nil
}
