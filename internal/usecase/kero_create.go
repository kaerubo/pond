package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/repository"
	"time"
)

type keroCreator struct {
	repo repository.KeroInserter
}

func NewKeroCreator(repo repository.KeroInserter) KeroCreator {
	return &keroCreator{repo: repo}
}

func (u *keroCreator) Create(ctx context.Context, k *entity.Kero) error {
	if k.Title == "" {
		return errors.New("title is required")
	}
	if k.Content == "" {
		return errors.New("content is required")
	}

	k.ID = uuid.New().String()
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now

	return u.repo.Insert(ctx, k)
}
