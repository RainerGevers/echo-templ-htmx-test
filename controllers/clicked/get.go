package clicked

import (
	"echo-templ-htmx-test/templates/partials"
	"github.com/labstack/echo/v4"
)

var clicked int

func ClickedGet(c echo.Context) error {
	component := partials.CountButton(clicked)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
