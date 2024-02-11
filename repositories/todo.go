package repositories

import (
	"template_app/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll() []models.Todo
	FindByID() error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db,
	}
}

func (r *todoRepository) FindAll() []models.Todo {
	session := r.db.Table("t_todos")
	todos := []models.Todo{}

	session.Find(&todos)

	return todos
}

func (r *todoRepository) FindByID() error {
	return nil
}
