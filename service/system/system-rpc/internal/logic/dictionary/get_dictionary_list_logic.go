package dictionary

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/dictionary"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryListLogic {
	return &GetDictionaryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryListLogic) GetDictionaryList(in *system.DictionaryListReq) (*system.DictionaryListResp, error) {
	var predicates []predicate.Dictionary
	if in.Title != "" {
		predicates = append(predicates, dictionary.TitleContains(in.Title))
	}
	if in.Name != "" {
		predicates = append(predicates, dictionary.NameContains(in.Name))
	}
	if in.Desc != "" {
		predicates = append(predicates, dictionary.DescContains(in.Desc))
	}
	result, err := l.svcCtx.DB.Dictionary.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.DictionaryListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.DictionaryInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Status:	uint32(v.Status),
			Title:	v.Title,
			Name:	v.Name,
			Desc:	v.Desc,
		})
	}

	return resp, nil
}
