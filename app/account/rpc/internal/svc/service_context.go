package svc

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/luxun9527/gex/app/account/rpc/internal/config"
	"github.com/luxun9527/gex/app/account/rpc/internal/dao/query"
	"github.com/luxun9527/gex/common/pkg/logger"
	pulsarConfig "github.com/luxun9527/gex/common/pkg/pulsar"
	"github.com/luxun9527/gex/common/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"sync"
)

type ServiceContext struct {
	Config            config.Config
	Query             *query.Query
	MatchConsumerList []pulsar.Consumer
	JwtClient         *utils.JWT
	RedisClient       *redis.Redis
	Symbols           *sync.Map
	Coins             *sync.Map
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()

	client, err := c.PulsarConfig.BuildClient()
	if err != nil {
		logx.Severef("init pulsar client failed %v", err)
	}
	consumers := make([]pulsar.Consumer, 0, len(c.SymbolListConf))
	for _, v := range c.SymbolListConf {
		topic := pulsarConfig.Topic{
			Tenant:    pulsarConfig.PublicTenant,
			Namespace: pulsarConfig.GexNamespace,
			Topic:     pulsarConfig.MatchResultTopic + "_" + v.SymbolName,
		}
		consumer, err := client.Subscribe(pulsar.ConsumerOptions{
			Topic:            topic.BuildTopic(),
			SubscriptionName: pulsarConfig.MatchResultAccountSub,
			Type:             pulsar.Shared,
		})
		if err != nil {
			logx.Severef("init pulsar consumer failed %v", err)
		}
		consumers = append(consumers, consumer)
	}
	q := query.Use(c.GormConf.MustNewGormClient())
	sc := &ServiceContext{
		Config:            c,
		Query:             q,
		MatchConsumerList: consumers,
		JwtClient:         utils.NewJWT(),
		RedisClient:       redis.MustNewRedis(c.RedisConf),
		Symbols:           c.SymbolListConf.CastToSyncMap(),
		Coins:             c.CoinListConf.CastToSyncMap(),
	}
	return sc
}
