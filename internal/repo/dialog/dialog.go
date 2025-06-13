package dialog

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
	"github.com/mrbelka12000/ai_hack/pkg/gorm/postgres"
)

type Repo struct {
	db *postgres.Gorm
}

func New(db *postgres.Gorm) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Create(ctx context.Context, obj internal.DialogCU) error {
	return r.db.WithContext(ctx).Table("dialogs").Create(&obj).Error
}

func (r *Repo) Update(ctx context.Context, obj internal.Dialog) error {
	return r.db.WithContext(ctx).Table("dialogs").Model(&internal.Dialog{}).Where("id = ?", obj.ID).Updates(obj).Error
}

func (r *Repo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Table("dialogs").Where("id = ?", id).Delete(&internal.Dialog{}).Error
}

func (r *Repo) Get(ctx context.Context, id uuid.UUID) (internal.Dialog, error) {
	var d internal.Dialog
	err := r.db.WithContext(ctx).Table("dialogs").Where("id = ?", id).First(&d).Error
	return d, err
}

func (r *Repo) List(ctx context.Context, pars internal.DialogPars) ([]internal.Dialog, error) {
	var dialogs []internal.Dialog
	query := r.db.WithContext(ctx).Table("dialogs").Model(&internal.Dialog{})

	if pars.ClientID != 0 {
		query = query.Where("client_id = ?", pars.ClientID)
	}
	if pars.OperatorID != 0 {
		query = query.Where("operator_id = ?", pars.OperatorID)
	}
	if pars.Status != "" {
		query = query.Where("status = ?", pars.Status)
	}
	if !pars.CreatedBefore.IsZero() {
		query = query.Where("created_at < ?", pars.CreatedBefore)
	}
	if !pars.CreatedAfter.IsZero() {
		query = query.Where("created_at > ?", pars.CreatedAfter)
	}
	if pars.Limit > 0 {
		query = query.Limit(pars.Limit)
	}
	if pars.Offset > 0 {
		query = query.Offset(pars.Offset)
	}

	err := query.Order("created_at DESC").Find(&dialogs).Error
	return dialogs, err
}
