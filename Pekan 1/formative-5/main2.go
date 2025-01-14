package main

import "fmt"

func introduce(name string, gender string, job string, age string) (result string) {

	if gender == "laki-laki" {
		// gender diubah menjadi Pak
		gender = "Pak"
	} else if gender == "perempuan" {
		gender = "Bu"
	}

	result = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", gender, name, job, age)
	return
}

func main() {
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

}
