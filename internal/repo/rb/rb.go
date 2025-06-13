package rb

import "github.com/mrbelka12000/ai_hack/pkg/gorm/postgres"

type Repo struct {
	db *postgres.Gorm
}

func New(db *postgres.Gorm) *Repo {
	return &Repo{db: db}
}
