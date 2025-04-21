package controllers

import (
	"net/http"
	"strconv"

	"github.com/devesh121/userAuth/internals/dto"
	"github.com/devesh121/userAuth/internals/services"
	"github.com/gin-gonic/gin"
)

// UserController  struct: depends on UserService abstraction
type UserController struct {
	userService services.UserService
}

// NewUserController returns a new controller with injected service
func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// RegisterUser handles incoming registration requests from client
func (uc *UserController) RegisterUser(c *gin.Context) {
	var req dto.RegisterRequest

	// Step 1: Bind JSON body to DTO and validate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide all required fields: name, email, and password "})
		return
	}

	// Step 2: Call the service layer to register the user
	user, err := uc.userService.RegisterUserService(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 3: Return the response to client
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"data":    user,
	})
}

// LoginUser handles incoming login requests from client
func (uc *UserController) LoginUser(c *gin.Context) {
	var req dto.LoginRequest

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request fields"})
		return
	}

	// Call service
	resp, err := uc.userService.LoginUserService(c, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return token
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data":    resp,
	})
}

// GetAllUser handles incoming getallusers request from client
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID handles incoming GetUserByID request from client
func (uc *UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.userService.GetUserByIDService(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByEmail handles incoming GetUserByEmail request from client
func (uc *UserController) GetUserByEmail(c *gin.Context) {
	var req dto.GetUserByEmailRequest

	// Bind JSON and validate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request email ID format"})
		return
	}

	user, err := uc.userService.GetUserByEmailService(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserByID handles incoming UpdateUserByID request from client
func (uc *UserController) UpdateUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // converting string id into integer
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// request body ko dto.UpdateRequest me bind kiya ja raha hai
	var req dto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// calling updateUser service layer and passing userid as unsigned int with updateUser data in dto form.
	updatedUser, err := uc.userService.UpdateUserService(req, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUserByID handles incoming DeleteUserByID request from client
func (uc *UserController) DeleteUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = uc.userService.DeleteUserService(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
