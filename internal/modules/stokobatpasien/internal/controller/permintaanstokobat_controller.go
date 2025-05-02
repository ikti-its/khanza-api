package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/exception"
	web "github.com/ikti-its/khanza-api/internal/app/model"
	"github.com/ikti-its/khanza-api/internal/modules/stokobatpasien/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/stokobatpasien/internal/usecase"
)

type PermintaanStokObatController struct {
	UseCase *usecase.PermintaanStokObatUseCase
}

func NewPermintaanStokObatController(useCase *usecase.PermintaanStokObatUseCase) *PermintaanStokObatController {
	return &PermintaanStokObatController{
		UseCase: useCase,
	}
}

func (c *PermintaanStokObatController) Create(ctx *fiber.Ctx) error {
	var request model.PermintaanStokObatRequest
	fmt.Println("📥 Received POST /permintaan-stok-obat")

	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println("❌ Error parsing body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid request body",
		})
	}

	response, err := c.UseCase.Create(&request)
	if err != nil {
		fmt.Println("❌ Error in usecase.Create():", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *PermintaanStokObatController) GetAll(ctx *fiber.Ctx) error {
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

func (c *PermintaanStokObatController) GetByNoPermintaan(ctx *fiber.Ctx) error {
	noPermintaan := ctx.Params("no_permintaan")

	response, err := c.UseCase.GetByNoPermintaan(noPermintaan)
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

func (c *PermintaanStokObatController) Update(ctx *fiber.Ctx) error {
	var request model.PermintaanStokObatRequest

	if err := ctx.BodyParser(&request); err != nil {
		panic(&exception.BadRequestError{Message: "Invalid request body"})
	}

	response, err := c.UseCase.Update(&request)
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

func (c *PermintaanStokObatController) Delete(ctx *fiber.Ctx) error {
	noPermintaan := ctx.Params("no_permintaan")

	if noPermintaan == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   400,
			"status": "Bad Request",
			"data":   "no_permintaan is required",
		})
	}

	err := c.UseCase.Delete(noPermintaan)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   500,
			"status": "Error",
			"data":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   200,
		"status": "Success",
		"data":   "Permintaan stok obat deleted successfully",
	})
}

func (c *PermintaanStokObatController) GetByNomorRawat(ctx *fiber.Ctx) error {
	nomorRawat := ctx.Params("nomor_rawat")
	permintaans, err := c.UseCase.GetByNomorRawat(nomorRawat)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":   404,
			"status": "Not Found",
			"data":   fmt.Sprintf("permintaan_stok_obat not found: %v", err),
		})
	}
	return ctx.JSON(fiber.Map{
		"code":   200,
		"status": "OK",
		"data":   permintaans,
	})
}

func (c *PermintaanStokObatController) UpdateValidasi(ctx *fiber.Ctx) error {
	noPermintaan := ctx.Params("no_permintaan")

	var payload struct {
		TglValidasi string `json:"tgl_validasi"`
		JamValidasi string `json:"jam_validasi"`
	}

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "invalid JSON",
		})
	}

	err := c.UseCase.UpdateValidasi(ctx.Context(), noPermintaan, payload.TglValidasi, payload.JamValidasi)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success", "message": "Validasi updated",
	})
}

func (c *PermintaanStokObatController) CreateWithDetail(ctx *fiber.Ctx) error {
	// 1. Use your existing model.PermintaanStokObatRequest
	//    which already has a StokObat []StokObatPasienRequest field
	var req model.PermintaanStokObatRequest
	details := make([]model.StokObatPasienRequest, len(req.StokObat))
	for i, d := range req.StokObat {
		details[i] = model.StokObatPasienRequest{
			Tanggal:     req.TglPermintaan,
			Jam:         req.Jam,
			NoRawat:     req.NoRawat,
			KodeBrng:    d.KodeBarang,
			Jumlah:      float64(d.Jumlah),
			KdBangsal:   d.KdBangsal, // only if you added this in StokObatRequest
			NoBatch:     d.NoBatch,
			NoFaktur:    d.NoFaktur,
			AturanPakai: d.AturanPakai,
			// optionally set Jam00~Jam23 based on d.JamObat here
		}
	}

	fmt.Println("📥 Received POST /permintaan-stok-obat/detail")

	// 2. Parse the flat JSON directly into it
	if err := ctx.BodyParser(&req); err != nil {
		fmt.Println("❌ Error parsing body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(web.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid JSON body",
		})
	}

	// 3. Now req.TglPermintaan, req.Jam, and req.StokObat are populated
	if err := c.UseCase.CreateWithDetail(ctx.Context(), c.UseCase.DB, &req, details); err != nil {
		fmt.Println("❌ Error inserting data:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
			Code:   fiber.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   "Permintaan dan stok obat berhasil disimpan.",
	})
}
