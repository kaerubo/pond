package usecase

import (
	"context"
	"errors"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/repository"
)

type keroByIDFinder struct {
	repo repository.KeroByIDFinder
}

func NewKeroByIDFinder(repo repository.KeroByIDFinder) KeroByIDFinder {
	return &keroByIDFinder{repo: repo}
}

func (r *keroByIDFinder) FindByID(ctx context.Context, id string) (*entity.Kero, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return r.repo.FindByID(ctx, id)
}
