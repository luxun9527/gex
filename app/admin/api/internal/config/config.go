package config

import (
	"github.com/luxun9527/gex/common/pkg/etcd"
	commongorm "github.com/luxun9527/gex/common/pkg/gorm"
	logger "github.com/luxun9527/zaplog"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	LoggerConfig     logger.Config
	EtcdConf         etcd.EtcdConfig
	GormConf         commongorm.GormConf
	LanguageEtcdConf etcd.EtcdConfig
}
