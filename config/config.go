package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string
	ServiceName string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	ElasticSearchUrls     []string
	ElasticSearchUser     string
	ElasticSearchPassword string

	MinioAccessKeyID string
	MinioSecretKey   string
	MinioEndpoint    string

	LogLevel string

	HttpHost string
	HttpPort string
	KafkaUrl string
}

func Load() Config {
	envFileName := cast.ToString(getOrReturnDefault("ENV_FILE_PATH", "./app/.env"))

	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "info"))
	config.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "billz_catalog_service_v2"))

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "root"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "invan_catalog_service"))

	config.ElasticSearchUrls = strings.Split(cast.ToString(getOrReturnDefault("ELASTIC_SEARCH_URLS", "http://localhost:9200")), ",")
	config.ElasticSearchUser = cast.ToString(getOrReturnDefault("ELASTIC_SEARCH_USER", "elastic"))
	config.ElasticSearchPassword = cast.ToString(getOrReturnDefault("ELASTIC_SEARCH_PASSWORD", "changeme"))

	config.KafkaUrl = cast.ToString(getOrReturnDefault("KAFKA_URL", "localhost:9092"))

	config.HttpPort = cast.ToString(getOrReturnDefault("GRPC_PORT", ":8008"))
	config.HttpHost = cast.ToString(getOrReturnDefault("LISTEN_HOST", "localhost"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
