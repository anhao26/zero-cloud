package attribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/logic/attribute"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
)

// swagger:route post /attribute attribute GetAttributeById
//
// Get attribute by ID | 通过ID获取Attribute
//
// Get attribute by ID | 通过ID获取Attribute
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: AttributeInfoResp

func GetAttributeByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attribute.NewGetAttributeByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetAttributeById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
