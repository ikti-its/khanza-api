package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/controller"
)

func AmbulansRoute(app *fiber.App, ambulansController *controller.AmbulansController) {
	ambulans := app.Group("/v1/ambulans")

	ambulans.Get("/request/pending", ambulansController.GetPendingRequests)
	ambulans.Post("/request", ambulansController.RequestAmbulans)
	ambulans.Put("/terima/:no_ambulans", ambulansController.AcceptAmbulansRequest)
	ambulans.Post("/", ambulansController.Create)
	ambulans.Get("/", ambulansController.GetAll)
	ambulans.Get("/:no_ambulans", ambulansController.GetByNoAmbulans)
	ambulans.Put("/:no_ambulans", ambulansController.Update)
	ambulans.Delete("/:no_ambulans", ambulansController.Delete)
}
