package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/golongan/internal/controller"

)

func Route(app *fiber.App, Controller *controller.Controller) {
	modul := app.Group("/v1/golongan")
	modul.Get("/",       Controller.GetAll)
	modul.Get("/:id",    Controller.GetById)
	modul.Post("/",      Controller.Create)
	modul.Put("/:id",    Controller.Update)
	modul.Delete("/:id", Controller.Delete)
}
