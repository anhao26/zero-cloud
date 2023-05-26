package attributeset

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeset"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeSetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttributeSetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeSetListLogic {
	return &GetAttributeSetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttributeSetListLogic) GetAttributeSetList(in *ticket.AttributeSetListReq) (*ticket.AttributeSetListResp, error) {
	var predicates []predicate.AttributeSet
	if in.AttributeSetName != "" {
		predicates = append(predicates, attributeset.AttributeSetNameContains(in.AttributeSetName))
	}
	result, err := l.svcCtx.DB.AttributeSet.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &ticket.AttributeSetListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &ticket.AttributeSetInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			EntityId:	v.EntityID,
			AttributeSetName:	v.AttributeSetName,
		})
	}

	return resp, nil
}
