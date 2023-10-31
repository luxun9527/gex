package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdConfig struct {
	Endpoints   []string
	DialTimeout int
}

func (c EtcdConfig) NewEtcdClient() (*clientv3.Client, error) {

	config := clientv3.Config{Endpoints: c.Endpoints, DialTimeout: time.Second * time.Duration(c.DialTimeout)}
	return clientv3.New(config)
}
