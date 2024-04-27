package api

import (
	"account-service/internal/service"
	"log/slog"
	"net/http"
)

type HTTPErrorHandler interface {
	Handle(err error) (int, interface{})
}

type echoErrorHandlerImpl struct {
	logger *slog.Logger
}

func NewHTTPErrorHandler(logger *slog.Logger) HTTPErrorHandler {
	return echoErrorHandlerImpl{logger}
}

func (h echoErrorHandlerImpl) Handle(err error) (int, interface{}) {
	var code int
	switch err.(type) {
	case service.NotFoundError:
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}

	return code, struct {
		Message string
	}{
		Message: err.Error(),
	}
}
