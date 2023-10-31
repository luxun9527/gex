package main

import (
	"flag"
	"fmt"
	"github.com/luxun9527/gex/common/pkg/validatorx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/luxun9527/gex/app/admin/api/internal/config"
	"github.com/luxun9527/gex/app/admin/api/internal/handler"
	"github.com/luxun9527/gex/app/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/admin/api/etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	httpx.SetValidator(validatorx.NewValidator())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
