package department

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepartmentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentByIdLogic {
	return &GetDepartmentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepartmentByIdLogic) GetDepartmentById(in *system.IDReq) (*system.DepartmentInfo, error) {
	result, err := l.svcCtx.DB.Department.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.DepartmentInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			Status:	uint32(result.Status),
			Sort:	result.Sort,
			Name:	result.Name,
			Ancestors:	result.Ancestors,
			Leader:	result.Leader,
			Phone:	result.Phone,
			Email:	result.Email,
			Remark:	result.Remark,
			ParentId:	result.ParentID,
	}, nil
}

