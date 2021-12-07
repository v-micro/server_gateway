package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"server_gateway/server_common/comutil"
	"server_gateway/server_common/protobuf/serverorder"
)

type Order struct {}


func (o Order) Get(ctx echo.Context) error {
	commonResponse := comutil.Response{}

	//链路追踪
	tracer, closer, err := comutil.InitJaeger("server_gateway","192.168.59.131:6831")	// 初始化jaeger，创建一条新tracer
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	//创建父标签
	span := tracer.StartSpan("url = /order/ping/get")
	defer span.Finish()
	//ctx传入链路新生成的
	c := opentracing.ContextWithSpan(context.Background(), span)

	//服务发现获取服务地址
	var endpoints = []string{"192.168.59.131:2379"}
	ser := comutil.NewServiceDiscovery(endpoints)
	defer ser.Close()
	err = ser.WatchService("server_order")
	if err != nil {
		return nil
	}


	//rpc 调用
	client := serverorder.GetClient(ser.GetServices()[0],tracer)
	defer client.Conn.Close()

	//链路调用
	response,err := client.Ping.Get(c,&serverorder.GetRequest{})
	if err != nil {
		return ctx.JSON(400, fmt.Sprintf("error: %s",err.Error()))
	}

	commonResponse.Code = 200
	commonResponse.Message = "请求成功"
	commonResponse.Details = response;
	return ctx.JSON(200,commonResponse)
}
