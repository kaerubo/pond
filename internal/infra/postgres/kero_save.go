package postgres

import (
	"context"
	"database/sql"
	"github.com/kaerubo/pond/internal/db/models"
	"github.com/kaerubo/pond/internal/entity"
	"github.com/kaerubo/pond/internal/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type KeroSave struct {
	db *sql.DB
}

func NewKeroSave(db *sql.DB) repository.KeroSaver {
	return &KeroSave{db: db}
}

func (s *KeroSave) Save(ctx context.Context, k *entity.Kero) error {
	model := &models.Kero{
		ID:        k.ID,
		Title:     k.Title,
		Content:   k.Content,
		CreatedAt: k.CreatedAt,
		UpdatedAt: k.UpdatedAt,
	}

	return model.Insert(ctx, s.db, boil.Infer())
}
