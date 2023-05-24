package dictionarydetail

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/dictionarydetail"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
)

// swagger:route post /dictionary_detail/list dictionarydetail GetDictionaryDetailList
//
// Get dictionary detail list | 获取字典键值列表
//
// Get dictionary detail list | 获取字典键值列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryDetailListReq
//
// Responses:
//  200: DictionaryDetailListResp

func GetDictionaryDetailListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryDetailListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionarydetail.NewGetDictionaryDetailListLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryDetailList(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
