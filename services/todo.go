package services

import (
	"template_app/models"
	"template_app/repositories"
)

type TodoService interface {
	FindAll() []models.Todo
}

type todoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &todoService{
		todoRepository,
	}
}

func (s *todoService) FindAll() []models.Todo {
	return s.todoRepository.FindAll()
}
