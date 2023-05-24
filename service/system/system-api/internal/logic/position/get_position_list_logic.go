package position

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionListLogic {
	return &GetPositionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPositionListLogic) GetPositionList(req *types.PositionListReq) (resp *types.PositionListResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetPositionList(l.ctx,
		&system.PositionListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
			Code:     req.Code,
			Remark:   req.Remark,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.PositionListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.PositionInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Trans:  l.svcCtx.Trans.Trans(l.ctx, v.Name),
				Status: v.Status,
				Sort:   v.Sort,
				Name:   v.Name,
				Code:   v.Code,
				Remark: v.Remark,
			})
	}
	return resp, nil
}
