package repositories

import (
	"github.com/Sahil2k07/kakfa/internal/connections"
	errz "github.com/Sahil2k07/kakfa/internal/errors"
	"github.com/Sahil2k07/kakfa/internal/interfaces"
	"github.com/Sahil2k07/kakfa/internal/models"
)

type authRepository struct{}

func (r *authRepository) CheckUserExist(email string) (bool, error) {
	var count int64

	err := connections.WDB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return true, err
	}

	if count > 0 {
		return true, errz.NewAlreadyExists("user already exists")
	}

	return false, nil
}

func (r *authRepository) GetUser(email string) (models.RUser, error) {
	panic("not implemeted")
}

func (r *authRepository) AddUser(user models.User) error {
	return connections.WDB.Create(&user).Error
}

func (r *authRepository) UpdatePassword(email, newPassword string) error {
	return connections.WDB.Model(&models.User{}).Where("email = ?", email).Update("password", newPassword).Error
}

func AuthRepository() interfaces.AuthRepository {
	return &authRepository{}
}
