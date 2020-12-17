package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Answer ... represent Answer detail
type Answer struct {
	Key string `gorm:"not null" json:"key,omitempty"`
	Val string `gorm:"not null" json:"value,omitempty"`
}

// TableName retrieve Table Name
func (a *Answer) TableName() string {
	return "answer"
}

// Value ... Make the Answer struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Answer) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan ... Make the Answer struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Answer) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// AnswerPayload ... http request type
type AnswerPayload struct {
	Value string
}
