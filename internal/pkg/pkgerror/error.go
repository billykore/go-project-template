package pkgerror

import (
	"github.com/billykore/go-service-tmpl/internal/pkg/codes"
)

const DefaultMsg = "Please try again later. If the problem persists, please contact support."

// Error represents domain error.
type Error struct {
	// Code is the error code.
	Code codes.Code
	// Msg is the custom error message that will be displayed to the client.
	// If no message is provided, the DefaultMsg will be used.
	Msg string
}

// New returns new Error.
func New(c codes.Code) *Error {
	return &Error{
		Code: c,
		Msg:  DefaultMsg,
	}
}

// NotFound returns a new Error with NotFound code.
func NotFound() *Error {
	return New(codes.NotFound)
}

// InternalServerError returns a new Error with InternalServerError code.
func InternalServerError() *Error {
	return New(codes.Internal)
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
	return err.Msg
}
