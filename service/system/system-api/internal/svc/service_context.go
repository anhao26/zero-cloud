package svc

import (
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/config"
	i18n2 "github.com/anhao26/zero-cloud/service/system/system-api/internal/i18n"
	"github.com/anhao26/zero-cloud/service/system/system-api/internal/middleware"
	"github.com/anhao26/zero-cloud/service/system/system-rpc/systemclient"
	"github.com/mojocn/base64Captcha"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
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
	Captcha   *base64Captcha.Captcha
	SystemRpc systemclient.System
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:     trans,
		Redis:     rds,
		Captcha:   captcha.MustNewRedisCaptcha(c.Captcha, rds),
		SystemRpc: systemclient.NewSystem(zrpc.NewClientIfEnable(c.SystemRpc)),
	}
}
