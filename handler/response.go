package handler

import (
	"errors"
	"net/http"
)

type ErrorStatus struct {
	Code         int
	ErrorMessage error
}

const (
	EmptyPathMessage             = "error occured: expected recieve 1 path"
	NoSuchFileOrDirectoryMessage = "error occured: no such file or directory: "
	NotFoundApplicationMessage   = "error occured: not found application: "
	CanNotExecuteMessage         = "error occured: can not execute open command: "
)

func NotFound(message string) *ErrorStatus {
	return &ErrorStatus{http.StatusNotFound, errors.New(message)}
}

func InternalServerError(message string) *ErrorStatus {
	return &ErrorStatus{http.StatusInternalServerError, errors.New(message)}
}

func BadRequest(message string) *ErrorStatus {
	return &ErrorStatus{http.StatusBadRequest, errors.New(message)}
}
