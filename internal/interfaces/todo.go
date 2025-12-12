package interfaces

import (
	"context"

	"github.com/Sahil2k07/kakfa/internal/graphql/generated"
	"github.com/Sahil2k07/kakfa/internal/models"
)

type TodoRepository interface {
	CreateTodo(todo *models.Todo) error
	UpdateTodo(id uint, todo *models.Todo) error
	DeleteTodo(id uint, userID uint) error
	GetTodoByID(id uint, userID uint) (*models.RTodo, error)
	GetAllTodos(userID uint, page *int, limit *int) ([]models.RTodo, int64, error)
}

type TodoService interface {
	CreateTodo(ctx context.Context, input generated.CreateTodoInput) (*generated.Todo, error)

	UpdateTodo(ctx context.Context, input generated.UpdateTodoInput) (*generated.Todo, error)

	DeleteTodo(ctx context.Context, id uint) (bool, error)

	GetTodos(ctx context.Context, page *int, limit *int) (*generated.TodoResponse, error)
}
