package storage

import (
	"github.com/Invan2/invan_marketing_service/pkg/logger"
	"github.com/Invan2/invan_marketing_service/storage/elastic"
	"github.com/Invan2/invan_marketing_service/storage/repo"
	"github.com/elastic/go-elasticsearch/v8"
)

type storageES struct {
	db  *elasticsearch.Client
	log logger.Logger

	clientRepo repo.ClientESI
}

type StorageES interface {
	Client() repo.ClientESI
}

func NewStorageES(log logger.Logger, db *elasticsearch.Client) StorageES {
	return &storageES{
		db:         db,
		log:        log,
		clientRepo: elastic.NewClientRepo(db, log),
	}
}

func (s *storageES) Client() repo.ClientESI {
	return s.clientRepo
}
