package model

import "github.com/jinzhu/gorm"

type Field struct {
	Model
	Name string `gorm:"not null" json:"name"`
	Type string `gorm:"not null" json:"type"`
	UIName string `gorm:"not null" json:"ui_name"`
	Entity   Entity `gorm:"foreignkey:entity_id" json:"entity"`
	EntityID int
}


func dbpreloadField(db *gorm.DB) *gorm.DB {
	return db.Preload("Field")
}

func AddField(db *gorm.DB, field *Field) {
	db.Create(&field)
}
func UpdateField(db *gorm.DB, field *Field) *Field {
	db.Save(&field)
	return field
}
func DeleteField(db *gorm.DB, field *Field) *Field {
	db.Delete(&field)
	return field
}
func FindField(db *gorm.DB, id int) *Field {
	field := Field{}
	dbpreloadField(db).First(&field, id)
	return &field
}
func FindAllField(db *gorm.DB) []*Field {
	var field []*Field
	dbpreloadField(db).Find(&field)
	return field
}
