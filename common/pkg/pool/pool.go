package pool

import (
	"context"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"strings"
	"sync"
	"time"
)

type RpcClients struct {
	Clients    *sync.Map
	KeyPrefix  string
	EtcdClient *clientv3.Client
	etcdConf   discov.EtcdConf
}

func NewRpcClients(etcdConf discov.EtcdConf) *RpcClients {
	cli := mustNewEtcdClient(etcdConf.Hosts)
	resp, err := cli.Get(context.Background(), etcdConf.Key, clientv3.WithPrefix())
	if err != nil {
		logx.Severef("etcd get key failed %v", err)
	}
	clis := &RpcClients{Clients: &sync.Map{}, KeyPrefix: etcdConf.Key, EtcdClient: cli}
	for _, v := range resp.Kvs {
		clis.addConn(v)
	}

	logx.Infow("init etcd client prefix", logx.Field("prefix", etcdConf.Key), logx.Field("data", resp.Kvs), logx.Field("data", clis.Clients))
	go clis.Watch()
	return clis
}
func (r *RpcClients) addConn(kv *mvccpb.KeyValue) {
	logx.Infow("add conn detail", logx.Field("key", kv.Key), logx.Field("value", kv.Value))
	//注册后的key service_order_rpc.BTC_USDT/78232932937927000的格式。
	d := strings.Split(string(kv.Key), "/")
	if len(d) != 2 {
		return
	}

	r1 := strings.Split(d[0], ".")
	if len(r1) != 2 {
		return
	}
	symbol := r1[1]

	//etcd的配置。
	etcdConfig := r.etcdConf
	etcdConfig.Key = string(kv.Key)
	rpcConfig := zrpc.RpcClientConf{
		Etcd:     etcdConfig,
		NonBlock: true,
	}
	//通过etcd创建连接有负载均衡的作用。
	conn := zrpc.MustNewClient(rpcConfig).Conn()

	r.Clients.Store(symbol, conn)
}
func (r *RpcClients) Watch() {
	rch := r.EtcdClient.Watch(context.Background(), r.KeyPrefix, clientv3.WithPrefix())
	for resp := range rch {
		logx.Sloww("before clients  changed", logx.Field("detail", resp), logx.Field("conn", r.Clients))
		for _, ev := range resp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				r.addConn(ev.Kv)
			case mvccpb.DELETE:
				r.delConn(ev.Kv)
			}
		}
		logx.Sloww("after rpc clients  changed", logx.Field("detail", resp), logx.Field("conn", r.Clients))

	}
}

func (r *RpcClients) delConn(kv *mvccpb.KeyValue) {
	d := strings.Split(string(kv.Key), "/")
	if len(d) != 2 {
		return
	}

	r1 := strings.Split(d[0], ".")
	if len(r1) != 2 {
		return
	}
	symbol := r1[1]
	_, ok := r.Clients.Load(symbol)
	if !ok {
		return
	}
	r.Clients.Delete(symbol)
}

func (r *RpcClients) GetConn(symbol string) (*grpc.ClientConn, bool) {
	v, ok := r.Clients.Load(symbol)
	if ok {
		return v.(*grpc.ClientConn), ok
	}
	return nil, ok
}

func mustNewEtcdClient(endpoints []string) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logx.Severef("init etcd client failed", logger.ErrorField(err))
	}
	return cli
}
