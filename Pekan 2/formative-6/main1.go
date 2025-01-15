//soal 1

package main

import (
	"fmt"
	"math"
)

func main() {
	var luasLingkaran float64
	var kelilingLingkaran float64

	updateLingkaran := func(luas *float64, keliling *float64, jariJari float64) {
		*luas = math.Pi * math.Pow(jariJari, 2)
		*keliling = 2 * math.Pi * jariJari
	}

	updateLingkaran(&luasLingkaran, &kelilingLingkaran, 7
	)

	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

}
