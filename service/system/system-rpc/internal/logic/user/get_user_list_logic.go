package user

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/user"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *system.UserListReq) (*system.UserListResp, error) {
	var predicates []predicate.User
	if in.Username != "" {
		predicates = append(predicates, user.UsernameContains(in.Username))
	}
	if in.Password != "" {
		predicates = append(predicates, user.PasswordContains(in.Password))
	}
	if in.Nickname != "" {
		predicates = append(predicates, user.NicknameContains(in.Nickname))
	}
	result, err := l.svcCtx.DB.User.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &system.UserListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &system.UserInfo{
			Id:          v.ID.String(),
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			Status:	uint32(v.Status),
			Username:	v.Username,
			Password:	v.Password,
			Nickname:	v.Nickname,
			Description:	v.Description,
			HomePath:	v.HomePath,
			Mobile:	v.Mobile,
			Email:	v.Email,
			Avatar:	v.Avatar,
			DepartmentId:	v.DepartmentID,
		})
	}

	return resp, nil
}
