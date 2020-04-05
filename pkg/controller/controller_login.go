package controller

import (
	"../model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

type Payload struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberme"`
}

func LoginController(g *echo.Echo, contextRoot string) {

	g.POST(contextRoot+"/authenticate", login)
}
func login(c echo.Context) error {

	loginRequest := new(Payload)
	er1 := c.Bind(loginRequest)

	if er1 != nil {
		return er1
	}
	users := model.FindAllUser(db)
	var targetUser *model.User
	if len(users) == 0 {
		targetUser = &model.User{Username: "admin", Password: "admin", Firstname: "Administrator", Lastname: "Administrator"}
		model.AddUser(db, targetUser)
	}
	for _, user := range users {
		if strings.EqualFold(user.Username, loginRequest.Username) && strings.EqualFold(user.Password, loginRequest.Password) {
			targetUser = user
		}
	}
	if targetUser == nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	claims := JwtCustomClaims{"admin", "ROLE_ADMIN,ROLE_USER", jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	c.Response().Header().Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", t))
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Origin, X-Requested-With, Content-Type, Accept")

	if err != nil {
		return err
	}
	targetUser.Token = t
	return c.JSON(http.StatusOK, targetUser)
}
