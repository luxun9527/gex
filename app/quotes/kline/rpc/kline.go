package main

import (
	"flag"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/config"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/logic"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/server"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/svc"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/pb"
	"github.com/luxun9527/gex/common/pkg/flagx"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/quotes/kline/rpc/etc/kline.yaml", "the config file")

func main() {
	flagx.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(&c)
	logic.InitKlineHandler(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterKlineServiceServer(grpcServer, server.NewKlineServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// mustNewServer会将全局的logx的writer重新设置
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.Infof("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
