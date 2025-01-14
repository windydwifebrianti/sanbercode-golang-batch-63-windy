//soal4
package main

import "fmt"

func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	var film = map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilm = append(*dataFilm, film)
}

func main() {

	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d. ", i+1)
		for key, value := range item {
			fmt.Printf("%s : %s\n", key, value)
		}
		fmt.Println()
	}
}
