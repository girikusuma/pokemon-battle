package main

import (
	"encoding/json"
	"farmacare/helper"
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

func Stories1() int {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/?offset=0&limit=2000")
	helper.ErrorPrint(err)

	body, err := ioutil.ReadAll(response.Body)
	helper.ErrorPrint(err)

	var object Response
	var count int

	json.Unmarshal(body, &object)

	count = object.Count

	return count
}