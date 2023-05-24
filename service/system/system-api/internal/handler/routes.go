// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/api"
	base "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/base"
	captcha "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/captcha"
	department "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/department"
	dictionary "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/dictionary"
	dictionarydetail "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/dictionarydetail"
	menu "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/menu"
	oauthprovider "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/oauthprovider"
	position "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/position"
	role "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/role"
	token "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/token"
	user "github.com/anhao26/zero-cloud/service/system/system-api/internal/handler/user"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: captcha.GetCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/perm",
					Handler: user.GetUserPermCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/create",
					Handler: user.CreateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/update",
					Handler: user.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/delete",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user",
					Handler: user.GetUserByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/change_password",
					Handler: user.ChangePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/profile",
					Handler: user.GetUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/profile",
					Handler: user.UpdateUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/menu/role/list",
					Handler: menu.GetMenuListByRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/list",
					Handler: menu.GetMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/create",
					Handler: menu.CreateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/update",
					Handler: menu.UpdateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/delete",
					Handler: menu.DeleteMenuHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role/create",
					Handler: role.CreateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/update",
					Handler: role.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/delete",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/list",
					Handler: role.GetRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.GetRoleByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/department/create",
					Handler: department.CreateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/update",
					Handler: department.UpdateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/delete",
					Handler: department.DeleteDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/list",
					Handler: department.GetDepartmentListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department",
					Handler: department.GetDepartmentByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/position/create",
					Handler: position.CreatePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/update",
					Handler: position.UpdatePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/delete",
					Handler: position.DeletePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/list",
					Handler: position.GetPositionListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position",
					Handler: position.GetPositionByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/token/create",
					Handler: token.CreateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/update",
					Handler: token.UpdateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/delete",
					Handler: token.DeleteTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/list",
					Handler: token.GetTokenListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token",
					Handler: token.GetTokenByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/logout",
					Handler: token.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/create",
					Handler: api.CreateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/update",
					Handler: api.UpdateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/delete",
					Handler: api.DeleteApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/list",
					Handler: api.GetApiListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api",
					Handler: api.GetApiByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/oauth/login",
				Handler: oauthprovider.OauthLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/create",
					Handler: oauthprovider.CreateOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/update",
					Handler: oauthprovider.UpdateOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/delete",
					Handler: oauthprovider.DeleteOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/list",
					Handler: oauthprovider.GetOauthProviderListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider",
					Handler: oauthprovider.GetOauthProviderByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/create",
					Handler: dictionary.CreateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/update",
					Handler: dictionary.UpdateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/delete",
					Handler: dictionary.DeleteDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/list",
					Handler: dictionary.GetDictionaryListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary",
					Handler: dictionary.GetDictionaryByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail/create",
					Handler: dictionarydetail.CreateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail/update",
					Handler: dictionarydetail.UpdateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail/delete",
					Handler: dictionarydetail.DeleteDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail/list",
					Handler: dictionarydetail.GetDictionaryDetailListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail",
					Handler: dictionarydetail.GetDictionaryDetailByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/dict/:name",
					Handler: dictionarydetail.GetDictionaryDetailByDictionaryNameHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
