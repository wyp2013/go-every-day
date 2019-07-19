## micro学习例子
### proto文件
- 生成代码工具
```
go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u -v github.com/micro/protoc-gen-micro
```
- hello.proto
```
syntax = "proto3";

//option java_multiple_files = true;
//option java_package = "io.grpc.examples.helloworld";
// option java_outer_classname = "HelloWorldProto";

package helloworld;

// The greeting service definition.
service Greeter {
    // Sends a greeting
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
    string name = 1;
    int32  age  = 2;
}

// The response message containing the greetings
message HelloReply {
    string message = 1;
    int32  id      = 2;
    int32  age     = 3;
}

```
- 用工具生成代码
```
# 在hello.proto的所在文件下运行
protoc --micro_out=. --go_out=. hello.proto
# 命令产生两个文件
hello.micro.go
hello.pb.go
```