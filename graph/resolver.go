package graph

import "github.com/redjoker011/poke-api-go/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	pokemons *model.Resource
}
