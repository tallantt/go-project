package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id uint) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByID(id uint) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user models.User) (models.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
