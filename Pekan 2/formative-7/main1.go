package main

import "fmt"

type Buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

func main() {
	nanas := Buah{
		nama:       "Nanas",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      9000,
	}

	jeruk := Buah{
		nama:       "Jeruk",
		warna:      "Oranye",
		adaBijinya: true,
		harga:      8000,
	}

	semangka := Buah{
		nama:       "Semangka",
		warna:      "Hijau & Merah",
		adaBijinya: true,
		harga:      10000,
	}

	pisang := Buah{
		nama:       "Pisang",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      5000,
	}

	fmt.Println("Daftar Buah:")

	fmt.Printf("\n1. %s\n", nanas.nama)
	fmt.Printf("   Warna: %s\n", nanas.warna)
	fmt.Printf("   Ada Biji: %s\n", getAdaBiji(nanas.adaBijinya))
	fmt.Printf("   Harga: Rp%d\n", nanas.harga)

	fmt.Printf("\n2. %s\n", jeruk.nama)
	fmt.Printf("   Warna: %s\n", jeruk.warna)
	fmt.Printf("   Ada Biji: %s\n", getAdaBiji(jeruk.adaBijinya))
	fmt.Printf("   Harga: Rp%d\n", jeruk.harga)

	fmt.Printf("\n3. %s\n", semangka.nama)
	fmt.Printf("   Warna: %s\n", semangka.warna)
	fmt.Printf("   Ada Biji: %s\n", getAdaBiji(semangka.adaBijinya))
	fmt.Printf("   Harga: Rp%d\n", semangka.harga)

	fmt.Printf("\n4. %s\n", pisang.nama)
	fmt.Printf("   Warna: %s\n", pisang.warna)
	fmt.Printf("   Ada Biji: %s\n", getAdaBiji(pisang.adaBijinya))
	fmt.Printf("   Harga: Rp%d\n", pisang.harga)
}

func getAdaBiji(ada bool) string {
	if ada {
		return "Ada"
	}
	return "Tidak"
}
