package helper

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func GenerateJWT(userId uuid.UUID, role int, cfg *config.Config) (string, error) {
	var (
		expire = cfg.GetInt("JWT_EXPIRE", 24)
		secret = cfg.Get("JWT_SECRET")
		iat    = time.Now().Unix()
		exp    = time.Now().Add(time.Hour * time.Duration(expire)).Unix()
	)

	sub := userId.String()
	claims := jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"iat":  iat,
		"exp":  exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(secret))
}
