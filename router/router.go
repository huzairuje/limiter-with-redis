package router

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/test_cache_CQRS/boot"
	"github.com/test_cache_CQRS/infrastructure/httplib"

	"github.com/labstack/echo/v4"
)

type HandlerRouter struct {
	Setup boot.HandlerSetup
}

func NewHandlerRouter(setup boot.HandlerSetup) InterfaceRouter {
	return &HandlerRouter{
		Setup: setup,
	}
}

type InterfaceRouter interface {
	RouterWithMiddleware() *echo.Echo
}

func (hr *HandlerRouter) RouterWithMiddleware() *echo.Echo {
	c := echo.New()
	c.Use(middleware.Logger())

	echo.NotFoundHandler = func(c echo.Context) error {
		// render 404 custom response
		return httplib.SetErrorResponse(c, http.StatusNotFound, "Not Matching of Any Routes")
	}

	//grouping on "/api/v1"
	api := c.Group("/api/v1")

	//module health
	prefixHealth := api.Group("/health")
	hr.Setup.HealthHttp.GroupHealth(prefixHealth)

	//module article
	prefixArticle := api.Group("/articles")
	hr.Setup.ArticleHttp.GroupArticle(prefixArticle)

	return c
}
