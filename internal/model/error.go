package model

import (
	"errors"
	"runtime/debug"
	"time"
)

type Error struct {
	Message    string    `json:"message"`
	Type       ErrorType `json:"type"`
	Timestamp  time.Time `json:"timestamp"`
	StackTrace string    `json:"-"`
	Err        error     `json:"-"`
}

func (error Error) Error() string {
	return error.Message
}

func ForbiddenError(message string, optionalError ...error) Error {
	return NewError(message, ForbiddenErrorType, optionalError...)
}

func InternalError(message string, optionalError ...error) Error {
	return NewError(message, InternalErrorType, optionalError...)
}

func NotFoundError(message string, optionalError ...error) Error {
	return NewError(message, NotFoundErrorType, optionalError...)
}

func UnauthorizedError(message string, optionalError ...error) Error {
	return NewError(message, UnauthorizedErrorType, optionalError...)
}

func ValidationError(message string, optionalError ...error) Error {
	return NewError(message, ValidationErrorType, optionalError...)
}

func NewError(message string, errorType ErrorType, optionalError ...error) Error {
	var err error
	if len(optionalError) > 0 {
		err = optionalError[0]
	}

	var customErr Error
	if err != nil && errors.As(err, &customErr) {
		return Error{
			Message:    customErr.Message,
			Type:       customErr.Type,
			Timestamp:  customErr.Timestamp,
			StackTrace: customErr.StackTrace,
			Err:        customErr.Err,
		}
	}

	return Error{
		Message:    message,
		Type:       errorType,
		Timestamp:  time.Now().UTC(),
		StackTrace: string(debug.Stack()),
		Err:        err,
	}
}
