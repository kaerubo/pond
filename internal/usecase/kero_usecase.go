package usecase

import (
	"context"

	"github.com/kaerubo/pond/internal/entity"
)

type KeroCreator interface {
	Create(ctx context.Context, k *entity.Kero) error
}

type KeroReader interface {
	GetByID(ctx context.Context, id int64) (*entity.Kero, error)
	List(ctx context.Context) ([]*entity.Kero, error)
}

type KeroUpdater interface {
	Update(ctx context.Context, k *entity.Kero) error
}

type KeroDeleter interface {
	Delete(ctx context.Context, id int64) error
}
