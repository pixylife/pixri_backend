package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	Db string
	User string
	Password string
	Url string
	Port string
}

func main(){
	content, _ := ioutil.ReadFile("properties.yml")
	p := Property{}
	err := yaml.Unmarshal([]byte(content), &p)
	db, err := gorm.Open("mysql", p.User+":"+p.Password+"@tcp("+p.Url+":"+p.Port+")/"+p.Db+"?charset=utf8&parseTime=True&loc=Local")
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
	controller.ThemeController(r, "api")
	controller.EntityController(r, "api")
	controller.FieldController(r, "api")

	e.Logger.Fatal(e.Start(":5001"))

}
