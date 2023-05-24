package user

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.SystemRpc.CreateUser(l.ctx,
		&system.UserInfo{
			Status:       req.Status,
			Username:     req.Username,
			Password:     req.Password,
			Nickname:     req.Nickname,
			Description:  req.Description,
			HomePath:     req.HomePath,
			RoleIds:      req.RoleIds,
			Mobile:       req.Mobile,
			Email:        req.Email,
			Avatar:       req.Avatar,
			DepartmentId: req.DepartmentId,
			PositionIds:  req.PositionIds,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
