package attributegroup

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributegroup"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttributeGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeGroupListLogic {
	return &GetAttributeGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttributeGroupListLogic) GetAttributeGroupList(in *ticket.AttributeGroupListReq) (*ticket.AttributeGroupListResp, error) {
	var predicates []predicate.AttributeGroup
	if in.AttributeGroupName != "" {
		predicates = append(predicates, attributegroup.AttributeGroupNameContains(in.AttributeGroupName))
	}
	result, err := l.svcCtx.DB.AttributeGroup.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &ticket.AttributeGroupListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &ticket.AttributeGroupInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			AttributeSetId:	v.AttributeSetID,
			AttributeGroupName:	v.AttributeGroupName,
			Sequence:	uint32(v.Sequence),
		})
	}

	return resp, nil
}
