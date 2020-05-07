package controller

import (
	"github.com/jinzhu/gorm"
	"pixri_backend/config"
)

var db *gorm.DB

func BaseController(env *config.Env) {
	db = env.DB
}

