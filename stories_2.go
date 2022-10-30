package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Pokemon struct {
	Abilities []Abilities `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms []Forms `json:"forms"`
	GameIndices []GameIndices `json:"game_indices"`
	Height int `json:"height"`
	HeldItems []string `json:"held_items"`
	Id int `json:"id"`
	IsDefault bool `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves []Moves `json:"moves"`
	Name string `json:"name"`
	Order int `json:"order"`
	PastTypes []string `json:"past_types"`
	Species Species `json:"species"`
	Sprites Sprites `json:"sprites"`
	Stats []Stats `json:"stats"`
	Weight int `json:"weight"`
}

type PokemonInBattle struct {
	Name string
	Hp int
	Attack int
	Defense int
	Point int
}

var all_pokemon []PokemonUrl = GetAllPokemon()

func GetAllPokemon() []PokemonUrl {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/?offset=0&limit=2000")
	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var object Response
	json.Unmarshal(body, &object)
	
	return object.Pokemon
}

func GeneratePokemon() PokemonInBattle {
	rand.Seed(time.Now().UTC().UnixNano())
	countPokemon, _ := Stories1()
	randNumber := rand.Intn(countPokemon)

	response, err := http.Get(all_pokemon[randNumber].Url)
	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var object Pokemon
	json.Unmarshal(body, &object)

	var hp int
	var attack int
	var defense int

	for _, stat := range object.Stats {
		if stat.Stat.Name == "hp" {
			hp = stat.BaseStat
		}
		if stat.Stat.Name == "attack" {
			attack = stat.BaseStat
		}
		if stat.Stat.Name == "defense" {
			defense = stat.BaseStat
		}
	}

	pokemon := PokemonInBattle{
		Name: object.Name,
		Hp: hp,
		Attack: attack,
		Defense: defense,
		Point: 1,
	}

	return pokemon
}

func Stories2() []PokemonInBattle {
	var pokemons = make([]PokemonInBattle, 5)
	
	for i := 0; i < 5; i++{
		pokemons[i] = GeneratePokemon()
	}

	Battle(0, pokemons)

	return pokemons
}

func Battle(index int, pokemons []PokemonInBattle) {
	indexDef := index + 1
	if indexDef == len(pokemons) {
		indexDef = 0
	}

	for pokemons[index].Hp <= 0 {
		index += 1
		if index >= len(pokemons) {
			index = 0
		}
	}

	for index == indexDef || pokemons[indexDef].Hp <= 0 {
		indexDef += 1
		if indexDef >= len(pokemons) {
			indexDef = 0
		}
	}
	Attack(&pokemons[index], &pokemons[indexDef], pokemons)	

	for _, pokemon := range pokemons {
		if pokemon.Point == 5 {
			return
		}
	}
	index += 1
	if index == len(pokemons) {
		index = 0
	}

	Battle(index, pokemons)

}

// func Battle(attacker *PokemonInBattle, defender *PokemonInBattle) PokemonInBattle {
// 	Attack(attacker, defender)
// 	if attacker.Hp > 0 && defender.Hp > 0 {
// 		Battle(defender, attacker)
// 	}
// 	if attacker.Hp <= 0 {
// 		return *defender
// 	}

// 	return *attacker
// }

func Attack(attacker *PokemonInBattle, defender *PokemonInBattle, pokemons []PokemonInBattle) {
	if attacker.Hp > 0 && defender.Hp > 0 {
		History = append(History, "Pokemon " + attacker.Name + " menyerang pokemon " + defender.Name + " dengan serangan " + strconv.Itoa(attacker.Attack))
	}
	DamageTaken(attacker.Attack, defender, attacker, pokemons)
	if attacker.Hp > 0 && defender.Hp > 0 {
		History = append(History, "Pokemon " + defender.Name + " menyerang pokemon " + attacker.Name + " dengan serangan " + strconv.Itoa(defender.Attack))
	}
	DamageTaken(defender.Attack, attacker, defender, pokemons)
}

func DamageTaken(attack int, defender *PokemonInBattle, attacker *PokemonInBattle, pokemons []PokemonInBattle) {
	if attacker.Hp > 0 && defender.Hp > 0 {
		damage := attack / (defender.Defense / 10)
		if damage > defender.Hp {
			defender.Hp = 0
		} else {
			defender.Hp -= damage
		}
		History = append(History, "Pokemon " + defender.Name + " menerima serangan sebesar " + strconv.Itoa(damage) + " dari serangan " + attacker.Name)
		History = append(History, "Sisa HP Pokemon " + defender.Name + " adalah " + strconv.Itoa(defender.Hp))
		if defender.Hp <= 0 {
			for i := 0; i < len(pokemons); i++ {
				if pokemons[i].Name != defender.Name && pokemons[i].Hp > 0 {
					pokemons[i].Point += 1
				}
			}
		}
	}
}