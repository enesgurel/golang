package loops

import "fmt"

func Demo3() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 10
	tahminSayisi := 0

	fmt.Println("Bir sayi tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi++

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük sayi tahmin et")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi++
		}

		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha kucuk bir sayi tahmin et")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi++
		}
	}

	basariDurumu := ""
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Super"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Guzel"
	} else {
		basariDurumu = "Fena Değil"
	}

	fmt.Printf("Bravo, Bildiniz. %v tahmin: %v", tahminSayisi, basariDurumu)

}
