package client

import (
	"log"

	"github.com/involvestecnologia/go-statuspage-client/api"
	v1 "github.com/involvestecnologia/go-statuspage-client/api/v1"
	"github.com/involvestecnologia/statuspage/models"
)

var (
	defaultAPI = v1.NewAPIV1
)

type client struct {
	api api.API
}

func DefaultClient(addr string) *client {
	return &client{
		api: defaultAPI(addr),
	}
}

func NewClient(v api.Version, addr string) *client {
	var a api.API
	switch v {
	case api.V1:
		a = v1.NewAPIV1(addr)
	default:
		log.Panicf("Version %s not available or implemented yet", v)
	}

	return &client{
		api: a,
	}
}

func (c *client) CreateClient(client models.Client) (string, error) {
	return c.api.CreateClient(client)
}

func (c *client) CreateComponent(component models.Component) (string, error) {
	return c.api.CreateComponent(component)
}
