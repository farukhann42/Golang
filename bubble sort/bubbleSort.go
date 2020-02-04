package main

import "fmt"

func main() {
	var sayi int
	var dizi [10]int

	// sayıları diziye yerleştirelim
	for k:=0; k < len(dizi) ; k++ {
		fmt.Scan(&sayi)
		dizi[k] = sayi
	}

	// dizideki elemanların sıralanmamış hali
	for i := 0; i < len(dizi); i++ {
		fmt.Print(dizi[i], " ")
	}
	fmt.Print("\n")

	for {
		dizilim := true
		for i := 0; i < len(dizi)-1; i++ {
			if dizi[i] > dizi[i+1] {
				tmp := dizi[i]
				dizi[i] = dizi[i+1]
				dizi[i+1] = tmp
				dizilim = false
			}
		}
		if dizilim == true {
			break
		}
	}
	// dizideki elemanların sıralanmış hali
	for i := 0; i < len(dizi); i++ {
		fmt.Print(dizi[i], " ")
	}
	fmt.Print("\n")
}
