package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_backend/pkg/model"
	"strconv"
)

func CreateUser(c echo.Context) error {

	user := model.User{}
	if error := c.Bind(&user); error != nil {
		return error
	}
	model.AddUser(db, &user)
	return c.JSON(http.StatusOK, user)
}
func UpdateUser(c echo.Context) error {

	user := model.User{}
	if error := c.Bind(&user); error != nil {
		return error
	}
	model.UpdateUser(db, &user)
	return c.JSON(http.StatusOK, user)
}
func DeleteUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	user := model.FindUser(db, id)
	model.DeleteUser(db, user)
	return c.JSON(http.StatusOK, user)
}
func FindUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	user := model.FindUser(db, id)
	return c.JSON(http.StatusOK, user)
}
func FindAllUser(c echo.Context) error {

	user := model.FindAllUser(db)
	return c.JSON(http.StatusOK, user)
}
func UserController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/users", CreateUser)
	g.PUT(contextRoot+"/users", UpdateUser)
	g.DELETE(contextRoot+"/users/:id", DeleteUser)
	g.GET(contextRoot+"/users/:id", FindUser)
	g.GET(contextRoot+"/users", FindAllUser)
}
