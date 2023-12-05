package provider

import (
	"github.com/fathoor/simkes-api/module/akun/role/controller"
	"github.com/fathoor/simkes-api/module/akun/role/repository"
	"github.com/fathoor/simkes-api/module/akun/role/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProvideModule(app *fiber.App, db *gorm.DB) {
	ProvideRole(app, db)
}

func ProvideRole(app *fiber.App, db *gorm.DB) {
	roleRepository := repository.ProvideRoleRepository(db)
	roleService := service.ProvideRoleService(&roleRepository)
	roleController := controller.ProvideRoleController(&roleService)

	roleController.Route(app)
}
