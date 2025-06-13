package personal_data

import (
	"context"

	"github.com/mrbelka12000/ai_hack/internal"
)

type (
	repo interface {
		Create(ctx context.Context, obj internal.PersonalData) error
		GetByCustID(ctx context.Context, custID string) ([]internal.PersonalData, error)
	}
)
