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

type UpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenLogic {
	return &UpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTokenLogic) UpdateToken(in *system.TokenInfo) (*system.BaseResp, error) {
	err := l.svcCtx.DB.Token.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
		SetNotEmptyStatus(uint8(in.Status)).
		SetUUID(uuidx.ParseUUIDString(in.Uuid)).
		SetToken(in.Token).
		SetSource(in.Source).
		SetExpiredAt(time.Unix(in.ExpiredAt, 0)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.BaseResp{Msg: i18n.CreateSuccess}, nil
}
