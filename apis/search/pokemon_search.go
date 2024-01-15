package general

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	pokemonTypes "pykadex-backend/types"

	"github.com/gorilla/mux"
)

func PokemonSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	paramsInt, paramsIntErr := strconv.Atoi(params["id"])
	if paramsIntErr != nil {
		fmt.Printf("Id %v conversion failed", paramsIntErr)
	}

	// load json
	jsonFile, err := os.Open("data/pokemon.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var pokemonData []pokemonTypes.Pokemon
	json.Unmarshal(byteValue, &pokemonData)

	// search for pokemon
	var targetPokemon *pokemonTypes.Pokemon
	for _, pokemon := range pokemonData {
		if pokemon.Id == paramsInt {
			targetPokemon = &pokemon
		}
	}

	// handle found / not found
	if targetPokemon != nil {
		json.NewEncoder(w).Encode(targetPokemon)
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		resp := make(map[string]string)
		resp["message"] = "Pokemon Not Found"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}
