package service

import (
	"rest-project/internal/models"
	"rest-project/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	// Здесь можно добавить валидации или хэширование пароля
	return s.userRepo.Create(user)
}

func (s *userService) UpdateUser(user models.User) (models.User, error) {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
