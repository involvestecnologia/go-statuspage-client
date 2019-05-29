package v1_test

import (
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"
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
		BodyString("testclient-ref")

	client.Name = "TestClient"

	ref, err = metroV1.CreateClient(client)
	assert.NoError(t, err)
	assert.NotEmpty(t, ref)

}

func TestFindComponent(t *testing.T) {
	defer gock.OffAll()

	metroV1 := v1.NewAPIV1(uri)
	componentName := "component-test"

	gock.New(uri).
		Get(v1.FindComponentEndpoint+componentName+"fail").
		MatchParam("search", "name").
		Reply(http.StatusBadRequest)

	ref, err := metroV1.FindComponent(componentName + "fail")
	assert.Error(t, err)
	assert.Empty(t, ref)

	gock.Off()
	gock.New(uri).
		Get(v1.FindComponentEndpoint+componentName).
		MatchParam("search", "name").
		Reply(http.StatusOK).
		JSON(models.Component{Ref: "test-ref", Name: componentName})

	comp, err := metroV1.FindComponent("component-test")
	assert.NoError(t, err)
	assert.Equal(t, componentName, comp.Name)

}

func TestGetComponentsWithLabels(t *testing.T) {
	defer gock.OffAll()

	metroV1 := v1.NewAPIV1(uri)

	gock.New(uri).
		Post(v1.SearchComponentByLabelEndpoint).
		JSON(map[string]interface{}{"labels": []string{"invalid-ref"}}).
		Reply(http.StatusBadRequest)

	ref, err := metroV1.GetComponentsWithLabels("invalid-ref")
	assert.Error(t, err)
	assert.Empty(t, ref)

	gock.Off()
	gock.New(uri).
		Post(v1.SearchComponentByLabelEndpoint).
		JSON(map[string]interface{}{"labels": []string{"valid-ref"}}).
		Reply(http.StatusOK).
		JSON([]models.Component{models.Component{Ref: "valid-ref"}})

	comps, err := metroV1.GetComponentsWithLabels("valid-ref")
	assert.NoError(t, err)
	assert.Equal(t, "valid-ref", comps[0].Ref)

}
