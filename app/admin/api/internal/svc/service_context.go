package svc

import (
	"github.com/luxun9527/gex/app/admin/api/internal/config"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceContext struct {
	Config  config.Config
	EtcdCli *clientv3.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()
	cli, err := c.EtcdConf.NewEtcdClient()
	if err != nil {
		logx.Severef("init etcd client failed %v", err)
	}
	return &ServiceContext{
		Config:  c,
		EtcdCli: cli,
	}
}
