package internal_errors

import "net/http"

type InternalErrors struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewInternalError(message, err string, code int, causes []Causes) *InternalErrors {
	return &InternalErrors{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes}
}

func NewBadRequestError(message string) *InternalErrors {
	return &InternalErrors{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *InternalErrors {
	return &InternalErrors{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *InternalErrors {
	return &InternalErrors{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *InternalErrors {
	return &InternalErrors{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewUnauthorizedError(message string) *InternalErrors {
	return &InternalErrors{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewForbiddenError(message string) *InternalErrors {
	return &InternalErrors{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
