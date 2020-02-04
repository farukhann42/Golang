package main

import "fmt"

func sayiAl() [10]int  {
	var sayi int
	var arr[10] int

	fmt.Println("sayıları giriniz : ")

	// sayıları diziye yerleştirelim
	for i:=0; i < 10 ; i++  {
		fmt.Scan(&sayi)
		arr[i] = sayi
	}
	return arr
}

func kontrol(dizi [10]int) [10]int {

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
	return dizi
}

func main() {
	arr := sayiAl()

	// dizideki elemanların sıralanmamış hali
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Print("\n")

	// dizideki elemanların sıralanmış hali
	fmt.Print(kontrol(arr))

}
