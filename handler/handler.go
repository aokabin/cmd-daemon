package handler

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

func Open(c echo.Context) error {
	file := c.QueryParam("file")
	cmd := exec.Command("open", file)
	err := cmd.Run()
	if err != nil {
		errStr := fmt.Sprintf("error occured by open execution: %v", err)
		return c.String(http.StatusInternalServerError, errStr)
	}
	return c.String(http.StatusOK, "ok")
}
