package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// TodoFindAll は全件取得を行う
func TodoFindAll(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
