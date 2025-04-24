package repositories

import (
	"testing"

	"github.com/devesh121/userAuth/internals/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Define the mockUserRepo struct
// This struct is used to mock the behavior of the User repository for testing purposes.
type mockUserRepo struct {
	mock.Mock // Embedding testify's Mock struct to enable mocking.
}

// CreateUser mocks the behavior of the CreateUser method in the repository.
func (m *mockUserRepo) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user) // Register the method call and return mocked arguments.
	if args.Get(0) == nil {
		return nil, args.Error(1) // Return nil and the mocked error if the first argument is nil.
	}
	return args.Get(0).(*models.User), args.Error(1) // Return the mocked user and error.
}

// GetUserByEmail mocks the behavior of the GetUserByEmail method in the repository.
func (m *mockUserRepo) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email) // Register the method call and return mocked arguments.
	if args.Get(0) == nil {
		return nil, args.Error(1) // Return nil and the mocked error if the first argument is nil.
	}
	return args.Get(0).(*models.User), args.Error(1) // Return the mocked user and error.
}

// GetAllUsers mocks the behavior of the GetAllUsers method in the repository.
func (m *mockUserRepo) GetAllUsers() ([]models.User, error) {
	args := m.Called()                                // Register the method call and return mocked arguments.
	return args.Get(0).([]models.User), args.Error(1) // Return the mocked list of users and error.
}

// GetUserByID mocks the behavior of the GetUserByID method in the repository.
func (m *mockUserRepo) GetUserByID(id uint) (*models.User, error) {
	args := m.Called(id)                             // Register the method call and return mocked arguments.
	return args.Get(0).(*models.User), args.Error(1) // Return the mocked user and error.
}

// UpdateUser mocks the behavior of the UpdateUser method in the repository.
func (m *mockUserRepo) UpdateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)                           // Register the method call and return mocked arguments.
	return args.Get(0).(*models.User), args.Error(1) // Return the mocked updated user and error.
}

// DeleteUser mocks the behavior of the DeleteUser method in the repository.
func (m *mockUserRepo) DeleteUser(id uint) error {
	args := m.Called(id) // Register the method call and return mocked arguments.
	return args.Error(0) // Return the mocked error.
}

// TestCreateUser tests the CreateUser method of the repository.
func TestCreateUser(t *testing.T) {
	mockRepo := new(mockUserRepo) // Create a new instance of the mock repository.
	// Create a dummy user object to simulate input.
	user := &models.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Age:      30,
		Role:     "admin",
	}

	// Mock the behavior of the CreateUser method.
	mockRepo.On("CreateUser", user).Return(user, nil)

	// Call the mocked CreateUser method.
	createdUser, err := mockRepo.CreateUser(user)

	// Assertions to verify the behavior.
	assert.NoError(t, err)                                    // Ensure no error was returned.
	assert.NotNil(t, createdUser)                             // Ensure the created user is not nil.
	assert.Equal(t, "John Doe", createdUser.Name)             // Verify the user's name matches.
	assert.Equal(t, "johndoe@example.com", createdUser.Email) // Verify the user's email matches.

	// Verify that all expectations on the mock were met.
	mockRepo.AssertExpectations(t)
}

// TestGetUserByEmail tests the GetUserByEmail method of the repository.
func TestGetUserByEmail(t *testing.T) {
	mockRepo := new(mockUserRepo) // Create a new instance of the mock repository.

	// Create a dummy user object to simulate the database record.
	email := "johndoe@example.com"
	user := &models.User{
		Name:     "John Doe",
		Email:    email,
		Password: "password123",
		Age:      30,
		Role:     "admin",
	}

	// Mock the behavior of the GetUserByEmail method.
	mockRepo.On("GetUserByEmail", email).Return(user, nil)

	// Call the mocked GetUserByEmail method.
	fetchedUser, err := mockRepo.GetUserByEmail(email)

	// Assertions to verify the behavior.
	assert.NoError(t, err)                        // Ensure no error was returned.
	assert.NotNil(t, fetchedUser)                 // Ensure the fetched user is not nil.
	assert.Equal(t, "John Doe", fetchedUser.Name) // Verify the user's name matches.
	assert.Equal(t, email, fetchedUser.Email)     // Verify the user's email matches.

	// Verify that all expectations on the mock were met.
	mockRepo.AssertExpectations(t)
}

// TestGetUserByID tests the GetUserByID method of the repository.
func TestGetUserByID(t *testing.T) {
	mockRepo := new(mockUserRepo) // Create a new instance of the mock repository.

	// Create a dummy user object to simulate the database record.
	id := uint(1)
	user := &models.User{
		Model:    gorm.Model{ID: id},
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Age:      30,
		Role:     "admin",
	}

	// Mock the behavior of the GetUserByID method.
	mockRepo.On("GetUserByID", id).Return(user, nil)

	// Call the mocked GetUserByID method.
	fetchedUser, err := mockRepo.GetUserByID(id)

	// Assertions to verify the behavior.
	assert.NoError(t, err)                        // Ensure no error was returned.
	assert.NotNil(t, fetchedUser)                 // Ensure the fetched user is not nil.
	assert.Equal(t, id, fetchedUser.ID)           // Verify the user's ID matches.
	assert.Equal(t, "John Doe", fetchedUser.Name) // Verify the user's name matches.

	// Verify that all expectations on the mock were met.
	mockRepo.AssertExpectations(t)
}

// TestGetAllUsers tests the GetAllUsers method of the repository.
func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mockUserRepo) // Create a new instance of the mock repository.
	// Create a dummy list of users to simulate the database records.
	users := []models.User{
		{
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Password: "password123",
			Age:      30,
			Role:     "admin",
		},
		{
			Name:     "Jane Doe",
			Email:    "janedoe@example.com",
			Password: "password456",
			Age:      25,
			Role:     "user",
		},
	}
	// Mock the behavior of the GetAllUsers method.
	mockRepo.On("GetAllUsers").Return(users, nil)

	// Call the mocked GetAllUsers method.
	fetchedUsers, err := mockRepo.GetAllUsers()

	// Assertions to verify the behavior.
	assert.NoError(t, err)                            // Ensure no error was returned.
	assert.NotNil(t, fetchedUsers)                    // Ensure the fetched users are not nil.
	assert.Len(t, fetchedUsers, 2)                    // Verify the number of users matches.
	assert.Equal(t, "John Doe", fetchedUsers[0].Name) // Verify the first user's name matches.

	// Verify that all expectations on the mock were met.
	mockRepo.AssertExpectations(t)
}

// TestUpdateUser tests the UpdateUser method of the repository.
func TestUpdateUser(t *testing.T) {
	mockRepo := new(mockUserRepo) // Create a new instance of the mock repository.

	// Create a dummy user object to simulate the database record.
	user := &models.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Age:      30,
		Role:     "admin",
	}

	// Mock the behavior of the UpdateUser method.
	mockRepo.On("UpdateUser", user).Return(user, nil)

	// Call the mocked UpdateUser method.
	updatedUser, err := mockRepo.UpdateUser(user)

	// Assertions to verify the behavior.
	assert.NoError(t, err)                        // Ensure no error was returned.
	assert.NotNil(t, updatedUser)                 // Ensure the updated user is not nil.
	assert.Equal(t, "John Doe", updatedUser.Name) // Verify the user's name matches.

	// Verify that all expectations on the mock were met.
	mockRepo.AssertExpectations(t)
}

// TestDeleteUser tests the DeleteUser method of the repository.
func TestDeleteUser(t *testing.T) {
	mockRepo := new(mockUserRepo) // Create a new instance of the mock repository.

	// Define the user ID to delete.
	id := uint(1)

	// Mock the behavior of the DeleteUser method.
	mockRepo.On("DeleteUser", id).Return(nil)

	// Call the mocked DeleteUser method.
	err := mockRepo.DeleteUser(id)

	// Assertions to verify the behavior.
	assert.NoError(t, err) // Ensure no error was returned.

	// Verify that all expectations on the mock were met.
	mockRepo.AssertExpectations(t)
}
