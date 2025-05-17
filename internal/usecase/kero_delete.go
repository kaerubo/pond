package usecase

import (
	"context"
	"errors"

	"github.com/kaerubo/pond/internal/repository"
)

type keroDeleter struct {
	repo repository.KeroDeleter
}

func NewKeroDeleter(repo repository.KeroDeleter) KeroDeleter {
	return &keroDeleter{repo: repo}
}

func (d *keroDeleter) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return d.repo.Delete(ctx, id)
}
