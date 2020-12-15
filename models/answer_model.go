package models

import "time"

//Answer ... represent Answer detail
type Answer struct {
	Key string `gorm:"not null"`
	Value string `gorm:"not null"`
	CreateDate time.Time `gorm:"not null"`
}

//TableName retrieve Table Name
func (a *Answer) TableName() string {
	return "answer"
}