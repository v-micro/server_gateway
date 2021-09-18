package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"server_gateway/server_common/config"
	"server_gateway/server_common/protobuf/serverorder"
)

type Order struct {}


// 微服务启动
func (o Order) GetGrpcClient() (*grpc.ClientConn, error) {
	Grpc, err := grpc.Dial(config.ORDERPROT, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return Grpc, err
	}
	return Grpc, nil
}

func (o Order) Get(ctx echo.Context) error {
	//grpc 链接
	conn, _ := o.GetGrpcClient()
	defer conn.Close()

	//客户端
	client := serverorder.NewPingClient(conn)

	//方法请求
	response,err := client.Get(context.Background(),&serverorder.GetRequest{})
	if err != nil {
		return ctx.JSON(400, fmt.Sprintf("error: %s",err.Error()))
	}

	//打印返回值
	fmt.Println(response)

	return ctx.String(200,"输出成功")
}
