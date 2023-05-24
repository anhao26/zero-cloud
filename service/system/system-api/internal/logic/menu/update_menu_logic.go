package menu

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.MenuPlainInfo) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.SystemRpc.UpdateMenu(l.ctx, &system.MenuInfo{
		Id:        req.Id,
		MenuType:  req.MenuType,
		ParentId:  req.ParentId,
		Path:      req.Path,
		Name:      req.Name,
		Redirect:  req.Redirect,
		Component: req.Component,
		Sort:      req.Sort,
		Disabled:  req.Disabled,
		Meta: &system.Meta{
			Title:              req.Title,
			Icon:               req.Icon,
			HideMenu:           req.HideMenu,
			HideBreadcrumb:     req.HideBreadcrumb,
			IgnoreKeepAlive:    req.IgnoreKeepAlive,
			HideTab:            req.HideTab,
			FrameSrc:           req.FrameSrc,
			CarryParam:         req.CarryParam,
			HideChildrenInMenu: req.HideChildrenInMenu,
			Affix:              req.Affix,
			DynamicLevel:       req.DynamicLevel,
			RealPath:           req.RealPath,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
