package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/role"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
)

// swagger:route post /role/list role GetRoleList
//
// Get role list | 获取角色列表
//
// Get role list | 获取角色列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: RoleListReq
//
// Responses:
//  200: RoleListResp

func GetRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewGetRoleListLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleList(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
