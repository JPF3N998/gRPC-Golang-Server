package pokemon

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct{}

const URL = "https://pokeapi.co/api/v2/pokemon/"

func (s *Server) GetPokemon(ctx context.Context, in *SearchRequest) (*SearchResponse, error) {
	log.Printf("Getting info for pokemon: " + in.Name)

	return &SearchResponse{Pokemon: extractInfo(fetch(URL + in.Name))}, nil
}

// Make get requests ref: https://blog.logrocket.com/making-http-requests-in-go/
func fetch(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	bodyAsString := string(body)

	return bodyAsString
}

func extractInfo(body string) *Pokemon {
	pkmn := Pokemon{}
	json.Unmarshal([]byte(body), &pkmn)
	return &pkmn
}
