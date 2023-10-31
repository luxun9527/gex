package svc

import (
	"encoding/json"
	"github.com/luxun9527/gex/app/account/api/internal/config"
	"github.com/luxun9527/gex/app/account/api/internal/middleware"
	"github.com/luxun9527/gex/app/account/rpc/accountservice"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	AccountRpcClient accountservice.AccountService
	Auth             rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()
	sc := &ServiceContext{
		Config:           c,
		Auth:             middleware.NewAuthMiddleware(accountservice.NewAccountService(zrpc.MustNewClient(c.AccountRpcConf))).Handle,
		AccountRpcClient: accountservice.NewAccountService(zrpc.MustNewClient(c.AccountRpcConf)),
	}
	d, _ := json.Marshal(c.LanguageEtcdConf)
	errs.InitTranslatorFromEtcd(string(d))

	return sc
}
