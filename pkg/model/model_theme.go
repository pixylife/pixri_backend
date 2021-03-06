package model

import "github.com/jinzhu/gorm"

type Theme struct {
	Model
	PrimaryColor     string `json:"primary_color"`
	SecondaryColor   string `json:"secondary_color"`
	PrimaryDarkColor string `json:"primary_dark_color"`
	BodyColor string `json:"body_color"`
	TextColorBody string `json:"text_color_body"`
	TextColorAppBar string `json:"text_color_appBar"`
	Application   Application `gorm:"foreignkey:application_id"  json:"-"`
	ApplicationID int `json:"application_id"`
}

func dbpreloadTheme(db *gorm.DB) *gorm.DB {
	return db.Preload("Theme")
}

func AddTheme(db *gorm.DB, theme *Theme) {
	db.Create(&theme)
}
func UpdateTheme(db *gorm.DB, theme *Theme) *Theme {
	db.Save(&theme)
	return theme
}
func DeleteTheme(db *gorm.DB, theme *Theme) *Theme {
	db.Delete(&theme)
	return theme
}
func FindTheme(db *gorm.DB, id int) *Theme {
	theme := Theme{}
	dbpreloadTheme(db).First(&theme, id)
	return &theme
}
func FindAllTheme(db *gorm.DB) []*Theme {
	var theme []*Theme
	dbpreloadTheme(db).Find(&theme)
	return theme
}

func GetThemeCount(db *gorm.DB, application_id int) int{
	var theme []*Theme
	var count int
	db.Where("application_id = ?", application_id).Find(&theme).Count(&count)
	return count
}

func FindAllThemeForApplication(db *gorm.DB, application_id int) []*Theme {
	var theme []*Theme
	dbpreloadTheme(db).Where("application_id = ?", application_id).Find(&theme)
	return theme
}

func DeleteAllThemesByApplication(db *gorm.DB,application_id int) []*Theme {
	var entity []*Theme
	db.Where("application_id = ?", application_id).Delete(&entity)
	return entity
}


func GetDefaultTheme() *Theme{
	var theme= Theme{}
	theme.ID = 0
	theme.PrimaryColor = "#008577"
	theme.PrimaryDarkColor = "#00574b"
	theme.SecondaryColor = "#d81b60"
	theme.BodyColor = "#ffffff"
	theme.TextColorAppBar = "#ffffff"
	theme.TextColorBody = "#000000"
	return &theme
}