# 生成gRPC代码

```
protoc -I grpc grpc/hello.proto  --go_out=plugins=grpc:grpc
```


```
1. generate grpc stub

protoc -I/usr/local/include -I. \
                 -I$GOPATH/src \
                 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
                 --go_out=plugins=grpc:. \
                 --grpc-gateway_out=logtostderr=true:. \
                 ./hello.proto


2. Generate reverse-proxy

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  ./hello.proto


protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  ./hello.proto

```