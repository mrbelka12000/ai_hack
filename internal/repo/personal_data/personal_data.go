package personal_data

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/ai_hack/internal"
	"github.com/mrbelka12000/ai_hack/pkg/gorm/postgres"
)

type Repo struct {
	db *postgres.Gorm
}

func New(db *postgres.Gorm) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, obj internal.PersonalData) error {
	return r.db.WithContext(ctx).Table("personal_data").Create(obj).Error
}

func (r *Repo) GetPersonalDataForResponse(ctx context.Context, obj internal.PersonalDataRequest) (string, error) {
	var (
		validColumns = map[string]string{
			"call-id":        "call_id",
			"phone-number":   "phone_number",
			"br":             "br",
			"currency":       "currency",
			"beg-date":       "beg_date",
			"end-date":       "end_date",
			"prol-date":      "prol_date",
			"prol-count":     "prol_count",
			"amt":            "amt",
			"amt-tng":        "amt_tng",
			"od":             "od",
			"pr-od":          "pr_od",
			"day-pr-od":      "day_pr_od",
			"pog":            "pog",
			"stav":           "stav",
			"sht":            "sht",
			"br-vyd":         "br_vyd",
			"flwork":         "flwork",
			"rate_effective": "rate_effective",
		}

		result string
	)

	column, ok := validColumns[obj.DataType]
	if !ok {
		return "", fmt.Errorf("invalid column: %s", obj.DataType)
	}

	var whereQuery string

	if obj.CallID != "" {
		whereQuery = fmt.Sprintf("WHERE call_id = '%s'", obj.CallID)
	} else if obj.PhoneNumber != "" {
		whereQuery = fmt.Sprintf("WHERE phone_number = '%s'", obj.PhoneNumber)
	}

	err := r.db.WithContext(ctx).
		Table("personal_data").
		Raw(fmt.Sprintf("SELECT %v FROM personal_data %s", column, whereQuery)).
		Scan(&result).Error

	return result, err
}
