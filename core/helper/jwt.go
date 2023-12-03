package helper

import (
	"github.com/fathoor/simkes-api/core/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(nip string, role int) (string, error) {
	var (
		cfg       = config.ProvideConfig()
		exp       = cfg.GetInt("JWT_EXPIRE")
		jwtSecret = cfg.Get("JWT_SECRET")
		jwtExpire = time.Now().Add(time.Hour * time.Duration(exp))
	)

	claims := jwt.MapClaims{
		"nip":  nip,
		"role": role,
		"exp":  jwtExpire.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}
