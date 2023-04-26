package repo

import (
	"context"
	"genproto/common"
	"genproto/marketing_service"
)

type ClientI interface {
	Create(ctx context.Context, in *marketing_service.CreateClientRequest) (*common.ResponseID, string, error)
	GetAll(ctx context.Context, in *marketing_service.GetAllClientsRequest) (*marketing_service.GetAllClientsResponse, error)
	Update(ctx context.Context, in *marketing_service.UpdateClientRequest) (*common.ResponseID, string, error)
	Delete(ctx context.Context, in *common.RequestID) (*common.ResponseID, error)
	GetById(ctx context.Context, in *common.RequestID) (*marketing_service.GetClientByIDResponse, error)
}
