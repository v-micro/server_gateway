package router

import (
	"github.com/labstack/echo/v4"
	"server_gateway/server"
)

func Init(r *echo.Echo)  {
	// order 服务
	order := r.Group("/order")
	{
		//订单服务
		order.GET("/ping/get",server.Order{}.Get)
	}
}
