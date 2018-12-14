# Go代码总结

总结一些常用业务模型，方便快速创建项目

使用 grpc-ecosystem 快速创建api,rpc,swagger

## roadmap

### models

- [x] user
- [ ] user_permission
- [ ] product
- [ ] order

### api

- [ ] account
- [ ] oauth
- [ ] product
- [ ] order
- [ ] payment

## 添加代码

### 1、在 api/protobuf 目录下添加新的目录，定义你的proto文件

### 2、运行 make gen 命令， 将proto文件生成对应go代码，swagger文件

### 3、在 internal/app 目录下创建你的目录，实现对应的server

### 4、在 cmd/swagger/main.go 文件中将实现的server注册进grpc.Server, 对应的gateway注册进mux,如下

```go
s := grpc.NewServer(
    grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
    grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
    grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
)

// 注册 grpc
user.RegisterRPCServiceServer(s, userRPCServer)

gwmux := runtime.NewServeMux(
    runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
    // This is necessary to get error details properly
    // marshalled in unary requests.
    runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
)

// 注册 gateway
err = user.RegisterRPCServiceHandler(context.Background(), gwmux, conn)
if err != nil {
    log.Fatalln("Failed to register gateway:", err)
}
```

### 5、编译 make gb

### 6、运行 make up



## 依赖

```go
  go get -v -u github.com/gogo/protobuf/protoc-gen-gogo
  go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
  go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
  go get -v -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
  go get -v -u github.com/rakyll/statik
```