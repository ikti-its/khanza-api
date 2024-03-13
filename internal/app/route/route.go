package route

import (
	akunController "github.com/fathoor/simkes-api/internal/akun/controller"
	"github.com/fathoor/simkes-api/internal/app/middleware"
	authController "github.com/fathoor/simkes-api/internal/auth/controller"
	cutiController "github.com/fathoor/simkes-api/internal/cuti/controller"
	departemenController "github.com/fathoor/simkes-api/internal/departemen/controller"
	fileController "github.com/fathoor/simkes-api/internal/file/controller"
	jabatanController "github.com/fathoor/simkes-api/internal/jabatan/controller"
	jadwalPegawaiController "github.com/fathoor/simkes-api/internal/jadwal-pegawai/controller"
	kehadiranController "github.com/fathoor/simkes-api/internal/kehadiran/controller"
	pegawaiController "github.com/fathoor/simkes-api/internal/pegawai/controller"
	roleController "github.com/fathoor/simkes-api/internal/role/controller"
	shiftController "github.com/fathoor/simkes-api/internal/shift/controller"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App                     *fiber.App
	AkunController          akunController.AkunController
	AuthController          authController.AuthController
	CutiController          cutiController.CutiController
	DepartemenController    departemenController.DepartemenController
	FileController          fileController.FileController
	JabatanController       jabatanController.JabatanController
	JadwalPegawaiController jadwalPegawaiController.JadwalPegawaiController
	KehadiranController     kehadiranController.KehadiranController
	PegawaiController       pegawaiController.PegawaiController
	RoleController          roleController.RoleController
	ShiftController         shiftController.ShiftController
}

func (r *Route) Setup() {
	akun := r.App.Group("/v1/akun", middleware.Authenticate("Public"))
	auth := r.App.Group("/v1/auth")
	cuti := r.App.Group("/v1/cuti", middleware.Authenticate("Pegawai"))
	departemen := r.App.Group("/v1/departemen", middleware.Authenticate("Admin"))
	file := r.App.Group("/v1/file", middleware.Authenticate("Public"))
	jabatan := r.App.Group("/v1/jabatan", middleware.Authenticate("Admin"))
	jadwalPegawai := r.App.Group("/v1/jadwal-pegawai", middleware.Authenticate("Admin"))
	kehadiran := r.App.Group("/v1/kehadiran", middleware.Authenticate("Admin"))
	pegawai := r.App.Group("/v1/pegawai", middleware.Authenticate("Pegawai"))
	role := r.App.Group("/v1/role", middleware.Authenticate("Admin"))
	shift := r.App.Group("/v1/shift", middleware.Authenticate("Admin"))

	akun.Post("/", r.AkunController.Create, middleware.Authenticate("Admin"))
	akun.Get("/", r.AkunController.Get, middleware.Authenticate("Pegawai"))
	akun.Get("/:nip", r.AkunController.GetByNIP, middleware.Authenticate("Pegawai"))
	akun.Put("/:nip", r.AkunController.Update, middleware.Authenticate("Pegawai"), middleware.AuthorizeNIP())
	akun.Delete("/:nip", r.AkunController.Delete, middleware.Authenticate("Admin"))

	auth.Post("/login", r.AuthController.Login)

	cuti.Post("/", r.CutiController.Create)
	cuti.Get("/", r.CutiController.Get)
	cuti.Get("/:id", r.CutiController.GetByID)
	cuti.Put("/:id", r.CutiController.Update)
	cuti.Delete("/:id", r.CutiController.Delete)

	departemen.Post("/", r.DepartemenController.Create)
	departemen.Get("/", r.DepartemenController.Get)
	departemen.Get("/:departemen", r.DepartemenController.GetByNama)
	departemen.Put("/:departemen", r.DepartemenController.Update)
	departemen.Delete("/:departemen", r.DepartemenController.Delete)

	file.Post("/", r.FileController.Upload)
	file.Get("/:filetype/:filename/download", r.FileController.Download)
	file.Get("/:filetype/:filename", r.FileController.View)
	file.Delete("/:filetype/:filename", r.FileController.Delete)

	jabatan.Post("/", r.JabatanController.Create)
	jabatan.Get("/", r.JabatanController.Get)
	jabatan.Get("/:jabatan", r.JabatanController.GetByNama)
	jabatan.Put("/:jabatan", r.JabatanController.Update)
	jabatan.Delete("/:jabatan", r.JabatanController.Delete)

	jadwalPegawai.Post("/", r.JadwalPegawaiController.Create)
	jadwalPegawai.Get("/", r.JadwalPegawaiController.Get)
	jadwalPegawai.Get("/:tahun/:bulan/:hari/:nip", r.JadwalPegawaiController.GetByPK)
	jadwalPegawai.Put("/:tahun/:bulan/:hari/:nip", r.JadwalPegawaiController.Update)
	jadwalPegawai.Delete("/:tahun/:bulan/:hari/:nip", r.JadwalPegawaiController.Delete)

	kehadiran.Post("/checkin", r.KehadiranController.CheckIn)
	kehadiran.Post("/checkout", r.KehadiranController.CheckOut)
	kehadiran.Get("/", r.KehadiranController.Get)
	kehadiran.Get("/:id", r.KehadiranController.GetByID)
	kehadiran.Put("/:id", r.KehadiranController.Update)
	kehadiran.Delete("/:id", r.KehadiranController.Delete)

	pegawai.Post("/", r.PegawaiController.Create)
	pegawai.Get("/", r.PegawaiController.Get)
	pegawai.Get("/:nip", r.PegawaiController.GetByNIP)
	pegawai.Put("/:nip", r.PegawaiController.Update, middleware.AuthorizeNIP())
	pegawai.Delete("/:nip", r.PegawaiController.Delete, middleware.Authenticate("Admin"))

	role.Post("/", r.RoleController.Create)
	role.Get("/", r.RoleController.Get)
	role.Get("/:role", r.RoleController.GetByNama)
	role.Put("/:role", r.RoleController.Update)
	role.Delete("/:role", r.RoleController.Delete)

	shift.Post("/", r.ShiftController.Create)
	shift.Get("/", r.ShiftController.Get)
	shift.Get("/:shift", r.ShiftController.GetByNama)
	shift.Put("/:shift", r.ShiftController.Update)
}
