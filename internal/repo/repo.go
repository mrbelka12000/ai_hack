package repo

import (
	"github.com/mrbelka12000/ai_hack/internal/repo/dialog"
	dialogmessages "github.com/mrbelka12000/ai_hack/internal/repo/dialog_messages"
	"github.com/mrbelka12000/ai_hack/internal/repo/personal_data"
	"github.com/mrbelka12000/ai_hack/internal/repo/user"
	"github.com/mrbelka12000/ai_hack/pkg/gorm/postgres"
)

type Repo struct {
	UserRepo        *user.Repo
	DialogRepo      *dialog.Repo
	DialogsMessages *dialogmessages.Repo
	Suggestions     *personal_data.Repo
}

func New(db *postgres.Gorm) *Repo {
	return &Repo{
		UserRepo:        user.New(db),
		DialogRepo:      dialog.New(db),
		DialogsMessages: dialogmessages.New(db),
		Suggestions:     personal_data.New(db),
	}
}
