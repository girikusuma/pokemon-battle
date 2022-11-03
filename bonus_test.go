package main

import (
	"reflect"
	"testing"
)

func TestBonusBattle(t *testing.T) {
	type args struct {
		pokemons []PokemonInBattle
		anulirName	string
	}
	tests := []struct {
		name string
		args args
		result []PokemonInBattle
	}{
		{
			name: "normal",
			args: args{
				pokemons: []PokemonInBattle{
					{
						Name: "pikachu",
						Hp: 60,
						Attack: 76,
						Defense: 80,
						Point: 1,
					},
					{
						Name: "pikachu2",
						Hp: 2,
						Attack: 1,
						Defense: 10,
						Point: 1,
					},
					{
						Name: "pikachu3",
						Hp: 3,
						Attack: 1,
						Defense: 10,
						Point: 1,
					},
					{
						Name: "pikachu4",
						Hp: 4,
						Attack: 1,
						Defense: 10,
						Point: 1,
					},
					{
						Name: "pikachu5",
						Hp: 5,
						Attack: 1,
						Defense: 10,
						Point: 1,
					},
				},
				anulirName: "pikachu5",
			},
			result: []PokemonInBattle{
				{
					Name: "pikachu",
					Hp: 60,
					Attack: 76,
					Defense: 80,
					Point: 5,
				},
				{
					Name: "pikachu2",
					Hp: 0,
					Attack: 1,
					Defense: 10,
					Point: 2,
				},
				{
					Name: "pikachu3",
					Hp: 0,
					Attack: 1,
					Defense: 10,
					Point: 4,
				},
				{
					Name: "pikachu4",
					Hp: 0,
					Attack: 1,
					Defense: 10,
					Point: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pokemons := DisqualifiedPokemon(tt.args.pokemons, tt.args.anulirName)
			BonusBattle(0, pokemons)
			if ok := reflect.DeepEqual(tt.result, pokemons); !ok {
				t.Errorf("not equal\ngot: %+v\nwant: %+v", pokemons, tt.result)
			}
		})
	}
}
