package model

import "github.com/jinzhu/gorm"

type Application struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	AgeGroup AgeGroup  `json:"age-group"`
	Purpose string  `json:"purpose"`
	BaseURL string`json:"baseURL"`
	Company string `json:"company"`

}

type AgeGroup struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func dbpreloadApplication(db *gorm.DB) *gorm.DB {
	return db.Preload("Application")
}

func AddApplication(db *gorm.DB, application *Application) {
	db.Create(&application)
}
func UpdateApplication(db *gorm.DB, application *Application) *Application {
	db.Save(&application)
	return application
}
func DeleteApplication(db *gorm.DB, application *Application) *Application {
	db.Delete(&application)
	return application
}
func FindApplication(db *gorm.DB, id int) *Application {
	application := Application{}
	dbpreloadApplication(db).First(&application, id)
	return &application
}
func FindAllApplication(db *gorm.DB) []*Application {
	var application []*Application
	dbpreloadApplication(db).Find(&application)
	return application
}
