package service

import (
	"rest-project/internal/models"
)

// Интерфейс репозитория машин
type CarRepository interface {
	GetAll() ([]models.Car, error)
	GetById(id int) (*models.Car, error)
	Create(car *models.Car) error
	Update(id int, car *models.CarEdit) error
	Delete(carID int) error
}

// Структура сервиса машин
type CarService struct {
	repo CarRepository
}

// Конструктор CarService
func NewCarService(carRepo CarRepository) *CarService {
	return &CarService{repo: carRepo}
}

// Получение всех машин
func (s *CarService) GetAllCars() ([]models.Car, error) {
	return s.repo.GetAll()
}

// Получение машины по ID
func (s *CarService) GetCarByID(id int) (*models.Car, error) {
	return s.repo.GetById(id)
}

// Создание новой машины
func (s *CarService) Create(brand, model string, year int, color string) (*models.Car, error) {
	car := &models.Car{
		Brand: brand,
		Model: model,
		Year:  year,
		Color: color,
	}
	err := s.repo.Create(car)
	return car, err
}

// Обновление машины
func (s *CarService) Update(id int, carEdit *models.CarEdit) (*models.Car, error) {
	err := s.repo.Update(id, carEdit)
	if err != nil {
		return nil, err
	}
	return s.GetCarByID(id)
}

// Удаление машины
func (s *CarService) DeleteCar(carID int) error {
	return s.repo.Delete(carID)
}
