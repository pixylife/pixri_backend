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
	application.ThemeID = 0
	model.AddApplication(db, &application)
	themes := GenerateTheme(application)
	model.DeleteAllThemesByApplication(db,application.ID)
	for _,theme := range themes {
		model.AddTheme(db, &theme)
	}
	return c.JSON(http.StatusOK, application)
}
func UpdateApplication(c echo.Context) error {

	application := model.Application{}
	if error := c.Bind(&application); error != nil {
		return error
	}

	application.ThemeID = 0
	application.GitHubUrl = ""
	model.UpdateApplication(db, &application)
	themes := GenerateTheme(application)
	model.DeleteAllThemesByApplication(db,application.ID)
	for _,theme := range themes {
		model.AddTheme(db, &theme)
	}	
	return c.JSON(http.StatusOK, application)
}
func DeleteApplication(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	application := model.FindApplication(db, id)
	model.DeleteApplication(db, application)
	model.DeleteAllThemesByApplication(db,application.ID)
	return c.JSON(http.StatusOK, application)
}
func FindApplication(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	application := model.FindApplication(db, id)
	if application.ThemeID != 0 {
		application.Theme = model.FindTheme(db, application.ThemeID)
	}else{
		application.Theme = model.GetDefaultTheme()
	}
	return c.JSON(http.StatusOK, application)
}
func FindAllApplication(c echo.Context) error {
	var applicationsList []model.Application
	application := model.FindAllApplication(db)
	for _,app := range application{
		if app.ThemeID != 0 {
			app.Theme = model.FindTheme(db, app.ThemeID)
		}else{
			app.Theme = model.GetDefaultTheme()
		}
		applicationsList = append(applicationsList, *app)
	}

	return c.JSON(http.StatusOK, applicationsList)
}

func GetAppDataCount(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	application := model.FindApplication(db, id)
	themeCount := model.GetThemeCount(db,application.ID)
	entityCount := model.GetEntityCount(db,application.ID)
	 appData := model.ApplicationData{ID: application.ID, ThemeCount: themeCount,EntityCount:entityCount}
	return c.JSON(http.StatusOK, appData)
}

func SelectApplicationTheme(c echo.Context) error {
	application := model.Application{}
	if error := c.Bind(&application); error != nil {
		return error
	}
	model.UpdateApplication(db, &application)
	return c.JSON(http.StatusOK, application)
}


func ApplicationController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/applications", CreateApplication)
	g.PUT(contextRoot+"/applications", UpdateApplication)
	g.DELETE(contextRoot+"/applications/:id", DeleteApplication)
	g.GET(contextRoot+"/applications/:id", FindApplication)
	g.GET(contextRoot+"/applications", FindAllApplication)
	g.GET(contextRoot+"/applications/info/:id", GetAppDataCount)
	g.PUT(contextRoot+"/applications/theme", SelectApplicationTheme)


}

