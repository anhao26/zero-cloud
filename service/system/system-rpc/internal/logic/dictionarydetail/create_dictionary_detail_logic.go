package dictionarydetail

import (
	"context"


	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
    "github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryDetailLogic {
	return &CreateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDictionaryDetailLogic) CreateDictionaryDetail(in *system.DictionaryDetailInfo) (*system.BaseIDResp, error) {
    result, err := l.svcCtx.DB.DictionaryDetail.Create().
			SetStatus(uint8(in.Status)).
			SetSort(in.Sort).
			SetTitle(in.Title).
			SetKey(in.Key).
			SetValue(in.Value).
			SetDictionaryID(in.DictionaryId).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &system.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
