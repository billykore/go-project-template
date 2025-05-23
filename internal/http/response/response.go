package response

import (
	"errors"
	"net/http"
	"time"

	"github.com/billykore/go-service-tmpl/pkg/utils/pkgerror"
)

// Response represents the response structure for HTTP responses.
type Response struct {
	Status     string `json:"status,omitempty"`
	Error      error  `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	ServerTime string `json:"serverTime,omitempty"`
}

// Success returns status code 200 and success response with data.
func Success(data any) (int, *Response) {
	return http.StatusOK, &Response{
		Status:     "OK",
		Data:       data,
		ServerTime: currentServerTime(),
	}
}

// Error returns error status code and error.
func Error(err error) (int, *Response) {
	var e *pkgerror.Error
	if errors.As(err, &e) {
		return responseCode[e.Code], &Response{
			Status:     responseStatus[e.Code],
			Message:    e.Msg,
			ServerTime: currentServerTime(),
		}
	}
	return InternalServerError(err)
}

// BadRequest returns status code 400 and error response.
func BadRequest(err error) (int, *Response) {
	return http.StatusBadRequest, &Response{
		Status:     "BAD_REQUEST",
		Message:    err.Error(),
		ServerTime: currentServerTime(),
	}
}

// Unauthorized returns status code 401 and error response.
func Unauthorized(err error) (int, *Response) {
	return http.StatusUnauthorized, &Response{
		Status:     "UNAUTHORIZED",
		Message:    err.Error(),
		ServerTime: currentServerTime(),
	}
}

// Forbidden returns status code 403 and error response.
func Forbidden(err error) (int, *Response) {
	return http.StatusForbidden, &Response{
		Status:     "FORBIDDEN",
		Message:    err.Error(),
		ServerTime: currentServerTime(),
	}
}

// InternalServerError returns status code 500 and error response.
func InternalServerError(err error) (int, *Response) {
	return http.StatusInternalServerError, &Response{
		Status:     "INTERNAL_SERVER_ERROR",
		Message:    err.Error(),
		ServerTime: currentServerTime(),
	}
}

// currentServerTime provides the current server time using the default time layout.
func currentServerTime() string {
	return time.Now().String()
}

// responseCode is a slice of integer HTTP status codes used for error response mapping.
var responseCode = []int{
	http.StatusOK,
	http.StatusBadRequest,
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusNotFound,
	http.StatusConflict,
	http.StatusInternalServerError,
}

// responseStatus is a list of string representations for HTTP response status codes.
var responseStatus = []string{
	"OK",
	"BAD_REQUEST",
	"UNAUTHORIZED",
	"FORBIDDEN",
	"NOT_FOUND",
	"CONFLICT",
	"INTERNAL_SERVER_ERROR",
}
