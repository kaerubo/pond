package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/kaerubo/pond/internal/entity"
	"github.com/kaerubo/pond/internal/repository"
	"time"
)

type keroCreator struct {
	repo repository.KeroSaver
}

func NewKeroCreator(repo repository.KeroSaver) KeroCreator {
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

	return u.repo.Save(ctx, k)
}
