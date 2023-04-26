package listeners

import (
	"context"
	"genproto/common"
	"genproto/order_service"
)

func (m *marketingService) AddOrder(ctx context.Context, in *order_service.CreateOrderCopyRequest) (*common.ResponseID, error) {

	return m.strgES.Client().AddOrder(ctx, in)
}
