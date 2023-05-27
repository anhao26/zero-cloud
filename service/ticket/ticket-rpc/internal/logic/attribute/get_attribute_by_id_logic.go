package attribute

import (
	"context"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttributeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeByIdLogic {
	return &GetAttributeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttributeByIdLogic) GetAttributeById(in *ticket.IDReq) (*ticket.AttributeInfo, error) {
	result, err := l.svcCtx.DB.Attribute.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	res := &ticket.AttributeInfo{
		Id:                    result.ID,
		CreatedAt:             result.CreatedAt.UnixMilli(),
		UpdatedAt:             result.UpdatedAt.UnixMilli(),
		EntityId:              result.EntityID,
		AttributeCode:         result.AttributeCode,
		BackendClass:          result.BackendClass,
		BackendType:           result.BackendType,
		BackendTable:          result.BackendTable,
		FrontendClass:         result.FrontendClass,
		FrontendType:          result.FrontendType,
		FrontendLabel:         result.FrontendLabel,
		SourceClass:           result.SourceClass,
		DefaultValue:          result.DefaultValue,
		IsFilterable:          uint32(result.IsFilterable),
		IsSearchable:          uint32(result.IsSearchable),
		IsRequired:            uint32(result.IsRequired),
		RequiredValidateClass: result.RequiredValidateClass,
	}

	// 获取相关信息
	if result.FrontendType == "select" {
		option, err := result.QueryOptionID().All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}
		for _, v := range option {
			res.OptionData = append(res.OptionData, &ticket.Options{
				Label: v.Label,
				Value: v.Value,
			})
		}
	}
	return res, nil
}
