package main

import "fmt"

func main() {
    
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