package attribute

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/logic/attribute"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/types"
)

// swagger:route post /attribute/update attribute UpdateAttribute
//
// Update attribute information | 更新Attribute
//
// Update attribute information | 更新Attribute
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AttributeInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateAttributeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AttributeInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attribute.NewUpdateAttributeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAttribute(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
