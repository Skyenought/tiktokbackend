package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DSN string
	}
	Minio struct {
		EndPoint   string
		UploadPath string
	}
	Oss struct {
		BucketURL  string
		UpdatePath string
	}
}
