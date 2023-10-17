package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"online-ticketing/app/config"
)

func AuthenticationError(msg string, err error) *Response {
	response := &Response{
		Code:        http.StatusUnauthorized,
		Message:     msg,
		MessageCode: "UNAUTHENTICATED",
	}

	if config.App.AppDebug && err != nil {
		response.Errors = err.Error()
	}

	return response
}

func AuthorizationError(msg string, err error) *Response {
	response := &Response{
		Code:        http.StatusForbidden,
		Message:     msg,
		MessageCode: "UNAUTHORIZED",
	}

	if config.App.AppDebug && err != nil {
		response.Errors = err.Error()
	}

	return response
}

func ClientError(msg string, err error) *Response {
	response := &Response{
		Code:        http.StatusBadRequest,
		Message:     msg,
		MessageCode: "BAD_REQUEST",
	}

	if config.App.AppDebug && err != nil {
		response.Errors = err.Error()
	}

	return response
}

func InvariantError(msg string, err error) *Response {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Handle the case when the record is not found
		return NotFoundError(msg, err)
	} else {
		// Handle other errors
	}

	response := &Response{
		Code:        http.StatusInternalServerError,
		MessageCode: "INTERNAL_SERVER_ERROR",
	}

	if config.App.AppDebug && err != nil {
		response.Message = msg
		response.Errors = err.Error()
	} else {
		response.Message = "Internal Server Error"
	}

	return response
}

func NotFoundError(msg string, err error) *Response {
	response := &Response{
		Code:        http.StatusNotFound,
		Message:     fmt.Sprintf("%s not found", msg),
		MessageCode: "DATA_NOT_FOUND",
	}

	if config.App.AppDebug && err != nil {
		response.Errors = err.Error()
	}

	return response
}
