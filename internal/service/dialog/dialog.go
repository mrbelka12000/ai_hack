package dialog

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

type (
	Service struct {
		repo  repo
		cache cache

		log *slog.Logger
	}
)

const (
	defaultTTL  = time.Minute * 5
	cachePrefix = "dialog-"
)

func NewService(repo repo, cache cache, log *slog.Logger) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
		log:   log,
	}
}

func (s *Service) Create(ctx context.Context, obj internal.DialogCU) error {
	return s.repo.Create(ctx, obj)
}

func (s *Service) Update(ctx context.Context, obj internal.Dialog) error {
	s.cache.Delete(cachePrefix + obj.ID.String())

	return s.repo.Update(ctx, obj)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (obj internal.Dialog, err error) {
	rawObject, ok := s.cache.Get(cachePrefix + id.String())
	if ok {

		if err = json.Unmarshal([]byte(rawObject), &obj); err != nil {
			return obj, err
		}

		return obj, nil
	}

	obj, err = s.repo.Get(ctx, id)
	if err != nil {
		return internal.Dialog{}, err
	}

	if err := json.Unmarshal(obj.RawData, &obj.Data); err != nil {
		return internal.Dialog{}, err
	}

	raw, err := json.Marshal(obj)
	if err == nil {
		if err = s.cache.Set(cachePrefix+id.String(), string(raw), defaultTTL); err != nil {
			s.log.With("error", err).Info("cache set")
		}
	}

	return obj, nil
}

func (s *Service) List(ctx context.Context, pars internal.DialogPars) ([]internal.Dialog, error) {
	return s.repo.List(ctx, pars)
}
