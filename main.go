package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"server_gateway/router"
)

func main()  {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//路由注入
	router.Init(e)

	// Start server
	_ = e.Start(":1000")
}

