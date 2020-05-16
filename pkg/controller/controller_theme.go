package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"net/http"
	"pixri_backend/pkg/model"
	"strconv"
)

const baseUrl = "http://localhost:5002/api/themes/generate"


func CreateTheme(c echo.Context) error {

	theme := model.Theme{}
	if error := c.Bind(&theme); error != nil {
		return error
	}
	model.AddTheme(db, &theme)
	app := model.FindApplication(db,theme.ApplicationID)
	theme.Application = *app
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

func FindApplicationThemes(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	theme := model.FindAllThemeForApplication(db,id)
	return c.JSON(http.StatusOK, theme)
}

func GenerateTheme(application model.Application) []model.Theme{
	client := resty.New()
	resp,_ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody("{\n        \"id\": 1,\n    " +
			"    \"CreatedAt\": \"2020-05-08T21:10:47+05:30\",\n    " +
			"    \"UpdatedAt\": \"2020-05-08T21:10:47+05:30\",\n      " +
			"  \"DeletedAt\": null,\n       " +
			" \"name\": \"Kinrgize\",\n       " +
			" \"type\": \"Health & Fitness\",\n  " +
			"      \"description\": \"Happy life\",\n     " +
			"   \"age-group\": {\n            \"min\": 0,\n            \"max\": 0\n        },\n       " +
			" \"purpose\": \"Kinrgize was developed to help people lead healthier lives using their smartphones.\",\n   " +
			"     \"baseURL\": \"127.0.0.1\",\n    " +
			"    \"company\": \"kinrgize\"\n    }").
		Post(baseUrl)


	if resp.StatusCode() == 200 {
		var data []model.Theme
		_ = json.Unmarshal([]byte(resp.String()), &data)

		fmt.Println(data)
		return data
	} else {
		return []model.Theme{}
	}
}





func ThemeController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/themes", CreateTheme)
	g.PUT(contextRoot+"/themes", UpdateTheme)
	g.DELETE(contextRoot+"/themes/:id", DeleteTheme)
	g.GET(contextRoot+"/themes/:id", FindTheme)
	g.GET(contextRoot+"/themes", FindAllTheme)
	g.GET(contextRoot+"/themes/app/:id", FindApplicationThemes)

}
