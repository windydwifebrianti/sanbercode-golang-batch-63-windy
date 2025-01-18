package main

import (
	"fmt"
	"math"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

type persegiPanjang struct {
	panjang, lebar int
}

func (p persegiPanjang) luas() int {
	return (p.panjang * p.lebar)
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {

	return 2*math.Pi*math.Pow(t.jariJari, 2) + 2*math.Pi*t.jariJari*t.tinggi
}

type balok struct {
	panjang, lebar, tinggi float64
}

func (b balok) volume() float64 {
	return b.panjang * b.lebar * b.tinggi
}

func (b balok) luasPermukaan() float64 {

	return 2 * ((b.panjang * b.lebar) + (b.panjang * b.tinggi) + (b.lebar * b.tinggi))
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func main() {

	var bangunDatar hitungBangunDatar

	bangunDatar = segitigaSamaSisi{14,7}
	fmt.Println("1. segitigaSamaSisi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = persegiPanjang{10,5}
	fmt.Println("2. persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	var bangunRuang hitungBangunRuang

	bangunRuang = tabung{7, 10}
	fmt.Println("3. tabung")
	fmt.Println("volume         :", bangunRuang.volume())
	fmt.Println("luas permukaan :", bangunRuang.luasPermukaan())

	bangunRuang = balok{5, 8, 9}
	fmt.Println("4. balok")
	fmt.Println("volume         :", bangunRuang.volume())
	fmt.Println("luas permukaan :", bangunRuang.luasPermukaan())


}
