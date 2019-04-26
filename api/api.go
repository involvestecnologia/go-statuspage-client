package api

import "github.com/involvestecnologia/statuspage/models"

const (
	V1 Version = "v1"
)

type API interface {
	CreateClient(models.Client) (string, error)
	CreateComponent(models.Component) (string, error)
	FindComponent(componentName string) (models.Component, error)
	GetComponentsWithLabels(labels ...string) ([]models.Component, error)
}

type Version string
