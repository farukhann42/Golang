# Golang / Go Programlama Dili
Go, Google'da 2007 yılından itibaren geliştirilmeye başlayan açık kaynak programlama dilidir. Daha çok sistem programlama için tasarlanmış olup, derlenmiş ve statik tipli bir dildir. Kasım 2009'da çıkmıştır.

# Tarihçe
Dil, Kasım 2009'da duyrulmuştur. Google'ın bazı ürünlerinin sistemlerinde olduğu gibi diğer firmalar tarafından da kullanılmaktadır.

Go, Google mühendisleri Robert Griesemer, Rob Pike, ve Ken Thompson tarafından bir deney olarak ortaya çıkarılmış, diğer dillerin bilinen eleştirilerini çözecek ve olumlu özelliklerini koruyacak şekilde tasarlanmıştır. Yeni dil aşağıdaki özellikleri içermekteydi:

Statik yazılmış, büyük sistemlere ölçeklenebilir olması (Java ve C++ gibi)
Üretken ve okunabilir olması, çok fazla zorunlu anahtar kelime ve tekrarlamaların kullanılmaması
Tümleşik geliştirme ortamına ihtiyaç duymaması ancak desteklemesi
Ağ (networking) ve çoklu işlemleri (multiprocessing) desteklemesi

# Dil Tasarımı
Go, C'nin bilinen özelliklerini taşımaktadır ancak yapılan değişiklikler dili basit, kısa ve güvenli hale getirmiştir. Aşağıda Go'yu kısaca tanımlayan özellikler yer almaktadır:

Söz dizimi ve çevresel kalıpları benimsemesi daha çok dinamik dillere benzemektedir:

Değişken tanımındaki tür belirtimi isteğe bağlıdır. (int x = 0; yerine x := 0;).
Hızlı derleme süresi.
Uzak paket yöneticisi (go get) ve online paket dökümantasyonu.
Belirli problemlere ayırt edici yaklaşımlar:
Dahili eş zamanlılık ilkelleri: light-weight process'lar (goroutines), kanallar, ve select ifadesi.
Varsayılan olarak statik olarak bağlanmış native binary'ler, ekstra bağımlılıklara ihtiyaç duyulmadan üretilir.

# Söz Dizimi
Go'nun söz dizimi, C'den yapılan değişiklikleri içerir, kodu kısa ve okunabilir kılmayı amaçlar. Birleştirilmiş declaration ve başlatma operatörü, programcılara i := 3 ya da s := "bazı kelimeler" yazarak, herhangi bir tür belirtimine ihtiyaç duyulmadan değişken tanımlamasını sağlar. Noktalı virgüller hala ifadeleri sonlandırır fakat satır sonlarında kullanılmasına gerek yoktur. Fonksiyonlar birden fazla değerler döndürebilir (return result, err).