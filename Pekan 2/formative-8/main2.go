package main

import "fmt"

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type smartphone interface {
	dataSmartphone() string
}

func (p phone) dataSmartphone() string {
	return fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.name, p.brand, p.year, p.colors)
}

func main() {
	myPhone := phone {
		name : "Samsung Galaxy Note 20",
		brand : "Samsung Galaxy Note 20",
		year : 2020,
		colors : []string{"Mystic Bronze", "Mystic White", "Mystic Black"},

	}

var hape smartphone = myPhone
fmt.Println(hape.dataSmartphone())
}