package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
)

var (
	appDir = "/Applications/"
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

	args := []string{}
	app := c.QueryParam("app")
	if app != "" {
		errStatus = appCheck(app)
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

func appCheck(appName string) *ErrorStatus {
	return findApp(appName)
}

func findApp(appName string) *ErrorStatus {
	apps := appwalk(appDir)
	appFile := fmt.Sprintf("%s.app", appName)
	for _, app := range apps {
		if app == appFile {
			return nil
		}
	}
	errStr := fmt.Sprintf("error occured: not found application: %v.", appName)
	return &ErrorStatus{http.StatusBadRequest, errors.New(errStr)}
}

func appwalk(appdir string) []string {
	if fileAccessDenyed(appdir) {
		return nil
	}

	files, err := ioutil.ReadDir(appdir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		fileIsApp := strings.HasSuffix(file.Name(), ".app")
		if fileIsApp {
			paths = append(paths, file.Name())
			continue
		}
		if file.IsDir() {
			paths = append(paths, appwalk(filepath.Join(appdir, file.Name()))...)
			continue
		}
	}

	return paths
}

func openCmd(args []string) *ErrorStatus {
	cmd := "open"
	fmt.Println(cmd, args)
	err := exec.Command(cmd, args...).Run()
	if err != nil {
		errStr := fmt.Sprintf("error occured by open execution: %v", err)
		return &ErrorStatus{http.StatusInternalServerError, errors.New(errStr)}
	}
	return nil
}

func fileAccessDenyed(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	m := info.Mode()
	return m&(1<<2) == 0
}
