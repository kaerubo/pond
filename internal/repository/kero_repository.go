//go:generate mockgen -source=kero_repository.go -destination=mock/kero_repository.go -package=mock
package repository

import (
	"context"
	"github.com/kaerubo/kaeruashi/internal/entity"
)

type KeroInserter interface {
	Insert(ctx context.Context, k *entity.Kero) error
}

type KeroByIDFinder interface {
	FindByID(ctx context.Context, id string) (*entity.Kero, error)
}

type KeroFinder interface {
	FindAll(ctx context.Context) ([]*entity.Kero, error)
}

type KeroUpdater interface {
	Update(ctx context.Context, k *entity.Kero) error
}

type KeroDeleter interface {
	Delete(ctx context.Context, id string) error
}
