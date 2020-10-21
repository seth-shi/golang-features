package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang-functions/utils"
	"time"
)

type Model interface {
	Mapping() string
	IndexName() string
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