package models

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"golang-functions/utils"
	"time"
)

type Feature struct {
	Id          string
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	Code        string `mapstructure:"code"`

	Published bool `mapstructure:"published"`
	ViewCount int `mapstructure:"view_count"`

	CreatedAt time.Time `mapstructure:"created_at"`
	UpdatedAt time.Time `mapstructure:"updated_at"`
}

func (f Feature) IndexName() string {
	return "features"
}

func (f Feature) Mapping() string {

	return `{
	"settings": {
		"number_of_replicas": 0
	},
	"mappings": {
		"properties": {
			"title": {
				"type": "text",
				"analyzer": "ik_max_word",
				"search_analyzer": "ik_smart",
				"index": true
			},
			"description": {
				"type": "text",
				"analyzer": "ik_max_word",
				"search_analyzer": "ik_smart",
				"index": true
			},
			"code": {
				"type": "text",
				"index": false
			},
			"view_count": {
				"type": "integer"
			},
			"published": {
				"type": "boolean"
			},
			"created_at": {
				"type": "date",
				"format": "yyyy-MM-dd HH:mm:ss"
			},
			"updated_at": {
				"type": "date",
				"format": "yyyy-MM-dd HH:mm:ss"
			}
		}
	}
}`
}

func (f Feature) Search(offset, limit int) (models []*Feature, count int, err error) {

	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{

			},
		},
	}
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(f.IndexName()),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithFrom(offset),
		es.Search.WithSize(limit),
		es.Search.WithPretty(),
	)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := parseEsResponse(res)
	if err != nil {
		return
	}

	// Print the response status, number of results, and request duration.
	count = int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {

		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		var f Feature
		err := utils.Decode(source, &f)
		if err != nil {
			continue
		}
		f.Id = hit.(map[string]interface{})["_id"].(string)

		models = append(models, &f)
	}

	return
}

func (f Feature) GetId() string {
	return f.Id
}


func (f Feature) GetCreatedAt() time.Time {
	return f.CreatedAt
}

func (f Feature) Find() (*Feature, error) {

	m, err := Find(f)
	if err != nil {
		return nil, err
	}

	source, ok := m["_source"].(map[string]interface{})
	if ! ok {
		return nil, errors.New("no match type")
	}

	var newF Feature
	err = utils.Decode(source, &newF)
	if err != nil {
		return nil, errors.New("decode fail")
	}

	newF.Id = f.Id

	return &newF, nil
}
