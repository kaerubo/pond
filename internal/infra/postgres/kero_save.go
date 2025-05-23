package postgres

import (
	"context"
	"github.com/kaerubo/kaeruashi/internal/db/models"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *keroRepository) Save(ctx context.Context, k *entity.Kero) error {
	model := &models.Kero{
		ID:        k.ID,
		Title:     k.Title,
		Content:   k.Content,
		CreatedAt: k.CreatedAt,
		UpdatedAt: k.UpdatedAt,
	}
	return model.Insert(ctx, r.db, boil.Infer())
}
