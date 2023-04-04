package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayi giriniz")
	}

	if tahmin > aklimdakiSayi {
		return "Daha kücük bir sayi giriniz", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayi giriniz", nil
	}

	return "Bildiniz", nil
}

func Demo2() {
	mesaj, hata := TahminEt(101)
	fmt.Println(mesaj)
	fmt.Println(hata)
}
