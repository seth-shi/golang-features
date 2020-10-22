package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"os"
)

var es *elasticsearch.Client
var indexModels = []Model{
	&Feature{},
}

func InitEs() {

	var err error

	es, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_HOST"),
		},
	})

	if err != nil {
		panic(err)
	}

	_, err = es.Ping()
	if err != nil {
		panic(err)
	}

	//// create all index
	//for _, model := range indexModels {
	//
	//	exists, err := es.IndexExists(model.IndexName()).Do(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	if !exists {
	//		_, err = es.CreateIndex(model.IndexName()).BodyString(model.Mapping()).Do(ctx)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//}
}

func parseEsResponse(res *esapi.Response) (map[string]interface{}, error) {

	var err error
	if res.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		}

		// 可能是未找到
		if _, exists := e["found"]; exists {
			return nil, errors.New("no found model")
		}

		// Print the response status and error information.
		err = errors.New(fmt.Sprintf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		))
		return nil, err
	}

	var r map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r, nil
}