package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"server_gateway/server_common/comutil"
	"server_gateway/server_common/protobuf/serverorder"
)

type Order struct {}


func (o Order) Get(ctx echo.Context) error {
	commonResponse := comutil.Response{}

	//rpc 调用
	client := serverorder.GetClient()
	defer client.Conn.Close()
	response,err := client.Ping.Get(context.Background(),&serverorder.GetRequest{})
	if err != nil {
		return ctx.JSON(400, fmt.Sprintf("error: %s",err.Error()))
	}

	commonResponse.Code = 200
	commonResponse.Message = "请求成功"
	commonResponse.Details = response;
	return ctx.JSON(200,commonResponse)
}
