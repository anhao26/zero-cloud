package user

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/user"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *system.UsernameReq) (*system.UserInfo, error) {
	result, err := l.svcCtx.DB.User.Query().Where(user.UsernameEQ(in.Username)).WithRoles().First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &system.UserInfo{
		Nickname:     result.Nickname,
		Avatar:       result.Avatar,
		Password:     result.Password,
		RoleIds:      GetRoleIds(result.Edges.Roles),
		RoleCodes:    GetRoleCodes(result.Edges.Roles),
		Mobile:       result.Mobile,
		Email:        result.Email,
		Status:       uint32(result.Status),
		Id:           result.ID.String(),
		Username:     result.Username,
		HomePath:     result.HomePath,
		Description:  result.Description,
		DepartmentId: result.DepartmentID,
		CreatedAt:    result.CreatedAt.Unix(),
		UpdatedAt:    result.UpdatedAt.Unix(),
	}, nil
}
