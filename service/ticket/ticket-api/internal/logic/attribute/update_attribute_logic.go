package attribute

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAttributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAttributeLogic {
	return &UpdateAttributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAttributeLogic) UpdateAttribute(req *types.AttributeInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.TicketRpc.UpdateAttribute(l.ctx,
		&ticket.AttributeInfo{
			Id:                    req.Id,
			EntityId:              req.EntityId,
			AttributeCode:         req.AttributeCode,
			BackendClass:          req.BackendClass,
			BackendType:           req.BackendType,
			BackendTable:          req.BackendTable,
			FrontendClass:         req.FrontendClass,
			FrontendType:          req.FrontendType,
			FrontendLabel:         req.FrontendLabel,
			SourceClass:           req.SourceClass,
			DefaultValue:          req.DefaultValue,
			IsFilterable:          req.IsFilterable,
			IsSearchable:          req.IsSearchable,
			IsRequired:            req.IsRequired,
			RequiredValidateClass: req.RequiredValidateClass,
			//OptionData: req.OptionData,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
