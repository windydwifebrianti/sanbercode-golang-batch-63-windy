package main

import "fmt"

func main() {
	var dataFilm = []map[string]string{}
	var tambahDataFilm = func(title, jam, genre, tahun string) {
		var film = map[string]string{
			"genre": genre,
			"jam":   jam,
			"tahun": tahun,
			"title": title,
		}
		dataFilm = append(dataFilm, film)
	}
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
