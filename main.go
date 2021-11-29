package main

import (
	"github.com/labstack/echo/v4"
	"server_gateway/router"
)

func main()  {
	//go 编译linux 环境可执行文件 darwin，freebsd，linux，windows
	//SET CGO_ENABLED=0
	//SET GOOS=linux
	//SET GOARCH=amd64
	//go build test.go

	e := echo.New()
	//路由配置
	router.RouterInit(e)
	//启动端口
	_ = e.Start(":10000")
}

