package attributeset

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeSetByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttributeSetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeSetByIdLogic {
	return &GetAttributeSetByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttributeSetByIdLogic) GetAttributeSetById(in *ticket.IDReq) (*ticket.AttributeSetInfo, error) {
	result, err := l.svcCtx.DB.AttributeSet.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.AttributeSetInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			EntityId:	result.EntityID,
			AttributeSetName:	result.AttributeSetName,
	}, nil
}

