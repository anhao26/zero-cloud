package captcha

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/captcha"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
)

// swagger:route get /captcha captcha GetCaptcha
//
// Get captcha | 获取验证码
//
// Get captcha | 获取验证码
//
// Responses:
//  200: CaptchaResp

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := captcha.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha()
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
