package token

import (
	"context"
     "time"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTokenLogic) CreateToken(in *system.TokenInfo) (*system.BaseUUIDResp, error) {
    result, err := l.svcCtx.DB.Token.Create().
			SetStatus(uint8(in.Status)).
			SetUUID(uuidx.ParseUUIDString(in.Uuid)).
			SetToken(in.Token).
			SetSource(in.Source).
			SetExpiredAt(time.Unix(in.ExpiredAt, 0)).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess }, nil
}
