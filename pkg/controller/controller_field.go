package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_backend/pkg/model"
	"strconv"
)

func CreateField(c echo.Context) error {

	field := model.Field{}
	if error := c.Bind(&field); error != nil {
		return error
	}
	model.AddField(db, &field)
	return c.JSON(http.StatusOK, field)
}
func UpdateField(c echo.Context) error {

	field := model.Field{}
	if error := c.Bind(&field); error != nil {
		return error
	}
	model.UpdateField(db, &field)
	return c.JSON(http.StatusOK, field)
}
func DeleteField(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	field := model.FindField(db, id)
	model.DeleteField(db, field)
	return c.JSON(http.StatusOK, field)
}
func FindField(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	field := model.FindField(db, id)
	return c.JSON(http.StatusOK, field)
}
func FindAllField(c echo.Context) error {

	field := model.FindAllField(db)
	return c.JSON(http.StatusOK, field)
}
func FieldController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/fields", CreateField)
	g.PUT(contextRoot+"/fields", UpdateField)
	g.DELETE(contextRoot+"/fields/:id", DeleteField)
	g.GET(contextRoot+"/fields/:id", FindField)
	g.GET(contextRoot+"/fields", FindAllField)
}
