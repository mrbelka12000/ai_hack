package user

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

type (
	repo interface {
		Create(ctx context.Context, obj internal.UserCU) error
		Update(ctx context.Context, obj internal.UserCU) error
		Delete(ctx context.Context, id uuid.UUID) error
		Get(ctx context.Context, pars internal.UserGetPars) (internal.User, error)
		List(ctx context.Context, pars internal.UserPars) ([]internal.User, error)
	}
)
