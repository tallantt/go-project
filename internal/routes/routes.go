package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/auth"
	"rest-project/internal/db"
	"rest-project/internal/delivery"
	"rest-project/internal/middleware"
	"rest-project/internal/repository"
	service "rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine) {
	// Auth routes
	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	// Protected routes
	protected := r.Group("api/v1")
	protected.Use(middleware.AuthRequired("admin"))

	{
		protected.GET("/me", auth.Me)

		// Car dependencies
		carRepo := repository.NewCarRepository(db.DB)
		carService := service.NewCarService(carRepo)
		carHandler := delivery.NewCarHandler(carService)

		// Car routes
		cars := protected.Group("/cars")
		{
			cars.GET("/", carHandler.GetAllCars)
			cars.GET("/:id", carHandler.GetCar)

			cars.POST("/", middleware.AuthRequired("admin"), carHandler.CreateCar)
			cars.PUT("/:id", middleware.AuthRequired("admin"), carHandler.UpdateCar)
			cars.DELETE("/:id", middleware.AuthRequired("admin"), carHandler.DeleteCar)
		}
	}
}
