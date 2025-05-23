package usecase

import (
	"context"
	"errors"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/repository"
	"time"
)

type keroUpdater struct {
	repo repository.KeroUpdater
}

func NewKeroUpdater(repo repository.KeroUpdater) KeroUpdater {
	return &keroUpdater{repo: repo}
}

func (u *keroUpdater) Update(ctx context.Context, k *entity.Kero) error {
	if k.ID == "" {
		return errors.New("id is required")
	}
	if k.Title == "" {
		return errors.New("title is required")
	}
	if k.Content == "" {
		return errors.New("content is required")
	}

	k.UpdatedAt = time.Now()
	return u.repo.Update(ctx, k)
}
