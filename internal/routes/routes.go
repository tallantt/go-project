package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rest-project/internal/delivery"
	"rest-project/internal/repository"
	"rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Инициализация репозитория
	studentRepo := repository.NewStudentRepository(db)

	// Инициализация сервиса
	studentService := service.NewStudentService(studentRepo)

	// Инициализация обработчика
	studentHandler := delivery.NewStudentHandler(studentService)

	// Роуты
	students := r.Group("api/v1/students")
	{
		students.GET("/", studentHandler.GetAllStudents)
		students.POST("/", studentHandler.CreateStudent)
		students.GET("/:id", studentHandler.GetStudent)
		students.PUT("/:id", studentHandler.UpdateStudent)
		students.DELETE("/:id", studentHandler.DeleteStudent)
	}
}
