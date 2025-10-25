package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" validate:"required" db:"first_name"`
	LastName  string    `json:"last_name"  validate:"required" db:"last_name"`
	Email     string    `json:"email"    validate:"required,email" db:"email"`
	Password  string    `json:"password,omitempty" validate:"required,min=6" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type LoginRequest struct{
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password"   validate:"required,min=6"`
}