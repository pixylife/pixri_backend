package model

import "github.com/jinzhu/gorm"

type Entity struct {
	Model
	Name          string      `gorm:"not null" json:"name"`
	Description   string      ` json:"description"`
	Application   Application `gorm:"foreignkey:application_id" json:"-"`
	ApplicationID int         `json:"application_id"`
}



type EntityData struct {
	ID        int `json:"id"`
	FieldCount int `json:"field_count"`
}


func dbpreloadEntity(db *gorm.DB) *gorm.DB {
	return db.Preload("Entity")
}

func AddEntity(db *gorm.DB, entity *Entity) {
	db.Create(&entity)
}
func UpdateEntity(db *gorm.DB, entity *Entity) *Entity {
	db.Save(&entity)
	return entity
}
func DeleteEntity(db *gorm.DB, entity *Entity) *Entity {
	db.Delete(&entity)
	return entity
}
func FindEntity(db *gorm.DB, id int) *Entity {
	entity := Entity{}
	dbpreloadEntity(db).First(&entity, id)
	return &entity
}
func FindAllEntity(db *gorm.DB) []*Entity {
	var entity []*Entity
	dbpreloadEntity(db).Find(&entity)
	return entity
}

func FindAllEntityByApplication(db *gorm.DB,application_id int) []*Entity {
	var entity []*Entity
	db.Where("application_id = ?", application_id).Find(&entity)
	return entity
}
