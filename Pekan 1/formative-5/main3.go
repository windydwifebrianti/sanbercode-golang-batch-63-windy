package main

import "fmt"

func buahFavorit(name string, buah ...string) (result string) {

	var buahList string
	for i, b := range buah {

		// penambahan quote sebelum nama buah awal "semangka
		if i == 0 {
			buahList += "\""
		} 


		// penambahan nama buah 
		buahList += b
		
		// penambahan quote, koma, dan spasi setelah buah awal  semangka", 
		if i < len(buah) - 1 {
			buahList += "\", \""
		}

		// penambahan quote pada setelah buah akhir pepaya"
		
		if i == len(buah) - 1 {
			buahList += "\""
		} 

	}
	result = fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", name, buahList)
	return
}
func main() {
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)
}

// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
