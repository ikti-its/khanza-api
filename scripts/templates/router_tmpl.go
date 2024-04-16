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
	{{.ModuleName}} := app.Group("/v1/{{.ModuleName}}")
	{
		{{.ModuleName}}.Post("/", middleware.Authenticate([]int{1337, 1}), {{.ModuleName}}Controller.Create)
		{{.ModuleName}}.Get("/", middleware.Authenticate([]int{1337, 1, 2}), {{.ModuleName}}Controller.Get)
		{{.ModuleName}}.Get("/:id", middleware.Authenticate([]int{1337, 1, 2}), {{.ModuleName}}Controller.GetById)
		{{.ModuleName}}.Put("/:id", middleware.Authenticate([]int{1337, 1}), {{.ModuleName}}Controller.Update)
		{{.ModuleName}}.Delete("/:id", middleware.Authenticate([]int{1337, 1}), {{.ModuleName}}Controller.Delete)
	}
}
`
