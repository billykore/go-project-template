package response

import (
	"errors"
	"net/http"

	"github.com/billykore/go-service-tmpl/internal/pkg/pkgerror"
)

// Response represents the response structure for HTTP responses.
type Response struct {
	Success bool           `json:"success"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Data    any            `json:"data,omitempty"`
}

type ErrorResponse struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
}

// Success returns status code 200 and success response with data.
func Success(data any) (int, *Response) {
	return http.StatusOK, &Response{
		Success: true,
		Data:    data,
	}
}

// Error returns error status code and error.
func Error(err error) (int, *Response) {
	var e *pkgerror.Error
	if errors.As(err, &e) {
		return responseCode[e.Code], &Response{
			Error: &ErrorResponse{
				Name:    responseName[e.Code],
				Message: err.Error(),
			},
		}
	}
	return InternalServerError(err)
}

// BadRequest returns status code 400 and error response.
func BadRequest(err error) (int, *Response) {
	return http.StatusBadRequest, &Response{
		Success: false,
		Error: &ErrorResponse{
			Name:    "BadRequest",
			Message: err.Error(),
		},
	}
}

// Unauthorized returns status code 401 and error response.
func Unauthorized(err error) (int, *Response) {
	return http.StatusUnauthorized, &Response{
		Success: false,
		Error: &ErrorResponse{
			Name:    "Unauthorized",
			Message: err.Error(),
		},
	}
}

// Forbidden returns status code 403 and error response.
func Forbidden(err error) (int, *Response) {
	return http.StatusForbidden, &Response{
		Success: false,
		Error: &ErrorResponse{
			Name:    "Forbidden",
			Message: err.Error(),
		},
	}
}

// InternalServerError returns status code 500 and error response.
func InternalServerError(err error) (int, *Response) {
	return http.StatusInternalServerError, &Response{
		Success: false,
		Error: &ErrorResponse{
			Name:    "InternalServerError",
			Message: err.Error(),
		},
	}
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

// responseName is a list of string representations for HTTP response status codes.
var responseName = []string{
	"Ok",
	"BadRequest",
	"Unauthorized",
	"Forbidden",
	"NotFound",
	"Conflict",
	"InternalServerError",
}
