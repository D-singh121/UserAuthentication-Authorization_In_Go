// internals/routes/user_routes.go
package routes

import (
	"github.com/devesh121/userAuth/internals/controllers"
	"github.com/devesh121/userAuth/internals/repositories"
	"github.com/devesh121/userAuth/internals/services"
	"github.com/devesh121/userAuth/pkg/config"
	"github.com/gin-gonic/gin"
)

// UserRoutes sets up all routes under /api/users
func UserRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	users := v1.Group("/users") // api versioning and grouping

	// Step 1: Inject DB
	db := config.DB

	// Step 2: Create repo, service, controller
	userRepo := repositories.NewPostgresUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	{
		users.POST("/register", userController.RegisterUser) // Register new user
		// users.POST("/login", controllers.LoginUser)            // Login
		// users.GET("/", controllers.GetAllUsers)                // Get all users
		// users.GET("/:id", controllers.GetUserByID)             // Get user by ID
		// users.GET("/email/:email", controllers.GetUserByEmail) // Get user by Email
		// users.PUT("/:id", controllers.UpdateUserByID)          // Update user
		// users.DELETE("/:id", controllers.DeleteUserByID)       // Delete user
	}
}
