package usecase

import (
	"context"
	"github.com/kaerubo/pond/internal/entity"
	"github.com/kaerubo/pond/internal/repository"
)

type keroLister struct {
	repo repository.KeroLister
}

func NewKeroLister(repo repository.KeroLister) KeroLister {
	return &keroLister{repo: repo}
}

func (l *keroLister) List(ctx context.Context) ([]*entity.Kero, error) {
	return l.repo.FindAll(ctx)
}
