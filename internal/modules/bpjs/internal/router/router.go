package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/bpjs/internal/controller"

)

func Route(app *fiber.App, Controller *controller.Controller) {
	var roles = []int{1337, 0, 1, 2, 3}
	modul := app.Group("/v1/bpjs")
	modul.Get("/",       middleware.Authenticate(roles), Controller.GetAll)
	modul.Get("/:id",    middleware.Authenticate(roles), Controller.GetById)
	modul.Post("/",      middleware.Authenticate(roles), Controller.Create)
	modul.Put("/:id",    middleware.Authenticate(roles), Controller.Update)
	modul.Delete("/:id", middleware.Authenticate(roles), Controller.Delete)	
}
