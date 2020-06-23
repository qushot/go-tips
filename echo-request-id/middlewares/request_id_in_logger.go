package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// RequestIDInLogger returns a middleware that logs HTTP requests contained RequestID.
func RequestIDInLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			c.SetLogger(log.New(id))
			return next(c)
		}
	}
}
