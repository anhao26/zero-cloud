package role

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *system.RoleInfo) (*system.BaseResp, error) {
    err := l.svcCtx.DB.Role.UpdateOneID(in.Id).
			SetNotEmptyStatus(uint8(in.Status)).
			SetNotEmptyName(in.Name).
			SetNotEmptyCode(in.Code).
			SetNotEmptyDefaultRouter(in.DefaultRouter).
			SetNotEmptyRemark(in.Remark).
			SetNotEmptySort(in.Sort).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseResp{Msg: i18n.CreateSuccess }, nil
}
