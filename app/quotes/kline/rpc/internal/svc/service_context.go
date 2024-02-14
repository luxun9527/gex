package svc

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/config"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/dao/query"
	"github.com/luxun9527/gex/common/pkg/logger"
	pulsarConfig "github.com/luxun9527/gex/common/pkg/pulsar"
	"github.com/luxun9527/gex/common/proto/define"
	gpushPb "github.com/luxun9527/gpush/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"time"
)

type ServiceContext struct {
	Config        *config.Config
	Query         *query.Query
	RedisClient   *redis.Redis
	MatchConsumer pulsar.Consumer
	WsClient      gpushPb.ProxyClient
	SymbolInfo    define.SymbolInfo
}

func NewServiceContext(c *config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()
	var symbolInfo define.SymbolInfo

	c.Etcd.Key += "." + c.Symbol

	//todo 从etcd中取交易对的配置
	define.InitSymbolConfig(c.Symbol, c.EtcdConfig, &symbolInfo)
	go func() {
		for {
			time.Sleep(time.Second)
			log.Println(symbolInfo)
		}
	}()
	client, err := c.PulsarConfig.BuildClient()
	if err != nil {
		logx.Severef("init pulsar client failed %v", err)
	}
	topic := pulsarConfig.Topic{
		Tenant:    pulsarConfig.PublicTenant,
		Namespace: pulsarConfig.GexNamespace,
		Topic:     pulsarConfig.MatchResultTopic + "_" + c.Symbol,
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic.BuildTopic(),
		SubscriptionName: pulsarConfig.MatchResultKlineSub,
		Type:             pulsar.Shared,
	})
	if err != nil {
		logx.Severef("init pulsar consumer failed %v", err)
	}
	sc := &ServiceContext{
		Config:        c,
		Query:         query.Use(c.GormConf.MustNewGormClient()),
		RedisClient:   redis.MustNewRedis(c.RedisConf),
		MatchConsumer: consumer,
		WsClient:      gpushPb.NewProxyClient(zrpc.MustNewClient(c.WsConf).Conn()),
	}
	return sc
}
