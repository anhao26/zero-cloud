package oauthprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/oauthprovider"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
)

// swagger:route post /oauth_provider/update oauthprovider UpdateOauthProvider
//
// Update oauth provider information | 更新第三方
//
// Update oauth provider information | 更新第三方
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: OauthProviderInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateOauthProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OauthProviderInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oauthprovider.NewUpdateOauthProviderLogic(r.Context(), svcCtx)
		resp, err := l.UpdateOauthProvider(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
