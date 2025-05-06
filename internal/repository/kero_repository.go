package repository

import (
	"context"

	"github.com/kaerubo/pond/internal/entity"
)

type KeroSaver interface {
	Save(ctx context.Context, k *entity.Kero) error
}

type KeroFinder interface {
	FindByID(ctx context.Context, id int64) (*entity.Kero, error)
	FindAll(ctx context.Context) ([]*entity.Kero, error)
}

type KeroUpdater interface {
	Update(ctx context.Context, k *entity.Kero) error
}

type KeroDeleter interface {
	Delete(ctx context.Context, id int64) error
}
