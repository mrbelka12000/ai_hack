package dialog_messages

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
	"github.com/mrbelka12000/ai_hack/pkg/gorm/postgres"
)

type Repo struct {
	db *postgres.Gorm
}

func New(db *postgres.Gorm) *Repo {
	return &Repo{db: db}
}

func (r *Repo) AddMessage(ctx context.Context, obj internal.DialogMessage) error {
	return r.db.WithContext(ctx).Table("dialogs_messages").Create(&obj).Error
}

func (r *Repo) GetMessagesByDialogID(ctx context.Context, dialogID uuid.UUID) ([]internal.DialogMessage, error) {
	var messages []internal.DialogMessage
	err := r.db.WithContext(ctx).
		Table("dialogs_messages").
		Where("dialog_id = ?", dialogID).
		Order("created_at ASC").
		Find(&messages).Error
	return messages, err
}
