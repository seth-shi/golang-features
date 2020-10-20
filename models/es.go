package models

import (
	"github.com/elastic/go-elasticsearch/v7"
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