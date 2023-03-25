package main

import (
	"flag"
	"fmt"

	"github.com/Skyenought/tiktokbackend/pkg/middleware"

	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/config"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/handler"
	"github.com/Skyenought/tiktokbackend/cmd/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.Use(middleware.NewCorsMiddleware().Handle)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
