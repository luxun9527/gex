package svc

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/luxun9527/gex/app/match/rpc/internal/config"
	"github.com/luxun9527/gex/app/match/rpc/internal/dao/query"
	"github.com/luxun9527/gex/app/match/rpc/internal/engine"
	"github.com/luxun9527/gex/app/order/rpc/orderservice"
	pulsarConfig "github.com/luxun9527/gex/common/pkg/pulsar"
	"github.com/luxun9527/gex/common/proto/define"
	ws "github.com/luxun9527/gpush/proto"
	logger "github.com/luxun9527/zaplog"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	MatchConsumer pulsar.Consumer
	Config        *config.Config
	MatchEngine   *engine.MatchEngine
	OrderClient   orderservice.OrderService
	Query         *query.Query
	RedisClient   *redis.Redis
}

func NewServiceContext(c *config.Config) *ServiceContext {
	logger.InitZapLogger(&c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.GetZapLogger()))
	logx.DisableStat()

	var symbolInfo define.SymbolInfo
	define.InitSymbolConfig(define.EtcdSymbolPrefix+c.Symbol, c.SymbolEtcdConfig, &symbolInfo)
	c.SymbolInfo = &symbolInfo
	c.Etcd.Key += "." + c.Symbol
	client, err := c.PulsarConfig.BuildClient()
	if err != nil {
		logx.Severef("init pulsar client failed err %v", err)
	}
	topic := pulsarConfig.Topic{
		Tenant:    pulsarConfig.PublicTenant,
		Namespace: pulsarConfig.GexNamespace,
		Topic:     pulsarConfig.MatchResultTopic + "_" + c.Symbol,
	}
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:           topic.BuildTopic(),
		SendTimeout:     10 * time.Second,
		DisableBatching: true,
	})
	if err != nil {
		logx.Severef("init pulsar producer failed %v", logger.ErrorField(err))
	}
	topic = pulsarConfig.Topic{
		Tenant:    pulsarConfig.PublicTenant,
		Namespace: pulsarConfig.GexNamespace,
		Topic:     pulsarConfig.MatchSourceTopic + "_" + c.Symbol,
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic.BuildTopic(),
		SubscriptionName: pulsarConfig.MatchSourceSub,
		Type:             pulsar.Shared,
	})
	if err != nil {
		logx.Severef("init pulsar consumer failed %v", logger.ErrorField(err))
	}
	sc := &ServiceContext{
		MatchConsumer: consumer,
		Config:        c,
		OrderClient:   orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRpcConf)),
		MatchEngine:   engine.NewMatchEngine(c, producer, ws.NewProxyClient(zrpc.MustNewClient(c.WsConf).Conn())),
		Query:         query.Use(c.GormConf.MustNewGormClient()),
		RedisClient:   redis.MustNewRedis(c.RedisConf),
	}
	return sc
}
