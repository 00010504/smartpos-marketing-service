package config

import "github.com/golang/protobuf/jsonpb"

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	ConsumerGroupID     = "invan_marketing_service"
	ElasticClientsIndex = "clients"
)

var (
	JSONPBMarshaler = jsonpb.Marshaler{EmitDefaults: true, OrigName: true}

	SexTypeMap = map[string]string{
		"1fe92aa8-2a61-4bf1-b907-182b497584ad": "Male",
		"9fb3ada6-a73b-4b81-9295-5c1605e54552": "Female",
	}
)
