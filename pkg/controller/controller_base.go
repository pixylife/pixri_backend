package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"pixri_backend/config"
)

type JwtCustomClaims struct {
	Sub  string
	Auth string

	jwt.StandardClaims
}

var db *gorm.DB

func BaseController(env *config.Env) {

	db = env.DB
}
func Getclaims(c echo.Context) *JwtCustomClaims {

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JwtCustomClaims)
	return claims
}
