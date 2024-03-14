package config

import (
	"github.com/luxun9527/gex/common/pkg/etcd"
	commongorm "github.com/luxun9527/gex/common/pkg/gorm"
	logger "github.com/luxun9527/zaplog"
"github.com/luxun9527/gex/common/pkg/pulsar"
"github.com/luxun9527/gex/common/proto/define"
"github.com/zeromicro/go-zero/zrpc"
)
type Config struct {
	zrpc.RpcServerConf
	AccountRpcConf   zrpc.RpcClientConf
	OrderRpcConf     zrpc.RpcClientConf
	DtmConf          zrpc.RpcClientConf
	PulsarConfig     pulsar.PulsarConfig
	LoggerConfig     logger.Config
	GormConf         commongorm.GormConf
	SymbolInfo       *define.SymbolInfo `json:",optional"`
	SnowFlakeWorkID  int64
	WsConf           zrpc.RpcClientConf
	Symbol           string
	SymbolEtcdConfig etcd.EtcdConfig
}
