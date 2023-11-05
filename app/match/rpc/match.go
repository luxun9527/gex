package main

import (
	"flag"
	"github.com/luxun9527/gex/app/match/rpc/internal/bootstrap"
	"github.com/luxun9527/gex/app/match/rpc/internal/config"
	"github.com/luxun9527/gex/app/match/rpc/internal/server"
	"github.com/luxun9527/gex/app/match/rpc/internal/svc"
	"github.com/luxun9527/gex/app/match/rpc/pb"
	"github.com/luxun9527/gex/common/pkg/confx"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	symbol     = flag.String("s", "BTC_USDT", "symbol 交易对")
	etcdConfig = flag.String("e", `{"Endpoints":["etcd:2379"],"DialTimeout":5}`, "symbol 交易对")
)

func main() {
	flag.Parse()

	var c config.Config
	//初始化配置
	confx.MustLoadFromEtcd(confx.Match.BuildKey(*symbol), *etcdConfig, &c, confx.WithDefaultInitLoadFunc())
	ctx := svc.NewServiceContext(&c)
	//初始化
	bootstrap.Start(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMatchServiceServer(grpcServer, server.NewMatchServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.Infof("Starting rpc server at %s...", c.ListenOn)
	s.Start()
}
