package menu

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMenuLogic) CreateMenu(in *system.MenuInfo) (*system.BaseIDResp, error) {
    result, err := l.svcCtx.DB.Menu.Create().
			SetSort(in.Sort).
			SetParentID(in.ParentId).
			SetMenuLevel(in.MenuLevel).
			SetMenuType(in.MenuType).
			SetPath(in.Path).
			SetName(in.Name).
			SetRedirect(in.Redirect).
			SetComponent(in.Component).
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
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
