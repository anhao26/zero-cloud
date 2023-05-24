package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/dictionary"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
)

// swagger:route post /dictionary/list dictionary GetDictionaryList
//
// Get dictionary list | 获取字典列表
//
// Get dictionary list | 获取字典列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryListReq
//
// Responses:
//  200: DictionaryListResp

func GetDictionaryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewGetDictionaryListLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryList(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
