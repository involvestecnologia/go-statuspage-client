package v1_test

import (
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/involvestecnologia/go-statuspage-client/api"
	v1 "github.com/involvestecnologia/go-statuspage-client/api/v1"
	"github.com/involvestecnologia/statuspage/models"
	"github.com/stretchr/testify/assert"
)

const uri = "http://localhost:8080"

func TestNewAPIV1(t *testing.T) {
	metroV1 := v1.NewAPIV1(uri)

	assert.Implements(t, (*api.API)(nil), metroV1)
}

func TestCreateComponent(t *testing.T) {
	defer gock.OffAll()

	gock.New(uri).
		Post(v1.CreateComponentEndpoint).
		Reply(http.StatusBadRequest)

	metroV1 := v1.NewAPIV1(uri)

	component := models.Component{
		Name:    "",
		Address: "test",
		Labels:  []string{"Applications"},
	}

	ref, err := metroV1.CreateComponent(component)
	assert.Error(t, err)
	assert.Empty(t, ref)

	gock.Off()
	gock.New(uri).
		Post(v1.CreateComponentEndpoint).
		Reply(http.StatusCreated).
		AddHeader("Content-Type", "application/json").
		BodyString("testcomponent-ref")

	component.Name = "Test"
	ref, err = metroV1.CreateComponent(component)
	assert.NoError(t, err)
	assert.NotEmpty(t, ref)
}

func TestCreateClient(t *testing.T) {
	defer gock.OffAll()

	metroV1 := v1.NewAPIV1(uri)

	gock.New(uri).
		Post(v1.CreateClientEndpoint).
		Reply(http.StatusBadRequest)

	client := models.Client{
		Name:      "",
		Resources: make([]string, 0),
	}

	ref, err := metroV1.CreateClient(client)
	assert.Error(t, err)
	assert.Empty(t, ref)

	gock.Off()
	gock.New(uri).
		Post(v1.CreateClientEndpoint).
		Reply(http.StatusCreated).
		AddHeader("Content-Type", "application/json").
		BodyString("testclient-ref")

	client.Name = "TestClient"

	ref, err = metroV1.CreateClient(client)
	assert.NoError(t, err)
	assert.NotEmpty(t, ref)

}
