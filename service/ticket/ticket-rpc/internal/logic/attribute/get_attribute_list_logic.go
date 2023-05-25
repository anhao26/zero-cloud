package attribute

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attribute"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeListLogic {
	return &GetAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttributeListLogic) GetAttributeList(in *ticket.AttributeListReq) (*ticket.AttributeListResp, error) {
	var predicates []predicate.Attribute
	if in.AttributeCode != "" {
		predicates = append(predicates, attribute.AttributeCodeContains(in.AttributeCode))
	}
	if in.BackendClass != "" {
		predicates = append(predicates, attribute.BackendClassContains(in.BackendClass))
	}
	if in.BackendType != "" {
		predicates = append(predicates, attribute.BackendTypeContains(in.BackendType))
	}
	result, err := l.svcCtx.DB.Attribute.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &ticket.AttributeListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &ticket.AttributeInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			EntityId:	v.EntityID,
			AttributeCode:	v.AttributeCode,
			BackendClass:	v.BackendClass,
			BackendType:	v.BackendType,
			BackendTable:	v.BackendTable,
			FrontendClass:	v.FrontendClass,
			FrontendType:	v.FrontendType,
			FrontendLabel:	v.FrontendLabel,
			SourceClass:	v.SourceClass,
			DefaultValue:	v.DefaultValue,
			IsFilterable:	uint32(v.IsFilterable),
			IsSearchable:	uint32(v.IsSearchable),
			IsRequired:	uint32(v.IsRequired),
			RequiredValidateClass:	v.RequiredValidateClass,
		})
	}

	return resp, nil
}
