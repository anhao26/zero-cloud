package department

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDepartmentLogic {
	return &CreateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDepartmentLogic) CreateDepartment(in *system.DepartmentInfo) (*system.BaseIDResp, error) {
    result, err := l.svcCtx.DB.Department.Create().
			SetStatus(uint8(in.Status)).
			SetSort(in.Sort).
			SetName(in.Name).
			SetAncestors(in.Ancestors).
			SetLeader(in.Leader).
			SetPhone(in.Phone).
			SetEmail(in.Email).
			SetRemark(in.Remark).
			SetParentID(in.ParentId).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
