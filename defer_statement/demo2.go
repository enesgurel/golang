package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		return "Çift Sayidir"
	}

	if sayi%2 != 0 {
		fmt.Println("Tek sayi calisti")
		return "Tek Sayidir"
	}
	return "Belli Degil"
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println("Sonuç :", sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
