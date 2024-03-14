package config

import (
	"github.com/luxun9527/gex/common/pkg/etcd"
	commongorm "github.com/luxun9527/gex/common/pkg/gorm"
	logger "github.com/luxun9527/zaplog"
"github.com/luxun9527/gex/common/pkg/pulsar"
"github.com/zeromicro/go-zero/core/stores/redis"
"github.com/zeromicro/go-zero/zrpc"
)
type Config struct {
	PulsarConfig     pulsar.PulsarConfig
	LoggerConfig     logger.Config
	GormConf         commongorm.GormConf
	WsConf           zrpc.RpcClientConf
	RedisConf        redis.RedisConf
	Symbol           string
	SymbolEtcdConfig etcd.EtcdConfig
}
