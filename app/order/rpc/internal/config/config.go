package config

import (
	commongorm "github.com/luxun9527/gex/common/pkg/gorm"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/pkg/pulsar"
	"github.com/luxun9527/gex/common/proto/define"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AccountRpcConf  zrpc.RpcClientConf
	OrderRpcConf    zrpc.RpcClientConf
	DtmConf         zrpc.RpcClientConf
	PulsarConfig    pulsar.PulsarConfig
	LoggerConfig    logger.Config
	GormConf        commongorm.GormConf
	SymbolInfo      define.SymbolInfo
	SnowFlakeWorkID int64
	WsConf          zrpc.RpcClientConf
}
