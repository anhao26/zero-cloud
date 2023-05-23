package menu

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/rpc/ent/menu"
	"github.com/anhao26/zero-cloud/service/system/rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/rpc/types/system"

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

func (l *GetMenuListLogic) GetMenuList(in *system.MenuListReq) (*system.MenuListResp, error) {
	var predicates []predicate.Menu
	if in.Path != "" {
		predicates = append(predicates, menu.PathContains(in.Path))
	}
	if in.Name != "" {
		predicates = append(predicates, menu.NameContains(in.Name))
	}
	if in.Redirect != "" {
		predicates = append(predicates, menu.RedirectContains(in.Redirect))
	}
	result, err := l.svcCtx.DB.Menu.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.MenuListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.MenuInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Sort:	v.Sort,
			ParentId:	v.ParentID,
			MenuLevel:	v.MenuLevel,
			MenuType:	v.MenuType,
			Path:	v.Path,
			Name:	v.Name,
			Redirect:	v.Redirect,
			Component:	v.Component,
			Disabled:	v.Disabled,
			Title:	v.Title,
			Icon:	v.Icon,
			HideMenu:	v.HideMenu,
			HideBreadcrumb:	v.HideBreadcrumb,
			IgnoreKeepAlive:	v.IgnoreKeepAlive,
			HideTab:	v.HideTab,
			FrameSrc:	v.FrameSrc,
			CarryParam:	v.CarryParam,
			HideChildrenInMenu:	v.HideChildrenInMenu,
			Affix:	v.Affix,
			DynamicLevel:	v.DynamicLevel,
			RealPath:	v.RealPath,
		})
	}

	return resp, nil
}
