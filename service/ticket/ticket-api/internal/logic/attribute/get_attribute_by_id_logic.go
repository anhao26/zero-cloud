package attribute

import (
	"context"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/types/ticket"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAttributeByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAttributeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttributeByIdLogic {
	return &GetAttributeByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAttributeByIdLogic) GetAttributeById(req *types.IDReq) (resp *types.AttributeInfoResp, err error) {
	data, err := l.svcCtx.TicketRpc.GetAttributeById(l.ctx, &ticket.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	resp = &types.AttributeInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.AttributeInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			EntityId:              data.EntityId,
			AttributeCode:         data.AttributeCode,
			BackendClass:          data.BackendClass,
			BackendType:           data.BackendType,
			BackendTable:          data.BackendTable,
			FrontendClass:         data.FrontendClass,
			FrontendType:          data.FrontendType,
			FrontendLabel:         data.FrontendLabel,
			SourceClass:           data.SourceClass,
			DefaultValue:          data.DefaultValue,
			IsFilterable:          data.IsFilterable,
			IsSearchable:          data.IsSearchable,
			IsRequired:            data.IsRequired,
			RequiredValidateClass: data.RequiredValidateClass,
		},
	}

	// 下拉特殊处理
	if "select" == data.FrontendType {
		for _, v := range data.OptionData {
			resp.Data.OptionData = append(resp.Data.OptionData,
				types.Options{
					Label: v.Label,
					Value: v.Value,
				})
		}
	}
	return resp, nil
}
