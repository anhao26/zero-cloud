package dictionarydetail

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailByIdLogic {
	return &GetDictionaryDetailByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryDetailByIdLogic) GetDictionaryDetailById(req *types.IDReq) (resp *types.DictionaryDetailInfoResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetDictionaryDetailById(l.ctx, &system.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DictionaryDetailInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.DictionaryDetailInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:       data.Status,
			Title:        data.Title,
			Key:          data.Key,
			Value:        data.Value,
			DictionaryId: data.DictionaryId,
			Sort:         data.Sort,
		},
	}, nil
}
