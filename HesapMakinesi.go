package main

import "fmt"

func toplama(x, y int) (toplam int) {
	toplam = x + y
	return toplam
}
func cikarma(x, y int) (cikarma int) {
	cikarma = x - y
	return cikarma
}
func carpma(x, y int) (carpma int) {
	carpma = x * y
	return carpma
}
func bolme(x, y int) (bolme int) {
	bolme = x / y
	return bolme
}

func main() {
	var islem int
	var sayi1 int
	var sayi2 int
basla:
	fmt.Print("İlk Sayıyı Giriniz : ")
	fmt.Scanf("%d", &sayi1)
	fmt.Print("İkinci Sayıyı Giriniz : ")
	fmt.Scanf("%d", &sayi2)
	fmt.Println("Toplama İşlemi İçin : 1 \n Çıkarma İşlemi İçin : 2 \n Çarpma İşlemi İçin : 3 \n Bölme İşlemi İçin : 4 giriniz")
	fmt.Scanf("%d", &islem)
	switch islem {
	case 1:
		fmt.Println("Toplama İşleminizin Sonucu :", toplama(sayi1, sayi2))
	case 2:
		fmt.Println("Çıkarma İşleminizin Sonucu :", cikarma(sayi1, sayi2))
	case 3:
		fmt.Println("Çarpma İşleminizin Sonucu :", carpma(sayi1, sayi2))
	case 4:
		fmt.Println("Bölme İşleminizin Sonucu :", bolme(sayi1, sayi2))
	default:
		fmt.Println("Hatalı Sayı Girdiniz. Lütfen Yapmak İstediğiniz Hesaplama Türünü DOğru Seçiniz..")
		goto basla
	}
	var tekrar int
	fmt.Println("Başka işlem yapmak ister misiniz? Evet İçin 1, Hayır İçin 2 Giriniz")
	fmt.Scanf("%d", &tekrar)
	if tekrar == 1 {
		goto basla
	}

}
