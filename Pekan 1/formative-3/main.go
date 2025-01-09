package main

import (
	"fmt"
	"strconv"
)

// soal1
func Soal1() {

var panjangPersegiPanjang string = "8"
var lebarPersegiPanjang string = "5"
var alasSegitiga string = "6"
var tinggiSegitiga string = "7"

	var num1, err1 = strconv.Atoi(panjangPersegiPanjang)

	if err1 == nil {
		fmt.Println(num1)
	}

	var num2, err2 = strconv.Atoi(lebarPersegiPanjang )

	if err2 == nil {
		fmt.Println(num2)
	}

	var num3, err3 = strconv.Atoi(alasSegitiga)

	if err3 == nil {
		fmt.Println(num3)
	}

	var num4, err4 = strconv.Atoi(tinggiSegitiga)

	if err4 == nil {
		fmt.Println(num4)
	}

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int
	
	//luasPersegiPanjang
	luasPersegiPanjang = num1*num2
	fmt.Println("Luas Persegi Panjang adalah: ", luasPersegiPanjang)

	//kelilingPersegiPanjang
	kelilingPersegiPanjang = 2*(num1 + num2)
	fmt.Println("Keliling Persegi Panjang adalah:", kelilingPersegiPanjang)

	//luasSegitiga
	luasSegitiga = (1/2)*num3*num4
	fmt.Println("Luas Segitiga adalah:", luasSegitiga)

}

//soal 2
func Soal2() {
	var nilaiJohn = 80
	var nilaiDoe = 50

	//nilaiJohn
	fmt.Printf("John mendapatkan nilai dengan ")
	if nilaiJohn >= 80 {
		fmt.Println("Indeks: A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Indeks: B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Indeks: C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Indeks: D")
	} else {
		fmt.Println("Indeks: E")
	}

	//nilaiDoe
	fmt.Printf("Doe mendapatkan nilai dengan ")
	if nilaiDoe >= 80 {
		fmt.Println("Indeks: A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Indeks: B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Indeks: C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Indeks: D")
	} else {
		fmt.Println("Indeks: E")
	}

}

//soal 3
func Soal3() {
	
	var tanggal = 25
	var bulan = 2
	var tahun = 2006
	var x string 

	switch bulan {
		case 1:
			x = "Januari"
		case 2:
			x = "Februari"
		case 3:
			x = "Maret"
		case 4:
			x = "April"
		case 5:
			x = "Mei"
		case 6:
			x = "Juni"
		case 7:
			x = "Juli"
		case 8:
			x = "Agustus"
		case 9:
			x = "September"
		case 10:
			x = "Oktober"
		case 11:
			x = "November"
		case 12:
			x = "Desember"
	}

	fmt.Println(tanggal, x, tahun)

}

//soal 4
func Soal4() {

	var Kelahiran = 2006

	if Kelahiran >= 1944 && Kelahiran <= 1964 {
		fmt.Println("Baby boomer")
	} else if Kelahiran >= 1965 && Kelahiran <= 1979 {
		fmt.Println("Generasi X")
	} else if Kelahiran >= 1980 && Kelahiran <= 1994 {
		fmt.Println("Generasi Y (Millenials)")
	} else if Kelahiran >= 1995 &&  Kelahiran <= 2015 {
		fmt.Println("Generasi Z")
	}
}

func main() {

	fmt.Println("Soal 1")
	Soal1()
	fmt.Println()

	fmt.Println("Soal 2")
	Soal2()
	fmt.Println()

	fmt.Println("Soal 3")
	Soal3()
	fmt.Println()

	fmt.Println("Soal 4")
	Soal4()
	fmt.Println()
	
}

