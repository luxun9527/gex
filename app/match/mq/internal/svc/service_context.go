package svc

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/luxun9527/gex/app/match/mq/internal/config"
	"github.com/luxun9527/gex/app/match/mq/internal/dao/model"
	"github.com/luxun9527/gex/app/match/mq/internal/dao/query"
	"github.com/luxun9527/gex/app/order/rpc/orderservice"
	"github.com/luxun9527/gex/common/pkg/logger"
	pulsarConfig "github.com/luxun9527/gex/common/pkg/pulsar"
	gpushPb "github.com/luxun9527/gpush/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	MatchConsumer pulsar.Consumer
	Config        config.Config
	OrderClient   orderservice.OrderService
	Query         *query.Query
	RedisClient   *redis.Redis
	WsClient      gpushPb.ProxyClient
	MatchDataChan chan *model.MatchData
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()
	client, err := c.PulsarConfig.BuildClient()
	if err != nil {
		logx.Severef("init pulsar client failed err ", err)
	}
	topic := pulsarConfig.Topic{
		Tenant:    pulsarConfig.PublicTenant,
		Namespace: pulsarConfig.GexNamespace,
		Topic:     pulsarConfig.MatchResultTopic + "_" + c.SymbolInfo.SymbolName,
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic.BuildTopic(),
		SubscriptionName: pulsarConfig.MatchResultMatchSub,
		Type:             pulsar.Shared,
	})
	if err != nil {
		logx.Severef("init pulsar consumer failed", logger.ErrorField(err))
	}
	sc := &ServiceContext{
		MatchConsumer: consumer,
		Config:        c,
		Query:         query.Use(c.GormConf.MustNewGormClient()),
		WsClient:      gpushPb.NewProxyClient(zrpc.MustNewClient(c.WsConf).Conn()),
		RedisClient:   redis.MustNewRedis(c.RedisConf),
		MatchDataChan: make(chan *model.MatchData),
	}
	return sc
}
