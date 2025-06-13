package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, obj internal.UserCU) error {
	_, err := s.repo.Get(ctx, internal.UserGetPars{
		PhoneNumber: obj.PhoneNumber,
	})
	if err == nil {
		return errors.New("user already exists")
	}

	obj.CreatedAt = time.Now().UTC()

	return s.repo.Create(ctx, obj)
}

func (s *Service) Update(ctx context.Context, obj internal.UserCU) error {
	return s.repo.Update(ctx, obj)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Get(ctx context.Context, pars internal.UserGetPars) (internal.User, error) {
	return s.repo.Get(ctx, pars)
}

func (s *Service) List(ctx context.Context, pars internal.UserPars) ([]internal.User, error) {
	return s.repo.List(ctx, pars)
}
