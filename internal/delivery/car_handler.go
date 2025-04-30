package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "rest-project/internal/middleware" // Импортировать middleware
	"rest-project/internal/models"
	_ "rest-project/internal/services"
	service "rest-project/internal/services"
	"strconv"
)

type CarHandler struct {
	service *service.CarService
}

// Конструктор
func NewCarHandler(service *service.CarService) *CarHandler {
	return &CarHandler{service: service}
}

// Получение списка всех машин
func (h *CarHandler) GetAllCars(c *gin.Context) {
	cars, _ := h.service.GetAllCars()
	c.JSON(http.StatusOK, cars)
}

// Получение машины по ID
func (h *CarHandler) GetCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	car, err := h.service.GetCarByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, car)
}

// Создание новой машины (требуется роль администратора)
func (h *CarHandler) CreateCar(c *gin.Context) {
	// Проверка на роль администратора через middleware
	if role, exists := c.Get("role"); !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	var carCreate models.CarEdit
	if err := c.ShouldBindJSON(&carCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newCar, err := h.service.Create(carCreate.Brand, carCreate.Model, carCreate.Year, carCreate.Color)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}

	c.JSON(http.StatusCreated, newCar)
}

// Обновление машины (требуется роль администратора)
func (h *CarHandler) UpdateCar(c *gin.Context) {
	// Проверка на роль администратора через middleware
	if role, exists := c.Get("role"); !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	var carEdit models.CarEdit
	if err := c.ShouldBindJSON(&carEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedCar, err := h.service.Update(id, &carEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, updatedCar)
}

// Удаление машины (требуется роль администратора)
func (h *CarHandler) DeleteCar(c *gin.Context) {
	// Проверка на роль администратора через middleware
	if role, exists := c.Get("role"); !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	if err := h.service.DeleteCar(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
