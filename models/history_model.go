package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

//History ... represent access history
type History struct {
	Event string `gorm:"not null"`
	Key string `gorm:"not null"`
	Data postgres.Jsonb `gorm:"type:jsonb;not null"`
	CreateDate time.Time `gorm:"not null"`
}


//TableName retrieve Table Name
func (h *History) TableName() string {
	return "history"
}