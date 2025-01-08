package main

import (
	"fmt"
	"strconv"
	"strings"
)

// soal 1
func Soal1() {

	var text1 string = "Bootcamp"
	var text2 string = "Digital"
	var text3 string = "Skill"
	var text4 string = "Sanbercode"
	var text5 string = "Golang"

	fmt.Println(text1, text2, text3, text4, text5)

}

// soal 2
func Soal2() {

	halo := "Halo Dunia"
	find := "Dunia"
	var replaceWith = "Golang"

	var newText1 = strings.Replace(halo, find, replaceWith, 1)
	fmt.Println(newText1)

}

// soal3
func Soal3() {

	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)
}

// soal4
func Soal4() {

	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var num1, err1 = strconv.Atoi(angkaPertama)

	if err1 == nil {
		fmt.Println(num1)
	}

	var num2, err2 = strconv.Atoi(angkaKedua)

	if err2 == nil {
		fmt.Println(num2)
	}

	var num3, err3 = strconv.Atoi(angkaKetiga)

	if err3 == nil {
		fmt.Println(num3)
	}

	var num4, err4 = strconv.Atoi(angkaKeempat)

	if err4 == nil {
		fmt.Println(num4)
	}

	var sum int
	sum = num1 + num2 + num3 + num4
	fmt.Println("Hasil penjumlahannya adalah: ", sum)
}

// soal 5
func Soal5() {

	kalimat := "halo halo bandung"
	angka := 2021
	find := "halo halo"
	ReplaceWith := "Hi Hi"

	str := strconv.Itoa(angka)

	var newText1 = strings.Replace(kalimat, find, ReplaceWith, -1)
	var newText2 = "\"" + newText1 + "\"" + " - " + str

	fmt.Println(newText2)

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

	fmt.Println("Soal 5")
	Soal5()
	
}
