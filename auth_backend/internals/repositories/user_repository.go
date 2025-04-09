package repository

import (
	"github.com/devesh121/userAuth/internals/models" // Importing User model
	"gorm.io/gorm"                                   // GORM ORM package for DB access
)

// UserRepo interface declares methods to be implemented :- matlab inhe implement karna hai but kaise vo nahi batata hai.
type UserRepo interface {
	Create(user *models.User) (*models.User, error) // Method to create new user
	GetByEmail(email string) (*models.User, error)  // Method to find user by email
	GetByID(id uint) (*models.User, error)          // Method to find user by ID
	GetAll() ([]models.User, error)                 // Method to get all users
	Update(user *models.User) (*models.User, error) // Method to update user
	Delete(id uint) error                           // Method to delete user
}

// postgresUserRepository is the concrete implementation of UserRepo using PostgreSQL (via GORM)
// future me agar hume koi aur db use karna ho to sirf ye struct ko modify karna hoga. like "mongoUserRepository"
type postgresUserRepository struct {
	db *gorm.DB // DB instance injected from config
}

// NewPostgresUserRepo returns a new instance of postgresUserRepository as UserRepo in object form.
// ye ek constructor factory method hai jo ek interface type ka object de deta hai, aur Clean Architecture + Testability maintain karta hai.
// Is function ko hum "repository" ke upper layer 'service' me use karenge
func NewPostgresUserRepo(db *gorm.DB) UserRepo {
	return &postgresUserRepository{db: db}
}

// Create adds a new user to the database
func (r *postgresUserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err // Return error if insertion fails
	}
	return user, nil
}

// GetByEmail finds a user by email
func (r *postgresUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err // Return error if user not found
	}
	return &user, nil
}

// GetByID finds a user by ID
func (r *postgresUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err // Return error if user not found
	}
	return &user, nil
}

// GetAll fetches all users from the DB
func (r *postgresUserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err // Return error if DB read fails
	}
	return users, nil // users ek slice hai isliye uska refrence yaha return me pass hoga isliye use "&users" aisa likhne ki jarurat nahi hai.
}

// Update modifies an existing user record
func (r *postgresUserRepository) Update(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err // Return error if update fails
	}
	return user, nil
}

// Delete removes a user by ID
func (r *postgresUserRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err // Return error if delete fails
	}
	return nil
}

// NOTE:
// If struct is passed as pointer → return as it is
// If struct is local inside function → return its address (&user)
