//go:generate mockgen -source=kero_usecase.go -destination=mock/kero_usecase.go -package=mock
package usecase

import (
	"context"
	"github.com/kaerubo/kaeruashi/internal/entity"
)

type KeroCreator interface {
	Create(ctx context.Context, k *entity.Kero) error
}

type KeroByIDFinder interface {
	FindByID(ctx context.Context, id string) (*entity.Kero, error)
}

type KeroLister interface {
	List(ctx context.Context) ([]*entity.Kero, error)
}

type KeroUpdater interface {
	Update(ctx context.Context, k *entity.Kero) error
}

type KeroDeleter interface {
	Delete(ctx context.Context, id string) error
}
