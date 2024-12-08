package etcd

import (
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc/attributes"
)

type EtcdRegisterConf struct {
	EtcdConf EtcdConfig
	Key      string
	Value    string                 `json:",optional"`
	Port     int32                  `json:",optional"`
	MataData *attributes.Attributes `json:",optional"`
}

func Register(conf EtcdRegisterConf) {
	go func() {
		cli, err := conf.EtcdConf.NewEtcdClient()
		if err != nil {
			logx.Severef("etcd new client err: %v", err)
		}
		manager, err := endpoints.NewManager(cli, conf.Key)
		if err != nil {
			logx.Severef("etcd new manager err: %v", err)
		}
		//设置租约时间
		resp, err := cli.Grant(context.Background(), 30)
		if err != nil {
			logx.Severef("etcd grant err: %v", err)
		}
		if conf.Value == "" {
			conf.Value = netx.InternalIp() + ":" + cast.ToString(conf.Port)
		}
		if err := manager.AddEndpoint(context.Background(), conf.Key+"/"+cast.ToString(int64(resp.ID)), endpoints.Endpoint{Addr: conf.Value, Metadata: conf.MataData}, clientv3.WithLease(resp.ID)); err != nil {
			logx.Severef("etcd add endpoint err: %v", err)
		}
		c, err := cli.KeepAlive(context.Background(), resp.ID)
		if err != nil {
			logx.Severef("etcd keepalive err: %v", err)
		}
		logx.Infof("etcd register success,key: %v,value: %v", conf.Key, conf.Value)
		for {
			select {
			case _, ok := <-c:
				if !ok {
					logx.Errorf("etcd keepalive failed,please check etcd key %v existed", conf.Key)
					return
				}
			}
		}

	}()

}
