package pool

import (
	"context"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"sync"
	"time"
)

type RpcClients struct {
	Clients    *sync.Map
	KeyPrefix  string
	EtcdClient *clientv3.Client
}

func NewRpcClients(keyPrefix string, endpoints []string) *RpcClients {
	cli := mustNewEtcdClient(endpoints)
	resp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		logx.Severef("etcd get key failed", logger.ErrorField(err))
	}
	clis := &RpcClients{Clients: &sync.Map{}, KeyPrefix: keyPrefix, EtcdClient: cli}
	for _, v := range resp.Kvs {
		clis.addConn(v)
	}
	logx.Infow("init etcd client prefix", logx.Field("prefix", keyPrefix), logx.Field("endpoints", endpoints), logx.Field("data", resp.Kvs), logx.Field("data", clis.Clients))
	go clis.Watch()
	return clis
}
func (r *RpcClients) addConn(kv *mvccpb.KeyValue) {
	logx.Infow("add conn detail", logx.Field("key", kv.Key), logx.Field("value", kv.Value))
	d := strings.Split(string(kv.Key), "/")
	if len(d) != 2 {
		return
	}
	r1 := strings.Split(d[0], ".")
	if len(r1) != 2 {
		return
	}
	symbol := r1[1]
	conn, err := grpc.Dial(string(kv.Value), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logx.Errorw("connect to client failed", logger.ErrorField(err), logx.Field("addr", string(kv.Value)), logx.Field("key", string(kv.Key)))
	}
	var (
		c  interface{}
		ok bool
	)

	c, ok = r.Clients.Load(symbol)
	if !ok {
		c = make([]*grpc.ClientConn, 0, 2)
	}
	cs := c.([]*grpc.ClientConn)
	c = append(cs, conn)
	r.Clients.Store(symbol, c)
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
	if len(d) != 3 {
		return
	}
	symbol := d[1]
	c, ok := r.Clients.Load(symbol)
	if !ok {
		return
	}
	cs := c.([]*grpc.ClientConn)
	newConn := make([]*grpc.ClientConn, 0, 1)
	for _, v := range cs {
		if v.Target() != string(kv.Value) {
			newConn = append(newConn, v)
		}
	}

	r.Clients.Store(symbol, newConn)
}

func (r *RpcClients) GetConn(symbol string) (*grpc.ClientConn, bool) {
	v, ok := r.Clients.Load(symbol)
	if ok {
		cs := v.([]*grpc.ClientConn)
		i := time.Now().UnixNano() % int64(len(cs))
		return cs[i], true
	}
	return nil, false
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
