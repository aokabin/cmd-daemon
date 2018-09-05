package main

import (
	"github.com/aokabin/cmd-daemon/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/open", handler.Open)
	e.Start(":22222")
}
