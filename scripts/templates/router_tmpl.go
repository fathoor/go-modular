package templates

var RouterTmpl = `package router

import (
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/controller"
	"github.com/gofiber/fiber/v3"
)

func Route(
	app *fiber.App,
	controller *controller.{{.Name}}Controller,
) {
	_ = app.Group("/v1/{{.ModuleName}}")
	{}
}
`
