package attribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/logic/attribute"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
)

// swagger:route post /attribute/create attribute CreateAttribute
//
// Create attribute information | 创建Attribute
//
// Create attribute information | 创建Attribute
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AttributeInfo
//
// Responses:
//  200: BaseMsgResp

func CreateAttributeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AttributeInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attribute.NewCreateAttributeLogic(r.Context(), svcCtx)
		resp, err := l.CreateAttribute(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
