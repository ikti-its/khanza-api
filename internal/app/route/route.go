package route

import (
	akunController "github.com/fathoor/simkes-api/internal/akun/controller"
	"github.com/fathoor/simkes-api/internal/app/middleware"
	authController "github.com/fathoor/simkes-api/internal/auth/controller"
	fileController "github.com/fathoor/simkes-api/internal/file/controller"
	roleController "github.com/fathoor/simkes-api/internal/role/controller"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App            *fiber.App
	AkunController akunController.AkunController
	AuthController authController.AuthController
	FileController fileController.FileController
	RoleController roleController.RoleController
}

func (r *Route) Setup() {
	akun := r.App.Group("/v1/akun")
	auth := r.App.Group("/v1/auth")
	role := r.App.Group("/v1/role")
	file := r.App.Group("/v1/file", middleware.Authenticate("Public"))

	akun.Post("/", r.AkunController.Create, middleware.Authenticate("Admin"))
	akun.Get("/", r.AkunController.Get, middleware.Authenticate("Pegawai"))
	akun.Get("/:nip", r.AkunController.GetByNIP, middleware.Authenticate("Pegawai"))
	akun.Put("/:nip", r.AkunController.Update, middleware.Authenticate("Pegawai"), middleware.AuthorizeNIP())
	akun.Delete("/:nip", r.AkunController.Delete, middleware.Authenticate("Admin"))

	auth.Post("/", r.AuthController.Login)

	role.Post("/", r.RoleController.Create)
	role.Get("/", r.RoleController.GetAll)
	role.Get("/:role", r.RoleController.Get)
	role.Put("/:role", r.RoleController.Update)
	role.Delete("/:role", r.RoleController.Delete)

	file.Post("/", r.FileController.Upload)
	file.Get("/:filetype/:filename/download", r.FileController.Download)
	file.Get("/:filetype/:filename", r.FileController.View)
	file.Delete("/:filetype/:filename", r.FileController.Delete)
}
