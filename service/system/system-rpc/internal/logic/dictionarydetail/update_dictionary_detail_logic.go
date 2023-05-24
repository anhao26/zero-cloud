package dictionarydetail

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryDetailLogic {
	return &UpdateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryDetailLogic) UpdateDictionaryDetail(in *system.DictionaryDetailInfo) (*system.BaseResp, error) {
    err := l.svcCtx.DB.DictionaryDetail.UpdateOneID(in.Id).
			SetNotEmptyStatus(uint8(in.Status)).
			SetNotEmptySort(in.Sort).
			SetNotEmptyTitle(in.Title).
			SetNotEmptyKey(in.Key).
			SetNotEmptyValue(in.Value).
			SetNotEmptyDictionaryID(in.DictionaryId).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseResp{Msg: i18n.CreateSuccess }, nil
}
