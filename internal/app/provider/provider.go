package provider

import (
	akunController "github.com/fathoor/simkes-api/internal/akun/controller"
	akunRepository "github.com/fathoor/simkes-api/internal/akun/repository"
	akunService "github.com/fathoor/simkes-api/internal/akun/service"
	"github.com/fathoor/simkes-api/internal/app/route"
	authController "github.com/fathoor/simkes-api/internal/auth/controller"
	authService "github.com/fathoor/simkes-api/internal/auth/service"
	fileController "github.com/fathoor/simkes-api/internal/file/controller"
	fileService "github.com/fathoor/simkes-api/internal/file/service"
	roleController "github.com/fathoor/simkes-api/internal/role/controller"
	roleRepository "github.com/fathoor/simkes-api/internal/role/repository"
	roleService "github.com/fathoor/simkes-api/internal/role/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Provider struct {
	App *fiber.App
	DB  *gorm.DB
}

func (p *Provider) Provide() {
	repositoryAkun := akunRepository.NewAkunRepositoryProvider(p.DB)
	serviceAkun := akunService.NewAkunServiceProvider(&repositoryAkun)
	controllerAkun := akunController.NewAkunControllerProvider(&serviceAkun)

	serviceAuth := authService.NewAuthServiceProvider(&repositoryAkun)
	controllerAuth := authController.NewAuthControllerProvider(&serviceAuth)

	serviceFile := fileService.NewFileServiceProvider()
	controllerFile := fileController.NewFileControllerProvider(&serviceFile)

	repositoryRole := roleRepository.NewRoleRepositoryProvider(p.DB)
	serviceRole := roleService.NewRoleServiceProvider(&repositoryRole)
	controllerRole := roleController.NewRoleControllerProvider(&serviceRole)

	router := route.Route{
		App:            p.App,
		AkunController: controllerAkun,
		AuthController: controllerAuth,
		FileController: controllerFile,
		RoleController: controllerRole,
	}

	router.Setup()
}
