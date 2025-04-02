package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type StudentRepositoryImpl struct {
	db *gorm.DB
}

// NewStudentRepository - Constructor
func NewStudentRepository(db *gorm.DB) *StudentRepositoryImpl {
	return &StudentRepositoryImpl{db: db}
}

// GetAll - Retrieve all students
func (s StudentRepositoryImpl) GetAll() ([]models.Student, error) {
	var students []models.Student
	err := s.db.Find(&students).Error
	return students, err
}

// GetById - Retrieve a student by ID
func (s StudentRepositoryImpl) GetById(id int) (*models.Student, error) {
	var student models.Student
	err := s.db.First(&student, id).Error
	return &student, err
}

// Create - Add a new student
func (s StudentRepositoryImpl) Create(student *models.Student) error {
	return s.db.Create(student).Error
}

// Update - Modify an existing student
func (s StudentRepositoryImpl) Update(id int, student *models.StudentEdit) error {
	return s.db.Model(&models.Student{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(student).Error
}

// Delete - Remove a student by ID
func (s StudentRepositoryImpl) Delete(studentID int) error {
	return s.db.Delete(&models.Student{}, studentID).Error
}
