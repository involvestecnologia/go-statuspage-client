package api

import "github.com/involvestecnologia/statuspage/models"

const (
	V1 Version = "v1"
)

type API interface {
	CreateClient(models.Client) (string, error)
	CreateComponent(models.Component) (string, error)
}

type Version string
