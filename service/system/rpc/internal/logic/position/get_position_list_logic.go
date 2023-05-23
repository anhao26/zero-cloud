package position

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/rpc/ent/position"
	"github.com/anhao26/zero-cloud/service/system/rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/rpc/types/system"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetPositionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionListLogic {
	return &GetPositionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionListLogic) GetPositionList(in *system.PositionListReq) (*system.PositionListResp, error) {
	var predicates []predicate.Position
	if in.Name != "" {
		predicates = append(predicates, position.NameContains(in.Name))
	}
	if in.Code != "" {
		predicates = append(predicates, position.CodeContains(in.Code))
	}
	if in.Remark != "" {
		predicates = append(predicates, position.RemarkContains(in.Remark))
	}
	result, err := l.svcCtx.DB.Position.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.PositionListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.PositionInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Status:	uint32(v.Status),
			Sort:	v.Sort,
			Name:	v.Name,
			Code:	v.Code,
			Remark:	v.Remark,
		})
	}

	return resp, nil
}
