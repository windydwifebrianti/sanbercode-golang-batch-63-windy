package main

import "fmt"

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luasSegitiga() {
	luasSegitiga := (s.alas * s.tinggi) / 2
	fmt.Println("Luas Segitiga: ", luasSegitiga)
}

func (s persegi) luasPersegi() {
	luasPersegi := (s.sisi * s.sisi)
	fmt.Println("Luas Persegi: ", luasPersegi)
}

func (s persegiPanjang) luasPersegiPanjang() {
	luasPersegiPanjang := (s.panjang * s.lebar)
	fmt.Println("Luas Persegi Panjang: ", luasPersegiPanjang)
}

func main() {

  var segitigaA = segitiga{4, 10}
	segitigaA.luasSegitiga()

  var persegiA = persegi{8}
	persegiA.luasPersegi()

	var persegiPanjangA = persegiPanjang{5, 10}
	persegiPanjangA.luasPersegiPanjang()

}
