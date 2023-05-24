package dictionarydetail

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailListLogic {
	return &GetDictionaryDetailListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryDetailListLogic) GetDictionaryDetailList(req *types.DictionaryDetailListReq) (resp *types.DictionaryDetailListResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetDictionaryDetailList(l.ctx,
		&system.DictionaryDetailListReq{
			Page:         req.Page,
			PageSize:     req.PageSize,
			DictionaryId: req.DictionaryId,
			Key:          req.Key,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.DictionaryDetailListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.DictionaryDetailInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:       v.Status,
				Title:        v.Title,
				Key:          v.Key,
				Value:        v.Value,
				DictionaryId: v.DictionaryId,
				Sort:         v.Sort,
			})
	}

	return resp, nil
}
