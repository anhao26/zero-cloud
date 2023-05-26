package attribute

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAttributeLogic {
	return &UpdateAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAttributeLogic) UpdateAttribute(in *ticket.AttributeInfo) (*ticket.BaseResp, error) {
	err := l.svcCtx.DB.Attribute.UpdateOneID(in.Id).
		SetNotEmptyEntityID(in.EntityId).
		SetNotEmptyAttributeCode(in.AttributeCode).
		SetNotEmptyBackendClass(in.BackendClass).
		SetNotEmptyBackendType(in.BackendType).
		SetNotEmptyBackendTable(in.BackendTable).
		SetNotEmptyFrontendClass(in.FrontendClass).
		SetNotEmptyFrontendType(in.FrontendType).
		SetNotEmptyFrontendLabel(in.FrontendLabel).
		SetNotEmptySourceClass(in.SourceClass).
		SetNotEmptyDefaultValue(in.DefaultValue).
		SetNotEmptyIsFilterable(uint8(in.IsFilterable)).
		SetNotEmptyIsSearchable(uint8(in.IsSearchable)).
		SetNotEmptyIsRequired(uint8(in.IsRequired)).
		SetNotEmptyRequiredValidateClass(in.RequiredValidateClass).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// 下拉选项

	return &ticket.BaseResp{Msg: i18n.CreateSuccess}, nil
}
