package models

import "github.com/jinzhu/gorm"

//Init ... create table schema from model
func Init(db *gorm.DB) {
	db.AutoMigrate(&Answer{}, &History{})
}

//DropTableIfExists ... Drop table if exists
func DropTableIfExists(db *gorm.DB) {
	db.DropTableIfExists(&Answer{}, &History{})
}
