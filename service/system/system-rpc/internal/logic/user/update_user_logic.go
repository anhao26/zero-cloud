package user

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/token"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/entx"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *system.UserInfo) (*system.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		updateQuery := tx.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
			SetNotEmptyUsername(in.Username).
			SetNotEmptyNickname(in.Nickname).
			SetNotEmptyEmail(in.Email).
			SetNotEmptyMobile(in.Mobile).
			SetNotEmptyAvatar(in.Avatar).
			SetNotEmptyHomePath(in.HomePath).
			SetNotEmptyDescription(in.Description).
			SetNotEmptyDepartmentID(in.DepartmentId).
			SetNotEmptyStatus(uint8(in.Status))

		if in.Password != "" {
			updateQuery = updateQuery.SetNotEmptyPassword(encrypt.BcryptEncrypt(in.Password))
		}

		if in.RoleIds != nil {
			err := tx.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).ClearRoles().Exec(l.ctx)
			if err != nil {
				return err
			}

			updateQuery = updateQuery.AddRoleIDs(in.RoleIds...)
		}

		if in.PositionIds != nil {
			err := tx.User.UpdateOneID(uuidx.ParseUUIDString(in.Id)).ClearPositions().Exec(l.ctx)
			if err != nil {
				return err
			}

			_, err = token.NewBlockUserAllTokenLogic(l.ctx, l.svcCtx).BlockUserAllToken(&system.UUIDReq{Id: in.Id})
			if err != nil {
				return err
			}

			updateQuery = updateQuery.AddPositionIDs(in.PositionIds...)
		}

		return updateQuery.Exec(l.ctx)
	})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.BaseResp{
		Msg: i18n.Success,
	}, nil
}
