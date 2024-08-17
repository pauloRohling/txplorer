package json

import (
	"github.com/pauloRohling/txplorer/internal/model"
	"net/http"
	"time"
)

type Response struct {
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Details   string `json:"details,omitempty"`
	Timestamp string `json:"timestamp"`
}

func NewResponseFromError(err model.Error) Response {
	status := ErrorStatus(err)
	return Response{
		Status:    status,
		Error:     http.StatusText(status),
		Message:   ErrorMessage(status),
		Details:   err.Error(),
		Timestamp: err.Timestamp.Format(time.RFC3339),
	}
}

func ErrorStatus(err model.Error) int {
	switch err.Type {
	case model.InternalErrorType:
		return http.StatusInternalServerError
	case model.ForbiddenErrorType:
		return http.StatusForbidden
	case model.NotFoundErrorType:
		return http.StatusNotFound
	case model.UnauthorizedErrorType:
		return http.StatusUnauthorized
	case model.ValidationErrorType:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func ErrorMessage(status int) string {
	switch status {
	case http.StatusInternalServerError:
		return "The server encountered an unexpected condition. Please contact support or system administrators for assistance."
	case http.StatusForbidden:
		return "Access to the requested resource is forbidden. Please ensure that the client has the required permissions."
	case http.StatusNotFound:
		return "The requested resource was not found on the server. Please check the URL and resource identifiers for accuracy."
	case http.StatusUnauthorized:
		return "Authentication credentials are missing or invalid. Please provide valid credentials or acquire proper authorization."
	case http.StatusBadRequest:
		return "The request contains invalid syntax or missing parameters. Please review the request and provide valid data."
	default:
		return ""
	}
}
