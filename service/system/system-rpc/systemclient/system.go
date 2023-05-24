// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package systemclient

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiInfo                  = system.ApiInfo
	ApiListReq               = system.ApiListReq
	ApiListResp              = system.ApiListResp
	BaseIDResp               = system.BaseIDResp
	BaseMsg                  = system.BaseMsg
	BaseResp                 = system.BaseResp
	BaseUUIDResp             = system.BaseUUIDResp
	CallbackReq              = system.CallbackReq
	DepartmentInfo           = system.DepartmentInfo
	DepartmentListReq        = system.DepartmentListReq
	DepartmentListResp       = system.DepartmentListResp
	DictionaryDetailInfo     = system.DictionaryDetailInfo
	DictionaryDetailListReq  = system.DictionaryDetailListReq
	DictionaryDetailListResp = system.DictionaryDetailListResp
	DictionaryInfo           = system.DictionaryInfo
	DictionaryListReq        = system.DictionaryListReq
	DictionaryListResp       = system.DictionaryListResp
	Empty                    = system.Empty
	IDReq                    = system.IDReq
	IDsReq                   = system.IDsReq
	MenuInfo                 = system.MenuInfo
	MenuInfoList             = system.MenuInfoList
	MenuRoleInfo             = system.MenuRoleInfo
	MenuRoleListResp         = system.MenuRoleListResp
	Meta                     = system.Meta
	OauthLoginReq            = system.OauthLoginReq
	OauthProviderInfo        = system.OauthProviderInfo
	OauthProviderListReq     = system.OauthProviderListReq
	OauthProviderListResp    = system.OauthProviderListResp
	OauthRedirectResp        = system.OauthRedirectResp
	PageInfoReq              = system.PageInfoReq
	PositionInfo             = system.PositionInfo
	PositionListReq          = system.PositionListReq
	PositionListResp         = system.PositionListResp
	RoleInfo                 = system.RoleInfo
	RoleListReq              = system.RoleListReq
	RoleListResp             = system.RoleListResp
	TokenInfo                = system.TokenInfo
	TokenListReq             = system.TokenListReq
	TokenListResp            = system.TokenListResp
	UUIDReq                  = system.UUIDReq
	UUIDsReq                 = system.UUIDsReq
	UserInfo                 = system.UserInfo
	UserListReq              = system.UserListReq
	UserListResp             = system.UserListResp
	UsernameReq              = system.UsernameReq

	System interface {
		// API management
		CreateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetApiList(ctx context.Context, in *ApiListReq, opts ...grpc.CallOption) (*ApiListResp, error)
		GetApiById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*ApiInfo, error)
		DeleteApi(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Department management
		CreateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetDepartmentList(ctx context.Context, in *DepartmentListReq, opts ...grpc.CallOption) (*DepartmentListResp, error)
		GetDepartmentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DepartmentInfo, error)
		DeleteDepartment(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Dictionary management
		CreateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetDictionaryList(ctx context.Context, in *DictionaryListReq, opts ...grpc.CallOption) (*DictionaryListResp, error)
		GetDictionaryById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DictionaryInfo, error)
		DeleteDictionary(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// DictionaryDetail management
		CreateDictionaryDetail(ctx context.Context, in *DictionaryDetailInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateDictionaryDetail(ctx context.Context, in *DictionaryDetailInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetDictionaryDetailList(ctx context.Context, in *DictionaryDetailListReq, opts ...grpc.CallOption) (*DictionaryDetailListResp, error)
		GetDictionaryDetailById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DictionaryDetailInfo, error)
		DeleteDictionaryDetail(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Menu management
		CreateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetMenuListByRole(ctx context.Context, in *BaseMsg, opts ...grpc.CallOption) (*MenuInfoList, error)
		GetMenuList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*MenuInfoList, error)
		// OauthProvider management
		CreateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetOauthProviderList(ctx context.Context, in *OauthProviderListReq, opts ...grpc.CallOption) (*OauthProviderListResp, error)
		GetOauthProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OauthProviderInfo, error)
		DeleteOauthProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthRedirectResp, error)
		OauthCallback(ctx context.Context, in *CallbackReq, opts ...grpc.CallOption) (*UserInfo, error)
		// Position management
		CreatePosition(ctx context.Context, in *PositionInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdatePosition(ctx context.Context, in *PositionInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetPositionList(ctx context.Context, in *PositionListReq, opts ...grpc.CallOption) (*PositionListResp, error)
		GetPositionById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*PositionInfo, error)
		DeletePosition(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Role management
		CreateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetRoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		GetRoleById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleInfo, error)
		DeleteRole(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Token management
		CreateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetTokenList(ctx context.Context, in *TokenListReq, opts ...grpc.CallOption) (*TokenListResp, error)
		GetTokenById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*TokenInfo, error)
		DeleteToken(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		BlockUserAllToken(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*BaseResp, error)
		// User management
		CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetUserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfo, error)
		GetUserByUsername(ctx context.Context, in *UsernameReq, opts ...grpc.CallOption) (*UserInfo, error)
		DeleteUser(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultSystem struct {
		cli zrpc.Client
	}
)

func NewSystem(cli zrpc.Client) System {
	return &defaultSystem{
		cli: cli,
	}
}

// API management
func (m *defaultSystem) CreateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateApi(ctx, in, opts...)
}

func (m *defaultSystem) UpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateApi(ctx, in, opts...)
}

func (m *defaultSystem) GetApiList(ctx context.Context, in *ApiListReq, opts ...grpc.CallOption) (*ApiListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetApiList(ctx, in, opts...)
}

func (m *defaultSystem) GetApiById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*ApiInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetApiById(ctx, in, opts...)
}

func (m *defaultSystem) DeleteApi(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

func (m *defaultSystem) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Department management
func (m *defaultSystem) CreateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateDepartment(ctx, in, opts...)
}

func (m *defaultSystem) UpdateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateDepartment(ctx, in, opts...)
}

func (m *defaultSystem) GetDepartmentList(ctx context.Context, in *DepartmentListReq, opts ...grpc.CallOption) (*DepartmentListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetDepartmentList(ctx, in, opts...)
}

func (m *defaultSystem) GetDepartmentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DepartmentInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetDepartmentById(ctx, in, opts...)
}

func (m *defaultSystem) DeleteDepartment(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteDepartment(ctx, in, opts...)
}

// Dictionary management
func (m *defaultSystem) CreateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateDictionary(ctx, in, opts...)
}

func (m *defaultSystem) UpdateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateDictionary(ctx, in, opts...)
}

func (m *defaultSystem) GetDictionaryList(ctx context.Context, in *DictionaryListReq, opts ...grpc.CallOption) (*DictionaryListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetDictionaryList(ctx, in, opts...)
}

func (m *defaultSystem) GetDictionaryById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DictionaryInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetDictionaryById(ctx, in, opts...)
}

func (m *defaultSystem) DeleteDictionary(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteDictionary(ctx, in, opts...)
}

// DictionaryDetail management
func (m *defaultSystem) CreateDictionaryDetail(ctx context.Context, in *DictionaryDetailInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateDictionaryDetail(ctx, in, opts...)
}

func (m *defaultSystem) UpdateDictionaryDetail(ctx context.Context, in *DictionaryDetailInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateDictionaryDetail(ctx, in, opts...)
}

func (m *defaultSystem) GetDictionaryDetailList(ctx context.Context, in *DictionaryDetailListReq, opts ...grpc.CallOption) (*DictionaryDetailListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetDictionaryDetailList(ctx, in, opts...)
}

func (m *defaultSystem) GetDictionaryDetailById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DictionaryDetailInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetDictionaryDetailById(ctx, in, opts...)
}

func (m *defaultSystem) DeleteDictionaryDetail(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteDictionaryDetail(ctx, in, opts...)
}

// Menu management
func (m *defaultSystem) CreateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateMenu(ctx, in, opts...)
}

func (m *defaultSystem) UpdateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

func (m *defaultSystem) DeleteMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
}

func (m *defaultSystem) GetMenuListByRole(ctx context.Context, in *BaseMsg, opts ...grpc.CallOption) (*MenuInfoList, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetMenuListByRole(ctx, in, opts...)
}

func (m *defaultSystem) GetMenuList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*MenuInfoList, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetMenuList(ctx, in, opts...)
}

// OauthProvider management
func (m *defaultSystem) CreateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateOauthProvider(ctx, in, opts...)
}

func (m *defaultSystem) UpdateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateOauthProvider(ctx, in, opts...)
}

func (m *defaultSystem) GetOauthProviderList(ctx context.Context, in *OauthProviderListReq, opts ...grpc.CallOption) (*OauthProviderListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetOauthProviderList(ctx, in, opts...)
}

func (m *defaultSystem) GetOauthProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OauthProviderInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetOauthProviderById(ctx, in, opts...)
}

func (m *defaultSystem) DeleteOauthProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteOauthProvider(ctx, in, opts...)
}

func (m *defaultSystem) OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthRedirectResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.OauthLogin(ctx, in, opts...)
}

func (m *defaultSystem) OauthCallback(ctx context.Context, in *CallbackReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.OauthCallback(ctx, in, opts...)
}

// Position management
func (m *defaultSystem) CreatePosition(ctx context.Context, in *PositionInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreatePosition(ctx, in, opts...)
}

func (m *defaultSystem) UpdatePosition(ctx context.Context, in *PositionInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdatePosition(ctx, in, opts...)
}

func (m *defaultSystem) GetPositionList(ctx context.Context, in *PositionListReq, opts ...grpc.CallOption) (*PositionListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetPositionList(ctx, in, opts...)
}

func (m *defaultSystem) GetPositionById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*PositionInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetPositionById(ctx, in, opts...)
}

func (m *defaultSystem) DeletePosition(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeletePosition(ctx, in, opts...)
}

// Role management
func (m *defaultSystem) CreateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

func (m *defaultSystem) UpdateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultSystem) GetRoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetRoleList(ctx, in, opts...)
}

func (m *defaultSystem) GetRoleById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetRoleById(ctx, in, opts...)
}

func (m *defaultSystem) DeleteRole(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

// Token management
func (m *defaultSystem) CreateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateToken(ctx, in, opts...)
}

func (m *defaultSystem) UpdateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateToken(ctx, in, opts...)
}

func (m *defaultSystem) GetTokenList(ctx context.Context, in *TokenListReq, opts ...grpc.CallOption) (*TokenListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetTokenList(ctx, in, opts...)
}

func (m *defaultSystem) GetTokenById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*TokenInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetTokenById(ctx, in, opts...)
}

func (m *defaultSystem) DeleteToken(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteToken(ctx, in, opts...)
}

func (m *defaultSystem) BlockUserAllToken(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.BlockUserAllToken(ctx, in, opts...)
}

// User management
func (m *defaultSystem) CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultSystem) UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultSystem) GetUserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetUserList(ctx, in, opts...)
}

func (m *defaultSystem) GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultSystem) GetUserByUsername(ctx context.Context, in *UsernameReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetUserByUsername(ctx, in, opts...)
}

func (m *defaultSystem) DeleteUser(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}
