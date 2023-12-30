package controllers

import (
	"echo-templ-htmx-test/templates"
	"github.com/labstack/echo/v4"
)

func Root(c echo.Context) error {
	component := templates.Hello("Rainer")
	return component.Render(c.Request().Context(), c.Response().Writer)
}
