package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redjoker011/poke-api-go/client"
	"github.com/redjoker011/poke-api-go/graph/generated"
	"github.com/redjoker011/poke-api-go/graph/model"
)

func (r *queryResolver) Pokemons(ctx context.Context, limit *int, offset *int) (*model.Resource, error) {
	client := client.New()
	resp, err := client.Fetch("pokemon", *offset, *limit)

	if err != nil {
		log.Fatal(err)
	}

	resource := model.Resource{}

	res, err := json.Marshal(resp)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(res, &resource)

	return &resource, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
