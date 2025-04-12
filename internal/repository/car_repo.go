package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type CarRepositoryImpl struct {
	db *gorm.DB
}

// NewCarRepository - Конструктор
func NewCarRepository(db *gorm.DB) *CarRepositoryImpl {
	return &CarRepositoryImpl{db: db}
}

// GetAll - Получение всех машин
func (r CarRepositoryImpl) GetAll() ([]models.Car, error) {
	var cars []models.Car
	err := r.db.Find(&cars).Error
	return cars, err
}

// GetById - Получение машины по ID
func (r CarRepositoryImpl) GetById(id int) (*models.Car, error) {
	var car models.Car
	err := r.db.First(&car, id).Error
	return &car, err
}

// Create - Создание новой машины
func (r CarRepositoryImpl) Create(car *models.Car) error {
	return r.db.Create(car).Error
}

// Update - Обновление машины
func (r CarRepositoryImpl) Update(id int, car *models.CarEdit) error {
	return r.db.Model(&models.Car{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(car).Error
}

// Delete - Удаление машины
func (r CarRepositoryImpl) Delete(carID int) error {
	return r.db.Delete(&models.Car{}, carID).Error
}
