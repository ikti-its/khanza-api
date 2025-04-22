package resep

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/controller"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/repository"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/repository/postgres"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/router"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/usecase"
)

func ProvideResep(app *fiber.App, db *sqlx.DB) {
	// 🧾 Resep Obat setup
	var resepObatRepo repository.ResepObatRepository = postgres.NewResepObatRepository(db)
	resepObatUseCase := usecase.NewResepObatUseCase(resepObatRepo)
	resepObatController := controller.NewResepObatController(resepObatUseCase)

	// 💊 Resep Dokter setup
	var resepDokterRepo repository.ResepDokterRepository = postgres.NewResepDokterRepository(db)
	resepDokterUseCase := usecase.NewResepDokterUseCase(resepDokterRepo)
	resepDokterController := controller.NewResepDokterController(resepDokterUseCase)

	// Register both
	router.RegisterResepRoutes(app, resepObatController, resepDokterController)
}
