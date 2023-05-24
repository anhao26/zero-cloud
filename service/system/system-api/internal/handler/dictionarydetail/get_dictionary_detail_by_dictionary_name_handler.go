package dictionarydetail

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/dictionarydetail"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
)

// swagger:route get /dict/{name} dictionarydetail GetDictionaryDetailByDictionaryName
//
// Get dictionary detail by dictionary name | 通过字典名称获取字典内容
//
// Get dictionary detail by dictionary name | 通过字典名称获取字典内容
//
// Responses:
//  200: DictionaryDetailListResp

func GetDictionaryDetailByDictionaryNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryNameReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionarydetail.NewGetDictionaryDetailByDictionaryNameLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryDetailByDictionaryName(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
