package dictionary

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryByIdLogic {
	return &GetDictionaryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryByIdLogic) GetDictionaryById(req *types.IDReq) (resp *types.DictionaryInfoResp, err error) {
	data, err := l.svcCtx.SystemRpc.GetDictionaryById(l.ctx, &system.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DictionaryInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.DictionaryInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Title:  data.Title,
			Name:   data.Name,
			Status: data.Status,
			Desc:   data.Desc,
		},
	}, nil
}
