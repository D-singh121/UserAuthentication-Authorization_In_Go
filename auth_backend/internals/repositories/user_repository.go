package repositories

import (
	"github.com/devesh121/userAuth/internals/models" // Importing User model
	"gorm.io/gorm"                                   // GORM ORM package for DB access
)

// UserRepo interface declares methods to be implemented :- matlab inhe implement karna hai but kaise vo nahi batata hai.
type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error) // Method to create new user
	GetUserByEmail(email string) (*models.User, error)  // Method to find user by email
	GetUserByID(id uint) (*models.User, error)          // Method to find user by ID
	GetAllUsers() ([]models.User, error)                // Method to get all users
	UpdateUser(user *models.User) (*models.User, error) // Method to update user
	DeleteUser(id uint) error                           // Method to delete user
}

// postgresUserRepository is the concrete implementation of UserRepo using PostgreSQL (via GORM)
// future me agar hume koi aur db use karna ho to sirf ye struct ko modify karna hoga. like "mongoUserRepository"
type postgresUserRepository struct {
	db *gorm.DB // DB instance injected from config
}

// hum yaha multile repositories bana sakte hai like "mongoUserRepository" etc.
// Iska matlab ye hai ki agar hum future me koi aur DB use karna chahte hai to sirf ye struct ko modify karna hoga.

// NewPostgresUserRepo returns a new instance of postgresUserRepository as UserRepo in object form.
// ye ek constructor factory method hai jo ek interface type ka object de deta hai, aur Clean Architecture + Testability maintain karta hai.
// Is function ko hum "repository" ke upper layer 'service' me use karenge
func NewPostgresUserRepo(db *gorm.DB) UserRepo {
	return &postgresUserRepository{db: db}
}

// CreateUser adds a new user to the database
func (r *postgresUserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail finds a user by email
func (r *postgresUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID finds a user by ID
func (r *postgresUserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers fetches all users from the DB
func (r *postgresUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil // users ek slice hai isliye uska refrence yaha return me pass hoga isliye use "&users" aisa likhne ki jarurat nahi hai.
}

// UpdateUser modifies an existing user record
func (r *postgresUserRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err // Return error if update fails
	}
	return user, nil
}

// DeleteUser removes a user by ID
func (r *postgresUserRepository) DeleteUser(id uint) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err // Return error if delete fails
	}
	return nil
}

// NOTE for me:
// If struct is passed as pointer → return as it is
// If struct is local inside function → return its address (&user)
