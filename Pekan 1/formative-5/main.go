//soal1

package main

import "fmt"

func luasPersegiPanjang(panjang int, lebar int) int {
	luas := panjang * lebar
	return luas
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	keliling := 2 * (panjang + lebar)
	return keliling
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	volume := panjang * lebar * tinggi
	return volume
}

func main() {
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
}
