package listeners

import (
	"context"
	"genproto/common"
	"genproto/marketing_service"
	"genproto/order_service"

	"github.com/Invan2/invan_marketing_service/config"
	"github.com/Invan2/invan_marketing_service/events"
	"github.com/Invan2/invan_marketing_service/pkg/logger"
	"github.com/Invan2/invan_marketing_service/storage"
	"github.com/minio/minio-go/v7"
)

type marketingService struct {
	log    logger.Logger
	kafka  events.PubSubServer
	strg   storage.StoragePg
	strgES storage.StorageES
	minio  *minio.Client
	cfg    *config.Config
}

type MarketingService interface {
	CreateClient(ctx context.Context, in *marketing_service.CreateClientRequest) (*common.ResponseID, error)
	GetAllClients(ctx context.Context, in *marketing_service.GetAllClientsRequest) (*marketing_service.GetAllClientsResponse, error)
	UpdateClient(ctx context.Context, in *marketing_service.UpdateClientRequest) (*common.ResponseID, error)
	DeleteClient(ctx context.Context, in *common.RequestID) (*common.ResponseID, error)
	GetClientByID(ctx context.Context, in *common.RequestID) (*marketing_service.GetClientByIDResponse, error)

	// order

	AddOrder(ctx context.Context, in *order_service.CreateOrderCopyRequest) (*common.ResponseID, error)
}

func NewMarketingService(log logger.Logger, kafka events.PubSubServer, strg storage.StoragePg, elastic storage.StorageES, minio *minio.Client, cfg *config.Config) MarketingService {
	return &marketingService{
		log:    log,
		kafka:  kafka,
		strg:   strg,
		strgES: elastic,
		minio:  minio,
		cfg:    cfg,
	}
}
