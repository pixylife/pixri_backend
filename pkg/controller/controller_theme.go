package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_backend/pkg/model"
	"strconv"
)

func CreateTheme(c echo.Context) error {

	theme := model.Theme{}
	if error := c.Bind(&theme); error != nil {
		return error
	}
	model.AddTheme(db, &theme)
	return c.JSON(http.StatusOK, theme)
}
func UpdateTheme(c echo.Context) error {

	theme := model.Theme{}
	if error := c.Bind(&theme); error != nil {
		return error
	}
	model.UpdateTheme(db, &theme)
	return c.JSON(http.StatusOK, theme)
}
func DeleteTheme(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	theme := model.FindTheme(db, id)
	model.DeleteTheme(db, theme)
	return c.JSON(http.StatusOK, theme)
}
func FindTheme(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	theme := model.FindTheme(db, id)
	return c.JSON(http.StatusOK, theme)
}
func FindAllTheme(c echo.Context) error {

	theme := model.FindAllTheme(db)
	return c.JSON(http.StatusOK, theme)
}
func ThemeController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/themes", CreateTheme)
	g.PUT(contextRoot+"/themes", UpdateTheme)
	g.DELETE(contextRoot+"/themes/:id", DeleteTheme)
	g.GET(contextRoot+"/themes/:id", FindTheme)
	g.GET(contextRoot+"/themes", FindAllTheme)
}
