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
	KlineRpcConf     zrpc.RpcClientConf
	MatchRpcConf     zrpc.RpcClientConf
	SymbolList       []*define.SymbolInfo
	LoggerConfig     logger.Config
	LanguageEtcdConf etcd.EtcdConfig
}
