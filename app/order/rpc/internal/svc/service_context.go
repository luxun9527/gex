package svc

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"github.com/dtm-labs/client/dtmgrpc/dtmgpb"
	"github.com/luxun9527/gex/app/order/rpc/internal/config"
	"github.com/luxun9527/gex/app/order/rpc/internal/dao/query"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/pkg/snowflake"
	ws "github.com/luxun9527/gpush/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"time"

	pulsarConfig "github.com/luxun9527/gex/common/pkg/pulsar"
)

type ServiceContext struct {
	Config             *config.Config
	Query              *query.Query
	DtmClient          dtmgpb.DtmClient
	MatchConsumer      pulsar.Consumer
	MatchProducer      pulsar.Producer
	SnowflakeGenerator *snowflake.Worker
	WsClient           ws.ProxyClient
}

func NewServiceContext(c *config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	writer := logger.NewZapWriter(logger.L)
	logx.SetWriter(writer)
	logx.DisableStat()
	c.Etcd.Key += "." + c.SymbolInfo.SymbolName
	target, err := c.DtmConf.BuildTarget()
	if err != nil {
		logx.Severef("init dtm client failed", logger.ErrorField(err))
		return nil
	}

	s, err := snowflake.NewWorker(c.SnowFlakeWorkID)
	if err != nil {
		logx.Severef("init snowflake fail", logger.ErrorField(err))
	}
	client, err := c.PulsarConfig.BuildClient()
	if err != nil {
		logx.Severef("init pulsar client failed", logger.ErrorField(err))
	}
	topic := pulsarConfig.Topic{
		Tenant:    pulsarConfig.PublicTenant,
		Namespace: pulsarConfig.GexNamespace,
		Topic:     pulsarConfig.MatchSourceTopic + "_" + c.SymbolInfo.SymbolName,
	}
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:           topic.BuildTopic(),
		SendTimeout:     10 * time.Second,
		DisableBatching: true,
	})
	if err != nil {
		logx.Severef("init pulsar producer failed", logger.ErrorField(err))
	}
	topic = pulsarConfig.Topic{
		Tenant:    pulsarConfig.PublicTenant,
		Namespace: pulsarConfig.GexNamespace,
		Topic:     pulsarConfig.MatchResultTopic + "_" + c.SymbolInfo.SymbolName,
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic.BuildTopic(),
		SubscriptionName: pulsarConfig.MatchResultOrderSub,
		Type:             pulsar.Shared,
	})
	if err != nil {
		logx.Severef("init pulsar consumer failed", logger.ErrorField(err))
	}

	sc := &ServiceContext{
		Config:             c,
		Query:              query.Use(c.GormConf.MustNewGormClient()),
		DtmClient:          dtmgimp.MustGetDtmClient(target),
		MatchConsumer:      consumer,
		MatchProducer:      producer,
		SnowflakeGenerator: s,
		WsClient:           ws.NewProxyClient(zrpc.MustNewClient(c.WsConf).Conn()),
	}
	return sc
}
