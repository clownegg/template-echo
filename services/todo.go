package services

import (
	"template_app/models"
	"template_app/repositories"
)

type TodoService interface {
	FindAll(cond models.TodosCond) ([]models.Todo, int64)
	FindById(id int, cond models.TodoCond) models.Todo
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

func (s *todoService) FindById(id int, cond models.TodoCond) models.Todo {
	return s.todoRepository.FindById(id, cond)
}
