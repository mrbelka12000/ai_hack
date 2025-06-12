package internal

type (
	PaginationParams struct {
		Limit  int `schema:"limit,omitempty" validate:"required"`
		Offset int `schema:"offset,omitempty"`
		Page   int `schema:"page,omitempty"`
	}
)

func (p PaginationParams) CalculatePage() int {
	if p.Limit > 0 && p.Offset >= 0 {
		return (p.Offset / p.Limit) + 1
	}

	return 0
}
