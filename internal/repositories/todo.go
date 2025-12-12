package repositories

import (
	"errors"

	"github.com/Sahil2k07/kakfa/internal/database"
	"github.com/Sahil2k07/kakfa/internal/interfaces"
	"github.com/Sahil2k07/kakfa/internal/models"
)

type todoRepository struct{}

func (r *todoRepository) GetAllTodos(userID uint, page *int, limit *int) ([]models.RTodo, int64, error) {
	panic("unimplemented")
}

func (r *todoRepository) CreateTodo(todo *models.Todo) error {
	return database.WDB.Create(todo).Error
}

func (r *todoRepository) UpdateTodo(id uint, todo *models.Todo) error {
	return database.WDB.Model(&models.Todo{}).Where("id = ?", id).Updates(todo).Error
}

func (r *todoRepository) DeleteTodo(id uint, userID uint) error {
	res := database.WDB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{})
	if res.RowsAffected == 0 {
		return errors.New("todo not found or not owned by user")
	}
	return res.Error
}

func (r *todoRepository) GetTodoByID(id uint, userID uint) (*models.RTodo, error) {
	panic("not implemented")
}

func TodoRepository() interfaces.TodoRepository {
	return &todoRepository{}
}
