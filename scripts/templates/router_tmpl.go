package templates

var RouterTmpl = `package router

import (
	"{{.Package}}/internal/modules/{{.ModuleName}}/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func Route(
	app *fiber.App,
) {
	{{.ModuleName}} := app.Group("/v1/{{.ModuleName}}")
}
`
