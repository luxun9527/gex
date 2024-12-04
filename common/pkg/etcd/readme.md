## 需求

有部分grpc服务需要每个交易对至少启动一个实例，如k线服务，撮合服务。但是api服务是不区分交易对的。如何让api服务连接上区分交易对的服务。



![image-20241204233701793](C:\Users\dengyongcai\AppData\Roaming\Typora\typora-user-images\image-20241204233701793.png)

这就不能使用go-zero提供的 基于etcd 单个key负载均衡的方式。

## 方案1

比较容易想到到的方式

将grpc服务注册到etcd的key设置为三级如：

key:klineRpc/BTC_USDT/1 value: 192.168.2.159:9999

klineRpc/IKUN_USDT/1 value: 192.168.2.159:9998

监听以klineRpc开头的etcd的key的变化，当有新的交易对上线，及时建立连接。当实例下线的时候及时删除。自己手动维护一个

map[(交易对)] [] string(服务地址) 的结构，当请求指定交易对的数据。从map中获取连接。



## 方案2

使用grpc 可以自定义resolver和负载均衡的方式来实现，根据交易对自动选择连接

![image-20241204234225989](C:\Users\dengyongcai\AppData\Roaming\Typora\typora-user-images\image-20241204234225989.png)

```go
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
       resp, err := cli.Grant(context.Background(), 5)
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
```