package internal

import (
	"time"

	"github.com/google/uuid"

	aihack "github.com/mrbelka12000/ai_hack"
)

type (
	DialogCU struct {
		ID         uuid.UUID           `gorm:"column:id" json:"-"`
		ClientID   int64               `gorm:"column:client_id" json:"-"`
		OperatorID int64               `gorm:"column:operator_id" json:"-"`
		Status     aihack.DialogStatus `gorm:"column:status" json:"-"`
		Message    string              `gorm:"-" json:"message" validate:"required"`
		CreatedAt  time.Time           `gorm:"column:created_at" json:"-"`
	}

	Dialog struct {
		ID         uuid.UUID           `gorm:"column:id" json:"id,omitempty"`
		ClientID   int64               `gorm:"column:client_id" json:"client_id,omitempty"`
		OperatorID int64               `gorm:"column:operator_id" json:"operator_id,omitempty"`
		Status     aihack.DialogStatus `gorm:"column:status" json:"status,omitempty"`
		CreatedAt  time.Time           `gorm:"column:created_at" json:"created_at"`

		DialogsMessages []DialogMessage `gorm:"-" json:"dialogs_messages"`
	}

	DialogPars struct {
		ClientID      int64               `schema:"client_id,omitempty"`
		OperatorID    int64               `schema:"operator_id,omitempty"`
		Status        aihack.DialogStatus `schema:"status,omitempty"`
		CreatedBefore time.Time           `schema:"created_before"`
		CreatedAfter  time.Time           `schema:"created_after"`

		PaginationParams
	}

	DialogListResponse struct {
		Dialog
		PaginationParams
	}
)
