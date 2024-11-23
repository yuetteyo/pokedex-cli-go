package main

import (
	"fmt"
	"log"

	"github.com/yuetteyo/pokedex-cli-go/internal/pokeapi"
)

func callbackMap() error {

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}