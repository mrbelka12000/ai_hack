package internal

import (
	"time"

	aihack "github.com/mrbelka12000/ai_hack"
)

type (
	UserCU struct {
		Email     string      `gorm:"column:email" json:"email,omitempty" validate:"required,email"`
		Password  string      `gorm:"column:password" json:"password,omitempty" validate:"required"`
		Role      aihack.Role `gorm:"column:role" json:"role,omitempty" validate:"required,oneof=client operator"'`
		CreatedAt time.Time   `gorm:"column:created_at" json:"-"`
	}

	User struct {
		ID        int64       `gorm:"column:id" json:"id,omitempty"`
		Email     string      `gorm:"column:email" json:"email,omitempty"`
		Password  string      `gorm:"column:password" json:"-,omitempty"`
		Role      aihack.Role `gorm:"column:role" json:"role,omitempty"`
		CreatedAt time.Time   `gorm:"column:created_at" json:"created_at"`
	}

	UserPars struct {
		Email         string      `schema:"email,omitempty"`
		Role          aihack.Role `schema:"role,omitempty" validate:"oneof=client operator"`
		CreatedBefore time.Time   `schema:"created_before"`
		CreatedAfter  time.Time   `schema:"created_after"`
		PaginationParams
	}

	UserGetPars struct {
		ID    int64       `schema:"id,omitempty"`
		Email string      `schema:"email,omitempty"`
		Role  aihack.Role `schema:"role,omitempty"`
	}

	UserLogin struct {
		Email    string `json:"email,omitempty" validate:"required,email"`
		Password string `json:"password,omitempty" validate:"required"`
	}

	UserListResponse struct {
		User
		PaginationParams
	}
)
