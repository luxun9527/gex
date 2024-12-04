package etcd

import (
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdConfig struct {
	Endpoints []string
}

func (c *EtcdConfig) NewEtcdClient() (*clientv3.Client, error) {
	config := clientv3.Config{Endpoints: c.Endpoints, DialTimeout: time.Second * time.Duration(5)}
	return clientv3.New(config)
}
func (c *EtcdConfig) MustNewEtcdClient() *clientv3.Client {
	client, err := c.NewEtcdClient()
	if err != nil {
		logx.Severef("etcd client init failed, err: %v", err)
	}
	return client
}
