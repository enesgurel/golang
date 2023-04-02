package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Para Yetersiz!")
	} else if cekilmekIstenen == hesap {
		fmt.Printf("Paraniz Hazirlaniyor!")
		fmt.Printf("Dikkat Hesapta Para Kalmadi")
	} else {
		fmt.Println("Paraniz Hazirlaniyor!")
	}
}
