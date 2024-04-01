package provider

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/modules/akun"
	"github.com/fathoor/simkes-api/internal/modules/auth"
	"github.com/fathoor/simkes-api/internal/modules/file"
	"github.com/fathoor/simkes-api/internal/modules/kehadiran"
	"github.com/fathoor/simkes-api/internal/modules/pegawai"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Provider struct {
	App       *fiber.App
	Config    *config.Config
	DB        *gorm.DB
	Validator *config.Validator
}

func (p *Provider) Provide() {
	akun.ProvideAkun(p.App, p.Config, p.DB, p.Validator)
	auth.ProvideAuth(p.App, p.Config, p.DB, p.Validator)
	file.ProvideFile(p.App, p.Config, p.Validator)
	pegawai.ProvidePegawai(p.App, p.DB, p.Validator)
	kehadiran.ProvideKehadiran(p.App, p.DB, p.Validator)
}
