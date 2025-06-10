package controller

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/exception"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/usecase"
)

type AmbulansController struct {
	UseCase *usecase.AmbulansUseCase
}

func NewAmbulansController(useCase *usecase.AmbulansUseCase) *AmbulansController {
	return &AmbulansController{
		UseCase: useCase,
	}
}

func (c *AmbulansController) Create(ctx *fiber.Ctx) error {
	var request model.AmbulansRequest
	fmt.Println("Received a POST request to /ambulans") // Debug log

	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println("Failed to parse request body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid request body",
		})
	}

	response, err := c.UseCase.Create(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "OK",
		Data:   response,
	})
}

func (c *AmbulansController) GetAll(ctx *fiber.Ctx) error {
	response, err := c.UseCase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}
	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *AmbulansController) GetByNoAmbulans(ctx *fiber.Ctx) error {
	noAmbulans := ctx.Params("no_ambulans")
	response, err := c.UseCase.GetByNoAmbulans(noAmbulans)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.Response{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   err.Error(),
		})
	}
	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *AmbulansController) Update(ctx *fiber.Ctx) error {
	noAmbulans := ctx.Params("no_ambulans")
	var request model.AmbulansRequest

	if err := ctx.BodyParser(&request); err != nil {
		panic(&exception.BadRequestError{Message: "Invalid request body"})
	}

	response, err := c.UseCase.Update(noAmbulans, &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}
	return ctx.JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *AmbulansController) Delete(ctx *fiber.Ctx) error {
	noAmbulans := ctx.Params("no_ambulans")
	err := c.UseCase.Delete(noAmbulans)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(web.Response{
		Code:   fiber.StatusNoContent,
		Status: "Deleted",
	})
}

func (c *AmbulansController) RequestAmbulans(ctx *fiber.Ctx) error {
	var req entity.Ambulans
	fmt.Println("🚨 Ambulance request hit")

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.AmbulansResponse{
			Code: fiber.StatusBadRequest,
			Data: "Invalid JSON body",
		})
	}

	if req.NoAmbulans == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.AmbulansResponse{
			Code: fiber.StatusBadRequest,
			Data: "no_ambulans is required",
		})
	}

	err := c.UseCase.Notify(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.AmbulansResponse{
			Code: fiber.StatusInternalServerError,
			Data: err.Error(),
		})
	}

	// Return full ambulance + status response
	return ctx.Status(fiber.StatusOK).JSON(model.AmbulansResponse{
		Code:       fiber.StatusOK,
		NoAmbulans: req.NoAmbulans,
		Status:     req.Status,
		Supir:      req.Supir,
		Data:       "Permintaan ambulans berhasil dikirim",
	})
}

func (c *AmbulansController) GetPendingRequests(ctx *fiber.Ctx) error {
	data, err := c.UseCase.GetPendingRequests()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   fiber.StatusInternalServerError,
			"status": "Failed",
			"data":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   data,
	})

}

func (c *AmbulansController) AcceptAmbulansRequest(ctx *fiber.Ctx) error {
	noAmbulans := ctx.Params("no_ambulans")

	err := c.UseCase.MarkRequestAccepted(noAmbulans)
	if err != nil {
		log.Println("[AcceptAmbulansRequest] DB error:", err) // add this
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   500,
			"status": "Error",
			"data":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   200,
		"status": "OK",
		"data":   "Ambulans request marked as accepted",
	})
}

func (c *AmbulansController) UpdateStatus(ctx *fiber.Ctx) error {
	type Payload struct {
		NoAmbulans string `json:"no_ambulans"`
		Status     string `json:"status"`
		NomorRujuk string `json:"nomor_rujuk"` // optional, for logging
	}

	var body Payload
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}

	// update in DB
	err := c.UseCase.UpdateStatus(body.NoAmbulans, body.Status)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update status",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Status updated successfully",
	})
}
