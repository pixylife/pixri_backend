package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`

}

func dbpreloadUser(db *gorm.DB) *gorm.DB {
	return db
}
func AddUser(db *gorm.DB, user *User) {
	db.Create(&user)
}
func UpdateUser(db *gorm.DB, user *User) *User {
	db.Save(&user)
	return user
}
func DeleteUser(db *gorm.DB, user *User) *User {
	db.Delete(&user)
	return user
}
func FindUser(db *gorm.DB, id int) *User {
	user := User{}
	dbpreloadUser(db).First(&user, id)
	return &user
}
func FindAllUser(db *gorm.DB) []*User {
	user := []*User{}
	dbpreloadUser(db).Find(&user)
	return user
}
