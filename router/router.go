package router

import (
	"net/http"

	"github.com/test_cache_CQRS/boot"
	"github.com/test_cache_CQRS/config"
	"github.com/test_cache_CQRS/infrastructure/httplib"
	customMiddleware "github.com/test_cache_CQRS/infrastructure/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	if config.Conf.LogMode {
		c.Use(middleware.Logger())
	}

	echo.NotFoundHandler = func(c echo.Context) error {
		// render 404 custom response
		return httplib.SetErrorResponse(c, http.StatusNotFound, "Not Matching of Any Routes")
	}

	//grouping on root endpoint
	api := c.Group("/", customMiddleware.RateLimiterMiddleware(hr.Setup.Limiter))

	//grouping on "api/v1"
	v1 := api.Group("api/v1")

	//module health
	prefixHealth := v1.Group("/health")
	hr.Setup.HealthHttp.GroupHealth(prefixHealth)

	//module article
	prefixArticle := v1.Group("/articles")
	hr.Setup.ArticleHttp.GroupArticle(prefixArticle)

	return c
}
