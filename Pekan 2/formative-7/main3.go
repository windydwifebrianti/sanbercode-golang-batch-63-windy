package main

import "fmt"

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

func main() {

	myPhone := phone{
		name:   "Galaxy S21",
		brand:  "Samsung",
		year:   2021,
		colors: []string{"Black", "White"},
	}

	fmt.Println("Data Phone Awal:")
	fmt.Printf("Name: %s\n", myPhone.name)
	fmt.Printf("Brand: %s\n", myPhone.brand)
	fmt.Printf("Year: %d\n", myPhone.year)
	fmt.Printf("Colors: %v\n", myPhone.colors)

	myPhone.addColor("Gold")
	myPhone.addColor("Silver")

	fmt.Println("\nData Phone Setelah Penambahan Warna:")
	fmt.Printf("Name: %s\n", myPhone.name)
	fmt.Printf("Brand: %s\n", myPhone.brand)
	fmt.Printf("Year: %d\n", myPhone.year)
	fmt.Printf("Colors: %v\n", myPhone.colors)
}
