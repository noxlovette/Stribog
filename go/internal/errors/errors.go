package errors

import "errors"

var (
	ErrNotFound        = errors.New("resource not found")
	ErrInvalidUserId   = errors.New("invalid user id")
	ErrOperationFailed = errors.New("operation failed")
	ErrForbidden       = errors.New("forbidden")
)
