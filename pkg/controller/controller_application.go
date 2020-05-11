package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_backend/pkg/model"
	"strconv"
)

func CreateApplication(c echo.Context) error {

	application := model.Application{}
	if error := c.Bind(&application); error != nil {
		return error
	}
	model.AddApplication(db, &application)
	return c.JSON(http.StatusOK, application)
}
func UpdateApplication(c echo.Context) error {

	application := model.Application{}
	if error := c.Bind(&application); error != nil {
		return error
	}
	model.UpdateApplication(db, &application)
	return c.JSON(http.StatusOK, application)
}
func DeleteApplication(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	application := model.FindApplication(db, id)
	model.DeleteApplication(db, application)
	return c.JSON(http.StatusOK, application)
}
func FindApplication(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	application := model.FindApplication(db, id)
	return c.JSON(http.StatusOK, application)
}
func FindAllApplication(c echo.Context) error {

	application := model.FindAllApplication(db)
	return c.JSON(http.StatusOK, application)
}

func GetAppDataCount(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	application := model.FindApplication(db, id)
	themeCount := model.GetThemeCount(db,application.ID)
	 appData := model.ApplicationData{ID: application.ID, ThemeCount: themeCount}
	return c.JSON(http.StatusOK, appData)
}


func ApplicationController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/applications", CreateApplication)
	g.PUT(contextRoot+"/applications", UpdateApplication)
	g.DELETE(contextRoot+"/applications/:id", DeleteApplication)
	g.GET(contextRoot+"/applications/:id", FindApplication)
	g.GET(contextRoot+"/applications", FindAllApplication)
	g.GET(contextRoot+"/applications/info/:id", GetAppDataCount)

}

