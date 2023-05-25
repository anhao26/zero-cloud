package attribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/logic/attribute"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
)

// swagger:route post /attribute/list attribute GetAttributeList
//
// Get attribute list | 获取Attribute列表
//
// Get attribute list | 获取Attribute列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AttributeListReq
//
// Responses:
//  200: AttributeListResp

func GetAttributeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AttributeListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attribute.NewGetAttributeListLogic(r.Context(), svcCtx)
		resp, err := l.GetAttributeList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
