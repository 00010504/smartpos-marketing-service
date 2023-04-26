package listeners

import (
	"context"
	"errors"
	"genproto/common"
	"genproto/marketing_service"
	"time"

	"github.com/Invan2/invan_marketing_service/config"
)

func (m *marketingService) CreateClient(ctx context.Context, in *marketing_service.CreateClientRequest) (*common.ResponseID, error) {

	_, ok := config.SexTypeMap[in.SexId]
	if !ok {
		return nil, errors.New("invalid gender id")
	}

	tr, err := m.strg.WithTransaction()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, externalID, err := tr.Client().Create(ctx, in)
	if err != nil {
		return nil, err
	}

	client := &marketing_service.ShortClient{
		Id:        res.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Group:     &marketing_service.ClientGroup{},
		CompanyId: in.Request.CompanyId,
		Sex: &marketing_service.IdName{
			Id:   in.SexId,
			Name: config.SexTypeMap[in.SexId],
		},
		PhoneNumber: in.PhoneNumber,
		ExternalId:  externalID,
		CreatedAt:   time.Now().UTC().String(),
		CardNumber:  in.CardNumber,
	}

	err = m.kafka.Push("v1.marketing_service.client.upsert", client)
	if err != nil {
		return nil, err
	}

	err = m.strgES.Client().Create(client)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (m *marketingService) GetAllClients(ctx context.Context, in *marketing_service.GetAllClientsRequest) (*marketing_service.GetAllClientsResponse, error) {
	return m.strgES.Client().GetAll(ctx, in)
}
func (m *marketingService) UpdateClient(ctx context.Context, in *marketing_service.UpdateClientRequest) (*common.ResponseID, error) {
	_, ok := config.SexTypeMap[in.SexId]
	if !ok {
		return nil, errors.New("invalid sex id")
	}

	tr, err := m.strg.WithTransaction()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, externalID, err := tr.Client().Update(ctx, in)
	if err != nil {
		return nil, err
	}

	client := &marketing_service.ShortClient{
		Id:        res.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Group:     &marketing_service.ClientGroup{},
		CompanyId: in.Request.CompanyId,
		Sex: &marketing_service.IdName{
			Id:   in.SexId,
			Name: config.SexTypeMap[in.SexId],
		},
		PhoneNumber: in.PhoneNumber,
		ExternalId:  externalID,
		CreatedAt:   time.Now().UTC().String(),
		CardNumber:  in.CardNumber,
	}

	err = m.kafka.Push("v1.marketing_service.client.upsert", client)
	if err != nil {
		return nil, err
	}

	err = m.strgES.Client().Update(client)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (m *marketingService) DeleteClient(ctx context.Context, in *common.RequestID) (*common.ResponseID, error) {
	tr, err := m.strg.WithTransaction()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, err := tr.Client().Delete(ctx, in)
	if err != nil {
		return nil, err
	}

	err = m.kafka.Push("v1.marketing_service.client.deleted", in)
	if err != nil {
		return nil, err
	}

	err = m.strgES.Client().Delete(in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *marketingService) GetClientByID(ctx context.Context, in *common.RequestID) (*marketing_service.GetClientByIDResponse, error) {
	return m.strg.Client().GetById(ctx, in)
}
