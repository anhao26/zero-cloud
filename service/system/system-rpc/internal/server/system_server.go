// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package server

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/api"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/base"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/department"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/menu"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/oauthprovider"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/position"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/role"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/logic/token"
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

// API management
func (s *SystemServer) CreateApi(ctx context.Context, in *system.ApiInfo) (*system.BaseIDResp, error) {
	l := api.NewCreateApiLogic(ctx, s.svcCtx)
	return l.CreateApi(in)
}

func (s *SystemServer) UpdateApi(ctx context.Context, in *system.ApiInfo) (*system.BaseResp, error) {
	l := api.NewUpdateApiLogic(ctx, s.svcCtx)
	return l.UpdateApi(in)
}

func (s *SystemServer) GetApiList(ctx context.Context, in *system.ApiListReq) (*system.ApiListResp, error) {
	l := api.NewGetApiListLogic(ctx, s.svcCtx)
	return l.GetApiList(in)
}

func (s *SystemServer) GetApiById(ctx context.Context, in *system.IDReq) (*system.ApiInfo, error) {
	l := api.NewGetApiByIdLogic(ctx, s.svcCtx)
	return l.GetApiById(in)
}

func (s *SystemServer) DeleteApi(ctx context.Context, in *system.IDsReq) (*system.BaseResp, error) {
	l := api.NewDeleteApiLogic(ctx, s.svcCtx)
	return l.DeleteApi(in)
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

func (s *SystemServer) DeleteMenu(ctx context.Context, in *system.IDReq) (*system.BaseResp, error) {
	l := menu.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

func (s *SystemServer) GetMenuListByRole(ctx context.Context, in *system.BaseMsg) (*system.MenuInfoList, error) {
	l := menu.NewGetMenuListByRoleLogic(ctx, s.svcCtx)
	return l.GetMenuListByRole(in)
}

func (s *SystemServer) GetMenuList(ctx context.Context, in *system.PageInfoReq) (*system.MenuInfoList, error) {
	l := menu.NewGetMenuListLogic(ctx, s.svcCtx)
	return l.GetMenuList(in)
}

// OauthProvider management
func (s *SystemServer) CreateOauthProvider(ctx context.Context, in *system.OauthProviderInfo) (*system.BaseIDResp, error) {
	l := oauthprovider.NewCreateOauthProviderLogic(ctx, s.svcCtx)
	return l.CreateOauthProvider(in)
}

func (s *SystemServer) UpdateOauthProvider(ctx context.Context, in *system.OauthProviderInfo) (*system.BaseResp, error) {
	l := oauthprovider.NewUpdateOauthProviderLogic(ctx, s.svcCtx)
	return l.UpdateOauthProvider(in)
}

func (s *SystemServer) GetOauthProviderList(ctx context.Context, in *system.OauthProviderListReq) (*system.OauthProviderListResp, error) {
	l := oauthprovider.NewGetOauthProviderListLogic(ctx, s.svcCtx)
	return l.GetOauthProviderList(in)
}

func (s *SystemServer) GetOauthProviderById(ctx context.Context, in *system.IDReq) (*system.OauthProviderInfo, error) {
	l := oauthprovider.NewGetOauthProviderByIdLogic(ctx, s.svcCtx)
	return l.GetOauthProviderById(in)
}

func (s *SystemServer) DeleteOauthProvider(ctx context.Context, in *system.IDsReq) (*system.BaseResp, error) {
	l := oauthprovider.NewDeleteOauthProviderLogic(ctx, s.svcCtx)
	return l.DeleteOauthProvider(in)
}

func (s *SystemServer) OauthLogin(ctx context.Context, in *system.OauthLoginReq) (*system.OauthRedirectResp, error) {
	l := oauthprovider.NewOauthLoginLogic(ctx, s.svcCtx)
	return l.OauthLogin(in)
}

func (s *SystemServer) OauthCallback(ctx context.Context, in *system.CallbackReq) (*system.UserInfo, error) {
	l := oauthprovider.NewOauthCallbackLogic(ctx, s.svcCtx)
	return l.OauthCallback(in)
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

// Token management
func (s *SystemServer) CreateToken(ctx context.Context, in *system.TokenInfo) (*system.BaseUUIDResp, error) {
	l := token.NewCreateTokenLogic(ctx, s.svcCtx)
	return l.CreateToken(in)
}

func (s *SystemServer) UpdateToken(ctx context.Context, in *system.TokenInfo) (*system.BaseResp, error) {
	l := token.NewUpdateTokenLogic(ctx, s.svcCtx)
	return l.UpdateToken(in)
}

func (s *SystemServer) GetTokenList(ctx context.Context, in *system.TokenListReq) (*system.TokenListResp, error) {
	l := token.NewGetTokenListLogic(ctx, s.svcCtx)
	return l.GetTokenList(in)
}

func (s *SystemServer) GetTokenById(ctx context.Context, in *system.UUIDReq) (*system.TokenInfo, error) {
	l := token.NewGetTokenByIdLogic(ctx, s.svcCtx)
	return l.GetTokenById(in)
}

func (s *SystemServer) DeleteToken(ctx context.Context, in *system.UUIDsReq) (*system.BaseResp, error) {
	l := token.NewDeleteTokenLogic(ctx, s.svcCtx)
	return l.DeleteToken(in)
}

func (s *SystemServer) BlockUserAllToken(ctx context.Context, in *system.UUIDReq) (*system.BaseResp, error) {
	l := token.NewBlockUserAllTokenLogic(ctx, s.svcCtx)
	return l.BlockUserAllToken(in)
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
