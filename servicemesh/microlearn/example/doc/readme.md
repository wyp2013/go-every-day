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
### server端实现
```
type Greeter struct {
}

// Sends a greeting
func (s *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest, reply *pb.HelloReply) error {
	if in == nil {
		return errors.New("input is empty")
	}

	age := in.GetAge()
	name := in.GetName()
	id := rand.Int31()

	reply.Message = fmt.Sprintf("you name is %s, age is %d, give you id is %d", name, age, id)
	reply.Id = id
	reply.Age = age

	return nil
}

func (s *Greeter) SayHello2(ctx context.Context, stream pb.Greeter_SayHello2Stream) error {
	count := 0

	for {
		req, err := stream.Recv()
		count++

		if err == io.EOF {
			fmt.Println("strem is close...")
			return stream.Close()
		}

		if err != nil {
			return err
		}

		// do something
		fmt.Println(req)
	}



	return nil
}


func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379",
			//"http://127.0.0.1:12379",
			//"http://127.0.0.1:22379",
		}
	})

	service := micro.NewService(
		micro.Name("HelloGreeter"),
		micro.Version("latest"),
		micro.Registry(reg),
		// micro.Server(grpc.NewServer()), // 这里不用默认的rpcserver，改用grpcserver，注释也可以运行
	)

	service.Init()
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
```
### client端实现
```
func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379", "http://127.0.0.1:12379", "http://127.0.0.1:22379",
		}
	})

	service := micro.NewService(
		micro.Name("HelloGreeter1"),
		micro.Registry(reg),
		// micro.Server(grpc.NewServer()),  // 这里不用默认的rpcserver，改用grpcserver，注释也可以运行
	)

	service.Init()

	ctx := context.Background()

	// 服务提供方的名字
	helloClient := pb.NewGreeterService("HelloGreeter", service.Client())

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("xx-%d-yy", i)
		age := int32(i + 10)
		req := &pb.HelloRequest{Name: name, Age: age}

		reply, err := helloClient.SayHello(ctx, req)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("response: ", reply)
	}

	stremClient, err := helloClient.SayHello2(ctx)
	defer stremClient.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("sayhllo2-%d", i)
		age := int32(i + 10)
		req := &pb.HelloRequest{Name: name, Age: age}
		err := stremClient.Send(req)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
```