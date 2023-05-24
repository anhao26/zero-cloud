package position

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePositionLogic {
	return &CreatePositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePositionLogic) CreatePosition(in *system.PositionInfo) (*system.BaseIDResp, error) {
    result, err := l.svcCtx.DB.Position.Create().
			SetStatus(uint8(in.Status)).
			SetSort(in.Sort).
			SetName(in.Name).
			SetCode(in.Code).
			SetRemark(in.Remark).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
