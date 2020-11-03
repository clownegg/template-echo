package route

import (
	"template_app/src/controller"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	todoRoutes := e.Group("/todos")
	{
		todoRoutes.GET("", controller.TodoFindAll)
	}
}
