package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo"
)

var (
	cmd     = "open"
	appPath = "/Applications/"
)

type ErrorStatus struct {
	Code         int
	ErrorMessage error
}

func Open(c echo.Context) error {
	path := c.QueryParam("path")
	errStatus := pathCheck(path)
	if errStatus != nil {
		return c.String(errStatus.Code, errStatus.ErrorMessage.Error())
	}

	app := c.QueryParam("app")
	args := []string{}
	if app != "" {
		args = append(args, "-a", app)
	}
	args = append(args, path)

	errStatus = openCmd(cmd, args)
	if errStatus != nil {
		return c.String(errStatus.Code, errStatus.ErrorMessage.Error())
	}

	return c.String(http.StatusOK, "ok")
}

func pathCheck(path string) *ErrorStatus {
	if path == "" {
		errStr := fmt.Sprintf("error occured: expected recieve 1 path.")
		return &ErrorStatus{http.StatusBadRequest, errors.New(errStr)}
	}

	if !exists(path) {
		errStr := fmt.Sprintf("error occured: no such file or directory: %v.", path)
		return &ErrorStatus{http.StatusBadRequest, errors.New(errStr)}
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func openCmd(cmd string, args []string) *ErrorStatus {
	fmt.Println(cmd, args)
	err := exec.Command(cmd, args...).Run()
	if err != nil {
		errStr := fmt.Sprintf("error occured by open execution: %v", err)
		return &ErrorStatus{http.StatusInternalServerError, errors.New(errStr)}
	}
	return nil
}
