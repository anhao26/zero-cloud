package user

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *system.UUIDReq) (*system.UserInfo, error) {
	result, err := l.svcCtx.DB.User.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.UserInfo{
		Id:           result.ID.String(),
		CreatedAt:    result.CreatedAt.UnixMilli(),
		UpdatedAt:    result.UpdatedAt.UnixMilli(),
		Status:       uint32(result.Status),
		Username:     result.Username,
		Password:     result.Password,
		Nickname:     result.Nickname,
		Description:  result.Description,
		HomePath:     result.HomePath,
		Mobile:       result.Mobile,
		Email:        result.Email,
		Avatar:       result.Avatar,
		DepartmentId: result.DepartmentID,
	}, nil
}

func GetRoleIds(data []*ent.Role) []uint64 {
	var ids []uint64
	for _, v := range data {
		ids = append(ids, v.ID)
	}
	return ids
}

func GetRoleCodes(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Code)
	}
	return codes
}

func GetPositionIds(data []*ent.Position) []uint64 {
	var ids []uint64
	for _, v := range data {
		ids = append(ids, v.ID)
	}
	return ids
}
