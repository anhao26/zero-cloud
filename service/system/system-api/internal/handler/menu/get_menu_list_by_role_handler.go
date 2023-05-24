package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/menu"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
)

// swagger:route get /menu/role/list menu GetMenuListByRole
//
// Get menu list by role | 获取菜单列表
//
// Get menu list by role | 获取菜单列表
//
// Responses:
//  200: MenuListResp

func GetMenuListByRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetMenuListByRoleLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuListByRole()
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
