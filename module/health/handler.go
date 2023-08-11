package health

import (
	"context"
	logger "github.com/test_cache_CQRS/infrastructure/log"
	"github.com/test_cache_CQRS/module/primitive"
	"github.com/test_cache_CQRS/utils"
	"net/http"

	"github.com/test_cache_CQRS/infrastructure/httplib"

	"github.com/labstack/echo/v4"
)

type Http struct {
	serviceHealth InterfaceService
}

func NewHttp(serviceHealth InterfaceService) InterfaceHttp {
	return &Http{
		serviceHealth: serviceHealth,
	}
}

type InterfaceHttp interface {
	GroupHealth(group *echo.Group)
}

func (h *Http) GroupHealth(g *echo.Group) {
	g.GET("/ping", h.Ping)
	g.GET("/check", h.HealthCheckApi)
}

func (h *Http) Ping(c echo.Context) error {
	return httplib.SetSuccessResponse(c, http.StatusOK, http.StatusText(http.StatusOK), "pong")
}

func (h *Http) HealthCheckApi(c echo.Context) error {
	logCtx := "handler.HealthCheckApi"
	ctx := context.Background()
	resp, err := h.serviceHealth.CheckUpTime(ctx)
	if err != nil {
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "h.serviceHealth.CheckUpTime")
		return httplib.SetErrorResponse(c, http.StatusInternalServerError, primitive.SomethingWentWrong)
	}
	return httplib.SetSuccessResponse(c, http.StatusOK, http.StatusText(http.StatusOK), resp)
}
