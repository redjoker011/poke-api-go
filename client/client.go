package client

import (
	"encoding/json"
	"fmt"
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

func (pc *PokeClient) Fetch(resource string, params ...int) (Resource, error) {
	offset, limit := parseParams(params)
	resourceUrl := fmt.Sprintf("%s?offset=%d&limit=%d", pc.BuildURL(resource), offset, limit)
	result := Resource{}
	resp, err := pc.client.Get(resourceUrl)

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

func parseParams(params []int) (offset, limit int) {
	switch l := len(params); {
	case l == 2:
		limit = params[1]
		fallthrough
	case l >= 1:
		offset = params[0]
	}
	return
}
