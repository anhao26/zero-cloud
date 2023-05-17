package student

import (
	"context"

	"github.com/anhao26/zero-cloud/system/system-rpc/ent/student"
	"github.com/anhao26/zero-cloud/system/system-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/system/system-rpc/types/system"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentListLogic) GetStudentList(in *system.StudentListReq) (*system.StudentListResp, error) {
	var predicates []predicate.Student
	if in.Name != "" {
		predicates = append(predicates, student.NameContains(in.Name))
	}
	result, err := l.svcCtx.DB.Student.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.StudentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.StudentInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Name:	v.Name,
			Age:	int64(v.Age),
			AgeInt8:	int32(v.AgeInt8),
			AgeUint8:	uint32(v.AgeUint8),
			AgeInt16:	int32(v.AgeInt16),
			AgeUint16:	uint32(v.AgeUint16),
			AgeInt32:	v.AgeInt32,
			AgeUint32:	v.AgeUint32,
			AgeInt64:	v.AgeInt64,
			AgeUint64:	v.AgeUint64,
			AgeInt:	int64(v.AgeInt),
			AgeUint:	uint64(v.AgeUint),
			WeightFloat:	v.WeightFloat,
			WeightFloat32:	v.WeightFloat32,
			ClassId:	v.ClassID.String(),
			EnrollAt:	v.EnrollAt.UnixMilli(),
			StatusBool:	v.StatusBool,
		})
	}

	return resp, nil
}
