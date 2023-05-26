package attributeoption

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeoption"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeOptionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttributeOptionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeOptionListLogic {
	return &GetAttributeOptionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttributeOptionListLogic) GetAttributeOptionList(in *ticket.AttributeOptionListReq) (*ticket.AttributeOptionListResp, error) {
	var predicates []predicate.AttributeOption
	if in.Label != "" {
		predicates = append(predicates, attributeoption.LabelContains(in.Label))
	}
	result, err := l.svcCtx.DB.AttributeOption.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &ticket.AttributeOptionListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &ticket.AttributeOptionInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			AttributeId:	v.AttributeID,
			Label:	v.Label,
			Value:	v.Value,
		})
	}

	return resp, nil
}
