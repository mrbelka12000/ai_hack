package internal

import (
	"time"

	aihack "github.com/mrbelka12000/ai_hack"
)

type (
	UserCU struct {
		PhoneNumber string      `gorm:"column:phone_number" json:"phone_number" validate:"required"`
		Code        string      `gorm:"-" json:"code" validate:"required"`
		Role        aihack.Role `gorm:"column:role" json:"role,omitempty" validate:"required,oneof=client operator"`
		CreatedAt   time.Time   `gorm:"column:created_at" json:"-"`
	}

	User struct {
		ID          int64       `gorm:"column:id" json:"id,omitempty"`
		PhoneNumber string      `gorm:"column:phone_number" json:"phone_number"`
		Role        aihack.Role `gorm:"column:role" json:"role,omitempty"`
		CreatedAt   time.Time   `gorm:"column:created_at" json:"created_at"`
	}

	UserPars struct {
		PhoneNumber   string      `json:"phone_number"`
		Role          aihack.Role `schema:"role,omitempty"`
		CreatedBefore time.Time   `schema:"created_before"`
		CreatedAfter  time.Time   `schema:"created_after"`
		PaginationParams
	}

	UserGetPars struct {
		ID          int64       `schema:"id,omitempty"`
		PhoneNumber string      `schema:"phone_number"`
		Role        aihack.Role `schema:"role,omitempty"`
	}

	UserLogin struct {
		PhoneNumber string `gorm:"column:phone_number" json:"phone_number" validate:"required"`
		Code        string `gorm:"-" json:"code" validate:"required"`
	}

	UserListResponse struct {
		User
		PaginationParams
	}
)
