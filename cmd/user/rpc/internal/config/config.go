package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	VideoRpcConf zrpc.RpcClientConf
	Mysql        struct {
		DSN string
	}
}
