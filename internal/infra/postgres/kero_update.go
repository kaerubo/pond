package postgres

import (
	"context"
	"github.com/kaerubo/kaeruashi/internal/db/models"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *keroRepository) Update(ctx context.Context, k *entity.Kero) error {
	model, err := models.Keros(qm.Where("id = ?", k.ID)).One(ctx, r.db)
	if err != nil {
		return err
	}

	model.Title = k.Title
	model.Content = k.Content
	model.UpdatedAt = k.UpdatedAt

	_, err = model.Update(ctx, r.db, boil.Infer())
	return err
}
