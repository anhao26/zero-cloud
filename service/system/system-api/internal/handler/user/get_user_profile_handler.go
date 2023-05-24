package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/user"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
)

// swagger:route get /user/profile user GetUserProfile
//
// Get user's profile | 获取用户个人信息
//
// Get user's profile | 获取用户个人信息
//
// Responses:
//  200: ProfileResp

func GetUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetUserProfile()
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
