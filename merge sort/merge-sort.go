package main

import (
	"fmt"
)

func main() {

	dizi := diziElemanlariAl(10)
	fmt.Println("Sıralanmamış Dizi : ", dizi)

	fmt.Println(sort(dizi))

}

func diziElemanlariAl(sayi int) []int {
	var dizi []int = make([]int, sayi)
	for i := 0; i < sayi; i++ {
		fmt.Print(i+1, ". Sayıyı Giriniz : ")
		fmt.Scan(&dizi[i])
	}

	return dizi
}

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left, right := bol(arr)
	left = sort(left)
	right = sort(right)
	return merge(left, right)
}

// diziyi ikiye ayırmak için bol fonksiyonu kullanıldı
func bol(arr []int) ([]int, []int) {
	ortaSayi := int(len(arr) / 2)
	left := arr[:ortaSayi]
	right := arr[ortaSayi:]
	return left, right
}

func merge(solDizi, sagDizi []int) []int {
	arr := make([]int, len(solDizi)+len(sagDizi))

	j, k := 0, 0

	for i := 0; i < len(arr); i++ {
		if j >= len(solDizi) {
			arr[i] = sagDizi[k]
			k++
			continue
		} else if k >= len(sagDizi) {
			arr[i] = solDizi[j]
			j++
			continue
		}

		if solDizi[j] > sagDizi[k] {
			arr[i] = sagDizi[k]
			k++
		} else {
			arr[i] = solDizi[j]
			j++
		}
	}

	return arr
}
