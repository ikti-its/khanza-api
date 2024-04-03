package templates

var RouterTmpl = `package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fathoor/go-modular/internal/modules/{{.ModuleName}}/internal/controller"
)

func Route(
	app *fiber.App,
	controller *controller.{{.Name}}Controller,
) {
	_ = app.Group("/v1/{{.ModuleName}}")
	{}
}
`
