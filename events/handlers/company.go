package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (e *EventHandler) CreateCompany(ctx context.Context, event *kafka.Message) error {

	var (
		req common.CompanyCreatedModel
	)

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return errors.Wrap(err, "error while unmarshaling company")
	}

	if err := e.strgPG.Company().Upsert(&req); err != nil {

		return err
	}

	return nil

}
