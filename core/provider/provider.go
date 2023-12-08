package provider

import (
	akunController "github.com/fathoor/simkes-api/module/akun/akun/controller"
	akunRepository "github.com/fathoor/simkes-api/module/akun/akun/repository"
	akunService "github.com/fathoor/simkes-api/module/akun/akun/service"
	roleController "github.com/fathoor/simkes-api/module/akun/role/controller"
	roleRepository "github.com/fathoor/simkes-api/module/akun/role/repository"
	roleService "github.com/fathoor/simkes-api/module/akun/role/service"
	authController "github.com/fathoor/simkes-api/module/auth/controller"
	authService "github.com/fathoor/simkes-api/module/auth/service"
	fileController "github.com/fathoor/simkes-api/module/file/controller"
	fileService "github.com/fathoor/simkes-api/module/file/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProvideModule(app *fiber.App, db *gorm.DB) {
	ProvideRole(app, db)
	ProvideAkun(app, db)
	ProvideAuth(app, db)
	ProvideFile(app)
}

func ProvideRole(app *fiber.App, db *gorm.DB) {
	repository := roleRepository.ProvideRoleRepository(db)
	service := roleService.ProvideRoleService(&repository)
	controller := roleController.ProvideRoleController(&service)

	controller.Route(app)
}

func ProvideAkun(app *fiber.App, db *gorm.DB) {
	repository := akunRepository.ProvideAkunRepository(db)
	service := akunService.ProvideAkunService(&repository)
	controller := akunController.ProvideAkunController(&service)

	controller.Route(app)
}

func ProvideAuth(app *fiber.App, db *gorm.DB) {
	repository := akunRepository.ProvideAkunRepository(db)
	service := authService.ProvideAuthService(&repository)
	controller := authController.ProvideAuthController(&service)

	controller.Route(app)
}

func ProvideFile(app *fiber.App) {
	service := fileService.ProvideFileService()
	controller := fileController.ProvideFileController(&service)

	controller.Route(app)
}
