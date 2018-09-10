package handler

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

func Open(c echo.Context) error {
	path := c.QueryParam("path")
	errStatus := ExistCheck(path)
	if errStatus != nil {
		return c.String(errStatus.Code, errStatus.ErrorMessage.Error())
	}

	args := []string{}
	app := c.QueryParam("app")
	if app != "" {
		errStatus = AppExistCheck(app)
		if errStatus != nil {
			return c.String(errStatus.Code, errStatus.ErrorMessage.Error())
		}
		args = append(args, "-a", app)
	}
	args = append(args, path)

	errStatus = openCmd(args)
	if errStatus != nil {
		return c.String(errStatus.Code, errStatus.ErrorMessage.Error())
	}

	return c.String(http.StatusOK, "ok")
}

func openCmd(args []string) *ErrorStatus {
	cmd := "open"
	err := exec.Command(cmd, args...).Run()
	if err != nil {
		errStr := fmt.Sprintf("%v: %v", CanNotExecuteMessage, err)
		return InternalServerError(errStr)
	}
	return nil
}
