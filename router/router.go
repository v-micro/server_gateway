package router

import (
	"github.com/labstack/echo/v4"
	"server_gateway/server"
)

func RouterInit(r *echo.Echo)  {
	//order服务
	order(r)
}

//订单服务
func order(r *echo.Echo)  {
	order := r.Group("/order")
	{
		order.GET("/ping/get",server.Order{}.Get)
	}
}
