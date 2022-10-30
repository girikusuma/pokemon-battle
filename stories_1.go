package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Pokemon []PokemonUrl `json:"results"`
}

type PokemonUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Stories1() (int, error) {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/?offset=0&limit=2000")
	if err != nil {
		return 0, err
	}
	
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var object Response
	var count int

	json.Unmarshal(body, &object)

	count = object.Count

	return count, nil
}