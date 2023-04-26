package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"genproto/common"
	"genproto/marketing_service"
	"genproto/order_service"
	"io/ioutil"

	"github.com/Invan2/invan_marketing_service/config"
	"github.com/Invan2/invan_marketing_service/pkg/logger"
	"github.com/Invan2/invan_marketing_service/storage/repo"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pkg/errors"
)

type H map[string]interface{}

type clientRepo struct {
	db  *elasticsearch.Client
	log logger.Logger
}

func NewClientRepo(db *elasticsearch.Client, log logger.Logger) repo.ClientESI {
	return &clientRepo{
		db:  db,
		log: log,
	}
}

func (p *clientRepo) Create(req *marketing_service.ShortClient) error {

	if !exists(p.db, config.ElasticClientsIndex) {
		res, err := p.db.Indices.Create(config.ElasticClientsIndex)
		if err != nil {
			return errors.Wrap(err, "error while create index")
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		if res.IsError() {
			p.log.Error("errror while create index", logger.Any("res", string(data)))
			return errors.New("error while create product on elastic")
		}

	}

	var (
		body bytes.Buffer
	)

	err := config.JSONPBMarshaler.Marshal(&body, req)
	if err != nil {
		return errors.Wrap(err, "error while marshaling, jsonpb")
	}

	res, err := p.db.Create(config.ElasticClientsIndex, req.Id, bytes.NewReader(body.Bytes()), p.db.Create.WithRefresh("true"))
	if err != nil {
		return errors.Wrap(err, "Failed to bulk insert products")
	}
	defer res.Body.Close()

	if res.IsError() {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		p.log.Error("errror while create ", logger.Any("res", string(data)))
		return errors.New("error while create product on elastic")
	}

	return nil
}

func makeSearchRequest(req *marketing_service.GetAllClientsRequest) H {

	must := make([]H, 0)
	bool := make(H)
	query := make(H)

	must = append(must, H{
		"term": H{
			"company_id.keyword": H{
				"value": req.Request.Request.CompanyId,
			},
		},
	})

	if req.Request.Search != "" {
		must = append(must, H{
			"query_string": H{
				"query":            getSearchString(req.Request.Search),
				"default_operator": "AND",
				"fields":           []string{"first_name", "last_name", "phone_number"},
			},
		})
	}

	if len(must) > 0 {
		bool["must"] = must
	}

	if len(bool) > 0 {
		query["bool"] = bool
	}

	sort := make([]H, 0)

	if req.SortBy == "total_purchase_amount" {
		sort = append(sort, H{
			"total_purchase_amount": H{
				"order": func() string {
					if req.IsDesc {
						return "desc"
					}
					return "asc"
				}(),
			},
		})
	} else {
		sort = append(sort, H{
			"created_at.keyword": H{
				"order": "desc",
			},
		})
	}

	return H{
		"query": query,
		"sort":  sort,
	}
}

func (p *clientRepo) GetAll(ctx context.Context, in *marketing_service.GetAllClientsRequest) (*marketing_service.GetAllClientsResponse, error) {
	var (
		res = marketing_service.GetAllClientsResponse{
			Data:  make([]*marketing_service.ShortClient, 0),
			Total: 0,
		}
		r    map[string]interface{}
		buf  bytes.Buffer
		size = int(in.Request.Limit)
		from = int((in.Request.Page - 1) * in.Request.Limit)
	)

	req := makeSearchRequest(in)

	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return nil, err
	}

	response, err := p.db.Search(
		p.db.Search.WithContext(context.Background()),
		p.db.Search.WithIndex(config.ElasticClientsIndex),
		p.db.Search.WithBody(&buf),
		p.db.Search.WithTrackTotalHits(true),
		p.db.Search.WithFrom(from),
		p.db.Search.WithSize(size),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while get client documents on elastic")
	}

	if response.IsError() {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		p.log.Error("errror while get all clients ", logger.Any("res", string(data)))
		return nil, errors.New("error while get  cliens on elastic " + string(data))
	}
	if response.StatusCode != 200 {
		return nil, errors.New("error while get documents on elastic. code: " + response.Status())
	}

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return nil, errors.Wrap(err, "error while json.decode elastic res.Body")
	}

	for _, source := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {

		client := marketing_service.ShortClient{}

		jsonString, _ := json.Marshal(source.(map[string]interface{})["_source"])

		err = json.Unmarshal(jsonString, &client)
		if err != nil {
			return nil, errors.Wrap(err, "error while json.Unmarshal jsonString &product")
		}

		res.Data = append(res.Data, &client)
	}

	res.Total = int32(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	return &res, nil

}

func (c *clientRepo) Delete(in *common.RequestID) error {

	res, err := c.db.Delete(config.ElasticClientsIndex, in.Id)
	if err != nil {
		return errors.Wrap(err, "error while delte client from elastic")
	}

	if res.IsError() {

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		c.log.Error("error while  delete clients ", logger.Any("res", string(data)))

		return errors.New("errror while  delete clients")
	}

	return nil

}

func (c *clientRepo) Update(client *marketing_service.ShortClient) error {
	if !exists(c.db, config.ElasticClientsIndex) {
		res, err := c.db.Indices.Create(config.ElasticClientsIndex)
		if err != nil {
			return errors.Wrap(err, "error while create index")
		}

		if err := checkResponseCodeToSuccess(res.StatusCode); err != nil {
			return err
		}
	}

	var (
		updateReq = marketing_service.UpsertClientES{Doc: client, DocAsUpsert: true}
		body      bytes.Buffer
	)

	err := config.JSONPBMarshaler.Marshal(&body, &updateReq)
	if err != nil {
		return errors.Wrap(err, "error while marshaling jsonpb")
	}

	res, err := c.db.Update(config.ElasticClientsIndex, client.Id, bytes.NewReader(body.Bytes()), c.db.Update.WithRefresh("true"))
	if err != nil {
		return errors.Wrap(err, "error while update document on elastic")
	}

	if res.IsError() {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		c.log.Error("errror while update ", logger.Any("res", string(data)))
		return errors.New("error while update client on elastic")
	}

	return nil
}

func (c *clientRepo) AddOrder(ctx context.Context, in *order_service.CreateOrderCopyRequest) (*common.ResponseID, error) {

	body := H{
		"script": H{
			"source": "ctx._source.total_purchase_amount += params.purchase_price",
			"lang":   "painless",
			"params": H{
				"purchase_price": in.TotalPrice,
			},
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := c.db.Update(config.ElasticClientsIndex, in.Client.Id, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		c.log.Error("errror while update client", logger.Any("res", string(data)))
		return nil, errors.New("error while update client on elastic")
	}

	return &common.ResponseID{Id: in.Client.Id}, nil

}
