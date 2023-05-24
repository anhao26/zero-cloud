package menu

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/menu"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListLogic) GetMenuList(in *system.PageInfoReq) (*system.MenuInfoList, error) {
	menus, err := l.svcCtx.DB.Menu.Query().Page(l.ctx, in.Page, in.PageSize, func(pager *ent.MenuPager) {
		pager.Order = ent.Asc(menu.FieldSort)
	})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.MenuInfoList{}
	for _, v := range menus.List {
		resp.Data = append(resp.Data, &system.MenuInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
			MenuType:  v.MenuType,
			Level:     v.MenuLevel,
			ParentId:  v.ParentID,
			Path:      v.Path,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Component: v.Component,
			Sort:      v.Sort,
			Meta: &system.Meta{
				Title:              v.Title,
				Icon:               v.Icon,
				HideMenu:           v.HideMenu,
				HideBreadcrumb:     v.HideBreadcrumb,
				IgnoreKeepAlive:    v.IgnoreKeepAlive,
				HideTab:            v.HideTab,
				FrameSrc:           v.FrameSrc,
				CarryParam:         v.CarryParam,
				HideChildrenInMenu: v.HideChildrenInMenu,
				Affix:              v.Affix,
				DynamicLevel:       v.DynamicLevel,
				RealPath:           v.RealPath,
			},
		})
	}
	resp.Total = menus.PageDetails.Total
	return resp, nil
}
