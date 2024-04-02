package templates

var PostgresTmpl = `package postgres

import (
	"math"

	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/entity"
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/repository"
	"github.com/jmoiron/sqlx"
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
	query := "SELECT ... FROM ..."

	var records []struct {
	}

	err := r.DB.Select(&records, query)
	if err != nil {
		return nil, err
	}

	{{.ModuleName}} := make([]entity.{{.Name}}, len(records))
	for i, record := range records {
		{{.ModuleName}}[i] = entity.{{.Name}}{
		}
	}

	return {{.ModuleName}}, nil
}

func (r *{{.ModuleName}}RepositoryImpl) FindPage(page, size int) ([]entity.{{.Name}}, int, error) {
	query := "SELECT ... FROM ... LIMIT ? OFFSET ?"
	totalQuery := "SELECT COUNT(...) FROM ..."

	var (
		records []struct {
		}
		total int
	)

	err := r.DB.Get(&total, totalQuery)
	if err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	err = r.DB.Select(&records, query, size, offset)
	if err != nil {
		return nil, 0, err
	}

	{{.ModuleName}} := make([]entity.{{.Name}}, len(records))
	for i, record := range records {
		{{.ModuleName}}[i] = entity.{{.Name}}{
		}
	}

	return {{.ModuleName}}, totalPage, nil
}

func (r *{{.ModuleName}}RepositoryImpl) FindById(id string) (entity.{{.Name}}, error) {
	query := "SELECT ... FROM ... WHERE ... = ?"

	var record struct {
	}

	err := r.DB.Get(&record, query, ...)
	if err != nil {
		return entity.{{.Name}}{}, err
	}

	{{.ModuleName}} := entity.{{.Name}}{
	}

	return {{.ModuleName}}, nil
}

func (r *{{.ModuleName}}RepositoryImpl) Update({{.ModuleName}} *entity.{{.Name}}) error {
	query := "UPDATE ... SET ... WHERE ... = ?"

	_, err := r.DB.Exec(query, ...)

	return err
}

func (r *{{.ModuleName}}RepositoryImpl) Delete({{.ModuleName}} *entity.{{.Name}}) error {
	query := "DELETE FROM ... WHERE ... = ?"

	_, err := r.DB.Exec(query, ...)

	return err
}
`
