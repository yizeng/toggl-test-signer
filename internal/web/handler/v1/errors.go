package v1

import "go.uber.org/zap"

type ErrResponse struct {
	Err error `json:"-"` // low-level runtime error

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func ErrBadRequest(err error) *ErrResponse {
	zap.L().Info("bad request", zap.Error(err))

	return &ErrResponse{
		Err:        err,
		StatusText: "bad request",
		ErrorText:  err.Error(),
	}
}

func ErrInternalServerError(err error) *ErrResponse {
	zap.L().Error("internal server error", zap.Error(err))

	return &ErrResponse{
		Err:        err,
		StatusText: "internal server error",
		ErrorText:  err.Error(),
	}
}

func ErrNotFound() *ErrResponse {
	zap.L().Info("not found")

	return &ErrResponse{
		StatusText: "not found",
	}
}
