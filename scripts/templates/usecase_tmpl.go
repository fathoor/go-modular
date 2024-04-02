package templates

var UsecaseTmpl = `package usecase

import (
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/repository"
)

type {{.Name}}UseCase struct {
	Repository repository.{{.Name}}Repository
}

func New{{.Name}}UseCase(repository *repository.{{.Name}}Repository) *{{.Name}}UseCase {
	return &{{.Name}}UseCase{
		Repository: *repository,
	}
}
`
