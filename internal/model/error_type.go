package model

type ErrorType string

func (errorType ErrorType) String() string {
	return string(errorType)
}

const (
	ForbiddenErrorType    ErrorType = "ForbiddenError"
	InternalErrorType     ErrorType = "InternalError"
	NotFoundErrorType     ErrorType = "NotFoundError"
	UnauthorizedErrorType ErrorType = "UnauthorizedError"
	ValidationErrorType   ErrorType = "ValidationError"
)
