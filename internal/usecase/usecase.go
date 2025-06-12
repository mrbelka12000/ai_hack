package usecase

import (
	"log/slog"

	"github.com/mrbelka12000/ai_hack/internal/repo"
	"github.com/mrbelka12000/ai_hack/internal/service/dialog"
	dialogmessages "github.com/mrbelka12000/ai_hack/internal/service/dialog_messages"
	"github.com/mrbelka12000/ai_hack/internal/service/user"
)

type UseCase struct {
	userService            *user.Service
	dialogService          *dialog.Service
	dialogsMessagesService *dialogmessages.Service

	log *slog.Logger
}

func New(r *repo.Repo, log *slog.Logger) *UseCase {
	return &UseCase{
		userService:            user.NewService(r.UserRepo),
		dialogService:          dialog.NewService(r.DialogRepo),
		dialogsMessagesService: dialogmessages.NewService(r.DialogsMessages),
		
		log: log,
	}
}
