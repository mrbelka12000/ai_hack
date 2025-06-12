package internal

import (
	"time"

	"github.com/google/uuid"

	aihack "github.com/mrbelka12000/ai_hack"
)

type (
	DialogMessage struct {
		DialogUUID  uuid.UUID   `json:"dialog_id" gorm:"column:dialog_id"`
		Role        aihack.Role `json:"role" gorm:"column:role" validate:"required"`
		Message     string      `json:"message" gorm:"column:message" validate:"required"`
		IsAnonymous bool        `json:"-" gorm:"-"`
		CreatedAt   time.Time   `json:"-" gorm:"column:created_at"`
	}
)
