package main

import (

	"./pkg/model"
	"./pkg/controller"
	"./config"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	jwtConfig := middleware.JWTConfig{
		Claims:     &controller.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(jwtConfig))

	controller.BaseController(env)
	controller.LoginController(e, "/api")
	controller.AccountController(r, "api")
	controller.UserController(r, "api")
}
