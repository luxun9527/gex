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
	PulsarConfig pulsar.PulsarConfig
	LoggerConfig logger.Config
	SymbolInfo   define.SymbolInfo
	GormConf     commongorm.GormConf
	WsConf       zrpc.RpcClientConf
	RedisConf    redis.RedisConf
}
