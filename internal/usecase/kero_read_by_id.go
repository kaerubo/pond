package usecase

import (
	"context"
	"errors"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/repository"
)

type keroReader struct {
	repo repository.KeroByIDFinder
}

func NewKeroByIDReader(repo repository.KeroByIDFinder) KeroByIDReader {
	return &keroReader{repo: repo}
}

func (r *keroReader) GetByID(ctx context.Context, id string) (*entity.Kero, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return r.repo.FindByID(ctx, id)
}
