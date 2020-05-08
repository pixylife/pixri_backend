package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_backend/pkg/model"
	"strconv"
)

func CreateEntity(c echo.Context) error {

	entity := model.Entity{}
	if error := c.Bind(&entity); error != nil {
		return error
	}
	model.AddEntity(db, &entity)
	return c.JSON(http.StatusOK, entity)
}
func UpdateEntity(c echo.Context) error {

	entity := model.Entity{}
	if error := c.Bind(&entity); error != nil {
		return error
	}
	model.UpdateEntity(db, &entity)
	return c.JSON(http.StatusOK, entity)
}
func DeleteEntity(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	entity := model.FindEntity(db, id)
	model.DeleteEntity(db, entity)
	return c.JSON(http.StatusOK, entity)
}
func FindEntity(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	entity := model.FindEntity(db, id)
	return c.JSON(http.StatusOK, entity)
}
func FindAllEntity(c echo.Context) error {

	entity := model.FindAllEntity(db)
	return c.JSON(http.StatusOK, entity)
}
func EntityController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/entitys", CreateEntity)
	g.PUT(contextRoot+"/entitys", UpdateEntity)
	g.DELETE(contextRoot+"/entitys/:id", DeleteEntity)
	g.GET(contextRoot+"/entitys/:id", FindEntity)
	g.GET(contextRoot+"/entitys", FindAllEntity)
}
