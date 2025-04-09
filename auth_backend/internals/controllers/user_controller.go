package controllers

import (
	"net/http"

	"github.com/devesh121/userAuth/internals/dto"
	"github.com/devesh121/userAuth/internals/services"
	"github.com/gin-gonic/gin"
)

// userController struct: depends on UserService abstraction
type userController struct {
	userService services.UserService
}

// NewUserController returns a new controller with injected service
func NewUserController(service services.UserService) *userController {
	return &userController{
		userService: service,
	}
}

// RegisterUser handles incoming registration requests from client
func (uc *userController) RegisterUser(c *gin.Context) {
	var req dto.RegisterRequest

	// Step 1: Bind JSON body to DTO and validate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide all required fields: name, email, and password "})
		return
	}

	// Step 2: Call the service layer to register the user
	res, err := uc.userService.RegisterUserService(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 3: Return the response to client
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"data":    res,
	})
}
