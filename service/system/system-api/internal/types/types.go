// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The simplest message | 最简单的信息
// swagger:response SimpleMsg
type SimpleMsg struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// Required: true
	Page uint64 `json:"page" validate:"number"`
	// Page size | 单页数据行数
	// Required: true
	// Maximum: 100000
	PageSize uint64 `json:"pageSize" validate:"number,max=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// Required: true
	// Max length: 36
	Id string `json:"id" validate:"len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
}

// The base ID response data | 基础ID信息
// swagger:model BaseIDInfo
type BaseIDInfo struct {
	// ID
	Id uint64 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础UUID信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id string `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The information of captcha | 验证码数据
// swagger:model CaptchaInfo
type CaptchaInfo struct {
	CaptchaId string `json:"captchaId"`
	ImgPath   string `json:"imgPath"`
}

// The response data of captcha | 验证码返回数据
// swagger:model CaptchaResp
type CaptchaResp struct {
	BaseDataInfo
	// The menu authorization data | 菜单授权信息数据
	Data CaptchaInfo `json:"data"`
}

// Login request | 登录参数
// swagger:model LoginReq
type LoginReq struct {
	// User Name | 用户名
	// required : true
	// max length : 20
	Username string `json:"username" validate:"required,alphanum,max=20"`
	// Password | 密码
	// Required: true
	// Min length: 6
	// Max length: 30
	Password string `json:"password" validate:"required,max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// Required: true
	// Max length: 20
	CaptchaId string `json:"captchaId"  validate:"required,len=20"`
	// The Captcha which users input | 用户输入的验证码
	// Required: true
	// Max length: 5
	Captcha string `json:"captcha" validate:"required,len=5"`
}

// The log in response data | 登录返回数据
// swagger:model LoginResp
type LoginResp struct {
	BaseDataInfo
	// The log in information | 登陆返回的数据信息
	Data LoginInfo `json:"data"`
}

// The log in information | 登陆返回的数据信息
// swagger:model LoginInfo
type LoginInfo struct {
	// User's UUID | 用户的UUID
	UserId string `json:"userId"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	Expire uint64 `json:"expire"`
}

// The response data of user's basic information | 用户基本信息返回数据
// swagger:model UserBaseIDInfoResp
type UserBaseIDInfoResp struct {
	BaseDataInfo
	// The  data of user's basic information | 用户基本信息
	Data UserBaseIDInfo `json:"data"`
}

// The  data of user's basic information | 用户基本信息
// swagger:model UserBaseIDInfo
type UserBaseIDInfo struct {
	// User's UUID | 用户的UUID
	UUID string `json:"userId"`
	// User's name | 用户名
	Username string `json:"username"`
	// User's nickname | 用户的昵称
	Nickname string `json:"nickname"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// The home page that the user enters after logging in | 用户登陆后进入的首页
	HomePath string `json:"homePath"`
	// The description of user | 用户的描述信息
	Description string `json:"desc"`
}

// The permission code for front end permission control | 权限码： 用于前端权限控制
// swagger:model PermCodeResp
type PermCodeResp struct {
	BaseDataInfo
	// Permission code data | 权限码数据
	Data []string `json:"data"`
}

// Get user list request params | 用户列表请求参数
// swagger:model UserListReq
type UserListReq struct {
	PageInfo
	// User Name | 用户名
	// max length : 20
	Username string `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	// User's nickname | 用户的昵称
	// max length : 10
	Nickname string `json:"nickname,optional" validate:"omitempty,alphanumunicode,max=10"`
	// User's mobile phone number | 用户的手机号码
	// max length : 18
	Mobile string `json:"mobile,optional" validate:"omitempty,numeric,max=18"`
	// The user's email address | 用户的邮箱
	// max length : 100
	Email string `json:"email,optional" validate:"omitempty,email,max=100"`
	// User's role ID | 用户的角色ID
	RoleIds []uint64 `json:"roleIds,optional"`
	// The user's department ID | 用户所属部门ID
	DepartmentId uint64 `json:"departmentId,optional"`
	// User's position id | 用户的职位ID
	PositionId uint64 `json:"positionId,optional"`
}

// The response data of user list | 用户列表数据
// swagger:model UserListResp
type UserListResp struct {
	BaseDataInfo
	// User list data | User列表数据
	Data UserListInfo `json:"data"`
}

// User list data | 用户列表数据
// swagger:model UserListInfo
type UserListInfo struct {
	BaseListInfo
	// The API list data | User列表数据
	Data []UserInfo `json:"data"`
}

// The response data of user information | 用户信息
// swagger:model UserInfo
type UserInfo struct {
	BaseUUIDInfo
	// Status | 状态
	// max : 20
	Status uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	// Username | 用户名
	// max length : 50
	Username string `json:"username,optional" validate:"omitempty,max=50"`
	// Nickname | 昵称
	// max length : 40
	Nickname string `json:"nickname,optional" validate:"omitempty,max=40"`
	// Password | 密码
	// min length : 6
	Password string `json:"password,optional" validate:"omitempty,min=6"`
	// Description | 描述
	// max length : 100
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	// HomePath | 首页
	// max length : 70
	HomePath string `json:"homePath,optional" validate:"omitempty,max=70"`
	// RoleId | 角色ID
	RoleIds []uint64 `json:"roleIds,optional"`
	// Mobile | 手机号
	// max length : 18
	Mobile string `json:"mobile,optional" validate:"omitempty,max=18"`
	// Email | 邮箱
	// max length : 80
	Email string `json:"email,optional" validate:"omitempty,max=80"`
	// Avatar | 头像地址
	// max length : 300
	Avatar string `json:"avatar,optional" validate:"omitempty,max=300"`
	// Department ID | 部门ID
	DepartmentId uint64 `json:"departmentId,optional"`
	// Position ID | 职位ID
	PositionIds []uint64 `json:"positionId,optional"`
}

// User information response | 用户信息返回体
// swagger:model UserInfoResp
type UserInfoResp struct {
	BaseDataInfo
	// User information | User数据
	Data UserInfo `json:"data"`
}

// change user's password request | 修改密码请求参数
// swagger:model ChangePasswordReq
type ChangePasswordReq struct {
	// User's old password | 用户旧密码
	// required : true
	// max length : 30
	OldPassword string `json:"oldPassword" validate:"required,max=30"`
	// User's new password | 用户新密码
	// required : true
	// max length : 30
	NewPassword string `json:"newPassword" validate:"required,max=30"`
}

// The profile response data | 个人信息返回数据
// swagger:model ProfileResp
type ProfileResp struct {
	BaseDataInfo
	// The profile information | 个人信息
	Data ProfileInfo `json:"data"`
}

// The profile information | 个人信息
// swagger:model ProfileInfo
type ProfileInfo struct {
	// user's nickname | 用户的昵称
	// max length : 10
	Nickname string `json:"nickname" validate:"omitempty,alphanumunicode,max=10"`
	// The user's avatar path | 用户的头像路径
	// max length : 300
	Avatar string `json:"avatar" validate:"omitempty,max=300"`
	// User's mobile phone number | 用户的手机号码
	// max length : 18
	Mobile string `json:"mobile" validate:"omitempty,numeric,max=18"`
	// The user's email address | 用户的邮箱
	// max length : 100
	Email string `json:"email" validate:"omitempty,email,max=100"`
}

// register request | 注册参数
// swagger:model RegisterReq
type RegisterReq struct {
	// User Name | 用户名
	// required : true
	// max length : 20
	Username string `json:"username" validate:"required,alphanum,max=20"`
	// Password | 密码
	// required : true
	// max length : 30
	// min length : 6
	Password string `json:"password" validate:"required,max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// required : true
	// max length : 20
	// min length : 20
	CaptchaId string `json:"captchaId" validate:"required,len=20"`
	// The Captcha which users input | 用户输入的验证码
	// required : true
	// max length : 5
	// min length : 5
	Captcha string `json:"captcha" validate:"required,len=5"`
	// The user's email address | 用户的邮箱
	// required : true
	// max length : 100
	Email string `json:"email" validate:"required,email,max=100"`
}

// The response data of menu information | 菜单信息
// swagger:model MenuInfo
type MenuInfo struct {
	BaseIDInfo
	// Translated Name | 国际化展示名称
	Trans string `json:"trans,optional"`
	// Level | 菜单层级
	Level uint32 `json:"level,optional"`
	// ParentId | 父级菜单ID
	ParentId uint64 `json:"parentId,optional"`
	// Path | 菜单访问路径
	Path string `json:"path,optional"`
	// Menu name | 菜单名称
	Name string `json:"name,optional"`
	// Redirect | 跳转地址
	Redirect string `json:"redirect,optional"`
	// Component | 组件地址
	Component string `json:"component,optional"`
	// Sort | 排序
	Sort uint32 `json:"sort,optional"`
	// Disabled | 是否启用
	Disabled bool `json:"disabled,optional"`
	// Meta | 菜单meta数据
	Meta Meta `json:"meta"`
	// MenuType | 菜单类型
	MenuType uint32 `json:"menuType,optional"`
}

// The meta data of menu | 菜单的meta数据
// swagger:model Meta
type Meta struct {
	// Menu title show in page | 菜单显示名
	// Max length: 50
	Title string `json:"title" validate:"max=50"`
	// Menu Icon | 菜单图标
	// Max length: 50
	Icon string `json:"icon" validate:"max=50"`
	// Hide menu | 隐藏菜单
	HideMenu bool `json:"hideMenu" validate:"boolean"`
	// If hide the breadcrumb | 隐藏面包屑
	HideBreadcrumb bool `json:"hideBreadcrumb,optional" validate:"boolean"`
	// Do not keep alive the tab | 不缓存Tab
	IgnoreKeepAlive bool `json:"ignoreKeepAlive,optional" validate:"boolean"`
	// Hide the tab header | 当前路由不在标签页显示
	HideTab bool `json:"hideTab,optional" validate:"boolean"`
	// Iframe path | 内嵌iframe的地址
	FrameSrc string `json:"frameSrc,optional"`
	// The route carries parameters or not | 如果该路由会携带参数，且需要在tab页上面显示。则需要设置为true
	CarryParam bool `json:"carryParam,optional" validate:"boolean"`
	// Hide children menu or not | 隐藏所有子菜单
	HideChildrenInMenu bool `json:"hideChildrenInMenu,optional" validate:"boolean"`
	// Affix tab | 是否固定标签
	Affix bool `json:"affix,optional" validate:"boolean"`
	// The maximum number of pages the router can open | 动态路由可打开Tab页数
	DynamicLevel uint32 `json:"dynamicLevel" validate:"number,lt=30"`
	// The real path of the route without dynamic part | 动态路由的实际Path, 即去除路由的动态部分
	RealPath string `json:"realPath,optional"`
}

// The response data of menu list | 菜单列表返回数据
// swagger:model MenuListResp
type MenuListResp struct {
	BaseDataInfo
	// Menu list data | Menu列表数据
	Data MenuListInfo `json:"data"`
}

// Menu list data | Menu列表数据
// swagger:model MenuListInfo
type MenuListInfo struct {
	BaseListInfo
	// The menu list data | 菜单列表数据
	Data []MenuInfo `json:"data"`
}

// Menu list data response | 菜单列表数据返回体
// swagger:model MenuPlainInfoListResp
type MenuPlainInfoListResp struct {
	BaseDataInfo
	// Menu list data | Menu列表数据
	Data MenuPlainInfoList `json:"data"`
}

// Menu list data | 菜单列表数据
type MenuPlainInfoList struct {
	BaseListInfo
	// The menu list data | 菜单列表数据
	Data []MenuPlainInfo `json:"data"`
}

// Menu information plain | 菜单信息无嵌套
// swagger:model MenuPlainInfo
type MenuPlainInfo struct {
	Id uint64 `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
	// Translated Name | 国际化展示名称
	Trans string `json:"trans,optional"`
	// Level | 菜单层级
	// max : 20
	Level uint32 `json:"level,optional" validate:"omitempty,lt=20"`
	// ParentId | 父级菜单ID
	ParentId uint64 `json:"parentId,optional"`
	// Path | 菜单访问路径
	// max length : 200
	Path string `json:"path,optional" validate:"omitempty,max=200"`
	// Menu name | 菜单名称
	// max length : 50
	Name string `json:"name,optional" validate:"omitempty,max=50"`
	// Redirect | 跳转地址
	// max length : 300
	Redirect string `json:"redirect,optional" validate:"omitempty,max=300"`
	// Component | 组件地址
	// max length : 80
	Component string `json:"component,optional" validate:"omitempty,max=80"`
	// Sort | 排序
	// max : 10000
	Sort uint32 `json:"sort,optional" validate:"omitempty,lt=10000"`
	// Disabled | 是否启用
	Disabled bool `json:"disabled,optional"`
	// MenuType | 菜单类型
	// max : 10
	MenuType uint32 `json:"menuType,optional" validate:"omitempty,lt=10"`
	// Menu title show in page | 菜单显示名
	// max length : 50
	Title string `json:"title" validate:"omitempty,max=50"`
	// Menu Icon | 菜单图标
	// max length : 50
	Icon string `json:"icon" validate:"omitempty,max=50"`
	// Hide menu | 隐藏菜单
	HideMenu bool `json:"hideMenu" validate:"boolean"`
	// If hide the breadcrumb | 隐藏面包屑
	HideBreadcrumb bool `json:"hideBreadcrumb,optional" validate:"boolean"`
	// Do not keep alive the tab | 不缓存Tab
	IgnoreKeepAlive bool `json:"ignoreKeepAlive,optional" validate:"boolean"`
	// Hide the tab header | 当前路由不在标签页显示
	HideTab bool `json:"hideTab,optional" validate:"boolean"`
	// Iframe path | 内嵌iframe的地址
	// max length : 300
	FrameSrc string `json:"frameSrc,optional" validate:"omitempty,max=300"`
	// The route carries parameters or not | 如果该路由会携带参数，且需要在tab页上面显示。则需要设置为true
	CarryParam bool `json:"carryParam,optional" validate:"boolean"`
	// Hide children menu or not | 隐藏所有子菜单
	HideChildrenInMenu bool `json:"hideChildrenInMenu,optional" validate:"boolean"`
	// Affix tab | 是否固定标签
	Affix bool `json:"affix,optional" validate:"boolean"`
	// The maximum number of pages the router can open | 动态路由可打开Tab页数
	DynamicLevel uint32 `json:"dynamicLevel" validate:"number,lt=30"`
	// The real path of the route without dynamic part | 动态路由的实际Path, 即去除路由的动态部分
	RealPath string `json:"realPath,optional"`
}

// The response data of role information | 角色信息
// swagger:model RoleInfo
type RoleInfo struct {
	BaseIDInfo
	// Translated Name | 展示名称
	Trans string `json:"trans,optional"`
	// Status | 状态
	// max : 20
	Status uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	// Name | 角色名称
	// max length : 30
	Name string `json:"name,optional" validate:"omitempty,max=30"`
	// Role code | 角色码
	// max length : 20
	Code string `json:"code,optional" validate:"omitempty,max=20"`
	// DefaultRouter | 默认首页
	// max length : 80
	DefaultRouter string `json:"defaultRouter,optional" validate:"omitempty,max=80"`
	// Remark | 备注
	// max length : 200
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
	// Sort | 排序
	// max : 10000
	Sort uint32 `json:"sort,optional" validate:"omitempty,lt=10000"`
}

// The response data of role list | 角色列表数据
// swagger:model RoleListResp
type RoleListResp struct {
	BaseDataInfo
	// Role list data | 角色列表数据
	Data RoleListInfo `json:"data"`
}

// Role list data | 角色列表数据
// swagger:model RoleListInfo
type RoleListInfo struct {
	BaseListInfo
	// The API list data | 角色列表数据
	Data []RoleInfo `json:"data"`
}

// Get role list request params | 角色列表请求参数
// swagger:model RoleListReq
type RoleListReq struct {
	PageInfo
	// Name | 角色名称
	Name string `json:"name,optional"`
}

// Role information response | 角色信息返回体
// swagger:model RoleInfoResp
type RoleInfoResp struct {
	BaseDataInfo
	// Role information | 角色数据
	Data RoleInfo `json:"data"`
}

// The response data of department information | 部门信息
// swagger:model DepartmentInfo
type DepartmentInfo struct {
	BaseIDInfo
	// Translated Name | 展示名称
	Trans string `json:"trans,optional"`
	// Status | 状态
	// max : 20
	Status uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	// Sort | 排序
	// max : 10000
	Sort uint32 `json:"sort,optional" validate:"omitempty,lt=10000"`
	// Name | 部门名称
	// min length : 1
	// max length : 50
	Name string `json:"name,optional" validate:"omitempty,min=1,max=50"`
	// Ancestors | 父级部门列表
	// max length : 200
	Ancestors string `json:"ancestors,optional" validate:"omitempty,max=200"`
	// Leader | 部门负责人
	// max length : 20
	Leader string `json:"leader,optional" validate:"omitempty,max=20"`
	// Phone | 电话号码
	// max length : 18
	Phone string `json:"phone,optional" validate:"omitempty,max=18"`
	// Email | 邮箱
	// min length : 5
	// max length : 70
	Email string `json:"email,optional" validate:"omitempty,min=5,max=70"`
	// Remark | 备注
	// max length : 200
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
	// ParentId | 父级 ID
	ParentId uint64 `json:"parentId,optional"`
}

// The response data of department list | 部门列表数据
// swagger:model DepartmentListResp
type DepartmentListResp struct {
	BaseDataInfo
	// Department list data | 部门列表数据
	Data DepartmentListInfo `json:"data"`
}

// Department list data | 部门列表数据
// swagger:model DepartmentListInfo
type DepartmentListInfo struct {
	BaseListInfo
	// The API list data | 部门列表数据
	Data []DepartmentInfo `json:"data"`
}

// Get department list request params | 部门列表请求参数
// swagger:model DepartmentListReq
type DepartmentListReq struct {
	PageInfo
	// Name | 部门名称
	// min length : 1
	// max length : 50
	Name string `json:"name,optional" validate:"omitempty,min=1,max=50"`
	// Leader | 部门负责人
	// max length : 20
	Leader string `json:"leader,optional" validate:"omitempty,max=20"`
}

// Department information response | 部门信息返回体
// swagger:model DepartmentInfoResp
type DepartmentInfoResp struct {
	BaseDataInfo
	// Department information | 部门数据
	Data DepartmentInfo `json:"data"`
}

// The response data of position information | 职位信息
// swagger:model PositionInfo
type PositionInfo struct {
	BaseIDInfo
	// Translated Name | 展示名称
	Trans string `json:"trans,optional"`
	// Status | 状态
	// max : 20
	Status uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	// Sort | 排序
	// max : 10000
	Sort uint32 `json:"sort,optional" validate:"omitempty,lt=10000"`
	// Name | 职位名称
	// max length : 50
	Name string `json:"name,optional" validate:"omitempty,max=50"`
	// Code | 职位代码
	// max length : 20
	Code string `json:"code,optional" validate:"omitempty,max=20"`
	// Remark | 备注
	// max length : 200
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
}

// The response data of position list | 职位列表数据
// swagger:model PositionListResp
type PositionListResp struct {
	BaseDataInfo
	// Position list data | 职位列表数据
	Data PositionListInfo `json:"data"`
}

// Position list data | 职位列表数据
// swagger:model PositionListInfo
type PositionListInfo struct {
	BaseListInfo
	// The API list data | 职位列表数据
	Data []PositionInfo `json:"data"`
}

// Get position list request params | 职位列表请求参数
// swagger:model PositionListReq
type PositionListReq struct {
	PageInfo
	// Name | 职位名称
	// max length : 50
	Name string `json:"name,optional" validate:"omitempty,max=50"`
	// Code | 职位代码
	// max length : 20
	Code string `json:"code,optional" validate:"omitempty,max=20"`
	// Remark | 备注
	// max length : 200
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
}

// Position information response | 职位信息返回体
// swagger:model PositionInfoResp
type PositionInfoResp struct {
	BaseDataInfo
	// Position information | 职位数据
	Data PositionInfo `json:"data"`
}

// The response data of token information | Token信息
// swagger:model TokenInfo
type TokenInfo struct {
	BaseUUIDInfo
	// Status | 状态
	Status uint32 `json:"status,optional"`
	// User's UUID | 用户的UUID
	Uuid string `json:"uuid,optional"`
	// Token | 用户的Token
	Token string `json:"token,optional"`
	// Source | Token 来源
	Source string `json:"source,optional"`
	// ExpiredAt | 过期时间
	ExpiredAt int64 `json:"expiredAt,optional"`
}

// The response data of token list | Token列表数据
// swagger:model TokenListResp
type TokenListResp struct {
	BaseDataInfo
	// Token list data | Token列表数据
	Data TokenListInfo `json:"data"`
}

// Token list data | Token列表数据
// swagger:model TokenListInfo
type TokenListInfo struct {
	BaseListInfo
	// The API list data | Token列表数据
	Data []TokenInfo `json:"data"`
}

// Get token list request params | Token列表请求参数
// swagger:model TokenListReq
type TokenListReq struct {
	PageInfo
	// Username
	Username string `json:"username,optional"`
	// Nickname
	Nickname string `json:"nickname,optional"`
	// Email
	Email string `json:"email,optional"`
	// Uuid
	Uuid string `json:"uuid,optional"`
}

// Token information response | Token信息返回体
// swagger:model TokenInfoResp
type TokenInfoResp struct {
	BaseDataInfo
	// Token information | Token数据
	Data TokenInfo `json:"data"`
}

// The API information | API信息
// swagger:model ApiInfo
type ApiInfo struct {
	BaseIDInfo
	// Translated Name | 多语言名称
	Trans string `json:"trans,optional"`
	// API path | API路径
	// min length : 1
	// max length : 80
	Path string `json:"path,optional" validate:"omitempty,min=1,max=80"`
	// API Description | API 描述
	// max length : 100
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	// API group | API分组
	// min length : 1
	// max length : 20
	Group string `json:"group,optional" validate:"omitempty,min=1,max=20"`
	// API request method e.g. POST | API请求类型 如POST
	// min length : 3
	// max length : 4
	Method string `json:"method,optional" validate:"omitempty,uppercase,min=3,max=4"`
}

// The response data of API list | API列表数据
// swagger:model ApiListResp
type ApiListResp struct {
	BaseDataInfo
	// API list data | API 列表数据
	Data ApiListInfo `json:"data"`
}

// API list data | API 列表数据
// swagger:model ApiListInfo
type ApiListInfo struct {
	BaseListInfo
	// The API list data | API列表数据
	Data []ApiInfo `json:"data"`
}

// Get API list request params | API列表请求参数
// swagger:model ApiListReq
type ApiListReq struct {
	PageInfo
	// API path | API路径
	// max length : 200
	Path string `json:"path,optional" validate:"omitempty,max=200"`
	// API Description | API 描述
	// max length : 100
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	// API group | API分组
	// max length : 20
	Group string `json:"group,optional" validate:"omitempty,max=20"`
	// API request method e.g. POST | API请求类型 如POST
	// min length : 3
	// max length : 4
	Method string `json:"method,optional" validate:"omitempty,uppercase,min=3,max=4"`
}

// API information response | API信息返回体
// swagger:model ApiInfoResp
type ApiInfoResp struct {
	BaseDataInfo
	// API information | API数据
	Data ApiInfo `json:"data"`
}
