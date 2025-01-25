package main

import "fmt"

func logging(kalimat string, tahun int) (string, int) {
	fmt.Printf("Kalimat: %s\nTahun: %d\n", kalimat, tahun)
	return  "Golang Backend Development", 2021
}

func runApplication() {
	defer logging("Golang Backend Development", 2021)
}

func main() {
	runApplication()
}
