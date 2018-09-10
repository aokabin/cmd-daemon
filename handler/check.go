package handler

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	appDir = "/Applications/"
)

func ExistCheck(path string) *ErrorStatus {
	if isEmpty(path) {
		return BadRequest(EmptyPathMessage)
	}

	if isNotExists(path) {
		errStr := fmt.Sprintf("%v: %v", NoSuchFileOrDirectoryMessage, path)
		return BadRequest(errStr)
	}
	return nil
}

func isEmpty(path string) bool {
	return path == ""
}

func isNotExists(path string) bool {
	_, err := os.Stat(path)
	return err != nil
}

func AppExistCheck(appName string) *ErrorStatus {
	apps := applications(appDir)
	appName = fmt.Sprintf("%s.app", appName)
	for _, app := range apps {
		if app == appName {
			return nil
		}
	}
	errStr := fmt.Sprintf("%v: %v", NotFoundApplicationMessage, appName)
	return BadRequest(errStr)
}

func applications(dir string) []string {
	if isAccessDenied(dir) {
		return nil
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if isApp(file.Name()) {
			paths = append(paths, file.Name())
		} else if file.IsDir() {
			paths = append(paths, applications(filepath.Join(dir, file.Name()))...)
		}
	}

	return paths
}

func isApp(file string) bool {
	return strings.HasSuffix(file, ".app")
}

func isAccessDenied(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	m := info.Mode()
	return m&(1<<2) == 0
}
