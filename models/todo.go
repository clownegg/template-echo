package models

import "time"

type Todo struct {
	ID        uint      `gorm:"primary_key, json:id"`
	TITLE     string    `json:"title"`
	DONE      *bool     `json:"done"`
	IsDeleted *bool     `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
