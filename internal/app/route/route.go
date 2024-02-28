package route

import (
	akunController "github.com/fathoor/simkes-api/internal/akun/controller"
	"github.com/fathoor/simkes-api/internal/app/middleware"
	authController "github.com/fathoor/simkes-api/internal/auth/controller"
	departemenController "github.com/fathoor/simkes-api/internal/departemen/controller"
	fileController "github.com/fathoor/simkes-api/internal/file/controller"
	jabatanController "github.com/fathoor/simkes-api/internal/jabatan/controller"
	pegawaiController "github.com/fathoor/simkes-api/internal/pegawai/controller"
	roleController "github.com/fathoor/simkes-api/internal/role/controller"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App                  *fiber.App
	AkunController       akunController.AkunController
	AuthController       authController.AuthController
	DepartemenController departemenController.DepartemenController
	FileController       fileController.FileController
	JabatanController    jabatanController.JabatanController
	PegawaiController    pegawaiController.PegawaiController
	RoleController       roleController.RoleController
}

func (r *Route) Setup() {
	akun := r.App.Group("/v1/akun", middleware.Authenticate("Public"))
	auth := r.App.Group("/v1/auth")
	departemen := r.App.Group("/v1/departemen", middleware.Authenticate("Admin"))
	file := r.App.Group("/v1/file", middleware.Authenticate("Public"))
	jabatan := r.App.Group("/v1/jabatan", middleware.Authenticate("Admin"))
	pegawai := r.App.Group("/v1/pegawai", middleware.Authenticate("Pegawai"))
	role := r.App.Group("/v1/role", middleware.Authenticate("Admin"))

	akun.Post("/", r.AkunController.Create, middleware.Authenticate("Admin"))
	akun.Get("/", r.AkunController.Get, middleware.Authenticate("Pegawai"))
	akun.Get("/:nip", r.AkunController.GetByNIP, middleware.Authenticate("Pegawai"))
	akun.Put("/:nip", r.AkunController.Update, middleware.Authenticate("Pegawai"), middleware.AuthorizeNIP())
	akun.Delete("/:nip", r.AkunController.Delete, middleware.Authenticate("Admin"))

	auth.Post("/", r.AuthController.Login)

	departemen.Post("/", r.DepartemenController.Create)
	departemen.Get("/", r.DepartemenController.GetAll)
	departemen.Get("/:departemen", r.DepartemenController.Get)
	departemen.Put("/:departemen", r.DepartemenController.Update)
	departemen.Delete("/:departemen", r.DepartemenController.Delete)

	file.Post("/", r.FileController.Upload)
	file.Get("/:filetype/:filename/download", r.FileController.Download)
	file.Get("/:filetype/:filename", r.FileController.View)
	file.Delete("/:filetype/:filename", r.FileController.Delete)

	jabatan.Post("/", r.JabatanController.Create)
	jabatan.Get("/", r.JabatanController.GetAll)
	jabatan.Get("/:jabatan", r.JabatanController.Get)
	jabatan.Put("/:jabatan", r.JabatanController.Update)
	jabatan.Delete("/:jabatan", r.JabatanController.Delete)

	pegawai.Post("/", r.PegawaiController.Create)
	pegawai.Get("/", r.PegawaiController.Get)
	pegawai.Get("/:nip", r.PegawaiController.GetByNIP)
	pegawai.Put("/:nip", r.PegawaiController.Update, middleware.AuthorizeNIP())
	pegawai.Delete("/:nip", r.PegawaiController.Delete, middleware.Authenticate("Admin"))

	role.Post("/", r.RoleController.Create)
	role.Get("/", r.RoleController.GetAll)
	role.Get("/:role", r.RoleController.Get)
	role.Put("/:role", r.RoleController.Update)
	role.Delete("/:role", r.RoleController.Delete)
}
