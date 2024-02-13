package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"template_app/factory"
	"template_app/models"

	"github.com/labstack/echo"
)

func FindTodoAll(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()

	cond := models.TodosCond{
		IsDeleted: 0,
		Offset:    0,
		Limit:     30,
	}
	err := c.Bind(&cond)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	todos, count := service.FindAll(cond)
	c.Response().Header().Set("x-count", fmt.Sprint(count))
	return c.JSON(http.StatusOK, todos)
}

func FindTodoById(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()

	cond := models.TodoCond{
		IsDeleted: 0,
	}
	err := c.Bind(&cond)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Path Parameter")
	}
	todo := service.FindById(id, cond)
	return c.JSON(http.StatusOK, todo)
}

func CreateTodo(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()

	postData := models.TodoPost{}
	err := c.Bind(&postData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = service.Create(postData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}
