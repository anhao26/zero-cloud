package attribute

import (
	"context"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeListLogic {
	return &GetAttributeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAttributeListLogic) GetAttributeList(req *types.AttributeListReq) (resp *types.AttributeListResp, err error) {
	data, err := l.svcCtx.TicketRpc.GetAttributeList(l.ctx,
		&ticket.AttributeListReq{
			Page:          req.Page,
			PageSize:      req.PageSize,
			AttributeCode: req.AttributeCode,
			BackendClass:  req.BackendClass,
			BackendType:   req.BackendType,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.AttributeListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.AttributeInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				EntityId:              v.EntityId,
				AttributeCode:         v.AttributeCode,
				BackendClass:          v.BackendClass,
				BackendType:           v.BackendType,
				BackendTable:          v.BackendTable,
				FrontendClass:         v.FrontendClass,
				FrontendType:          v.FrontendType,
				FrontendLabel:         v.FrontendLabel,
				SourceClass:           v.SourceClass,
				DefaultValue:          v.DefaultValue,
				IsFilterable:          v.IsFilterable,
				IsSearchable:          v.IsSearchable,
				IsRequired:            v.IsRequired,
				RequiredValidateClass: v.RequiredValidateClass,
			})
	}
	return resp, nil
}
