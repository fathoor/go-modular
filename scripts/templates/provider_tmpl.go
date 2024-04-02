package templates

var ProviderTmpl = `package auth

import (
	"github.com/fathoor/go-modular/internal/app/config"
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/controller"
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/repository/postgres"
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/router"
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

func Provide(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	{{.ModuleName}}Repository := postgres.New{{.Name}}Repository(db)
	{{.ModuleName}}UseCase := usecase.New{{.Name}}UseCase(&{{.ModuleName}}Repository, cfg)
	{{.ModuleName}}Controller := controller.New{{.Name}}Controller({{.ModuleName}}UseCase, validator)

	router.Route(app, {{.ModuleName}}Controller)
}
`
