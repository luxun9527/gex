package svc

import (
	matchpb "github.com/luxun9527/gex/app/match/rpc/pb"
	"github.com/luxun9527/gex/app/quotes/api/internal/config"
	klinepb "github.com/luxun9527/gex/app/quotes/kline/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	logger "github.com/luxun9527/zaplog"
"github.com/luxun9527/gex/common/pkg/pool"
"github.com/zeromicro/go-zero/core/logx"
"google.golang.org/grpc"
)
type (
	GetKlineClientFunc func(cc grpc.ClientConnInterface) klinepb.KlineServiceClient
	GetMatchClientFunc func(cc grpc.ClientConnInterface) matchpb.MatchServiceClient
)

type ServiceContext struct {
	Config         config.Config
	KlineClients   *pool.RpcClients
	MatchClients   *pool.RpcClients
	GetKlineClient GetKlineClientFunc
	GetMatchClient GetMatchClientFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	errs.InitTranslatorFromEtcd(c.LanguageEtcdConf)
	sc := &ServiceContext{
		Config:         c,
		KlineClients:   pool.NewRpcClients(c.KlineRpcConf.Etcd),
		MatchClients:   pool.NewRpcClients(c.MatchRpcConf.Etcd),
		GetKlineClient: klinepb.NewKlineServiceClient,
		GetMatchClient: matchpb.NewMatchServiceClient,
	}

	return sc
}
