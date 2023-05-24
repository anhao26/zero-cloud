package department

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepartmentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentByIdLogic {
	return &GetDepartmentByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepartmentByIdLogic) GetDepartmentById(req *types.IDReq) (resp *types.DepartmentInfoResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetDepartmentById(l.ctx, &system.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DepartmentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.DepartmentInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:    data.Status,
			Sort:      data.Sort,
			Name:      data.Name,
			Ancestors: data.Ancestors,
			Leader:    data.Leader,
			Phone:     data.Phone,
			Email:     data.Email,
			Remark:    data.Remark,
			ParentId:  data.ParentId,
		},
	}, nil
}
