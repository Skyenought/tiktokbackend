package main

import (
	"flag"
	"fmt"

	"github.com/Skyenought/tiktokbackend/pkg/interceptor/rpcserver"

	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/config"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/server"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/internal/svc"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videorpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/video.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		videorpc.RegisterVideoServiceServer(grpcServer, server.NewVideoServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
