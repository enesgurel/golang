package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Enes"
	fmt.Println(s.HasPrefix(isim, "Ene"))
	fmt.Println(s.HasSuffix(isim, "nes"))
	fmt.Println(s.Index(isim, "es"))
	harfler := []string{"e", "n", "e", "s"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", 3))
	fmt.Println(s.Split(sonuc, "*"))
	fmt.Println(s.Repeat(sonuc, 5))
}
