// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package server

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/base"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/department"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/menu"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/position"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/role"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/user"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
)

type SystemServer struct {
	svcCtx *svc.ServiceContext
	system.UnimplementedSystemServer
}

func NewSystemServer(svcCtx *svc.ServiceContext) *SystemServer {
	return &SystemServer{
		svcCtx: svcCtx,
	}
}

func (s *SystemServer) InitDatabase(ctx context.Context, in *system.Empty) (*system.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// Department management
func (s *SystemServer) CreateDepartment(ctx context.Context, in *system.DepartmentInfo) (*system.BaseIDResp, error) {
	l := department.NewCreateDepartmentLogic(ctx, s.svcCtx)
	return l.CreateDepartment(in)
}

func (s *SystemServer) UpdateDepartment(ctx context.Context, in *system.DepartmentInfo) (*system.BaseResp, error) {
	l := department.NewUpdateDepartmentLogic(ctx, s.svcCtx)
	return l.UpdateDepartment(in)
}

func (s *SystemServer) GetDepartmentList(ctx context.Context, in *system.DepartmentListReq) (*system.DepartmentListResp, error) {
	l := department.NewGetDepartmentListLogic(ctx, s.svcCtx)
	return l.GetDepartmentList(in)
}

func (s *SystemServer) GetDepartmentById(ctx context.Context, in *system.IDReq) (*system.DepartmentInfo, error) {
	l := department.NewGetDepartmentByIdLogic(ctx, s.svcCtx)
	return l.GetDepartmentById(in)
}

func (s *SystemServer) DeleteDepartment(ctx context.Context, in *system.IDsReq) (*system.BaseResp, error) {
	l := department.NewDeleteDepartmentLogic(ctx, s.svcCtx)
	return l.DeleteDepartment(in)
}

// Menu management
func (s *SystemServer) CreateMenu(ctx context.Context, in *system.MenuInfo) (*system.BaseIDResp, error) {
	l := menu.NewCreateMenuLogic(ctx, s.svcCtx)
	return l.CreateMenu(in)
}

func (s *SystemServer) UpdateMenu(ctx context.Context, in *system.MenuInfo) (*system.BaseResp, error) {
	l := menu.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

func (s *SystemServer) GetMenuList(ctx context.Context, in *system.MenuListReq) (*system.MenuListResp, error) {
	l := menu.NewGetMenuListLogic(ctx, s.svcCtx)
	return l.GetMenuList(in)
}

func (s *SystemServer) GetMenuById(ctx context.Context, in *system.IDReq) (*system.MenuInfo, error) {
	l := menu.NewGetMenuByIdLogic(ctx, s.svcCtx)
	return l.GetMenuById(in)
}

func (s *SystemServer) DeleteMenu(ctx context.Context, in *system.IDsReq) (*system.BaseResp, error) {
	l := menu.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

// Position management
func (s *SystemServer) CreatePosition(ctx context.Context, in *system.PositionInfo) (*system.BaseIDResp, error) {
	l := position.NewCreatePositionLogic(ctx, s.svcCtx)
	return l.CreatePosition(in)
}

func (s *SystemServer) UpdatePosition(ctx context.Context, in *system.PositionInfo) (*system.BaseResp, error) {
	l := position.NewUpdatePositionLogic(ctx, s.svcCtx)
	return l.UpdatePosition(in)
}

func (s *SystemServer) GetPositionList(ctx context.Context, in *system.PositionListReq) (*system.PositionListResp, error) {
	l := position.NewGetPositionListLogic(ctx, s.svcCtx)
	return l.GetPositionList(in)
}

func (s *SystemServer) GetPositionById(ctx context.Context, in *system.IDReq) (*system.PositionInfo, error) {
	l := position.NewGetPositionByIdLogic(ctx, s.svcCtx)
	return l.GetPositionById(in)
}

func (s *SystemServer) DeletePosition(ctx context.Context, in *system.IDsReq) (*system.BaseResp, error) {
	l := position.NewDeletePositionLogic(ctx, s.svcCtx)
	return l.DeletePosition(in)
}

// Role management
func (s *SystemServer) CreateRole(ctx context.Context, in *system.RoleInfo) (*system.BaseIDResp, error) {
	l := role.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

func (s *SystemServer) UpdateRole(ctx context.Context, in *system.RoleInfo) (*system.BaseResp, error) {
	l := role.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

func (s *SystemServer) GetRoleList(ctx context.Context, in *system.RoleListReq) (*system.RoleListResp, error) {
	l := role.NewGetRoleListLogic(ctx, s.svcCtx)
	return l.GetRoleList(in)
}

func (s *SystemServer) GetRoleById(ctx context.Context, in *system.IDReq) (*system.RoleInfo, error) {
	l := role.NewGetRoleByIdLogic(ctx, s.svcCtx)
	return l.GetRoleById(in)
}

func (s *SystemServer) DeleteRole(ctx context.Context, in *system.IDsReq) (*system.BaseResp, error) {
	l := role.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

// User management
func (s *SystemServer) CreateUser(ctx context.Context, in *system.UserInfo) (*system.BaseUUIDResp, error) {
	l := user.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *SystemServer) UpdateUser(ctx context.Context, in *system.UserInfo) (*system.BaseResp, error) {
	l := user.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *SystemServer) GetUserList(ctx context.Context, in *system.UserListReq) (*system.UserListResp, error) {
	l := user.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(in)
}

func (s *SystemServer) GetUserById(ctx context.Context, in *system.UUIDReq) (*system.UserInfo, error) {
	l := user.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *SystemServer) GetUserByUsername(ctx context.Context, in *system.UsernameReq) (*system.UserInfo, error) {
	l := user.NewGetUserByUsernameLogic(ctx, s.svcCtx)
	return l.GetUserByUsername(in)
}

func (s *SystemServer) DeleteUser(ctx context.Context, in *system.UUIDsReq) (*system.BaseResp, error) {
	l := user.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}
