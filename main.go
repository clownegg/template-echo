package main

import (
	"github.com/labstack/echo"

	"template_app/src/config"
	"template_app/src/route"
)

func main() {
	e := echo.New()
	config.Init(e)
	route.Init(e)
	e.Start(":8080")
}
