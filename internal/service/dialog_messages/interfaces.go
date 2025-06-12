package dialog_messages

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

type (
	repo interface {
		AddMessage(ctx context.Context, obj internal.DialogMessage) error
		GetMessagesByDialogID(ctx context.Context, dialogID uuid.UUID) ([]internal.DialogMessage, error)
	}
)
