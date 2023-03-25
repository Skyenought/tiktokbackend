package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf    zrpc.RpcClientConf
	CommentRpcConf zrpc.RpcClientConf
	VideoRpcConf   zrpc.RpcClientConf
	Auth           struct {
		Secret string
		Expire int64
	}
}
