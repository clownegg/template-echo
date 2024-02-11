package handlers

import (
	"net/http"
	"template_app/factory"

	"github.com/labstack/echo"
)

func TodoFindAll(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()
	todos := service.FindAll()

	return c.JSON(http.StatusOK, todos)
}

func TodoFindById(c echo.Context) error {
	return c.JSON(200, "test")
}
