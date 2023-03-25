package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		Secret string
		Expire int64
	}
	VideoRpcConf zrpc.RpcClientConf
	UserRpcConf  zrpc.RpcClientConf
	Minio        struct {
		EndPoint        string
		BucketName      string
		AccessKeyID     string
		SecretAccessKey string
		UploadPath      string
		UseSSL          bool
	}
	Oss struct {
		BucketURL string
	}
}
