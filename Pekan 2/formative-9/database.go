package database

import (
	"fmt"
	"math"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
	return (p.Panjang * p.Lebar)
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {

	return 2*math.Pi*math.Pow(t.JariJari, 2) + 2*math.Pi*t.JariJari*t.Tinggi
}

type Balok struct {
	Panjang, Lebar, Tinggi float64
}

func (b Balok) Volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}

func (b Balok) LuasPermukaan() float64 {

	return 2 * ((b.Panjang * b.Lebar) + (b.Panjang * b.Tinggi) + (b.Lebar * b.Tinggi))
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type Smartphone interface {
	DataSmartphone() string
}

func (p Phone) DataSmartphone() string {
	return fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.Name, p.Brand, p.Year, p.Colors)
}

func LuasPersegi(Sisi int, WithMessage bool) interface{} {

	if Sisi == 0 {
		if WithMessage {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	Luas := Sisi * Sisi

	if WithMessage {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", Sisi, Luas)
	}
	return Luas
}