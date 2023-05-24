package department

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/anhao26/zero-cloud/service/system/system-api/internal/logic/department"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/types"
)

// swagger:route post /department/list department GetDepartmentList
//
// Get department list | 获取部门列表
//
// Get department list | 获取部门列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DepartmentListReq
//
// Responses:
//  200: DepartmentListResp

func GetDepartmentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DepartmentListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := department.NewGetDepartmentListLogic(r.Context(), svcCtx)
		resp, err := l.GetDepartmentList(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
