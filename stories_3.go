package main

import (
	"fmt"
	"strings"
)

var History []string

func Stories3(list_pokemon []string, name string, pokemons []PokemonInBattle) int {
	var check bool = false
	var ranking int
	for _, pokemon := range list_pokemon {
		if pokemon == name {
			check = true
		}
	}

	if check {
		for _, history := range History{
			if strings.Contains(history, name) && !strings.Contains(history, "HP") {
				fmt.Println(history)
			}
		}
		for _, pokemon := range pokemons {
			if pokemon.Name == name {
				ranking = len(pokemons) - pokemon.Point + 1
			}
		}
	} else {
		fmt.Println("Nama Pokemon tidak ada")
		ranking = 0
	}

	return ranking
}