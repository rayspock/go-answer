package models

import (
	"time"
)

//History ... represent access history
type History struct {
	Event string `gorm:"not null" json:"event,omitempty"`
	Key string `gorm:"not null" json:"key,omitempty"`
	Data Answer `gorm:"type:jsonb;not null" json:"data,omitempty"`
	CreateDate time.Time `gorm:"not null" json:"createDate,omitempty"`
}


//TableName retrieve Table Name
func (h *History) TableName() string {
	return "history"
}