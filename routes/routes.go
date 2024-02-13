package routes

import (
	"template_app/handlers"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	{
		v1.GET("/todos", handlers.FindTodoAll)
		v1.GET("/todos/:id", handlers.FindTodoById)
		v1.POST("/todos", handlers.CreateTodo)
	}
}
