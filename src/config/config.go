package config

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{},
		AllowHeaders: []string{},
	}))
}
