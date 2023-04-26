package listeners_test

import (
	"os"
	"testing"

	"github.com/Invan2/invan_marketing_service/config"
	"github.com/Invan2/invan_marketing_service/pkg/logger"
	"github.com/Invan2/invan_marketing_service/services/listeners"
	"github.com/Invan2/invan_marketing_service/storage"
)

var (
	marketingService listeners.MarketingService
)

func TestMain(m *testing.M) {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, cfg.Environment)

	marketingService = listeners.NewMarketingService(log, nil, storage.NewStoragePg(log, nil), storage.NewStorageES(log, nil), nil, &cfg)

	os.Exit(m.Run())
}
