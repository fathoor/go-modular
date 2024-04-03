package templates

var PostgresTmpl = `package postgres

import (
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/entity"
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"math"
)

type {{.ModuleName}}RepositoryImpl struct {
	DB *sqlx.DB
}

func New{{.Name}}Repository(db *sqlx.DB) repository.{{.Name}}Repository {
	return &{{.ModuleName}}RepositoryImpl{db}
}

func (r *{{.ModuleName}}RepositoryImpl) Insert({{.ModuleName}} *entity.{{.Name}}) error {
	query := "INSERT INTO ... VALUES ..."

	_, err := r.DB.Exec(query, ...)

	return err
}

func (r *{{.ModuleName}}RepositoryImpl) Find() ([]entity.{{.Name}}, error) {
	query := "SELECT ... FROM ... WHERE deleted_at IS NULL"

	var records []entity.{{.Name}}
	err := r.DB.Select(&records, query)

	return {{.ModuleName}}, err
}

func (r *{{.ModuleName}}RepositoryImpl) FindPage(page, size int) ([]entity.{{.Name}}, int, error) {
	query := "SELECT ... FROM ... WHERE deleted_at IS NULL LIMIT $1 OFFSET $2"
	totalQuery := "SELECT COUNT(*) FROM ... WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.{{.Name}}
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *{{.ModuleName}}RepositoryImpl) FindById(id uuid.UUID) (entity.{{.Name}}, error) {
	query := "SELECT ... FROM ... WHERE ... = $1 AND deleted_at IS NULL"

	var record entity.{{.Name}}
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *{{.ModuleName}}RepositoryImpl) Update({{.ModuleName}} *entity.{{.Name}}) error {
	query := "UPDATE ..., updated_at = ..., updater = ... SET ... WHERE ... AND deleted_at IS NULL"

	_, err := r.DB.Exec(query, ...)

	return err
}

func (r *{{.ModuleName}}RepositoryImpl) Delete({{.ModuleName}} *entity.{{.Name}}) error {
	query := "UPDATE ... SET deleted_at = $1 WHERE ... = $2"

	_, err := r.DB.Exec(query, ...)

	return err
}
`

var SubPostgresTmpl = `package postgres

import (
	"github.com/fathoor/go-modular/internal/modules/{{.Module}}/internal/entity"
	"github.com/fathoor/go-modular/internal/modules/{{.Module}}/internal/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"math"
)

type {{.ModuleName}}RepositoryImpl struct {
	DB *sqlx.DB
}

func New{{.Name}}Repository(db *sqlx.DB) repository.{{.Name}}Repository {
	return &{{.ModuleName}}RepositoryImpl{db}
}

func (r *{{.ModuleName}}RepositoryImpl) Insert({{.ModuleName}} *entity.{{.Name}}) error {
	query := "INSERT INTO ... VALUES ..."

	_, err := r.DB.Exec(query, ...)

	return err
}

func (r *{{.ModuleName}}RepositoryImpl) Find() ([]entity.{{.Name}}, error) {
	query := "SELECT ... FROM ... WHERE deleted_at IS NULL"

	var records []entity.{{.Name}}
	err := r.DB.Select(&records, query)

	return {{.ModuleName}}, err
}

func (r *{{.ModuleName}}RepositoryImpl) FindPage(page, size int) ([]entity.{{.Name}}, int, error) {
	query := "SELECT ... FROM ... WHERE deleted_at IS NULL LIMIT $1 OFFSET $2"
	totalQuery := "SELECT COUNT(*) FROM ... WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.{{.Name}}
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *{{.ModuleName}}RepositoryImpl) FindById(id uuid.UUID) (entity.{{.Name}}, error) {
	query := "SELECT ... FROM ... WHERE ... = $1 AND deleted_at IS NULL"

	var record entity.{{.Name}}
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *{{.ModuleName}}RepositoryImpl) Update({{.ModuleName}} *entity.{{.Name}}) error {
	query := "UPDATE ..., updated_at = ..., updater = ... SET ... WHERE ... AND deleted_at IS NULL"

	_, err := r.DB.Exec(query, ...)

	return err
}

func (r *{{.ModuleName}}RepositoryImpl) Delete({{.ModuleName}} *entity.{{.Name}}) error {
	query := "UPDATE ... SET deleted_at = $1 WHERE ... = $2"

	_, err := r.DB.Exec(query, ...)

	return err
}
`
