package svc

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Skyenought/tiktokbackend/cmd/user/rpc/userservice"
	"github.com/Skyenought/tiktokbackend/cmd/video/api/internal/config"
	"github.com/Skyenought/tiktokbackend/cmd/video/rpc/videoservice"
	"github.com/minio/minio-go"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	MinioClient *minio.Client
	VideoRpc    videoservice.VideoService
	UserRpc     userservice.UserService
	TencentOSS  *cos.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		MinioClient: newMinioClient(c),
		VideoRpc:    videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpcConf)),
		UserRpc:     userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
		TencentOSS:  newTencentOssClient(c),
	}
}

func newMinioClient(c config.Config) *minio.Client {
	oss := c.Minio
	client, err := minio.New(oss.EndPoint, oss.AccessKeyID, oss.SecretAccessKey, oss.UseSSL)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return client
}

func newTencentOssClient(c config.Config) (client *cos.Client) {
	bucketURL, _ := url.Parse(c.Oss.BucketURL)
	baseURL := &cos.BaseURL{BucketURL: bucketURL}
	client = cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("SECRETID"),
			SecretKey: os.Getenv("SECRETKEY"),
		},
	})
	return
}
