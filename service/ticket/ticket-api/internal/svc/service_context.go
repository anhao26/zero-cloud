package svc

import (
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/config"
	i18n2 "github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/i18n"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-api/internal/middleware"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ticketclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	Trans     *i18n.Translator
	TicketRpc ticketclient.Ticket
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:     trans,
		TicketRpc: ticketclient.NewTicket(zrpc.NewClientIfEnable(c.TicketRpc)),
	}
}
