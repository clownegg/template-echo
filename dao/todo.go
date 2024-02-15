package dao

import (
	"fmt"
	"template_app/models"

	"gorm.io/gorm"
)

func SearchTodo(db *gorm.DB, cond models.TodosCond) ([]models.Todo, int64) {
	session := db.Table("t_todos")
	todos := []models.Todo{}
	var count int64

	if cond.Done == 1 {
		session.Where("done = ?", true)
	}

	if cond.IsDeleted == 1 {
		session.Where("is_deleted = ?", true)
	}

	if cond.Title != "" {
		t := "%" + cond.Title + "%"
		session.Where("title LIKE ?", t)
	}

	if cond.Limit != 0 {
		session.Limit(cond.Limit)
	}

	if cond.Offset != 0 {
		session.Offset(cond.Offset)
	}

	session.Find(&todos).Count(&count)

	return todos, count
}

func FindTodoById(db *gorm.DB, id int, cond models.TodoCond) (*models.Todo, error) {
	session := db.Table("t_todos")
	todo := models.Todo{}

	if cond.IsDeleted == 1 {
		session.Where("is_deleted = ?", true)
	}

	session.Where("id = ?", id)
	result := session.Take(&todo)

	if result.Error != nil {
		fmt.Println("FindTodoByID SQL ERROR:", result.Error)
		return nil, result.Error
	}

	return &todo, nil
}

func CreateTodo(db *gorm.DB, postData models.TodoPost) error {
	session := db.Table("t_todos")
	result := session.Create(&postData)

	if result.Error != nil {
		fmt.Println("CreateTodo SQL ERROR:", result.Error)
		return result.Error
	}

	return nil
}
