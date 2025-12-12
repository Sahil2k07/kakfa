package services

import (
	"context"
	"fmt"
	"time"

	errz "github.com/Sahil2k07/kakfa/internal/errors"
	"github.com/Sahil2k07/kakfa/internal/graphql/generated"
	"github.com/Sahil2k07/kakfa/internal/interfaces"
	"github.com/Sahil2k07/kakfa/internal/models"
	"github.com/Sahil2k07/kakfa/internal/utils"
)

type todoService struct{ repo interfaces.TodoRepository }

func (s *todoService) CreateTodo(ctx context.Context, input generated.CreateTodoInput) (*generated.Todo, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return nil, errz.NewUnauthorized("unauthorized")
	}

	todo := &models.Todo{
		Title:       input.Title,
		Description: input.Description,
		UserID:      user.ID,
	}

	if err := s.repo.CreateTodo(todo); err != nil {
		return nil, errz.NewInternalError(fmt.Sprintf("failed to create todo: %v", err))
	}

	return &generated.Todo{
		ID:          fmt.Sprintf("%d", todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      generated.TodoStatus(todo.Status),
		CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, input generated.UpdateTodoInput) (*generated.Todo, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return nil, errz.NewUnauthorized("unauthorized")
	}

	todo, err := s.repo.GetTodoByID(uint(input.ID), user.ID)
	if err != nil {
		return nil, errz.NewNotFound("todo not found")
	}

	updatedTodo := models.Todo{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	if err := s.repo.UpdateTodo(todo.PrimaryID, &updatedTodo); err != nil {
		return nil, errz.NewInternalError(fmt.Sprintf("failed to update todo: %v", err))
	}

	return &generated.Todo{
		ID:          todo.ID,
		PrimaryID:   int(todo.PrimaryID),
		Title:       todo.Title,
		Status:      generated.TodoStatus(todo.Status),
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id uint) (bool, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return false, errz.NewUnauthorized("unauthorized")
	}

	if err := s.repo.DeleteTodo(id, user.ID); err != nil {
		return false, errz.NewInternalError(fmt.Sprintf("failed to delete todo: %v", err))
	}

	return true, nil
}

func (s *todoService) GetTodoByID(ctx context.Context, id uint) (*generated.Todo, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return nil, errz.NewUnauthorized("unauthorized")
	}

	todo, err := s.repo.GetTodoByID(id, user.ID)
	if err != nil {
		return nil, errz.NewNotFound("todo not found")
	}

	return &generated.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Status:      generated.TodoStatus(todo.Status),
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *todoService) GetTodos(ctx context.Context, page *int, limit *int) (*generated.TodoResponse, error) {
	// user, err := utils.GetUserClaims(ctx)
	// if err != nil {
	// 	return nil, errz.NewUnauthorized("unauthorized")
	// }
	panic("not implemented")
}

func TodoService(repo interfaces.TodoRepository) interfaces.TodoService {
	return &todoService{repo: repo}
}
