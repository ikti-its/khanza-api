package provider

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-api/internal/app/config"
	"github.com/ikti-its/khanza-api/internal/modules/akun"
	"github.com/ikti-its/khanza-api/internal/modules/auth"
	"github.com/ikti-its/khanza-api/internal/modules/file"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran"
	"github.com/ikti-its/khanza-api/internal/modules/pegawai"
	"github.com/jmoiron/sqlx"
)

type Provider struct {
	App       *fiber.App
	Config    *config.Config
	PG        *sqlx.DB
	Validator *config.Validator
}

func (p *Provider) Provide() {
	akun.ProvideAkun(p.App, p.Config, p.PG, p.Validator)
	auth.ProvideAuth(p.App, p.Config, p.PG, p.Validator)
	file.ProvideFile(p.App, p.Config, p.Validator)
	pegawai.ProvidePegawai(p.App, p.PG, p.Validator)
	kehadiran.ProvideKehadiran(p.App, p.PG, p.Validator)
}
