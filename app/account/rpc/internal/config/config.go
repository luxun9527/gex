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
	GormConf       commongorm.GormConf
	LoggerConfig   logger.Config
	PulsarConfig   pulsar.PulsarConfig
	RedisConf      redis.RedisConf
	SymbolListConf define.SymbolCoinConfig[string, *define.SymbolInfo]
	CoinListConf   define.SymbolCoinConfig[string, *define.CoinInfo]
}
