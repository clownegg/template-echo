package services

import (
	"template_app/models"
	"template_app/repositories"
)

type TodoService interface {
	FindAll(cond models.TodosCond) ([]models.Todo, int64)
	FindById(id int, cond models.TodoCond) (*models.Todo, error)
	Create(postData models.TodoPost) error
}

type todoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &todoService{
		todoRepository,
	}
}

func (s *todoService) FindAll(cond models.TodosCond) ([]models.Todo, int64) {
	return s.todoRepository.FindAll(cond)
}

func (s *todoService) FindById(id int, cond models.TodoCond) (*models.Todo, error) {
	return s.todoRepository.FindById(id, cond)
}

func (s *todoService) Create(postData models.TodoPost) error {
	return s.todoRepository.Create(postData)
}
