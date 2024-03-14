package config

import (
	"github.com/luxun9527/gex/common/pkg/etcd"
	commongorm "github.com/luxun9527/gex/common/pkg/gorm"
	logger "github.com/luxun9527/zaplog"
"github.com/luxun9527/gex/common/pkg/pulsar"
"github.com/luxun9527/gex/common/proto/define"
"github.com/zeromicro/go-zero/core/stores/redis"
"github.com/zeromicro/go-zero/zrpc"
)
type Config struct {
	zrpc.RpcServerConf
	GormConf         commongorm.GormConf
	RedisConf        redis.RedisConf
	WsConf           zrpc.RpcClientConf
	PulsarConfig     pulsar.PulsarConfig
	LoggerConfig     logger.Config
	Symbol           string
	SymbolEtcdConfig etcd.EtcdConfig
	SymbolInfo       *define.SymbolInfo `json:",optional"`
}
