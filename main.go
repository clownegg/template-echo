package main

import (
	"github.com/labstack/echo"

	"template_app/src/route"
)

func main() {
	e := echo.New()
	route.Init(e)
	e.Start(":8080")
}
