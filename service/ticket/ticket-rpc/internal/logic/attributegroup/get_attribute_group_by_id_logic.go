package attributegroup

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeGroupByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttributeGroupByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeGroupByIdLogic {
	return &GetAttributeGroupByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttributeGroupByIdLogic) GetAttributeGroupById(in *ticket.IDReq) (*ticket.AttributeGroupInfo, error) {
	result, err := l.svcCtx.DB.AttributeGroup.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.AttributeGroupInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			AttributeSetId:	result.AttributeSetID,
			AttributeGroupName:	result.AttributeGroupName,
			Sequence:	uint32(result.Sequence),
	}, nil
}

