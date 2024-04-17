package templates

var ProviderTmpl = `package {{.ModuleName}}

import (
	"{{.Package}}/internal/app/config"
	"{{.Package}}/internal/modules/{{.ModuleName}}/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Provide(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	router.Route(
		app,
	)
}
`
