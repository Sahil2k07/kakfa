package repositories

import (
	"github.com/Sahil2k07/kakfa/internal/database"
	"github.com/Sahil2k07/kakfa/internal/interfaces"
	"github.com/Sahil2k07/kakfa/internal/models"
)

type authRepository struct{}

func (r *authRepository) CheckUserExist(email string) (bool, error) {
	panic("not implemented")
}

func (r *authRepository) GetUser(email string) (models.RUser, error) {
	panic("not implemeted")
}

func (r *authRepository) AddUser(user models.User) error {
	return database.WDB.Create(&user).Error
}

func (r *authRepository) UpdatePassword(email, newPassword string) error {
	return database.WDB.Model(&models.User{}).Where("email = ?", email).Update("password", newPassword).Error
}

func AuthRepository() interfaces.AuthRepository {
	return &authRepository{}
}
