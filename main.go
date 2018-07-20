package main

import (
    "github.com/labstack/echo"
    "github.com/aokabin/cmd-daemon/handler"
)

func main() {
    e := echo.New()
    e.GET("/open", handler.Open)
    e.Start(":22222")
}