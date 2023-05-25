package attribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/logic/attribute"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
)

// swagger:route post /attribute/delete attribute DeleteAttribute
//
// Delete attribute information | 删除Attribute信息
//
// Delete attribute information | 删除Attribute信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteAttributeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attribute.NewDeleteAttributeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteAttribute(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
