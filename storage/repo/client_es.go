package repo

import (
	"context"
	"genproto/common"
	"genproto/marketing_service"
	"genproto/order_service"
)

type ClientESI interface {
	Create(*marketing_service.ShortClient) error
	GetAll(ctx context.Context, in *marketing_service.GetAllClientsRequest) (*marketing_service.GetAllClientsResponse, error)
	Delete(in *common.RequestID) error
	Update(client *marketing_service.ShortClient) error
	AddOrder(ctx context.Context, in *order_service.CreateOrderCopyRequest) (*common.ResponseID, error)
}
