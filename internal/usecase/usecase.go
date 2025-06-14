package usecase

import (
	"log/slog"
	"time"

	"github.com/mrbelka12000/ai_hack/internal/client/ml"
	"github.com/mrbelka12000/ai_hack/internal/repo"
	"github.com/mrbelka12000/ai_hack/internal/service/dialog"
	dialogmessages "github.com/mrbelka12000/ai_hack/internal/service/dialog_messages"
	"github.com/mrbelka12000/ai_hack/internal/service/personal_data"
	"github.com/mrbelka12000/ai_hack/internal/service/user"
	"github.com/mrbelka12000/ai_hack/pkg/redis"
)

type (
	UseCase struct {
		userService            *user.Service
		dialogService          *dialog.Service
		dialogsMessagesService *dialogmessages.Service
		personalData           *personal_data.Service
		cache                  cache

		log *slog.Logger
	}

	cache interface {
		Set(key string, value interface{}, dur time.Duration) error
		Get(key string) (string, bool)
		Delete(key string)
	}
)

func New(r *repo.Repo, log *slog.Logger, rds *redis.Cache, mlClient *ml.Client) *UseCase {
	return &UseCase{
		userService:            user.NewService(r.UserRepo),
		dialogService:          dialog.NewService(r.DialogRepo, log),
		dialogsMessagesService: dialogmessages.NewService(r.DialogsMessages, mlClient, rds),
		personalData:           personal_data.New(r.Suggestions),

		cache: rds,

		log: log,
	}
}

func (uc *UseCase) StartParseMB(filePath string) error {
	return uc.personalData.StartParseMB(filePath)
}

func (uc *UseCase) StartParseRB(filePath string) error {
	return uc.personalData.StartParseRB(filePath)
}
