package student

import (
	"context"

	"github.com/anhao26/zero-cloud/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentListLogic) GetStudentList(req *types.StudentListReq) (resp *types.StudentListResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetStudentList(l.ctx,
		&system.StudentListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.StudentListResp{}
	resp.Msg = data.String()
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.StudentInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Name:          v.Name,
				Age:           v.Age,
				AgeInt8:       v.AgeInt8,
				AgeUint8:      v.AgeUint8,
				AgeInt16:      v.AgeInt16,
				AgeUint16:     v.AgeUint16,
				AgeInt32:      v.AgeInt32,
				AgeUint32:     v.AgeUint32,
				AgeInt64:      v.AgeInt64,
				AgeUint64:     v.AgeUint64,
				AgeInt:        v.AgeInt,
				AgeUint:       v.AgeUint,
				WeightFloat:   v.WeightFloat,
				WeightFloat32: v.WeightFloat32,
				ClassId:       v.ClassId,
				EnrollAt:      v.EnrollAt,
				StatusBool:    v.StatusBool,
			})
	}
	return resp, nil
}
