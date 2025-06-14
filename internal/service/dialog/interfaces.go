package dialog

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

type (
	repo interface {
		Create(ctx context.Context, obj internal.DialogCU) error
		Update(ctx context.Context, obj internal.Dialog) error
		Delete(ctx context.Context, id uuid.UUID) error
		Get(ctx context.Context, id uuid.UUID) (internal.Dialog, error)
		List(ctx context.Context, pars internal.DialogPars) ([]internal.Dialog, error)
	}
)
