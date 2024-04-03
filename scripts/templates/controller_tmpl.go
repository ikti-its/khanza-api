package templates

var ControllerTmpl = `package controller

import (
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/{{.ModuleName}}/internal/usecase"
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

var SubControllerTmpl = `package controller

import (
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/{{.Module}}/internal/usecase"
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
