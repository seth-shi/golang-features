package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang-functions/utils"
	"time"
)

type Model interface {
	Mapping() string
	IndexName() string
	GetId() string
	GetCreatedAt() time.Time
}

func Create(model Model) error {

	var m map[string]interface{}

	var buf bytes.Buffer

	err := utils.Decode(model, &m)
	if err != nil {
		return err
	}

	delete(m, "Id")
	m["created_at"] = time.Now().Format("2006-01-02 15:04:05")
	m["updated_at"] = time.Now().Format("2006-01-02 15:04:05")

	if err = json.NewEncoder(&buf).Encode(m); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	_, err = es.Create(
		model.IndexName(),
		fmt.Sprintf("%s", id),
		&buf,
	)

	return err
}

func Update(model Model) error {

	if model.GetId() == "" {
		return errors.New("无效的模型")
	}

	var m map[string]interface{}

	var buf bytes.Buffer

	err := utils.Decode(model, &m)
	if err != nil {
		return err
	}

	delete(m, "Id")
	m["created_at"] = model.GetCreatedAt().Format("2006-01-02 15:04:05")
	m["updated_at"] = time.Now().Format("2006-01-02 15:04:05")

	if err = json.NewEncoder(&buf).Encode(m); err != nil {
		return err
	}

	res, err := es.Update(
		model.IndexName(),
		model.GetId(),
		&buf,
	)


	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = parseEsResponse(res)

	return err
}

func Find(model Model) (map[string]interface{}, error) {

	res, err := es.Get(model.IndexName(), model.GetId())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return parseEsResponse(res)
}
