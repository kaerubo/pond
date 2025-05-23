package postgres

import (
	"context"
	"github.com/kaerubo/kaeruashi/internal/db/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *keroRepository) Delete(ctx context.Context, id string) error {
	model, err := models.Keros(qm.Where("id = ?", id)).One(ctx, r.db)
	if err != nil {
		return err
	}
	_, err = model.Delete(ctx, r.db)
	return err
}
