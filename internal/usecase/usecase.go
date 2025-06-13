package usecase

import (
	"log/slog"

	"github.com/mrbelka12000/ai_hack/internal/client/ml"
	"github.com/mrbelka12000/ai_hack/internal/repo"
	"github.com/mrbelka12000/ai_hack/internal/service/dialog"
	dialogmessages "github.com/mrbelka12000/ai_hack/internal/service/dialog_messages"
	"github.com/mrbelka12000/ai_hack/internal/service/mb"
	"github.com/mrbelka12000/ai_hack/internal/service/user"
	"github.com/mrbelka12000/ai_hack/pkg/redis"
)

type UseCase struct {
	userService            *user.Service
	dialogService          *dialog.Service
	dialogsMessagesService *dialogmessages.Service
	mbService              *mb.Service

	log *slog.Logger
}

func New(r *repo.Repo, log *slog.Logger, rds *redis.Cache, mlClient *ml.Client) *UseCase {
	return &UseCase{
		userService:            user.NewService(r.UserRepo),
		dialogService:          dialog.NewService(r.DialogRepo),
		dialogsMessagesService: dialogmessages.NewService(r.DialogsMessages, mlClient, rds),
		mbService:              mb.New(r.Suggestions),

		log: log,
	}
}

func (uc *UseCase) StartParseMB(filePath string) error {
	return uc.mbService.StartParse(filePath)
}
