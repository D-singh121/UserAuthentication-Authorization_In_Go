package controllers

import (
	"net/http"

	"github.com/devesh121/userAuth/internals/dto"      // DTOs (Data Transfer Objects)
	"github.com/devesh121/userAuth/internals/services" // Service layer
	"github.com/gin-gonic/gin"                         // Gin web framework
)

type UserController struct {
	UserService services.UserService // Injecting service to interact with business logic
}

// Constructor function to initialize UserController with UserService
func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// RegisterUser handles POST /api/v1/users/register
func (uc *UserController) RegisterUser(c *gin.Context) {
	var userRequest dto.RegisterRequest // Create a variable to hold the incoming request payload

	// Bind and validate the incoming JSON request body to userRequest struct
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		// Return 400 Bad Request if validation or binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service layer's Register function, which handles the business logic
	userResponse, err := uc.UserService.Register(userRequest)
	if err != nil {
		// If service returns an error, respond with 500 or appropriate error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If everything is successful, return 201 Created with the user response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    userResponse,
	})
}
