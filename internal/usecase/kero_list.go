package usecase

import (
	"context"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/kaerubo/kaeruashi/internal/repository"
)

type keroLister struct {
	repo repository.KeroFinder
}

func NewKeroLister(repo repository.KeroFinder) KeroLister {
	return &keroLister{repo: repo}
}

func (l *keroLister) List(ctx context.Context) ([]*entity.Kero, error) {
	return l.repo.FindAll(ctx)
}
