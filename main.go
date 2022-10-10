package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testproj/proto"
)

func main() {
	//grpc链接客户端
	conn, err := grpc.Dial("127.0.0.1:8899", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	//程序完成不要忘记关闭链接
	defer conn.Close()
	//真正去链接服务端
	client := proto.NewUserSerivceClient(conn)
	//调用服务端 第一个是上线文  第二个是要传入的参数  传入的是结构体  返回的也是结构体
	ret, err := client.GetUser(context.Background(), &proto.UserRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.User.Id)
	fmt.Println(ret.User.Name)
}
