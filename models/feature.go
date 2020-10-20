package models

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type Feature struct {
	Id          string
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`

	ViewCount int `json:"view_count"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (f *Feature) IndexName() string {
	return "features"
}

func (f *Feature) Mapping() string {

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

func (f *Feature) Search(offset, limit int) ([]interface{}, error) {

	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{

			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
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
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			// Print the response status and error information.
			return nil, errors.New(fmt.Sprintf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			))
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	// Print the response status, number of results, and request duration.
	//hits := int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	//var models []*Feature
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//
	//	source := hit.(map[string]interface{})["source"].(map[string]interface{})
	//
	//	f := &Feature{
	//		Id:    hit.(map[string]interface{})["_id"].(string),
	//		Title: source["title"].(string),
	//		Description: source["description"].(string),
	//	}
	//
	//	// hit.(map[string]interface{})["_id"]
	//	log.Printf(" * ID=%s\n", hit.(map[string]interface{}))
	//}

	return nil, nil
}
