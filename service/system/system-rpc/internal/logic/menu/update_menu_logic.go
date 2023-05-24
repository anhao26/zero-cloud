package menu

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/menu"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/enum/common"
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
	// get parent level
	var menuLevel uint32
	if in.ParentId != common.DefaultParentId {
		m, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(in.ParentId)).First(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		menuLevel = m.MenuLevel + 1
	} else {
		menuLevel = 1
	}

	err := l.svcCtx.DB.Menu.UpdateOneID(in.Id).
		SetMenuLevel(menuLevel).
		SetMenuType(in.MenuType).
		SetParentID(in.ParentId).
		SetPath(in.Path).
		SetName(in.Name).
		SetRedirect(in.Redirect).
		SetComponent(in.Component).
		SetSort(in.Sort).
		SetDisabled(in.Disabled).
		// meta
		SetTitle(in.Meta.Title).
		SetIcon(in.Meta.Icon).
		SetHideMenu(in.Meta.HideMenu).
		SetHideBreadcrumb(in.Meta.HideBreadcrumb).
		SetIgnoreKeepAlive(in.Meta.IgnoreKeepAlive).
		SetHideTab(in.Meta.HideTab).
		SetFrameSrc(in.Meta.FrameSrc).
		SetCarryParam(in.Meta.CarryParam).
		SetHideChildrenInMenu(in.Meta.HideChildrenInMenu).
		SetAffix(in.Meta.Affix).
		SetDynamicLevel(in.Meta.DynamicLevel).
		SetRealPath(in.Meta.RealPath).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
