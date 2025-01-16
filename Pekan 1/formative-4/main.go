package main

import "fmt"

//soal 1
func Soal1() {

	var angka int

	for angka = 1; angka <= 20; angka++ {

		if (angka%2 == 1) && (angka%3 != 0) {
			fmt.Println(angka, "- Santai")
		} else if angka%2 == 0 {
			fmt.Println(angka, "- Berkualitas")
		} else if (angka%3 == 0) && (angka%2 == 1) {
			fmt.Println(angka, "- I Love Coding")
		}
	}


}

//soal 2
func Soal2() {

	for i := 0; i <= 7; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

//soal 3
func Soal3() {

	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var kalimatBaru = kalimat[2:]

	fmt.Println(kalimatBaru)
}

//soal 4
func Soal4() {

	var sayuran = []string{
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}

	// using for loop
	for i := 0; i < len(sayuran); i++ {
		fmt.Print(i+1, ".", " ", sayuran[i], "\n")
	}
}

//soal 5
func Soal5() {

	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	satuan["panjang"] = 7
	satuan["lebar"] = 4
	satuan["tinggi"] = 6

	var VolumeBalok int
	VolumeBalok = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	fmt.Println("panjang", "=", satuan["panjang"])
	fmt.Println("lebar", "=", satuan["lebar"])
	fmt.Println("tinggi", "=", satuan["tinggi"])
	fmt.Println("Volume Balok", "=", VolumeBalok)

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
	fmt.Println()

}
