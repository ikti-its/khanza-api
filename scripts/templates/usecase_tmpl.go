package templates

var UsecaseTmpl = `package usecase

import "github.com/ikti-its/khanza-api/internal/modules/{{.ModuleName}}/internal/repository"

type {{.Name}}UseCase struct {
	Repository repository.{{.Name}}Repository
}

func New{{.Name}}UseCase(repository *repository.{{.Name}}Repository) *{{.Name}}UseCase {
	return &{{.Name}}UseCase{
		Repository: *repository,
	}
}
`

var SubUsecaseTmpl = `package usecase

import "github.com/ikti-its/khanza-api/internal/modules/{{.Module}}/internal/repository"

type {{.Name}}UseCase struct {
	Repository repository.{{.Name}}Repository
}

func New{{.Name}}UseCase(repository *repository.{{.Name}}Repository) *{{.Name}}UseCase {
	return &{{.Name}}UseCase{
		Repository: *repository,
	}
}
`
