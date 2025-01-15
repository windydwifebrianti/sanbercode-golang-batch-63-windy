// soal 2

package main

import "fmt"

func introduce(sentence *string, name, gender, job, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, job, age)
	} else if gender == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, job, age)
	}
}

func main() {

	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

}
