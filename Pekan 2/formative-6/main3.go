//soal 3

package main

import "fmt"

func addFruits(fruits *[]string, newFruits ...string) {
	*fruits = append(*fruits, newFruits...)
}

func main() {
	var buah = []string{}
	addFruits(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i := 0; i < len(buah); i++ {
		fmt.Print(i+1, ".", " ", buah[i], "\n")
	}
}
