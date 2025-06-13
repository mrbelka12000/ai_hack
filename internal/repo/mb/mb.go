package mb

import (
	"context"

	"github.com/mrbelka12000/ai_hack/internal"
	"github.com/mrbelka12000/ai_hack/pkg/gorm/postgres"
)

type Repo struct {
	db *postgres.Gorm
}

func New(db *postgres.Gorm) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, obj internal.MB) error {
	return r.db.WithContext(ctx).Table("mb").Create(obj).Error
}

func (r *Repo) GetByCustID(ctx context.Context, custID string) ([]internal.MB, error) {
	var results []internal.MB
	err := r.db.WithContext(ctx).
		Table("mb").
		Where("cust_id = ?", custID).
		Find(&results).Error
	return results, err
}
