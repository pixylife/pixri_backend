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

const generatorUrl = "http://localhost:5003/api/generate"


type GenRequest struct {
	Application model.Application
	Entity []GenEntity
	Theme model.Theme
}


type GenEntity struct {
	Entity model.Entity
	Fields []*model.Field
}


func GenerateApplication(c echo.Context)  error{
	var request = GenRequest{}
	var entityRequest []GenEntity


	id, _ := strconv.Atoi(c.Param("id"))
	application := model.FindApplication(db, id)
	if application.ThemeID!=0 {
		theme := model.FindTheme(db, application.ThemeID)
		request.Theme = *theme

	}else {
		theme := model.GetDefaultTheme()
		request.Theme = *theme

	}



	request.Application = *application




	 entities := model.FindAllEntityByApplication(db,application.ID)

	 for _,entity := range entities {
	 	fields := model.FindAllFieldsByEntity(db,entity.ID)
		 genEntity := GenEntity{*entity, fields}
		 entityRequest = append(entityRequest,genEntity)
	 }

	 request.Entity = entityRequest


	 fmt.Println(request)

	client := resty.New()
	resp,_ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(generatorUrl)

	if resp.StatusCode() == 200 {
		var data model.GenResponse
		_ = json.Unmarshal([]byte(resp.String()), &data)

		application.GitHubUrl = data.GithubURL
		model.UpdateApplication(db,application)

		return c.JSON(http.StatusOK, data)
	} else {
		return c.JSON(http.StatusBadRequest, "failed")
	}


}



func GenerationController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/applications/generate/:id", GenerateApplication)

}
