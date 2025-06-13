package internal

import (
	"time"

	"github.com/google/uuid"

	aihack "github.com/mrbelka12000/ai_hack"
)

type (
	DialogMessage struct {
		DialogID    uuid.UUID   `json:"dialog_id" gorm:"column:dialog_id"`
		Role        aihack.Role `json:"role" gorm:"column:role" validate:"required"`
		Message     string      `json:"message" gorm:"column:message" validate:"required"`
		IsAnonymous bool        `json:"-" gorm:"-"`
		CreatedAt   time.Time   `json:"-" gorm:"column:created_at"`
	}

	DialogMessageResponse struct {
		Message           string
		RelativeQuestions []string `json:"relative_questions"`
		DatabaseFile      string   `json:"database_file"`
		DatabaseFilePart  string   `json:"database_file_part"`
		Confidence        int      `json:"confidence"`
	}
)
