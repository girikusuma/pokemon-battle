package main

import (
	"fmt"
	"strconv"
)

func Bonus() ([]PokemonInBattle, string) {
	var pokemons = make([]PokemonInBattle, 5)
	var anulirName string

	for i := 0; i < 5; i++ {
		pokemons[i] = GeneratePokemon()
	}

	fmt.Println("List Pokemon yang akan bertanding:")
	for index, pokemon := range pokemons {
		fmt.Println(strconv.Itoa(index + 1) + ". " + pokemon.Name)
	}
	fmt.Print("Masukan nama Pokemon yang akan di anulir: ")
	fmt.Scanln(&anulirName)

	resultPokemon := DisqualifiedPokemon(pokemons, anulirName)

	BonusBattle(0, resultPokemon)

	return resultPokemon, anulirName
}

func DisqualifiedPokemon(pokemons []PokemonInBattle, anulirName string) []PokemonInBattle {
	for i := 0; i < len(pokemons); i++ {
		if pokemons[i].Name == anulirName {
			pokemons[i].Point = 0
		} else {
			pokemons[i].Point += 1
		}
	}

	getIndex := 0
	for index, pokemon := range pokemons {
		if pokemon.Name == anulirName {
			getIndex = index
		}
	}

	copy(pokemons[getIndex:], pokemons[getIndex + 1:])
	pokemons[len(pokemons) - 1] = PokemonInBattle{
		Name: "",
		Attack: 0,
		Defense: 0,
		Hp: 0,
		Point: 0,
	}
	pokemons = pokemons[:len(pokemons) - 1]

	fmt.Println("List Pokemon yang akan bertanding setelah anulir:")
	for i := 0; i < len(pokemons); i++ {
		fmt.Println(strconv.Itoa(i + 1) + ". " + pokemons[i].Name)
	}

	return pokemons
}

func BonusBattle(index int, pokemons []PokemonInBattle) {
	indexDef := index + 1
	if indexDef == len(pokemons) {
		indexDef = 0
	}

	for pokemons[index].Hp <= 0{
		index += 1
		if index >= len(pokemons) {
			index = 0
		}
	}

	for index == indexDef || pokemons[indexDef].Hp <= 0{
		indexDef += 1
		if indexDef >= len(pokemons) {
			indexDef = 0
		}
	}
	BonusAttack(&pokemons[index], &pokemons[indexDef], pokemons)

	for _, pokemon := range pokemons {
		if pokemon.Point == 5 {
			return
		}
	}
	index += 1
	if index == len(pokemons) {
		index = 0
	}

	BonusBattle(index, pokemons)
}

func BonusAttack(attacker *PokemonInBattle, defender *PokemonInBattle, pokemons []PokemonInBattle) {
	if attacker.Hp > 0 && defender.Hp > 0 {
		fmt.Println("Pokemon " + attacker.Name + " menyerang pokemon " + defender.Name + " dengan serangan " + strconv.Itoa(attacker.Attack))
	}
	BonusDamageTaken(attacker.Attack, defender, attacker, pokemons)
	if attacker.Hp > 0 && defender.Hp > 0 {
		fmt.Println("Pokemon " + defender.Name + " menyerang pokemon " + attacker.Name + " dengan serangan " + strconv.Itoa(defender.Attack))
	}
	BonusDamageTaken(defender.Attack, attacker, defender, pokemons)
}

func BonusDamageTaken(attack int, defender *PokemonInBattle, attacker *PokemonInBattle, pokemons []PokemonInBattle) {
	if attacker.Hp > 0 && defender.Hp > 0 {
		damage := attack / (defender.Defense / 10)
		if damage > defender.Hp {
			defender.Hp = 0
		} else {
			defender.Hp -= damage
		}
		fmt.Println("Pokemon " + defender.Name + " menerima serangan sebesar " + strconv.Itoa(damage) + " dari serangan " + attacker.Name)
		fmt.Println("Sisa HP Pokemon " + defender.Name + " adalah " + strconv.Itoa(defender.Hp))
		if defender.Hp <= 0 {
			for i := 0; i < len(pokemons); i++ {
				if pokemons[i].Name != defender.Name && pokemons[i].Hp > 0 {
					pokemons[i].Point += 1
				}
			}
		}
	}
}