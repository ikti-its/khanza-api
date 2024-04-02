package provider

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/akun"
	"github.com/ikti-its/khanza-api/internal/modules/auth"
	"github.com/ikti-its/khanza-api/internal/modules/file"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran"
	"github.com/ikti-its/khanza-api/internal/modules/pegawai"
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
