package main

import (
	"fmt"
	searchPokemon "pykadex-backend/apis/search"
	uploadPokemon "pykadex-backend/apis/upload"
	"strconv"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/pokemon-upload", uploadPokemon.PokemonUploadHandler).Methods(http.MethodPost)
	r.HandleFunc("/pokemon-search/{id}", searchPokemon.PokemonSearchHandler).Methods(http.MethodGet)

	port := 2000
	portStr := strconv.Itoa(port)

	fmt.Printf("server is running on 127.0.0.1:%v", port)
	log.Fatal(http.ListenAndServe(":"+portStr, r))
}
