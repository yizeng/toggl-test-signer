package v1

import (
	"net/http"

	"go.uber.org/zap"
)

type ErrResponse struct {
	Err error `json:"-"` // low-level runtime error

	StatusText string `json:"status"` // user-level status message
	StatusCode int    `json:"status_code"`

	AppCode   int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText string `json:"error,omitempty"` // application-level error message, for debugging
}

func ErrBadRequest(err error) *ErrResponse {
	zap.L().Info("bad request", zap.Error(err))

	return &ErrResponse{
		Err:        err,
		ErrorText:  err.Error(),
		StatusCode: http.StatusBadRequest,
		StatusText: "bad request",
	}
}

func ErrInternalServerError(err error) *ErrResponse {
	zap.L().Error("internal server error", zap.Error(err))

	return &ErrResponse{
		Err:        err,
		ErrorText:  err.Error(),
		StatusCode: http.StatusInternalServerError,
		StatusText: "internal server error",
	}
}

func ErrNotFound(err error) *ErrResponse {
	zap.L().Info("not found", zap.Error(err))

	return &ErrResponse{
		Err:        err,
		ErrorText:  err.Error(),
		StatusCode: http.StatusNotFound,
		StatusText: "not found",
	}
}
