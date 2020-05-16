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

	fmt.Println("Start")

	id, _ := strconv.Atoi(c.Param("id"))
	application := model.FindApplication(db, id)
	theme := model.FindTheme(db,application.ThemeID)

	fmt.Println("Start A")


	request.Application = *application
	request.Theme = *theme

	fmt.Println("Start B")



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

	data := GenRequest{}
	_ = json.Unmarshal([]byte(resp.String()), &data)

	return c.JSON(http.StatusOK, data)
}



func GenerationController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/applications/generate/:id", GenerateApplication)

}
