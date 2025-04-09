package services

import (
	"errors"

	"github.com/devesh121/userAuth/internals/dto"          // Request and response DTOs
	"github.com/devesh121/userAuth/internals/models"       // DB models
	"github.com/devesh121/userAuth/internals/repositories" // Repository abstraction
	"golang.org/x/crypto/bcrypt"                           // Password hashing
	"gorm.io/gorm"
)

// UserService interface defines business logic layer functions
type UserService interface {
	RegisterUserService(userReq dto.RegisterRequest) (*dto.UserResponse, error)
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

	// Step 3: Map the request DTO to DB model
	newUser := &models.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: string(hashedPassword),
		Age:      userReq.Age,
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
	}

	return response, nil
}
