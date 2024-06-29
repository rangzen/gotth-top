package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// render creates a compatible interface for rendering templ components
// using the Echo web framework.
// Reference: https://github.com/danawoodman/go-echo-htmx-templ/blob/main/handlers/render.go
// Reference: https://github.com/emarifer/go-echo-templ-htmx/blob/main/handlers/auth.handlers.go
func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response().Writer)
}
