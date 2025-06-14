package dialog_messages

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
	"github.com/mrbelka12000/ai_hack/internal/client/ml"
)

type (
	repo interface {
		AddMessage(ctx context.Context, obj internal.DialogMessage) error
		GetMessagesByDialogID(ctx context.Context, dialogID uuid.UUID) ([]internal.DialogMessage, error)
	}

	mlClient interface {
		Analyze(ctx context.Context, req ml.AnalyzeRequest) (out ml.AnalyzeResponse, err error)
	}

	cache interface {
		Set(key string, value interface{}, dur time.Duration) error
		Get(key string) (string, bool)
	}
)
