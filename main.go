package main

import (
	"github.com/labstack/echo/v4"
	"server_gateway/router"
)

func main()  {
	e := echo.New()
	//路由配置
	router.RouterInit(e)
	//启动端口
	_ = e.Start(":10000")
}

