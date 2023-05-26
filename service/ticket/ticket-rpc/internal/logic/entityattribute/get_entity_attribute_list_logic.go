package entityattribute

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEntityAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEntityAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntityAttributeListLogic {
	return &GetEntityAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEntityAttributeListLogic) GetEntityAttributeList(in *ticket.EntityAttributeListReq) (*ticket.EntityAttributeListResp, error) {
	var predicates []predicate.EntityAttribute
	result, err := l.svcCtx.DB.EntityAttribute.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &ticket.EntityAttributeListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &ticket.EntityAttributeInfo{
			Id:               v.ID,
			CreatedAt:        v.CreatedAt.UnixMilli(),
			UpdatedAt:        v.UpdatedAt.UnixMilli(),
			AttributeId:      v.AttributeID,
			EntityId:         v.EntityID,
			AttributeSetId:   v.AttributeSetID,
			AttributeGroupId: v.AttributeGroupID,
			Sequence:         uint32(v.Sequence),
		})
	}

	return resp, nil
}
