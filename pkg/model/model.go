package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func InitModels(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Application{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Theme{})
}
