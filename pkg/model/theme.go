package model

import "github.com/jinzhu/gorm"

type Theme struct {
	PrimaryColor     string `json:"primary_color"`
	SecondaryColor   string `json:"secondary_color"`
	PrimaryDarkColor string `json:"primary_dark_color"`
	BodyColor string `json:"body_color"`
	TextColorBody string `json:"text_color_body"`
	TextColorAppBar string `json:"text_color_appBar"`
	Application   Application `gorm:"foreignkey:application_id" json:"application"`
	ApplicationID int
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

func FindAllThemeForApplication(db *gorm.DB, application_id int) []*Theme {
	var theme []*Theme
	dbpreloadTheme(db).Where("application_id = ?", application_id).Find(&theme)
	return theme
}