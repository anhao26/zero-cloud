package menu

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/menu"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/role"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"strings"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole(in *system.BaseMsg) (*system.MenuInfoList, error) {
	roles, err := l.svcCtx.DB.Role.Query().Where(role.CodeIn(strings.Split(in.Msg, ",")...)).WithMenus(func(query *ent.MenuQuery) {
		query.Order(ent.Asc(menu.FieldSort))
	}).All(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.MenuInfoList{}

	existMap := map[uint64]struct{}{}
	for _, r := range roles {
		for _, m := range r.Edges.Menus {
			if _, ok := existMap[m.ID]; !ok {
				resp.Data = append(resp.Data, &system.MenuInfo{
					Id:        m.ID,
					CreatedAt: m.CreatedAt.UnixMilli(),
					UpdatedAt: m.UpdatedAt.UnixMilli(),
					MenuType:  m.MenuType,
					Level:     m.MenuLevel,
					ParentId:  m.ParentID,
					Path:      m.Path,
					Name:      m.Name,
					Redirect:  m.Redirect,
					Component: m.Component,
					Sort:      m.Sort,
					Meta: &system.Meta{
						Title:              m.Title,
						Icon:               m.Icon,
						HideMenu:           m.HideMenu,
						HideBreadcrumb:     m.HideBreadcrumb,
						IgnoreKeepAlive:    m.IgnoreKeepAlive,
						HideTab:            m.HideTab,
						FrameSrc:           m.FrameSrc,
						CarryParam:         m.CarryParam,
						HideChildrenInMenu: m.HideChildrenInMenu,
						Affix:              m.Affix,
						DynamicLevel:       m.DynamicLevel,
						RealPath:           m.RealPath,
					},
				})
				existMap[m.ID] = struct{}{}
			}
		}
	}

	resp.Total = uint64(len(resp.Data))

	return resp, nil
}
