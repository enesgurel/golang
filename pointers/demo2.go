package pointers

import "fmt"

func Demo2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demoda ki sayi : ", sayilar[0])
}
