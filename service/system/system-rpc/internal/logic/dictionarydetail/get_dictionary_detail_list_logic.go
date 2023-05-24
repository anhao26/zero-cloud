package dictionarydetail

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/dictionarydetail"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailListLogic {
	return &GetDictionaryDetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryDetailListLogic) GetDictionaryDetailList(in *system.DictionaryDetailListReq) (*system.DictionaryDetailListResp, error) {
	var predicates []predicate.DictionaryDetail
	if in.Title != "" {
		predicates = append(predicates, dictionarydetail.TitleContains(in.Title))
	}
	if in.Key != "" {
		predicates = append(predicates, dictionarydetail.KeyContains(in.Key))
	}
	if in.Value != "" {
		predicates = append(predicates, dictionarydetail.ValueContains(in.Value))
	}
	result, err := l.svcCtx.DB.DictionaryDetail.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.DictionaryDetailListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.DictionaryDetailInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Status:	uint32(v.Status),
			Sort:	v.Sort,
			Title:	v.Title,
			Key:	v.Key,
			Value:	v.Value,
			DictionaryId:	v.DictionaryID,
		})
	}

	return resp, nil
}
