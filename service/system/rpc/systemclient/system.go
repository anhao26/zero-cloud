// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package systemclient

import (
	"context"

	"github.com/anhao26/zero-cloud/service/system/rpc/types/system"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp         = system.BaseIDResp
	BaseResp           = system.BaseResp
	BaseUUIDResp       = system.BaseUUIDResp
	DepartmentInfo     = system.DepartmentInfo
	DepartmentListReq  = system.DepartmentListReq
	DepartmentListResp = system.DepartmentListResp
	Empty              = system.Empty
	IDReq              = system.IDReq
	IDsReq             = system.IDsReq
	MenuInfo           = system.MenuInfo
	MenuListReq        = system.MenuListReq
	MenuListResp       = system.MenuListResp
	PageInfoReq        = system.PageInfoReq
	PositionInfo       = system.PositionInfo
	PositionListReq    = system.PositionListReq
	PositionListResp   = system.PositionListResp
	RoleInfo           = system.RoleInfo
	RoleListReq        = system.RoleListReq
	RoleListResp       = system.RoleListResp
	UUIDReq            = system.UUIDReq
	UUIDsReq           = system.UUIDsReq
	UserInfo           = system.UserInfo
	UserListReq        = system.UserListReq
	UserListResp       = system.UserListResp

	System interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Department management
		CreateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetDepartmentList(ctx context.Context, in *DepartmentListReq, opts ...grpc.CallOption) (*DepartmentListResp, error)
		GetDepartmentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DepartmentInfo, error)
		DeleteDepartment(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Menu management
		CreateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetMenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		GetMenuById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuInfo, error)
		DeleteMenu(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
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
		// User management
		CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetUserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfo, error)
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

// Menu management
func (m *defaultSystem) CreateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.CreateMenu(ctx, in, opts...)
}

func (m *defaultSystem) UpdateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

func (m *defaultSystem) GetMenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetMenuList(ctx, in, opts...)
}

func (m *defaultSystem) GetMenuById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuInfo, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.GetMenuById(ctx, in, opts...)
}

func (m *defaultSystem) DeleteMenu(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
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

func (m *defaultSystem) DeleteUser(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}
