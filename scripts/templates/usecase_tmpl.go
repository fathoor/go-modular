package templates

var UsecaseTmpl = `package usecase

import "{{.Package}}/internal/modules/{{.Module}}/internal/repository"

type {{.Name}}UseCase struct {
	Repository repository.{{.Name}}Repository
}

func New{{.Name}}UseCase(repository *repository.{{.Name}}Repository) *{{.Name}}UseCase {
	return &{{.Name}}UseCase{
		Repository: *repository,
	}
}
`
