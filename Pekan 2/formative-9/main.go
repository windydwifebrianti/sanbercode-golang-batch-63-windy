package main

import (
    "fmt"
    "SANBERCODE-GOLANG-BATCH-63-WINDY/Pekan-2/formative-9/database"
)

func main() {

	var bangunDatar database.HitungBangunDatar

	bangunDatar = database.SegitigaSamaSisi{14,7}
	fmt.Println("1. segitigaSamaSisi")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	bangunDatar = database.PersegiPanjang{10,5}
	fmt.Println("2. persegi")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	var bangunRuang database.HitungBangunRuang

	bangunRuang = database.Tabung{7, 10}
	fmt.Println("3. tabung")
	fmt.Println("volume         :", bangunRuang.Volume())
	fmt.Println("luas permukaan :", bangunRuang.LuasPermukaan())

	bangunRuang = database.Balok{5, 8, 9}
	fmt.Println("4. balok")
	fmt.Println("volume         :", bangunRuang.Volume())
	fmt.Println("luas permukaan :", bangunRuang.LuasPermukaan())

	myPhone := database.Phone {
		name : "Samsung Galaxy Note 20",
		brand : "Samsung Galaxy Note 20",
		year : 2020,
		colors : []string{"Mystic Bronze", "Mystic White", "Mystic Black"},

	}
	
	var hape database.Smartphone = myPhone
	fmt.Println(hape.DataSmartphone())

	fmt.Println(database.LuasPersegi(4, true))
	fmt.Println(database.LuasPersegi(8, false))
	fmt.Println(database.LuasPersegi(0, true))
    
    var prefix interface{} = "hasil penjumlahan dari "
    var kumpulanAngkaPertama interface{} = []int{6, 8}
    var kumpulanAngkaKedua interface{} = []int{12, 14}

    prefixStr := prefix.(string)

    angkaPertama := kumpulanAngkaPertama.([]int)
    angkaKedua := kumpulanAngkaKedua.([]int)

    allNumbers := append(angkaPertama, angkaKedua...)

    total := 0
    for _, num := range allNumbers {
        total += num
    }

    numbersStr := fmt.Sprintf("%d + %d + %d + %d", 
        allNumbers[0], allNumbers[1], allNumbers[2], allNumbers[3])

    fmt.Printf("%s%s = %d", prefixStr, numbersStr, total)
}
