package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/menu"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
)

// swagger:route post /menu/create menu CreateMenu
//
// Create menu information | 创建菜单
//
// Create menu information | 创建菜单
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MenuPlainInfo
//
// Responses:
//  200: BaseMsgResp

func CreateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuPlainInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewCreateMenuLogic(r.Context(), svcCtx)
		resp, err := l.CreateMenu(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
