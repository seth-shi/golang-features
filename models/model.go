package models

type Model interface {
	Mapping() string
	IndexName() string
}