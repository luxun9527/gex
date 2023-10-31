package config

import (
	"github.com/luxun9527/gex/common/pkg/etcd"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/proto/define"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	LoggerConfig     logger.Config
	OrderRpcConf     zrpc.RpcClientConf
	MatchRpcConf     zrpc.RpcClientConf
	AccountRpcConf   zrpc.RpcClientConf
	SymbolListConf   define.SymbolCoinConfig[string, *define.SymbolInfo]
	LanguageEtcdConf etcd.EtcdConfig
}
