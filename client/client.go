package client

import (
	"encoding/json"
	"net/http"
)

const BaseURL string = "https://pokeapi.co/api/v2/"

type PokeClient struct {
	client  *http.Client
	baseURL string
}

// Resource is a resource list for an endpoint.
type Resource struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

// Result is a resource list result.
type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func New() *PokeClient {
	return &PokeClient{
		client:  &http.Client{},
		baseURL: BaseURL,
	}
}

func (pc *PokeClient) Fetch(resource string) (Resource, error) {
	result := Resource{}
	resp, err := pc.client.Get(pc.BuildURL(resource))

	if err != nil {
		return Resource{}, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Resource{}, err
	}

	return result, nil
}

func (pc *PokeClient) BuildURL(resource string) string {
	return BaseURL + resource
}
