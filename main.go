package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	pb "proto"

	"google.golang.org/grpc"
)

// Declare Go map/dictionary
var APIMAP map[string]string

// Instantiante the global map and add values to it
func initMap() {
	// Initialize map otherwise APIMAP is nil
	APIMAP = make(map[string]string)
	setValues()
}

// Add mappings to new API URLs
func setValues() {
	APIMAP["pokeapi"] = "https://pokeapi.co/api/v2/pokemon/"
}

const (
	port = ":6809"
)

type pokemonServer struct {
	pb.UnimplementedPokedexServer
}

func newServer() *pokemonServer {
	s := &pokemonServer{}
	return s
}

const URL = "https://pokeapi.co/api/v2/pokemon/"

func (s *pokemonServer) GetPokemon(ctx context.Context, in *pb.SearchRequest) (*pb.Pokemon, error) {
	log.Printf("Getting info for pokemon: " + in.Name)

	return extractInfo(fetch(URL + in.Name)), nil
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

func extractInfo(body string) *pb.Pokemon {
	pkmn := pb.Pokemon{}
	json.Unmarshal([]byte(body), &pkmn)
	return &pkmn
}

func main() {
	initMap()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterPokedexServer(grpcServer, newServer())
	fmt.Println("Listening on ", port)
	grpcServer.Serve(listener)

}

/*
Lessons along the way:

fmt is not thread safe, log IS
Know this because when printing res with fmt in main function
VSCode + terminal would freeze whereas using log did not

When dealing with nested JSON objects, in proto, nested properties can have the same name

Generation of code updated compiler: https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable

*/
