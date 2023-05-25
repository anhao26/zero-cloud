package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/dictionary"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
)

// swagger:route post /dictionary/update dictionary UpdateDictionary
//
// Update dictionary information | 更新字典
//
// Update dictionary information | 更新字典
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateDictionaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewUpdateDictionaryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDictionary(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}