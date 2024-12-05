package svc

import (
	"github.com/luxun9527/gex/app/match/rpc/matchservice"
	matchpb "github.com/luxun9527/gex/app/match/rpc/pb"
	"github.com/luxun9527/gex/app/quotes/api/internal/config"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/klineservice"
	klinepb "github.com/luxun9527/gex/app/quotes/kline/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/etcd"
	logger "github.com/luxun9527/zaplog"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config       config.Config
	KlineClients klinepb.KlineServiceClient
	MatchClients matchpb.MatchServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitZapLogger(&c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.GetZapLogger()))
	errs.InitTranslatorFromEtcd(c.LanguageEtcdConf)
	//自定义负载均衡策略
	var clientOpts []zrpc.ClientOption
	serviceConfig := grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"symbol_lb"}`)

	//自定义resolver
	etcdConfig := etcd.EtcdConfig{Endpoints: c.MatchRpcConf.Etcd.Hosts}
	etcdCli := etcdConfig.MustNewEtcdClient()
	etcdResolver, err := resolver.NewBuilder(etcdCli)
	if err != nil {
		logx.Severef("NewBuilder error: %v", err)
	}
	r := grpc.WithResolvers(etcdResolver)

	clientOpts = append(clientOpts, zrpc.WithDialOption(r), zrpc.WithDialOption(serviceConfig))

	sc := &ServiceContext{
		Config:       c,
		KlineClients: klineservice.NewKlineService(zrpc.MustNewClient(c.KlineRpcConf, clientOpts...)),
		MatchClients: matchservice.NewMatchService(zrpc.MustNewClient(c.MatchRpcConf, clientOpts...)),
	}

	return sc
}
