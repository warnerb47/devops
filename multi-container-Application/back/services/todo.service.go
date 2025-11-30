package services

import (
	"warnerb47/todo/models"
)

type TodoService interface {
	GetTodoById(string) (*models.Todo, error)
	GetTodos() ([]*models.Todo, error)
	CreateTodo(*models.Todo) error
	DeleteTodo(string) error
	UpdateTodo(string, *models.Todo) error
}
