package repository

import (
	"fatura-yonetim-sistemi/entity/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) FindUserForLogin(identityNumber string) (*models.User, error) {
	var user models.User
	result := r.DB.Exec(`SELECT id,password,name,surname,email FROM users WHERE identity_number = ?`, identityNumber).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r UserRepository) SetSessionActive(userID string) error {
	result := r.DB.Exec(`UPDATE users SET session_active = 1 WHERE id = ?`, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r UserRepository) SetSessionInactive(userID string) error {
	result := r.DB.Exec(`UPDATE users SET session_active = 0 WHERE id = ?`, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r UserRepository) CreateUser(user *models.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r UserRepository) FindUserForSessionControl(id string) (*models.User, error) {
	var user models.User
	result := r.DB.Exec(`SELECT session_active FROM users WHERE id = ?`, id).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r UserRepository) FindUserForIsManager(id string) (*models.User, error) {
	var user models.User
	result := r.DB.Exec(`SELECT is_manager FROM users WHERE id = ?`, id).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	result := r.DB.Preload("Apartments").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
