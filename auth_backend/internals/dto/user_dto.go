package dto

// 📝 Request struct for user registration
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`     // Required field
	Email    string `json:"email" binding:"required"`    // Required + email format (can add custom validator)
	Password string `json:"password" binding:"required"` // Required field
	Age      int    `json:"age"`                         // Optional
}

// 🔐 Login request payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`    // Required
	Password string `json:"password" binding:"required"` // Required
}

// 📤 Response struct to return filtered user info (excluding sensitive data like password)
type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// 📤 Response struct to return filtered user info (excluding sensitive data like password)
