package models

import "time"

type Todo struct {
	ID        uint      `gorm:"primary_key, json:id"`
	Title     string    `json:"title"`
	Done      *bool     `json:"done"`
	IsDeleted *bool     `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodosCond struct {
	Title     string `query:"title"`
	Done      int    `query:"done"`
	IsDeleted int    `query:"is_deleted"`
	Sort      string `query:"sort"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
}

type TodoCond struct {
	IsDeleted int `query:"is_deleted"`
}
