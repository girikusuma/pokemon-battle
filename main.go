package main

import (
	"fmt"
	"strconv"
)

func main() {
	var periodSelect int
	var list_pokemon []string
	var selected string

	fmt.Println("Sroties")

	fmt.Println("Sebagai seorang operasional Tim Roket, saya perlu bisa mencari Pokemon dari API")
	fmt.Println("Pokedex: http://pokeapi.co/ yang saya ingin bisa melihat jumlah stok semua")
	fmt.Println("Pokemon yang dimiliki.")
	fmt.Println("Jawab:")

	stockPokemon, _ := Stories1()
	fmt.Println("Jumlah stok Pokemon yang dimiliki: " + strconv.Itoa(stockPokemon))
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("Sebagai pengedar gelap Tim Roket, saya perlu mencatat setiap pertandingan, berupa")
	fmt.Println("Pokemon mana saja yang bertanding (5 Pokemon bertanding sekaligus), dan")
	fmt.Println("Pokemon mana yang juara diurutkan berdasarkan Pokemon yang paling lama")
	fmt.Println("bertahan di ring. Pokemon yang paling lama bertahan mendapatkan skor 5 dan paling")
	fmt.Println("cepat kalah mendapat skor 1.")
	fmt.Println("Jawab:")

	Stories2()
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("Sebagai Bos, saya ingin bisa melihat semua data pertandingan yang pernah terjadi")
	fmt.Println("selama periode tertentu dan urutan peringkat Pokemonnya.")
	fmt.Println("Jawab:")
	fmt.Print("Masukan periode pertandingan yang ingin dilihat: ")
	fmt.Scanln(&periodSelect)

	for _, pokemon := range Periods[periodSelect - 1].Result {
		list_pokemon = append(list_pokemon, pokemon.Name)
	}
	fmt.Println("List Pokemon yang bertanding:")
	for i, name := range list_pokemon {
		fmt.Println(strconv.Itoa(i + 1) + ". " + name)
	}
	fmt.Print("Masukan nama Pokemon yang ingin dilihat data pertandingannya: ")
	fmt.Scanln(&selected)
	fmt.Println("Hasil Pertandingan:")
	ranking := Stories3(list_pokemon, selected, Periods[periodSelect - 1].Result, periodSelect - 1)
	fmt.Println("Urutan peringkat Pokemonnya: " + strconv.Itoa(ranking))
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("Sebagai Bos, saya ingin bisa melihat secara keseluruhan urutan Pokemon mana yang")
	fmt.Println("paling tinggi skornya untuk dipromosikan.")
	fmt.Println("Jawab:")
	pokemonPromote := Stories4()
	fmt.Println("Pokemon yang dipromosikan adalah: " + pokemonPromote)
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("BONUS: Sebagai Bos, karena beberapa anggota Tim Roket terkadang curang, saya")
	fmt.Println("ingin bisa menganulir salah satu Pokemon di pertandingan tersebut. Pokemon")
	fmt.Println("tersebut akan dianulir dan urutan pemenang di pertandingan tersebut naik satu")
	fmt.Println("peringkat menggantikan Pokemon yang dianulir.")
	fmt.Println("Jawab:")
	result_bonus, anulir := Bonus()
	for _, pokemon := range result_bonus {
		fmt.Println("Pokemon " + pokemon.Name + " mendapatkan skor: " + strconv.Itoa(pokemon.Point))
	}
	fmt.Println("Pokemon yang di anulir: " + anulir)
}