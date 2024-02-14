package svc

import (
	"github.com/luxun9527/gex/app/account/rpc/accountservice"
	matchpb "github.com/luxun9527/gex/app/match/rpc/pb"
	"github.com/luxun9527/gex/app/order/api/internal/config"
	"github.com/luxun9527/gex/app/order/api/internal/middleware"
	orderpb "github.com/luxun9527/gex/app/order/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/confx"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/pkg/pool"
	"github.com/luxun9527/gex/common/proto/define"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"sync"
)

type GetOrderClientFunc func(cc grpc.ClientConnInterface) orderpb.OrderServiceClient
type GetMatchClientFunc func(cc grpc.ClientConnInterface) matchpb.MatchServiceClient

type ServiceContext struct {
	Config           config.Config
	OrderClients     *pool.RpcClients
	MatchClients     *pool.RpcClients
	GetOrderClient   GetOrderClientFunc
	GetMatchClient   GetMatchClientFunc
	Auth             rest.Middleware
	AccountRpcClient accountservice.AccountService
	Symbols          *sync.Map
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()
	errs.InitTranslatorFromEtcd(c.LanguageEtcdConf)

	var symbolConfig sync.Map
	// 从etcd中取出交易对配置。
	confx.MustLoadFromEtcd(define.EtcdSymbolPrefix, c.SymbolEtcdConfig, &symbolConfig, confx.WithCustomInitLoadFunc(func(kvs []*mvccpb.KeyValue, target any) {
		for _, v := range kvs {
			var s define.SymbolInfo
			if err := yaml.Unmarshal(v.Value, &s); err != nil {
				logx.Severef("get symbol config failed symbolInfo = %v", define.EtcdSymbolPrefix)
			}
			s.QuoteCoinPrec.Store(s.QuoteCoinPrecValue)
			s.BaseCoinPrec.Store(s.BaseCoinPrecValue)
			symbolConfig.Store(s.SymbolName, &s)
		}
	}), confx.WithCustomWatchFunc(func(evs []*clientv3.Event, target any) {
		for _, v := range evs {
			switch v.Type {
			case mvccpb.PUT: //修改或者新增
				var s define.SymbolInfo
				if err := yaml.Unmarshal(v.Kv.Value, &s); err != nil {
					logx.Errorf("get symbol config failed symbolInfo =%v", s)
				}
				s.QuoteCoinPrec.Store(s.QuoteCoinPrecValue)
				s.BaseCoinPrec.Store(s.BaseCoinPrecValue)
				symbolConfig.Store(s.SymbolName, &s)
			case mvccpb.DELETE: //删除
				var s define.SymbolInfo
				if err := yaml.Unmarshal(v.Kv.Value, &s); err != nil {
					logx.Errorf("get symbol config failed symbolInfo =%v", s)
				}
				symbolConfig.Delete(s.SymbolName)
				logx.Sloww("warn symbol config deleted")
			}

		}
	}))

	accountRpcClient := accountservice.NewAccountService(zrpc.MustNewClient(c.AccountRpcConf))
	return &ServiceContext{
		Config:           c,
		Auth:             middleware.NewAuthMiddleware(accountRpcClient).Handle,
		OrderClients:     pool.NewRpcClients(c.OrderRpcConf.Etcd),
		MatchClients:     pool.NewRpcClients(c.MatchRpcConf.Etcd),
		GetOrderClient:   orderpb.NewOrderServiceClient,
		GetMatchClient:   matchpb.NewMatchServiceClient,
		AccountRpcClient: accountRpcClient,
		Symbols:          &symbolConfig,
	}
}
