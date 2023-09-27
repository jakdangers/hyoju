package cerrors

import (
	"errors"
	"net/http"
)

var (
	ErrInvalid    = &SentinelAPIError{status: http.StatusBadRequest, message: "Invalid"}
	ErrAuth       = &SentinelAPIError{status: http.StatusUnauthorized, message: "unauthorized"}
	ErrPermission = &SentinelAPIError{status: http.StatusForbidden, message: "forbidden"}
	ErrConflict   = &SentinelAPIError{status: http.StatusConflict, message: "conflict"}
	ErrNotFound   = &SentinelAPIError{status: http.StatusNotFound, message: "not found"}
	ErrInternal   = &SentinelAPIError{status: http.StatusInternalServerError, message: "server Error"}
)

type APIError interface {
	// APIError returns an HTTP status code and an API-safe error message.
	APIError() (int, string)
}

type SentinelAPIError struct {
	status      int
	message     string
	serviceCode string
}

func (e SentinelAPIError) Error() string {
	return e.message
}

func (e SentinelAPIError) APIError() (int, string) {
	return e.status, e.message
}

func ToSentinelAPIError(err error) (int, *SentinelAPIError) {
	var cErr *Error
	if errors.As(err, &cErr) {
		switch cErr.Kind {
		case Other, Internal, IO:
			return http.StatusInternalServerError, ErrInternal
		case Invalid:
			return http.StatusBadRequest, ErrInvalid
		case Auth:
			return http.StatusUnauthorized, ErrAuth
		case Permission:
			return http.StatusForbidden, ErrPermission
		case Exist:
			return http.StatusConflict, ErrConflict
		case NotExist:
			return http.StatusNotFound, ErrNotFound
		default:
			return http.StatusInternalServerError, ErrInternal
		}
	}
	return http.StatusBadRequest, &SentinelAPIError{
		status:  http.StatusBadRequest,
		message: err.Error(),
	}
}
