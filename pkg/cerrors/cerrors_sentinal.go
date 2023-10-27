package cerrors

import (
	"errors"
	"net/http"
)

type SentinelAPIError struct {
	status         int
	serviceMessage string
	serviceCode    string
}

func NewSentinelAPIError(status int, serviceMessage ServiceMessage, serviceCode ServiceCode) (int, SentinelAPIError) {
	return status, SentinelAPIError{
		status:         status,
		serviceMessage: string(serviceMessage),
		serviceCode:    string(serviceCode),
	}
}

func (e SentinelAPIError) Error() string {
	return e.serviceMessage
}

func ToSentinelAPIError(err error) (int, SentinelAPIError) {
	var cErr *Error
	if errors.As(err, &cErr) {
		switch cErr.Kind {
		case Other, Internal, IO:
			return NewSentinelAPIError(http.StatusInternalServerError, cErr.ServiceMessage, cErr.ServiceCode)
		case Invalid:
			return NewSentinelAPIError(http.StatusBadRequest, cErr.ServiceMessage, cErr.ServiceCode)
		case Auth:
			return NewSentinelAPIError(http.StatusUnauthorized, cErr.ServiceMessage, cErr.ServiceCode)
		case Permission:
			return NewSentinelAPIError(http.StatusForbidden, cErr.ServiceMessage, cErr.ServiceCode)
		case Exist:
			return NewSentinelAPIError(http.StatusConflict, cErr.ServiceMessage, cErr.ServiceCode)
		case NotExist:
			return NewSentinelAPIError(http.StatusNotFound, cErr.ServiceMessage, cErr.ServiceCode)
		default:
			return NewSentinelAPIError(http.StatusInternalServerError, cErr.ServiceMessage, cErr.ServiceCode)
		}
	}
	return http.StatusBadRequest, SentinelAPIError{
		status:         http.StatusBadRequest,
		serviceMessage: err.Error(),
	}
}
