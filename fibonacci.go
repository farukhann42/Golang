package main

import "fmt"

func main() {
	
	//ilk elemanın dizilerdeki karşılığı 0'dır yani fibonacci[5] = {0,1,2,3,4} diye bir 
		//dizideki 0 sayısına fibonacci[0] ile ulaşılır. 
	
	var sayi int
	var i int
	
	fmt.Print ("Fibonacci Dizisinin Eleman Sayısını Giriniz : ")
	fmt.Scan(&sayi)

	fibonacci := make([]int, sayi)  // klavyeden girilen sayı kadar elemanlı bir fibonacci dizisi tanımlandı  

	fibonacci[0] = 0  // fibonacci dizisinin ilk hanesinin değeri 0 olarak ayarlandı
	fibonacci[1] = 1  // fibonacci dizisinin ikinci hanesinin değeri 1 olarak ayarlandı

	fmt.Println (fibonacci[0]) // atanılan değerler tek seferlik olmak üzere ekrana yazdırılır
	fmt.Println (fibonacci[1])

	for i=2; i < sayi; i++ {  // for döngüsü ile fibonacci dizisinin ikinci hanesinden başlanması istenildi ve 
												//son hane olan fibonacci[girilen sayıya kadar] kadar gitmesi istendi

		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]  // 2.hanesinin değerini ondan önceki 2 hane ile toplayarak elde etmesi istenildi
		fmt.Println (fibonacci[i]) // ekrana yazdırıldı
	}

}