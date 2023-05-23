package department

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDepartmentLogic {
	return &UpdateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDepartmentLogic) UpdateDepartment(in *system.DepartmentInfo) (*system.BaseResp, error) {
    err := l.svcCtx.DB.Department.UpdateOneID(in.Id).
			SetNotEmptyStatus(uint8(in.Status)).
			SetNotEmptySort(in.Sort).
			SetNotEmptyName(in.Name).
			SetNotEmptyAncestors(in.Ancestors).
			SetNotEmptyLeader(in.Leader).
			SetNotEmptyPhone(in.Phone).
			SetNotEmptyEmail(in.Email).
			SetNotEmptyRemark(in.Remark).
			SetNotEmptyParentID(in.ParentId).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseResp{Msg: i18n.CreateSuccess }, nil
}
