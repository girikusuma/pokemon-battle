package main

import (
	"reflect"
	"testing"
)

func TestBonusBattle(t *testing.T) {
	type args struct {
		index    int
		pokemons []PokemonInBattle
	}
	tests := []struct {
		name string
		args args
		result []PokemonInBattle
	}{
		{
			name: "normal",
			args: args{
				index: 0,
				pokemons: []PokemonInBattle{
					{
						Name: "pikachu",
						Hp: 60,
						Attack: 76,
						Defense: 80,
						Point: 2,
					},
					{
						Name: "pikachu2",
						Hp: 2,
						Attack: 1,
						Defense: 10,
						Point: 2,
					},
					{
						Name: "pikachu3",
						Hp: 3,
						Attack: 1,
						Defense: 10,
						Point: 2,
					},
					{
						Name: "pikachu4",
						Hp: 4,
						Attack: 1,
						Defense: 10,
						Point: 2,
					},
				},
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
			BonusBattle(tt.args.index, tt.args.pokemons)
			if ok := reflect.DeepEqual(tt.result, tt.args.pokemons); !ok {
				t.Errorf("not equal\ngot: %+v\nwant: %+v", tt.args.pokemons, tt.result)
			}
		})
	}
}
