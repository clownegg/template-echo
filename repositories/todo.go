package repositories

import (
	"template_app/dao"
	"template_app/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll(cond models.TodosCond) ([]models.Todo, int64)
	FindById(id int, cond models.TodoCond) (*models.Todo, error)
	Create(postData models.TodoPost) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db,
	}
}

func (r *todoRepository) FindAll(cond models.TodosCond) ([]models.Todo, int64) {
	todos, count := dao.SearchTodo(r.db, cond)
	return todos, count
}

func (r *todoRepository) FindById(id int, cond models.TodoCond) (*models.Todo, error) {
	return dao.FindTodoById(r.db, id, cond)
}

func (r *todoRepository) Create(postData models.TodoPost) error {
	return dao.CreateTodo(r.db, postData)
}
