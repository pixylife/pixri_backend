package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"pixri_backend/config"
	"pixri_backend/pkg/controller"
	"pixri_backend/pkg/model"
)
type Property struct {
	Dburl string
}

func main(){
	content, _ := ioutil.ReadFile("properties.yml")
	p := Property{}
	err := yaml.Unmarshal([]byte(content), &p)
	db, err := gorm.Open("sqlite3", p.Dburl)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	model.InitModels(db)
	env := &config.Env{DB: db}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")
	controller.BaseController(env)
	controller.UserController(r, "api")
	controller.ApplicationController(r, "api")

	e.Logger.Fatal(e.Start(":1235"))

}
