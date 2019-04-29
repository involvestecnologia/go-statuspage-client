package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/involvestecnologia/go-statuspage-client/api"
	"github.com/involvestecnologia/statuspage/models"
)

const (
	defaultTimeout                 time.Duration = 15 * time.Second
	defaultContentType             string        = "application/json"
	CreateClientEndpoint           string        = "/v1/client"
	CreateComponentEndpoint        string        = "/v1/component"
	FindComponentEndpoint          string        = "/v1/component/"
	SearchComponentByLabelEndpoint string        = "/v1/component/label"
)

type v1 struct {
	URL        string
	httpClient http.Client
}

//NewAPIV1 return the V1 client implementation of the metronome API
func NewAPIV1(url string) api.API {
	return &v1{
		URL: url,
		httpClient: http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (api *v1) CreateClient(client models.Client) (string, error) {
	body, err := json.Marshal(client)
	if err != nil {
		return "", err
	}

	resp, err := api.httpClient.Post(api.URL+CreateClientEndpoint, defaultContentType, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("Unexpected response from statuspage: %d", resp.StatusCode)
	}

	ref, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ref), nil
}

func (api *v1) CreateComponent(component models.Component) (string, error) {
	body, err := json.Marshal(component)
	if err != nil {
		return "", err
	}

	resp, err := api.httpClient.Post(api.URL+CreateComponentEndpoint, defaultContentType, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("Unexpected response from statuspage: %d", resp.StatusCode)
	}

	ref, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ref), nil
}

func (api *v1) FindComponent(componentName string) (models.Component, error) {
	var comp models.Component

	nameSearchParam := "?search=name"
	resp, err := api.httpClient.Get(api.URL + FindComponentEndpoint + componentName + nameSearchParam)
	if err != nil {
		return comp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return comp, fmt.Errorf("Unexpected response from statuspage: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&comp)
	return comp, err
}

func (api *v1) GetComponentsWithLabels(labels ...string) ([]models.Component, error) {
	var components []models.Component
	l := models.ComponentLabels{
		Labels: labels,
	}

	body, err := json.Marshal(l)
	if err != nil {
		return components, err
	}

	resp, err := api.httpClient.Post(api.URL+SearchComponentByLabelEndpoint, defaultContentType, bytes.NewReader(body))
	if err != nil {
		return components, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return components, fmt.Errorf("Unexpected response from statuspage: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&components)
	return components, err
}
