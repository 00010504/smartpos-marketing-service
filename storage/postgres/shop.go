package postgres

import (
	"genproto/common"

	"github.com/Invan2/invan_marketing_service/models"
	"github.com/Invan2/invan_marketing_service/pkg/logger"
	"github.com/Invan2/invan_marketing_service/storage/repo"
	"github.com/pkg/errors"
)

type shopRepo struct {
	db  models.DB
	log logger.Logger
}

func NewShopRepo(log logger.Logger, db models.DB) repo.ShopI {
	return &shopRepo{
		db:  db,
		log: log,
	}
}

func (s *shopRepo) Upsert(entity *common.ShopCreatedModel) error {

	query := `
		INSERT INTO
			"shop"
		(
			id,
			name,
			company_id,
			created_by
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		) ON CONFLICT (id) DO
		UPDATE
			SET
			"name" = $2,
			"company_id" = $3
			;
	`

	_, err := s.db.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.Request.CompanyId,
		entity.Request.UserId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert shop")
	}

	return nil
}
