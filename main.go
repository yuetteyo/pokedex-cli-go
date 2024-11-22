package main

import (
	"fmt"
	"log"

	"github.com/yuetteyo/pokedex-cli-go/internal/pokeapi"
)

func main() {
	// startRepl()

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

}
