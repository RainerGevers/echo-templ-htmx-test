package clicked

import (
	"echo-templ-htmx-test/templates/partials"
	"github.com/labstack/echo/v4"
)

func ClickedPost(c echo.Context) error {
	clicked = clicked + 1
	component := partials.Clicked(clicked)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
