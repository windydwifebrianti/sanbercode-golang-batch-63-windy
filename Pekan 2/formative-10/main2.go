package main

import ( 
	"fmt"
	"errors"
)

func kelilingSegitigaSamaSisi (sisi int, message bool) (string, error) {
	if (sisi == 0) && (message == true) {
		return 0, errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if (sisi == 0) && (message == false) {
		return 0, errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}
}

fmt.Println(kelilingSegitigaSamaSisi(4, true))

fmt.Println(kelilingSegitigaSamaSisi(8, false))

fmt.Println(kelilingSegitigaSamaSisi(0, true))

fmt.Println(kelilingSegitigaSamaSisi(0, false))