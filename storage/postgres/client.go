package postgres

import (
	"context"
	"database/sql"
	"genproto/common"
	"genproto/marketing_service"

	"github.com/Invan2/invan_marketing_service/models"
	"github.com/Invan2/invan_marketing_service/pkg/logger"
	"github.com/Invan2/invan_marketing_service/storage/repo"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type clientRepo struct {
	db  models.DB
	log logger.Logger
}

func NewClientRepo(db models.DB, log logger.Logger) repo.ClientI {
	return &clientRepo{
		db:  db,
		log: log,
	}
}

func (c *clientRepo) Create(ctx context.Context, in *marketing_service.CreateClientRequest) (*common.ResponseID, string, error) {

	var (
		clientID   = uuid.NewString()
		externalID string
	)

	query := `
		INSERT INTO "client" (
			"id",
			"company_id",
			"sex_id",
			"first_name",
			"last_name",
			"email",
			"phone_number",
			"info",
			"birthday",
			"card_number",
			"created_by"
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11
		) RETURNING external_id
	`

	err := c.db.QueryRow(query,
		clientID,
		in.Request.CompanyId,
		in.SexId,
		in.FirstName,
		in.LastName,
		in.Email,
		in.PhoneNumber,
		in.Info,
		in.Birthday,
		in.CardNumber,
		in.Request.UserId).Scan(&externalID)
	if err != nil {
		return nil, externalID, err
	}

	return &common.ResponseID{Id: clientID}, externalID, nil
}
func (c *clientRepo) GetAll(ctx context.Context, in *marketing_service.GetAllClientsRequest) (*marketing_service.GetAllClientsResponse, error) {

	// var (
	// 	params = map[string]interface{}{
	// 		"company_id": in.Request.Request.CompanyId,
	// 		"search":     in.Request.Search,
	// 		"limit":      in.Request.Limit,
	// 		"offset":     (in.Request.Page - 1) * in.Request.Limit,
	// 	}
	// )

	// query := ``

	return nil, nil
}
func (c *clientRepo) Update(ctx context.Context, in *marketing_service.UpdateClientRequest) (*common.ResponseID, string, error) {

	var externalID string

	query := `
	UPDATE "client" SET 
		"sex_id"=$3,
		"first_name"=$4,
		"last_name"=$5,
		"email"=$6,
		"phone_number"=$7,
		"info"=$8,
		"birthday"=$9,
		"card_number"=$10
	WHERE id=$1 AND company_id=$2 AND deleted_at=0 RETURNING external_id

`

	err := c.db.QueryRow(query, in.Id, in.Request.CompanyId, in.SexId, in.FirstName, in.LastName, in.Email, in.PhoneNumber, in.Info, in.Birthday, in.CardNumber).Scan(&externalID)
	if err != nil {
		return nil, "", err
	}

	return &common.ResponseID{Id: in.Id}, externalID, nil
}

func (c *clientRepo) Delete(ctx context.Context, in *common.RequestID) (*common.ResponseID, error) {
	query := `
	UPDATE "client" SET 
		deleted_at=extract(epoch from now())::bigint
	WHERE id=$1 AND company_id=$2
	`

	res, err := c.db.Exec(query, in.Id, in.Request.CompanyId)
	if err != nil {
		return nil, err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return nil, sql.ErrNoRows
	}

	return &common.ResponseID{Id: in.Id}, nil
}

func (c *clientRepo) GetById(ctx context.Context, in *common.RequestID) (*marketing_service.GetClientByIDResponse, error) {
	var (
		res = marketing_service.GetClientByIDResponse{
			Sex: &marketing_service.IdName{},
		}
		user models.NullShortUser
	)

	query := `
		SELECT
			c."id",
			c."first_name",
			c."last_name",
			c."email",
			c."phone_number",
			c."info",
			c."birthday",
			c."card_number",
			u.id,
			u.first_name,
			u.last_name,
			s.id,
			s.name
		FROM "client" c
		LEFT JOIN "user" u ON u."id"=c."created_by" AND u."deleted_at"=0
		LEFT JOIN "sex" s ON s.id=c.sex_id
		WHERE c."id" = $1 AND c."company_id"=$2 AND c."deleted_at"=0	
	`

	err := c.db.QueryRow(query, in.Id, in.Request.CompanyId).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.PhoneNumber,
		&res.Info,
		&res.Birthday,
		&res.CardNumber,
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&res.Sex.Id,
		&res.Sex.Name,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while ")
	}

	if user.ID.Valid {
		res.CreatedBy = &common.ShortUser{
			Id:        user.ID.String,
			FirstName: user.FirstName.String,
			LastName:  user.LastName.String,
		}
	}
	return &res, nil
}
