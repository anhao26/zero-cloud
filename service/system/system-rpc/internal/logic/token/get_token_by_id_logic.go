package token

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenByIdLogic) GetTokenById(in *system.UUIDReq) (*system.TokenInfo, error) {
	result, err := l.svcCtx.DB.Token.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.TokenInfo{
		Id:          result.ID.String(),
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			Status:	uint32(result.Status),
			Uuid:	result.UUID.String(),
			Token:	result.Token,
			Source:	result.Source,
			ExpiredAt:	result.ExpiredAt.UnixMilli(),
	}, nil
}

