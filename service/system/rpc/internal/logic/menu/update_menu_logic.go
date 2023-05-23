package menu

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/rpc/types/system"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *system.MenuInfo) (*system.BaseResp, error) {
	err := l.svcCtx.DB.Menu.UpdateOneID(in.Id).
		SetNotEmptySort(in.Sort).
		SetNotEmptyParentID(in.ParentId).
		SetNotEmptyMenuLevel(in.MenuLevel).
		SetNotEmptyMenuType(in.MenuType).
		SetNotEmptyPath(in.Path).
		SetNotEmptyName(in.Name).
		SetNotEmptyRedirect(in.Redirect).
		SetNotEmptyComponent(in.Component).
		SetDisabled(in.Disabled).
		SetTitle(in.Title).
		SetIcon(in.Icon).
		SetHideMenu(in.HideMenu).
		SetHideBreadcrumb(in.HideBreadcrumb).
		SetIgnoreKeepAlive(in.IgnoreKeepAlive).
		SetHideTab(in.HideTab).
		SetFrameSrc(in.FrameSrc).
		SetCarryParam(in.CarryParam).
		SetHideChildrenInMenu(in.HideChildrenInMenu).
		SetAffix(in.Affix).
		SetDynamicLevel(in.DynamicLevel).
		SetRealPath(in.RealPath).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.BaseResp{Msg: i18n.CreateSuccess}, nil
}
