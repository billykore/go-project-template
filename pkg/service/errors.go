package service

import "errors"

var (
	// ErrGetUserFromContext indicates failure while attempting to get user from context.
	ErrGetUserFromContext = errors.New("failed to get user from context")

	// ErrEmptyGreets indicates no messages were stored in the repository.
	ErrEmptyGreets = errors.New("empty greets")

	// ErrSayHelloFailed indicates failure while attempting to execute a say hello operation.
	ErrSayHelloFailed = errors.New("failed to say hello")
)
