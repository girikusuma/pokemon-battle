package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Sroties")

	fmt.Println("Sebagai seorang operasional Tim Roket, saya perlu bisa mencari Pokemon dari API")
	fmt.Println("Pokedex: http://pokeapi.co/ yang saya ingin bisa melihat jumlah stok semua")
	fmt.Println("Pokemon yang dimiliki.")
	fmt.Println("Jawab:")

	stockPokemon := Stories1()
	fmt.Println("Jumlah stok Pokemon yang dimiliki: " + strconv.Itoa(stockPokemon))
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("Sebagai pengedar gelap Tim Roket, saya perlu mencatat setiap pertandingan, berupa")
	fmt.Println("Pokemon mana saja yang bertanding (5 Pokemon bertanding sekaligus), dan")
	fmt.Println("Pokemon mana yang juara diurutkan berdasarkan Pokemon yang paling lama")
	fmt.Println("bertahan di ring. Pokemon yang paling lama bertahan mendapatkan skor 5 dan paling")
	fmt.Println("cepat kalah mendapat skor 1.")
	fmt.Println("Jawab:")

	result := Stories2()
	list_pokemon := make([]string, len(result))

	fmt.Println("Nama Pokemon yang bertanding:")
	for index, pokemon := range result {
		fmt.Println(strconv.Itoa(index + 1) + ". " + pokemon.Name)
		list_pokemon[index] = pokemon.Name
	}
	for i := 0; i < len(result) - 1; i++ {
		for j := 0; j < len(result) - i - 1; j++ {
			if result[j].Point < result[j + 1].Point {
				result[j], result[j + 1] = result[j + 1], result[j]
			}
		}
	}
	fmt.Println("Hasil pertandingan berdasarkan Pokemon yang memiliki skor tertinggi:")
	for _, pokemon := range result {
		fmt.Println("Pokemon " + pokemon.Name + " mendapatkan skor: " + strconv.Itoa(pokemon.Point))
	}
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("Sebagai Bos, saya ingin bisa melihat semua data pertandingan yang pernah terjadi")
	fmt.Println("selama periode tertentu dan urutan peringkat Pokemonnya.")
	fmt.Println("Jawab:")
	fmt.Println("List Pokemon yang bertanding:")
	for i, name := range list_pokemon {
		fmt.Println(strconv.Itoa(i + 1) + ". " + name)
	}
	fmt.Print("Masukan nama Pokemon yang ingin dilihat data pertandingannya: ")
	var selected string
	fmt.Scanln(&selected)
	fmt.Println("Hasil Pertandingan:")
	ranking := Stories3(list_pokemon, selected, result)
	fmt.Println("Urutan peringkat Pokemonnya: " + strconv.Itoa(ranking))
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("Sebagai Bos, saya ingin bisa melihat secara keseluruhan urutan Pokemon mana yang")
	fmt.Println("paling tinggi skornya untuk dipromosikan.")
	fmt.Println("Jawab:")

	fmt.Println("Pokemon yang dipromosikan adalah: " + result[0].Name)
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("BONUS: Sebagai Bos, karena beberapa anggota Tim Roket terkadang curang, saya")
	fmt.Println("ingin bisa menganulir salah satu Pokemon di pertandingan tersebut. Pokemon")
	fmt.Println("tersebut akan dianulir dan urutan pemenang di pertandingan tersebut naik satu")
	fmt.Println("peringkat menggantikan Pokemon yang dianulir.")
	result_bonus := Bonus()
	for _, pokemon := range result_bonus {
		fmt.Println("Pokemon " + pokemon.Name + " mendapatkan skor: " + strconv.Itoa(pokemon.Point))
	}
}