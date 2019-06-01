package api

import (
	"time"

	"github.com/involvestecnologia/statuspage/models"
)

const (
	V1 Version = "v1"
)

type API interface {
	CreateClient(models.Client) (string, error)
	CreateComponent(models.Component) (string, error)
	FindComponent(componentName string) (models.Component, error)
	GetComponentsWithLabels(labels ...string) ([]models.Component, error)
	GetIncidentsFromPeriod(startDt, endDt time.Time, resolved bool) ([]models.Incident, error)
}

type Version string
