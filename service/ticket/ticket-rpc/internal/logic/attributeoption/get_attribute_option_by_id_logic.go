package attributeoption

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeOptionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttributeOptionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeOptionByIdLogic {
	return &GetAttributeOptionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttributeOptionByIdLogic) GetAttributeOptionById(in *ticket.IDReq) (*ticket.AttributeOptionInfo, error) {
	result, err := l.svcCtx.DB.AttributeOption.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &ticket.AttributeOptionInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			AttributeId:	result.AttributeID,
			Label:	result.Label,
			Value:	result.Value,
	}, nil
}

