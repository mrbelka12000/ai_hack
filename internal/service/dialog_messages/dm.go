package dialog_messages

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddMessage(ctx context.Context, obj internal.DialogMessage) error {
	return s.repo.AddMessage(ctx, obj)
}

func (s *Service) GetMessagesByDialogID(ctx context.Context, dialogID uuid.UUID) ([]internal.DialogMessage, error) {
	return s.repo.GetMessagesByDialogID(ctx, dialogID)
}
