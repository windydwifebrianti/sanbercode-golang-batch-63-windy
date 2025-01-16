package main

import "fmt"

type movie struct {
	title, genre   string
	duration, year int
}

var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, films *[]movie) {
	newMovie := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}

	*films = append(*films, newMovie)
}

func main() {
	tambahDataFilm("LOTR", 2, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 2, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 2, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 2, "horror", 2004, &dataFilm)

	fmt.Println("Data Film:")
	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s\n duration: %d jam\n genre: %s\n year: %d\n\n", i+1, film.title, film.duration, film.genre, film.year)
	}
}
