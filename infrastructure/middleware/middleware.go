package middleware

import (
	"net/http"

	"github.com/test_cache_CQRS/infrastructure/httplib"
	"github.com/test_cache_CQRS/infrastructure/limiter"

	"github.com/labstack/echo/v4"
)

func RateLimiterMiddleware(limiter *limiter.RateLimiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if limiter.Allow() {
				return next(c)
			}
			return httplib.SetErrorResponse(c, http.StatusTooManyRequests, "rate limit exceeded")
		}
	}
}
