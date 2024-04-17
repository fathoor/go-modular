package templates

var RepositoryTmpl = `package repository

import (
	"{{.Package}}/internal/modules/{{.Module}}/internal/entity"
	"github.com/google/uuid"
)

type {{.Name}}Repository interface {
	Insert({{.ModuleName}} *entity.{{.Name}}) error
	Find() ([]entity.{{.Name}}, error)
	FindPage(page, size int) ([]entity.{{.Name}}, int, error)
	FindById(id uuid.UUID) (entity.{{.Name}}, error)
	Update({{.ModuleName}} *entity.{{.Name}}) error
	Delete({{.ModuleName}} *entity.{{.Name}}) error
}
`
