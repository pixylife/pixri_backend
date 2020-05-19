package controller

import (
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
	theme := model.FindTheme(db,application.ThemeID)



	request.Application = *application
	request.Theme = *theme




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

	fmt.Println(resp.StatusCode())
	fmt.Println(resp.String())


	return c.JSON(http.StatusOK, resp.String())
}



func GenerationController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/applications/generate/:id", GenerateApplication)

}
