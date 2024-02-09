package routes

import (
	"template_app/handlers"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	todoRoutes := e.Group("/todos")
	{
		todoRoutes.GET("", handlers.TodoFindAll)
	}
}
