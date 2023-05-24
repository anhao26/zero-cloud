package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/user"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
)

// swagger:route get /user/logout user Logout
//
// Log out | 退出登陆
//
// Log out | 退出登陆
//
// Responses:
//  200: BaseMsgResp

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
