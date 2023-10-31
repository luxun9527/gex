package main

import (
	"flag"
	"github.com/luxun9527/gex/app/account/api/internal/config"
	"github.com/luxun9527/gex/app/account/api/internal/handler"
	"github.com/luxun9527/gex/app/account/api/internal/svc"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/account/api/etc/account.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
