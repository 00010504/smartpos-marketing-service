package storage

import (
	"context"
	"database/sql"

	"github.com/Invan2/invan_marketing_service/models"
	"github.com/Invan2/invan_marketing_service/pkg/logger"
	"github.com/Invan2/invan_marketing_service/storage/postgres"
	"github.com/Invan2/invan_marketing_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type repos struct {
	companyRepo repo.CompanyPgI
	userRepo    repo.UserPgI
	shopRepo    repo.ShopI
	clientRepo  repo.ClientI
}

type repoIs interface {
	Company() repo.CompanyPgI
	User() repo.UserPgI
	Shop() repo.ShopI
	Client() repo.ClientI
}

type storage struct {
	db  *sqlx.DB
	log logger.Logger
	repos
}

type storageTr struct {
	tr *sqlx.Tx
	repos
}

type StorageTrI interface {
	Commit() error
	Rollback() error
	repoIs
}

type StoragePg interface {
	WithTransaction() (StorageTrI, error)
	repoIs
}

func NewStoragePg(log logger.Logger, db *sqlx.DB) StoragePg {

	return &storage{
		db:    db,
		log:   log,
		repos: getRepos(log, db),
	}
}

func (s *storage) WithTransaction() (StorageTrI, error) {

	tr, err := s.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return &storageTr{
		tr:    tr,
		repos: getRepos(s.log, tr),
	}, nil
}

func getRepos(log logger.Logger, db models.DB) repos {
	return repos{

		companyRepo: postgres.NewCompanyRepo(log, db),
		userRepo:    postgres.NewUserRepo(log, db),
		shopRepo:    postgres.NewShopRepo(log, db),
		clientRepo:  postgres.NewClientRepo(db, log),
	}
}

func (s *storageTr) Commit() error {
	return s.tr.Commit()
}

func (s *storageTr) Rollback() error {
	return s.tr.Rollback()
}

func (r *repos) Company() repo.CompanyPgI {
	return r.companyRepo
}

func (r *repos) User() repo.UserPgI {
	return r.userRepo
}

func (r *repos) Shop() repo.ShopI {
	return r.shopRepo
}

func (r *repos) Client() repo.ClientI {
	return r.clientRepo
}
