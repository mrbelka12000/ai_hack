package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

func (uc *UseCase) UserCreate(ctx context.Context, obj internal.UserCU) error {
	return uc.userService.Create(ctx, obj)
}

func (uc *UseCase) UserUpdate(ctx context.Context, obj internal.UserCU) error {
	return uc.userService.Update(ctx, obj)
}

func (uc *UseCase) UserDelete(ctx context.Context, id uuid.UUID) error {
	return uc.userService.Delete(ctx, id)
}

func (uc *UseCase) UserGet(ctx context.Context, pars internal.UserGetPars) (internal.User, error) {
	return uc.userService.Get(ctx, pars)
}

func (uc *UseCase) UserList(ctx context.Context, pars internal.UserPars) ([]internal.User, error) {
	return uc.userService.List(ctx, pars)
}

func (uc *UseCase) UserLogin(ctx context.Context, obj internal.UserLogin) (out internal.User, err error) {
	user, err := uc.userService.Get(ctx, internal.UserGetPars{
		PhoneNumber: obj.PhoneNumber,
	})
	if err != nil {
		return out, err
	}

	return user, nil
}
