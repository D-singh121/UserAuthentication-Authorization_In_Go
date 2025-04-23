// internals/routes/user_routes.go
package routes

import (
	"github.com/devesh121/userAuth/internals/controllers"
	"github.com/devesh121/userAuth/internals/middlewares"
	"github.com/devesh121/userAuth/internals/repositories"
	"github.com/devesh121/userAuth/internals/services"
	"github.com/devesh121/userAuth/pkg/config"
	"github.com/gin-gonic/gin"
)

// UserRoutes sets up all routes under /api/users
func UserRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1") // api versioning
	users := v1.Group("/users")   // /api/v1/users

	db := config.DB // global DB variable in our config package

	// Initialize repositories, services, and controllers by injecting dependencies into each layer
	userRepo := repositories.NewPostgresUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Public routes
	users.POST("/register", userController.RegisterUser)
	users.POST("/login", userController.LoginUser)
	users.POST("/logout", userController.LogoutUser)

	// Protected routes
	protected := users.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())
	{
		protected.GET("/", userController.GetAllUsers)
		protected.GET("/:id", userController.GetUserByID)
		protected.POST("/email", userController.GetUserByEmail)
		protected.PUT("/:id", userController.UpdateUserByID)
		protected.DELETE("/:id", userController.DeleteUserByID)
	}
}
