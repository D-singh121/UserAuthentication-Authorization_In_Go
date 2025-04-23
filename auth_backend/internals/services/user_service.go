package services

import (
	"errors"
	"fmt"

	"github.com/devesh121/userAuth/internals/dto"          // Request and response DTOs
	"github.com/devesh121/userAuth/internals/models"       // DB models
	"github.com/devesh121/userAuth/internals/repositories" // Repository abstraction
	"github.com/devesh121/userAuth/internals/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt" // Password hashing
	"gorm.io/gorm"
)

// UserService interface defines business logic layer functions
type UserService interface {
	RegisterUserService(userReq dto.RegisterRequest) (*dto.UserResponse, error)
	LoginUserService(userReq dto.LoginRequest) (*dto.LoginResponse, string, error)
	LogoutUserService(c *gin.Context) error
	GetAllUsersService() ([]dto.UserResponse, error)
	GetUserByIDService(id uint) (*dto.UserResponse, error)
	GetUserByEmailService(email string) (*dto.UserResponse, error)
	UpdateUserService(userReq dto.UpdateRequest, id uint) (*dto.UserResponse, error)
	DeleteUserService(id uint) error
}

// userServiceImpl struct implements the UserService interface
type userServiceImpl struct {
	userRepo repositories.UserRepo // Depends on abstraction of repo layer
}

// NewUserService constructor returns implementation of UserService interface for future use in controller layer.
func NewUserService(repo repositories.UserRepo) UserService {
	return &userServiceImpl{userRepo: repo}
}

// RegisterUserService handles the business logic of registering a new user
func (s *userServiceImpl) RegisterUserService(userReq dto.RegisterRequest) (*dto.UserResponse, error) {
	// Step 1: Check if user already exists by email
	existingUser, err := s.userRepo.GetUserByEmail(userReq.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Step 2: Hash the password using bcrypt before saving it to DB
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Default role to "user" if not provided
	role := userReq.Role
	if role == "" {
		role = "user"
	}

	// Step 3: Map the request DTO to DB model
	newUser := &models.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: string(hashedPassword),
		Age:      userReq.Age,
		Role:     role,
	}

	// Step 4: Call repo to create the user in DB
	createdUser, err := s.userRepo.CreateUser(newUser)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	// Step 5: Return the response DTO (don't include password)
	response := &dto.UserResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
		Age:   createdUser.Age,
		Role:  createdUser.Role,
	}

	return response, nil
}

// LoginUserService handles the business logic of user login
func (s *userServiceImpl) LoginUserService(userReq dto.LoginRequest) (*dto.LoginResponse, string, error) {
	//  Find user by email
	user, err := s.userRepo.GetUserByEmail(userReq.Email)
	if err != nil {
		return nil, "", errors.New("no such user found")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		return nil, "", errors.New("invalid credentials password or email")
	}

	//  Generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	// Return token + user details
	return &dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
		Role:  user.Role,
	}, token, nil
}

// LogoutUserService handles the business logic of user logout
func (s *userServiceImpl) LogoutUserService(c *gin.Context) error {
	// Try to get the cookie value (auth_token)
	_, err := c.Cookie("auth_token")
	if err == nil {
		fmt.Println("Warning: Cookie still exists after attempted removal")
	} else {
		fmt.Println("Cookie successfully removed")
	}

	// Expire the cookie by setting maxAge to -1 and an empty value
	c.SetCookie("auth_token", "", -1, "/", "localhost", true, true)
	return nil
}

// GetAllUsersService retrieves all users and returns them as DTOs
func (s *userServiceImpl) GetAllUsersService() ([]dto.UserResponse, error) {
	// Step 1: Call the repository to get all users
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Step 2: Map DB models to DTOs
	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Age:   user.Age,
			Role:  user.Role,
		})
	}

	return userResponses, nil
}

// GetUserByIDService retrieves a user by ID
func (s *userServiceImpl) GetUserByIDService(id uint) (*dto.UserResponse, error) {
	// Step 1: Call the repository to get the user by ID
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Step 2: Return the user data as a DTO
	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
		Role:  user.Role,
	}, nil
}

// GetUserByEmailService retrieves a user by email
func (s *userServiceImpl) GetUserByEmailService(email string) (*dto.UserResponse, error) {
	// Call the repository to get the user by email
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Return the user data as a DTO
	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
		Role:  user.Role,
	}, nil
}

// UpdateUserService updates an existing user's details
func (s *userServiceImpl) UpdateUserService(userReq dto.UpdateRequest, id uint) (*dto.UserResponse, error) {
	// Step 1: Fetch the existing user from the repository
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Step 2: Update user fields (only update non-empty fields)
	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	if userReq.Email != "" {
		user.Email = userReq.Email
	}
	if userReq.Password != "" {
		// Hash the new password before saving
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		user.Password = string(hashedPassword)
	}
	if userReq.Age != 0 {
		user.Age = userReq.Age
	}

	// Step 3: Call the repository to save the updated user
	updatedUser, err := s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	// Step 4: Return the updated user as a DTO
	return &dto.UserResponse{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Age:   updatedUser.Age,
	}, nil
}

// DeleteUserService deletes a user by ID
func (s *userServiceImpl) DeleteUserService(id uint) error {
	// Check if user exists
	_, err := s.userRepo.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Proceed to delete
	err = s.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
