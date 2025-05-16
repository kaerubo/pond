//go:generate mockgen -source=kero_repository.go -destination=../../internal/test/mock/kero_repository.go -package=mock
package repository

import (
	"context"

	"github.com/kaerubo/pond/internal/entity"
)

type KeroSaver interface {
	Save(ctx context.Context, k *entity.Kero) error
}

type KeroByIDFinder interface {
	FindByID(ctx context.Context, id string) (*entity.Kero, error)
}

type KeroLister interface {
	FindAll(ctx context.Context) ([]*entity.Kero, error)
}

type KeroUpdater interface {
	Update(ctx context.Context, k *entity.Kero) error
}

type KeroDeleter interface {
	Delete(ctx context.Context, id string) error
}
