package token

import (
	"context"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/predicate"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/token"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/ent/user"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/svc"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/internal/utils/dberrorhandler"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/types/system"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenListLogic) GetTokenList(in *system.TokenListReq) (*system.TokenListResp, error) {
	var tokens *ent.TokenPageList
	var err error
	if in.Username == "" && in.Uuid == "" && in.Nickname == "" && in.Email == "" {
		tokens, err = l.svcCtx.DB.Token.Query().Page(l.ctx, in.Page, in.PageSize)

		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}
	} else {
		var predicates []predicate.User

		if in.Uuid != "" {
			predicates = append(predicates, user.IDEQ(uuidx.ParseUUIDString(in.Uuid)))
		}

		if in.Username != "" {
			predicates = append(predicates, user.Username(in.Username))
		}

		if in.Email != "" {
			predicates = append(predicates, user.EmailEQ(in.Email))
		}

		if in.Nickname != "" {
			predicates = append(predicates, user.NicknameEQ(in.Nickname))
		}

		u, err := l.svcCtx.DB.User.Query().Where(predicates...).First(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		tokens, err = l.svcCtx.DB.Token.Query().Where(token.UUIDEQ(u.ID)).Page(l.ctx, in.Page, in.PageSize)

		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}
	}

	resp := &system.TokenListResp{}
	resp.Total = tokens.PageDetails.Total

	for _, v := range tokens.List {
		resp.Data = append(resp.Data, &system.TokenInfo{
			Id:        v.ID.String(),
			Uuid:      v.UUID.String(),
			Token:     v.Token,
			Status:    uint32(v.Status),
			Source:    v.Source,
			ExpiredAt: v.ExpiredAt.UnixMilli(),
			CreatedAt: v.CreatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
