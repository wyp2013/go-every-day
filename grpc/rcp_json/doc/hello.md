## 证书
```
# 生成私钥
openssl ecparam -genkey -name secp384r1 -out server.key
# 根据私钥生成公钥，并且服务节点为localhost
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650 -nodes -subj '/CN=localhost'
```


## 编译proto文件

在编写完`.proto`文件后，我们需要对其进行编译，就能够在`server`中使用

进入`proto`目录，执行以下命令

```
# 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

#编译hello_http.proto为hello_http.pb.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=go-every-day/grpc/rcp_json/proto/google/api:. ./hello.proto

#编译hello_http.proto为hello_http.pb.gw.proto
protoc --grpc-gateway_out=logtostderr=true:. ./hello.proto
```

执行完毕后将生成`hello.pb.go`和`hello.gw.pb.go`，分别针对`grpc`和`grpc-gateway`的功能支持

## 测试
```
# curl 测试
curl -X POST -k https://localhost:8090/hello_world -d '{"name": "xx", "age":1}'
```