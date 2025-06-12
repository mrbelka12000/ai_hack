package user

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

func (r *Repo) Create(ctx context.Context, obj internal.UserCU) error {
	return r.db.WithContext(ctx).Table("users").Create(&obj).Error
}

func (r *Repo) Update(ctx context.Context, obj internal.UserCU) error {
	return r.db.WithContext(ctx).
		Model(&internal.User{}).
		Table("users").
		Where("email = ?", obj.Email).
		Updates(map[string]interface{}{
			"password":   obj.Password,
			"role":       obj.Role,
			"created_at": obj.CreatedAt,
		}).Error
}

func (r *Repo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&internal.User{}).Error
}

func (r *Repo) Get(ctx context.Context, pars internal.UserGetPars) (internal.User, error) {
	var u internal.User
	db := r.db.WithContext(ctx).Table("users")

	if pars.ID != 0 {
		db = db.Where("id = ?", pars.ID)
	}
	if pars.Email != "" {
		db = db.Where("email = ?", pars.Email)
	}
	if pars.Role != "" {
		db = db.Where("role = ?", pars.Role)
	}

	err := db.First(&u).Error
	return u, err
}

func (r *Repo) List(ctx context.Context, pars internal.UserPars) ([]internal.User, error) {
	var users []internal.User
	query := r.db.WithContext(ctx).Table("users").Model(&internal.User{})

	if pars.Email != "" {
		query = query.Where("email = ?", pars.Email)
	}
	if pars.Role != "" {
		query = query.Where("role = ?", pars.Role)
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

	err := query.Order("created_at DESC").Find(&users).Error
	return users, err
}
