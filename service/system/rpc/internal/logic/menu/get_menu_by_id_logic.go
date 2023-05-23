package menu

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuByIdLogic {
	return &GetMenuByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuByIdLogic) GetMenuById(in *system.IDReq) (*system.MenuInfo, error) {
	result, err := l.svcCtx.DB.Menu.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.MenuInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			Sort:	result.Sort,
			ParentId:	result.ParentID,
			MenuLevel:	result.MenuLevel,
			MenuType:	result.MenuType,
			Path:	result.Path,
			Name:	result.Name,
			Redirect:	result.Redirect,
			Component:	result.Component,
			Disabled:	result.Disabled,
			Title:	result.Title,
			Icon:	result.Icon,
			HideMenu:	result.HideMenu,
			HideBreadcrumb:	result.HideBreadcrumb,
			IgnoreKeepAlive:	result.IgnoreKeepAlive,
			HideTab:	result.HideTab,
			FrameSrc:	result.FrameSrc,
			CarryParam:	result.CarryParam,
			HideChildrenInMenu:	result.HideChildrenInMenu,
			Affix:	result.Affix,
			DynamicLevel:	result.DynamicLevel,
			RealPath:	result.RealPath,
	}, nil
}

