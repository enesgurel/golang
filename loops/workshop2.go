package loops

import "fmt"

func Demo4() {
	sayi := 0
	fmt.Println("Bir Sayi Giriniz : ")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi {
		fmt.Println("Asaldır")
	} else {
		fmt.Println("Asal Değildir")
	}
}
