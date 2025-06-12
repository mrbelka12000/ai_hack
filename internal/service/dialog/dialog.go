package dialog

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

type (
	Service struct {
		repo repo
	}
)

func NewService(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, obj internal.DialogCU) (uuid.UUID, error) {
	obj.ID = uuid.New()
	err := s.repo.Create(ctx, obj)
	return obj.ID, err
}

func (s *Service) Update(ctx context.Context, obj internal.Dialog) error {
	return s.repo.Update(ctx, obj)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (internal.Dialog, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) List(ctx context.Context, pars internal.DialogPars) ([]internal.Dialog, error) {
	return s.repo.List(ctx, pars)
}
