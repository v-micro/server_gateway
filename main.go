package main

import (
	"github.com/labstack/echo/v4"
	"server_gateway/router"
	"server_gateway/server_common/config"
)

func main()  {
	e := echo.New()
	router.Init(e)
	_ = e.Start(config.GATEWAYPROT)
}

