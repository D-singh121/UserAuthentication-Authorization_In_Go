package dto

// 📝 Request struct for user registration
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`              // Required field
	Email    string `json:"email" binding:"required"`             // Required + email format (can add custom validator)
	Password string `json:"password" binding:"required"`          // Required field
	Age      int    `json:"age" binding:"required,gte=5,lte=120"` // required
	Role     string `json:"role"`                                 // Required field
}

// 🔐 Login request payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`    // Required
	Password string `json:"password" binding:"required"` // Required
}

// login response struct dto
type LoginResponse struct {
	ID    uint   `json:"id"`
	Token string `json:"token,omitempty"` // JWT token optional
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
}

// 📤 Response struct to return filtered user info (excluding sensitive data like password)
type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
}

// UpdateRequest defines the expected payload for updating a user
type UpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=0,lte=120"`
	Password string `json:"password,omitempty"` // Optional
}

type GetUserByEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}
