package svc

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/luxun9527/gex/app/account/rpc/internal/config"
	"github.com/luxun9527/gex/app/account/rpc/internal/dao/query"
	"github.com/luxun9527/gex/common/pkg/confx"
	"github.com/luxun9527/gex/common/pkg/logger"
	pulsarConfig "github.com/luxun9527/gex/common/pkg/pulsar"
	"github.com/luxun9527/gex/common/proto/define"
	"github.com/luxun9527/gex/common/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
	"sync"
)

type ServiceContext struct {
	Config            config.Config
	Query             *query.Query
	MatchConsumerList []pulsar.Consumer
	JwtClient         *utils.JWT
	RedisClient       *redis.Redis
	Coins             *sync.Map
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()
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

	var coinConfig sync.Map
	// 从etcd中取出币种配置。。
	confx.MustLoadFromEtcd(define.EtcdCoinPrefix, c.SymbolEtcdConfig, &symbolConfig, confx.WithCustomInitLoadFunc(func(kvs []*mvccpb.KeyValue, target any) {
		for _, v := range kvs {
			var s define.CoinInfo
			if err := yaml.Unmarshal(v.Value, &s); err != nil {
				logx.Severef("get symbol config failed symbolInfo = %v", string(v.Value))
			}
			coinConfig.Store(s.CoinName, &s)
		}
	}), confx.WithCustomWatchFunc(func(evs []*clientv3.Event, target any) {
		for _, v := range evs {
			switch v.Type {
			case mvccpb.PUT: //修改或者新增
				var s define.CoinInfo
				if err := yaml.Unmarshal(v.Kv.Value, &s); err != nil {
					logx.Errorf("get symbol config failed symbolInfo =%v", s)
				}
				coinConfig.Store(s.CoinName, &s)
			case mvccpb.DELETE: //删除
				var s define.CoinInfo
				if err := yaml.Unmarshal(v.Kv.Value, &s); err != nil {
					logx.Errorf("get symbol config failed symbolInfo =%v", s)
				}
				coinConfig.Delete(s.CoinName)
				logx.Sloww("warn coin config deleted")
			}

		}
	}))

	client, err := c.PulsarConfig.BuildClient()
	if err != nil {
		logx.Severef("init pulsar client failed %v", err)
	}
	consumers := make([]pulsar.Consumer, 0, 10)
	symbolConfig.Range(func(key, value any) bool {
		symbolInfo := value.(*define.SymbolInfo)
		topic := pulsarConfig.Topic{
			Tenant:    pulsarConfig.PublicTenant,
			Namespace: pulsarConfig.GexNamespace,
			Topic:     pulsarConfig.MatchResultTopic + "_" + symbolInfo.SymbolName,
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
		return true
	})
	q := query.Use(c.GormConf.MustNewGormClient())
	sc := &ServiceContext{
		Config:            c,
		Query:             q,
		MatchConsumerList: consumers,
		JwtClient:         utils.NewJWT(),
		RedisClient:       redis.MustNewRedis(c.RedisConf),
		Coins:             &coinConfig,
	}
	return sc
}
