package svc

import (
	"encoding/json"
	matchpb "github.com/luxun9527/gex/app/match/rpc/pb"
	"github.com/luxun9527/gex/app/quotes/api/internal/config"
	klinepb "github.com/luxun9527/gex/app/quotes/kline/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
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
	d, _ := json.Marshal(c.LanguageEtcdConf)
	errs.InitTranslatorFromEtcd(string(d))
	sc := &ServiceContext{
		Config:         c,
		KlineClients:   pool.NewRpcClients(c.KlineRpcConf.Etcd.Key, c.KlineRpcConf.Etcd.Hosts),
		MatchClients:   pool.NewRpcClients(c.MatchRpcConf.Etcd.Key, c.MatchRpcConf.Etcd.Hosts),
		GetKlineClient: klinepb.NewKlineServiceClient,
		GetMatchClient: matchpb.NewMatchServiceClient,
	}

	return sc
}
