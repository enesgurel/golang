package maps

import "fmt"

func Demo1() {
	// key- value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayisi : ", len(sozluk))
	fmt.Println("Sozluk : ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman Sayisi : ", len(sozluk))
	fmt.Println("Sozluk : ", sozluk)

	deger, varmi := sozluk["pencil"]
	fmt.Println(deger)
	fmt.Println("Listede Olma Durumu : ", varmi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)
}
