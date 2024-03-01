package gormtable

import (
	"main.go/global"
	"main.go/models"
)

func CreateUserTable() {
	if !global.DB.Migrator().HasTable(&models.User{}) {
		global.DB.AutoMigrate(&models.User{})
	}
}
