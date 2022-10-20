package main

import (
	"context"
	"fmt"
	"greeter/proto/greeter"
	"net"

	"google.golang.org/grpc"
)

//1、定义远程调用的结构体和方法 我们这个结构体需要实现GreeterServer的接口

type HelloWorld struct{}

func (this HelloWorld) SayHelloWorld(c context.Context, req *greeter.HelloWorldReq) (*greeter.HelloWorldRes, error) {

	return &greeter.HelloWorldRes{
		Message: "你好" + req.Name,
	}, nil

}
func main() {
	//1. 初始一个 grpc 对象
	grpcServer := grpc.NewServer()
	//2. 注册服务
	greeter.RegisterGreeterServer(grpcServer, &HelloWorld{})
	//3.  设置监听， 指定 IP、port
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	// 4退出关闭监听
	defer listener.Close()
	//5、启动服务
	grpcServer.Serve(listener)

}
