package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/kaerubo/kaeruashi/internal/db/models"
	"github.com/kaerubo/kaeruashi/internal/entity"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *keroRepository) FindByID(ctx context.Context, id string) (*entity.Kero, error) {
	model, err := models.Keros(
		qm.Where("id = ?", id),
	).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &entity.Kero{
		ID:        model.ID,
		Title:     model.Title,
		Content:   model.Content,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}
