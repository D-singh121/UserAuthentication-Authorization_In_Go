package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // automatically handles ID, CreatedAt, UpdatedAt
	Name       string `json:"name" binding:"required"`                      // required
	Email      string `json:"email" gorm:"unique" binding:"required,email"` // required + must be a valid email
	Password   string `json:"password" binding:"required,min=6"`            // required + minimum 6 chars
	Age        int    `json:"age" binding:"required,gte=0,lte=120"`         // required + between 0 and 100
}
