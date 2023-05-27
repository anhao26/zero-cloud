package attribute

import (
	"context"


	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAttributeLogic {
	return &CreateAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAttributeLogic) CreateAttribute(in *ticket.AttributeInfo) (*ticket.BaseIDResp, error) {
    result, err := l.svcCtx.DB.Attribute.Create().
			SetEntityID(in.EntityId).
			SetAttributeCode(in.AttributeCode).
			SetBackendClass(in.BackendClass).
			SetBackendType(in.BackendType).
			SetBackendTable(in.BackendTable).
			SetFrontendClass(in.FrontendClass).
			SetFrontendType(in.FrontendType).
			SetFrontendLabel(in.FrontendLabel).
			SetSourceClass(in.SourceClass).
			SetDefaultValue(in.DefaultValue).
			SetIsFilterable(uint8(in.IsFilterable)).
			SetIsSearchable(uint8(in.IsSearchable)).
			SetIsRequired(uint8(in.IsRequired)).
			SetRequiredValidateClass(in.RequiredValidateClass).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &ticket.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
