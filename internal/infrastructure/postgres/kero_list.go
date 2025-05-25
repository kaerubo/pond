package postgres

import (
	"context"
	"github.com/kaerubo/kaeruashi/internal/db/models"
	"github.com/kaerubo/kaeruashi/internal/entity"
)

func (r *KeroRepository) FindAll(ctx context.Context) ([]*entity.Kero, error) {
	modelsList, err := models.Keros().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var result []*entity.Kero
	for _, m := range modelsList {
		result = append(result, &entity.Kero{
			ID:        m.ID,
			Title:     m.Title,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}
	return result, nil
}
