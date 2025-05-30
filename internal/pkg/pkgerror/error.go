package pkgerror

import (
	"github.com/billykore/go-service-tmpl/internal/pkg/codes"
)

const defaultMsg = "Please try again later. If the problem persists, please contact support."

// Error represents domain error.
type Error struct {
	// Code is the error code.
	Code codes.Code
	// Err is the error.
	Err error
	// Msg is the custom error message that will be displayed to the client.
	Msg string
}

// New returns new Error.
func New(c codes.Code, err error) *Error {
	return &Error{
		Code: c,
		Err:  err,
		Msg:  defaultMsg,
	}
}

// NotFound returns a new Error with NotFound code.
func NotFound(err error) *Error {
	return New(codes.NotFound, err)
}

// InternalServerError returns a new Error with InternalServerError code.
func InternalServerError(err error) *Error {
	return New(codes.Internal, err)
}

// SetMsg sets a custom error message for the client to display.
// If no message is provided, the default message will be used.
//
// Default message:
// "Please try again later. If the problem persists, please contact support".
func (err *Error) SetMsg(msg string) *Error {
	err.Msg = msg
	return err
}

func (err *Error) Error() string {
	return err.Err.Error()
}
