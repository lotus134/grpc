package main

import (
	"client/proto/greeter"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1、连接服务器
	/*
		credentials.NewClientTLSFromFile ：从输入的证书文件中为客户端构造TLS凭证。
		grpc.WithTransportCredentials ：配置连接级别的安全凭证（例如，TLS/SSL），返回一个
		DialOption，用于连接服务器。

	*/

	grpcClient, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
	}
	//2、注册客户端
	client := greeter.NewGreeterClient(grpcClient)

	res, _ := client.SayHelloWorld(context.Background(), &greeter.HelloWorldReq{
		Name: "小王",
	})

	fmt.Println(res.Message)
}
