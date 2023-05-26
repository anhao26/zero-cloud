package entityattribute

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEntityAttributeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEntityAttributeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntityAttributeByIdLogic {
	return &GetEntityAttributeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEntityAttributeByIdLogic) GetEntityAttributeById(in *ticket.IDReq) (*ticket.EntityAttributeInfo, error) {
	result, err := l.svcCtx.DB.EntityAttribute.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.EntityAttributeInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			AttributeId:	result.AttributeID,
			EntityId:	result.EntityID,
			AttributeSetId:	result.AttributeSetID,
			AttributeGroupId:	result.AttributeGroupID,
			Sequence:	uint32(result.Sequence),
	}, nil
}

