package templates

var ControllerTmpl = `package controller

import (
	"github.com/gofiber/fiber/v2"
	"{{.Package}}/internal/app/config"
	"{{.Package}}/internal/modules/{{.Module}}/internal/usecase"
)

type {{.Name}}Controller struct {
	UseCase   *usecase.{{.Name}}UseCase
	Validator *config.Validator
}

func New{{.Name}}Controller(useCase *usecase.{{.Name}}UseCase, validator *config.Validator) *{{.Name}}Controller {
	return &{{.Name}}Controller{
		UseCase:   useCase,
		Validator: validator,
	}
}
`
