package config

import (
	commongorm "github.com/luxun9527/gex/common/pkg/gorm"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/pkg/pulsar"
	"github.com/luxun9527/gex/common/proto/define"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PulsarConfig pulsar.PulsarConfig
	LoggerConfig logger.Config
	GormConf     commongorm.GormConf
	WsConf       zrpc.RpcClientConf
	SymbolInfo   define.SymbolInfo
	OrderRpcConf zrpc.RpcClientConf
	RedisConf    redis.RedisConf
}
