# tiktokbackend

为极简版抖音后端, 整体使用 golang 开发, 项目 layout 使用 go-zero 生成, 业务之间使用 grpc 进行沟通
前端使用 Android app 进行交互

## 项目结构

```
.            
├── Makefile     // 项目构建脚本
├── README.md    // 项目说明
├── cmd          // 项目主要代码
│   ├── comment  // 评论服务
│   ├── user     // 用户服务
│   └── video    // 视频服务
├── data         // deploy 产生的项目数据
├── deploy       // 项目部署脚本
│   ├── gen-dockerfile.sh // 生成 dockerfile 脚本
│   ├── gen-k8s-file.sh   // 生成 k8s 配置文件脚本
│   └── nginx             // nginx 配置目录
├── docker-compose.yaml   // docker-compose 配置文件            
├── etc                   // 项目配置文件
│   ├── dev               // 开发环境配置文件
│   └── prod              // 生产环境配置文件
├── go.mod
├── go.sum
├── modd.conf             // 项目热加载配置文件, 详情见 https://github.com/cortesi/modd
├── pkg                   // 项目公共代码
    ├── api               // go-zero api 文件  
    ├── bcryptx           // 密码加密
    ├── errno             // 错误码以及信息
    ├── idl               // 项目中使用的 idl 文件
    ├── interceptor       // grpc 拦截器
    ├── jwtx              // jwt 加密
    ├── middleware        // http middieware 
    ├── result            // http 返回结果
    └── tpl               // go-zero 生成代码模板
```

## 接口文档
[地址](https://www.apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50717106)

## 运行项目

1. 填写 cmd 下各个服务的配置文件, 将 mysql, etcd 的地址填写正确. 本项目使用 gorm AutoMigrate 进行数据库迁移, 所以不需要手动导入 SQL
```yaml
# cmd/comment/etc/comment-api.yaml

Name: video.api
Host: 0.0.0.0
Port: 9004
Mode: dev
MaxBytes: 31457280

Log:
  ServiceName: video.api
  Encoding: plain
# 用于 jwt 的生成和解析
Auth:
  Secret: # 自行配置
  Expire:

UserRpcConf:
  Etcd:
    Hosts:
      - 0.0.0.0:2379
    Key: user.rpc
VideoRpcConf:
  Etcd:
    Hosts:
      - 0.0.0.0:2379
    Key: video.rpc

Minio:
  EndPoint:
  BucketName: video-storage
  AccessKeyID: minio
  SecretAccessKey: minio
  UploadPath:
  UseSSL: false

# 腾讯云 OSS
Oss:
  BucketURL:
  ServiceURL:
 
```
2. 运行 `docker-compose -f docker-compose-env.yaml up -d` 启动所需的中间件
```bash
docker-compose -f docker-compose-env.yaml up -d
```
3. 若是想要热启动开发项目, 可以使用 `modd` 进行热启动
```bash
go install github.com/cortesi/modd/cmd/modd@latest
modd
```
