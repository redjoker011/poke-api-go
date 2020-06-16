package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/redjoker011/poke-api-go/client"
	"log"
	"net/http"
)

func pokemonsHandler(w http.ResponseWriter, _ *http.Request) {
	client := client.New()
	resp, err := client.Fetch("pokemon")
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(resp)
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/pokemons", pokemonsHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
